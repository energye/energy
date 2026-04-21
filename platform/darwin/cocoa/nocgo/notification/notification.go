//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package notification

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/platform/notification/types"
	"github.com/energye/lcl/rtl/version"
	"sync"
	"time"
	"unsafe"
)

const UNNotificationActionOptionDestructive = 1 << 0

const (
	appleDefaultActionID = "com.apple.UNNotificationDefaultActionIdentifier"
	authorizationTimeout = 180 * time.Second
	operationTimeout     = 5 * time.Second
)

var (
	once                                                                          sync.Once
	gNotification                                                                 *Notification
	notificationDelegateClass                                                     objc.Class
	notificationDelegate                                                          objc.ID
	sel_userNotificationCenterWillPresentNotificationWithCompletionHandler        objc.SEL
	sel_userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler objc.SEL
)

type Notification struct {
	notificationResultCallback TNotificationResponseEvent
	callbackLock               sync.RWMutex
	callbackChannels           map[int]chan callbackResult
	mu                         sync.Mutex
	nextID                     int
	isInitialized              bool
}

type callbackResult struct {
	ok  bool
	err error
}

func New() INotification {
	once.Do(func() {
		gNotification = &Notification{
			callbackChannels: make(map[int]chan callbackResult),
			nextID:           0,
		}
		if err := gNotification.Initialize(); err != nil {
			println("Notification service initialization warning:", err.Error())
		}
	})
	return gNotification
}

func init() {
	initSelectors()
	registerNotificationDelegateClass()
}

func initSelectors() {
	sel_userNotificationCenterWillPresentNotificationWithCompletionHandler = objc.RegisterName("userNotificationCenter:willPresentNotification:withCompletionHandler:")
	sel_userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler = objc.RegisterName("userNotificationCenter:didReceiveNotificationResponse:withCompletionHandler:")
}

func registerNotificationDelegateClass() {
	var err error
	protoUNUserNotificationCenterDelegate := objc.GetProtocol("UNUserNotificationCenterDelegate")
	notificationDelegateClass, err = objc.RegisterClass(
		"EnergyNotificationDelegate", objc.GetClass("NSObject"),
		[]*objc.Protocol{protoUNUserNotificationCenterDelegate},
		nil,
		[]objc.MethodDef{
			{
				Cmd: sel_userNotificationCenterWillPresentNotificationWithCompletionHandler,
				Fn:  userNotificationCenterWillPresentNotificationWithCompletionHandler,
			},
			{
				Cmd: sel_userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler,
				Fn:  userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler,
			},
		},
	)

	if err != nil {
		panic(err)
	}
}

// userNotificationCenterWillPresentNotificationWithCompletionHandler 处理通知即将展示时的回调
// 根据 macOS 版本选择合适的展示选项，并调用完成处理器以允许通知显示
func userNotificationCenterWillPresentNotificationWithCompletionHandler(self objc.ID, _cmd objc.SEL,
	center objc.ID, notification objc.ID, completionHandler objc.Block) {
	var options uintptr
	if available(11, 0) {
		const UNNotificationPresentationOptionList = 1 << 2
		const UNNotificationPresentationOptionBanner = 1 << 3
		const UNNotificationPresentationOptionSound = 1 << 1
		options = UNNotificationPresentationOptionList | UNNotificationPresentationOptionBanner | UNNotificationPresentationOptionSound
	} else {
		const UNNotificationPresentationOptionAlert = 1 << 0
		const UNNotificationPresentationOptionSound = 1 << 1
		options = UNNotificationPresentationOptionAlert | UNNotificationPresentationOptionSound
	}
	if _, err := BlockInvoke(completionHandler, options); err != nil {
		println("Failed to invoke presentNotification completionHandler:", err.Error())
	}
}

// userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler 处理用户与通知交互后的回调
// 提取通知响应数据并序列化为 JSON，通过 Go 层回调传递结果，最后调用完成处理器
func userNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(self objc.ID, _cmd objc.SEL,
	center objc.ID, response objc.ID, completionHandler objc.Block) {

	payload := buildNotificationResponsePayload(response)
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		errorMsg := fmt.Sprintf("JSON serialization failed: %s", err.Error())
		didReceiveNotificationResponse("", errorMsg)
	} else {
		didReceiveNotificationResponse(string(jsonBytes), "")
	}
	if _, err := BlockInvoke(completionHandler); err != nil {
		println("Failed to invoke didReceive completionHandler:", err.Error())
	}
}

// Initialize 初始化通知服务
// 执行环境检查、Bundle 标识验证以及通知中心代理的设置，确保通知功能可用
func (m *Notification) Initialize() error {
	if m.isInitialized {
		return nil
	}

	if !available(10, 15) {
		//Debug("notifications unavailable: requires macOS 10.15 or later")
		return errors.New("notifications unavailable: requires macOS 10.15 or later")
	}

	if !checkBundleIdentifier() {
		//Debug("missing bundle identifier: package app as .app bundle with Info.plist")
		return errors.New("missing bundle identifier: package app as .app bundle with Info.plist")
	}

	if !initializeNotificationCenter() {
		//Debug("failed to initialize notification delegate")
		return errors.New("failed to initialize notification delegate")
	}

	m.isInitialized = true
	return nil
}

// RequestNotificationAuthorization 请求用户授权以发送通知
// 该方法会触发系统授权弹窗，并等待用户响应或超时
func (m *Notification) RequestNotificationAuthorization() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authorizationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	requestNotificationAuthorization(id)

	select {
	case result := <-ch:
		//Debug("RequestNotificationAuthorization OK:", result.ok, "err:", result.err)
		return result.ok, result.err
	case <-ctx.Done():
		return false, fmt.Errorf("authorization request timeout: %w", ctx.Err())
	}
}

// CheckNotificationAuthorization 检查当前应用的通知授权状态
// 该方法不会触发系统弹窗，仅查询用户当前的授权设置
func (m *Notification) CheckNotificationAuthorization() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	checkNotificationAuthorization(id)

	select {
	case result := <-ch:
		// Debug("CheckNotificationAuthorization OK:", result.ok, "err:", result.err)
		return result.ok, result.err
	case <-ctx.Done():
		return false, fmt.Errorf("authorization check timeout: %w", ctx.Err())
	}
}

// SendNotification 发送本地通知
// 根据提供的选项构建并展示通知，支持标题、副标题、正文及附加数据
func (m *Notification) SendNotification(opts Options) error {
	// Debug("SendNotification:", opts)
	if err := validateNotificationOptions(opts); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	sendNotification(id, opts)

	select {
	case result := <-ch:
		if !result.ok {
			if result.err != nil {
				return result.err
			}
			return fmt.Errorf("notification delivery failed")
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("send operation timeout: %w", ctx.Err())
	}
}

// SendNotificationWithActions 发送带有交互动作的本地通知
// 该通知支持用户点击预设按钮或进行文本回复，需预先注册对应的 Category
func (m *Notification) SendNotificationWithActions(opts Options) error {
	if err := validateNotificationOptions(opts); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	sendNotificationWithActions(id, opts)

	select {
	case result := <-ch:
		if !result.ok {
			if result.err != nil {
				return result.err
			}
			return fmt.Errorf("notification with actions delivery failed")
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("send with actions timeout: %w", ctx.Err())
	}
}

// RegisterNotificationCategory 注册通知交互类别
// 定义通知的可用操作按钮或文本回复功能，需在发送带交互的通知前调用
func (m *Notification) RegisterNotificationCategory(category Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	registerNotificationCategory(id, category)

	select {
	case result := <-ch:
		if !result.ok {
			if result.err != nil {
				return result.err
			}
			return fmt.Errorf("category registration failed")
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("registration timeout: %w", ctx.Err())
	}
}

func (m *Notification) RemoveNotificationCategory(categoryID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	removeNotificationCategory(id, categoryID)

	select {
	case result := <-ch:
		if !result.ok {
			if result.err != nil {
				return result.err
			}
			return fmt.Errorf("category removal failed")
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("removal timeout: %w", ctx.Err())
	}
}

func (m *Notification) RemoveAllPendingNotifications() error {
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))
	center.Send(objc.RegisterName("removeAllPendingNotificationRequests"))
	return nil
}

func (m *Notification) RemovePendingNotification(identifier string) error {
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	cIdentifier := GoStringToNSString(identifier)

	arrayClass := objc.GetClass("NSArray")
	identifiers := objc.ID(arrayClass).Send(objc.RegisterName("arrayWithObject:"), cIdentifier)
	center.Send(objc.RegisterName("removePendingNotificationRequestsWithIdentifiers:"), identifiers)
	return nil
}

func (m *Notification) RemoveAllDeliveredNotifications() error {
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))
	center.Send(objc.RegisterName("removeAllDeliveredNotifications"))
	return nil
}

func (m *Notification) RemoveDeliveredNotification(identifier string) error {
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	cIdentifier := GoStringToNSString(identifier)

	arrayClass := objc.GetClass("NSArray")
	identifiers := objc.ID(arrayClass).Send(objc.RegisterName("arrayWithObject:"), cIdentifier)
	center.Send(objc.RegisterName("removeDeliveredNotificationsWithIdentifiers:"), identifiers)
	return nil
}

func (m *Notification) RemoveNotification(_ string) error {
	return nil
}

func (m *Notification) SetOnNotificationResponse(callback TNotificationResponseEvent) {
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.notificationResultCallback = callback
}

func (m *Notification) handleNotificationResult(result Result) {
	m.callbackLock.RLock()
	callback := m.notificationResultCallback
	m.callbackLock.RUnlock()

	if callback != nil {
		callback(result)
	}
}

func (m *Notification) allocateChannel() (int, chan callbackResult) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := m.nextID
	m.nextID++
	ch := make(chan callbackResult, 1)
	m.callbackChannels[id] = ch
	return id, ch
}

func (m *Notification) GetChannel(id int) (chan callbackResult, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	ch, exists := m.callbackChannels[id]
	if exists {
		delete(m.callbackChannels, id)
	}
	return ch, exists
}

func (m *Notification) releaseChannel(id int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if ch, exists := m.callbackChannels[id]; exists {
		delete(m.callbackChannels, id)
		close(ch)
	}
}

func didReceiveNotificationResponse(jsonPayload string, errMsg string) {
	result := Result{}
	if errMsg != "" {
		result.Error = fmt.Errorf("response error: %s", errMsg)
		dispatchResult(result)
		return
	}

	var response Response
	if parseErr := json.Unmarshal([]byte(jsonPayload), &response); parseErr != nil {
		result.Error = fmt.Errorf("response parsing failed: %w", parseErr)
		dispatchResult(result)
		return
	}

	if response.ActionIdentifier == appleDefaultActionID {
		response.ActionIdentifier = DefaultActionIdentifier
	}

	result.Response = response
	dispatchResult(result)
}

func dispatchResult(result Result) {
	if gNotification != nil {
		gNotification.handleNotificationResult(result)
	}
}

func createSendCompletionBlock(channelID int, identifier string) objc.Block {
	block := objc.NewBlock(func(self objc.Block, errorPtr unsafe.Pointer) {
		// Debug("SendCompletionBlock", "self:", self, "error:", uintptr(errorPtr))
		if errorPtr != nil {
			errorObj := objc.ID(errorPtr)
			localizedDescription := errorObj.Send(objc.RegisterName("localizedDescription"))
			errMsg := NSStringToGoString(localizedDescription)
			onCallbackResult(channelID, false, fmt.Sprintf("Send failed: %s", errMsg))
		} else {
			onCallbackResult(channelID, true, "")
		}
	})
	return block
}

func checkNotificationAuthorization(channelID int) {
	if !initializeNotificationCenter() {
		onCallbackResult(channelID, false, "checkNotificationAuthorization.initializeNotificationCenter: Failed to initialize the notification center")
		return
	}
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	completionBlock := createSettingsCompletionBlock(channelID)
	defer completionBlock.Release()

	center.Send(objc.RegisterName("getNotificationSettingsWithCompletionHandler:"), completionBlock)
}

func createSettingsCompletionBlock(channelID int) objc.Block {
	block := objc.NewBlock(
		func(self objc.Block, settingsPtr unsafe.Pointer) {
			settings := objc.ID(settingsPtr)
			authorizationStatus := settings.Send(objc.RegisterName("authorizationStatus"))

			const UNAuthorizationStatusAuthorized = 2
			authorized := authorizationStatus == UNAuthorizationStatusAuthorized

			onCallbackResult(channelID, authorized, "")
		},
	)
	return block
}

func onCallbackResult(channelID int, success bool, errMsg string) {
	if gNotification == nil {
		return
	}
	ch, exists := gNotification.GetChannel(channelID)
	if !exists {
		return
	}
	var opErr error
	if errMsg != "" {
		opErr = fmt.Errorf("%s", errMsg)
	}

	ch <- callbackResult{
		ok:  success,
		err: opErr,
	}
	close(ch)
}

func requestNotificationAuthorization(channelID int) {
	if !initializeNotificationCenter() {
		onCallbackResult(channelID, false, "requestNotificationAuthorization.initializeNotificationCenter: Failed to initialize the notification center")
		return
	}

	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	const UNAuthorizationOptionAlert = 1 << 0
	const UNAuthorizationOptionSound = 1 << 1
	const UNAuthorizationOptionBadge = 1 << 2
	requestedOptions := UNAuthorizationOptionAlert | UNAuthorizationOptionSound | UNAuthorizationOptionBadge

	completionBlock := createAuthorizationCompletionBlock(channelID)
	defer completionBlock.Release()

	center.Send(objc.RegisterName("requestAuthorizationWithOptions:completionHandler:"), requestedOptions, completionBlock)
}

func createAuthorizationCompletionBlock(channelID int) objc.Block {
	block := objc.NewBlock(func(self objc.Block, granted bool, errorPtr unsafe.Pointer) {
		var errMsg string
		if errorPtr != nil {
			errorObj := objc.ID(errorPtr)
			localizedDescription := errorObj.Send(objc.RegisterName("localizedDescription"))
			errMsg = NSStringToGoString(localizedDescription)
		}
		// Debug("AuthorizationCompletionBlock", "channelID:", channelID, "granted:", granted, "errMsg:", errMsg)
		onCallbackResult(channelID, granted, errMsg)
	})
	return block
}

func available(major, minor int) bool {
	if version.OSVersion.Major > major {
		return true
	}
	if version.OSVersion.Major == major && version.OSVersion.Minor >= minor {
		return true
	}
	return false
}

func checkBundleIdentifier() bool {
	nsBundleClass := objc.GetClass("NSBundle")
	mainBundle := objc.ID(nsBundleClass).Send(objc.RegisterName("mainBundle"))
	bundleID := mainBundle.Send(objc.RegisterName("bundleIdentifier"))
	//Debug("bundle identifier:", NSStringToGoString(bundleID))
	return bundleID != 0 && NSStringToGoString(bundleID) != ""
}

func initializeNotificationCenter() bool {
	if notificationDelegate == 0 {
		notificationDelegate = objc.ID(notificationDelegateClass).Send(objc.RegisterName("new"))
		if notificationDelegate == 0 {
			return false
		}
	}
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))
	if center == 0 {
		return false
	}
	center.Send(objc.RegisterName("setDelegate:"), notificationDelegate)
	return true
}

func buildNotificationResponsePayload(response objc.ID) map[string]any {
	payload := make(map[string]any)

	notification := response.Send(objc.RegisterName("notification"))
	request := notification.Send(objc.RegisterName("request"))
	content := request.Send(objc.RegisterName("content"))

	identifier := NSStringToGoString(request.Send(objc.RegisterName("identifier")))
	actionIdentifier := NSStringToGoString(response.Send(objc.RegisterName("actionIdentifier")))
	title := NSStringToGoString(content.Send(objc.RegisterName("title")))
	body := NSStringToGoString(content.Send(objc.RegisterName("body")))

	payload["id"] = identifier
	payload["actionIdentifier"] = actionIdentifier
	payload["title"] = title
	payload["body"] = body

	categoryIdentifier := NSStringToGoString(content.Send(objc.RegisterName("categoryIdentifier")))
	if categoryIdentifier != "" {
		payload["categoryIdentifier"] = categoryIdentifier
	}

	subtitle := NSStringToGoString(content.Send(objc.RegisterName("subtitle")))
	if subtitle != "" {
		payload["subtitle"] = subtitle
	}

	userInfo := content.Send(objc.RegisterName("userInfo"))
	if userInfo != 0 {
		count := userInfo.Send(objc.RegisterName("count"))
		if count > 0 {
			userInfoMap := convertNSDictionaryToGoMap(userInfo)
			if len(userInfoMap) > 0 {
				payload["userInfo"] = userInfoMap
			}
		}
	}

	textInputResponseClass := objc.GetClass("UNTextInputNotificationResponse")
	if response.Send(objc.RegisterName("isKindOfClass:"), textInputResponseClass) != 0 {
		userText := NSStringToGoString(response.Send(objc.RegisterName("userText")))
		if userText != "" {
			payload["userText"] = userText
		}
	}

	return payload
}

func validateNotificationOptions(options Options) error {
	if options.ID == "" {
		return fmt.Errorf("notification ID cannot be empty")
	}
	if options.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	return nil
}

func removeNotificationCategory(channelID int, categoryID string) {
	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	completionBlock := createRemoveCategoryCompletionBlock(channelID, categoryID)
	defer completionBlock.Release()

	center.Send(objc.RegisterName("getNotificationCategoriesWithCompletionHandler:"), completionBlock)
}

func createRemoveCategoryCompletionBlock(channelID int, categoryID string) objc.Block {
	block := objc.NewBlock(
		func(self objc.Block, categoriesPtr unsafe.Pointer) {
			existingCategories := objc.ID(categoriesPtr)
			setClass := objc.GetClass("NSMutableSet")
			updatedCategories := objc.ID(setClass).Send(objc.RegisterName("setWithSet:"), existingCategories)

			allObjects := updatedCategories.Send(objc.RegisterName("allObjects"))
			count := allObjects.Send(objc.RegisterName("count"))

			found := false
			for i := uintptr(0); i < uintptr(count); i++ {
				cat := allObjects.Send(objc.RegisterName("objectAtIndex:"), i)
				identifier := NSStringToGoString(cat.Send(objc.RegisterName("identifier")))
				if identifier == categoryID {
					updatedCategories.Send(objc.RegisterName("removeObject:"), cat)
					found = true
					break
				}
			}

			if found {
				unCenterClass := objc.GetClass("UNUserNotificationCenter")
				center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))
				center.Send(objc.RegisterName("setNotificationCategories:"), updatedCategories)
				onCallbackResult(channelID, true, "")
			} else {
				onCallbackResult(channelID, false, fmt.Sprintf("Category not found: %s", categoryID))
			}
		},
	)
	return block
}

func registerNotificationCategory(channelID int, category Category) {
	if !initializeNotificationCenter() {
		onCallbackResult(channelID, false, "registerNotificationCategory.initializeNotificationCenter: Failed to initialize the notification center")
		return
	}

	actionsArray := buildNotificationActions(category.Actions, category.HasReplyField, category.ReplyPlaceholder, category.ReplyButtonTitle)

	categoryClass := objc.GetClass("UNNotificationCategory")
	cCategoryID := GoStringToNSString(category.ID)
	categoryObj := objc.ID(categoryClass).Send(objc.RegisterName("categoryWithIdentifier:actions:intentIdentifiers:options:"),
		cCategoryID, actionsArray, 0, 0)

	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	completionBlock := createCategoriesCompletionBlock(channelID, category.ID, categoryObj)
	defer completionBlock.Release()

	center.Send(objc.RegisterName("getNotificationCategoriesWithCompletionHandler:"), completionBlock)
}

func buildNotificationActions(actions []Action, hasReplyField bool, replyPlaceholder, replyButtonTitle string) objc.ID {
	actionsArray := objc.ID(objc.GetClass("NSMutableArray")).Send(objc.RegisterName("array"))

	for _, action := range actions {
		if action.ID == "" || action.Title == "" {
			continue
		}
		actionClass := objc.GetClass("UNNotificationAction")
		cActionID := GoStringToNSString(action.ID)
		cActionTitle := GoStringToNSString(action.Title)
		opts := 0
		if action.Destructive {
			opts = UNNotificationActionOptionDestructive
		}
		newAction := objc.ID(actionClass).Send(objc.RegisterName("actionWithIdentifier:title:options:"),
			cActionID, cActionTitle, opts)
		actionsArray.Send(objc.RegisterName("addObject:"), newAction)
	}

	if hasReplyField && replyPlaceholder != "" && replyButtonTitle != "" {
		textInputClass := objc.GetClass("UNTextInputNotificationAction")
		const UNNotificationActionOptionNone = 0
		cActionWithIdentifier := GoStringToNSString("TEXT_REPLY")
		cReplyButtonTitle := GoStringToNSString(replyButtonTitle)
		cReplyPlaceholder := GoStringToNSString(replyPlaceholder)
		textInputAction := objc.ID(textInputClass).Send(objc.RegisterName("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"),
			cActionWithIdentifier, cReplyButtonTitle, UNNotificationActionOptionNone, cReplyButtonTitle, cReplyPlaceholder)
		actionsArray.Send(objc.RegisterName("addObject:"), textInputAction)
	} else {
		actionClass := objc.GetClass("UNNotificationAction")
		for _, action := range actions {
			var opts uintptr
			if action.Destructive {
				opts = UNNotificationActionOptionDestructive
			}
			cActionID := GoStringToNSString(action.ID)
			cActionTitle := GoStringToNSString(action.Title)
			actionObj := objc.ID(actionClass).Send(objc.RegisterName("actionWithIdentifier:title:options:"),
				cActionID, cActionTitle, opts)
			actionsArray.Send(objc.RegisterName("addObject:"), actionObj)
		}
	}

	return actionsArray
}

func createCategoriesCompletionBlock(channelID int, categoryID string, newCategory objc.ID) objc.Block {
	block := objc.NewBlock(
		func(self objc.Block, categoriesPtr unsafe.Pointer) {
			existingCategories := objc.ID(categoriesPtr)

			setClass := objc.GetClass("NSMutableSet")
			updatedCategories := objc.ID(setClass).Send(objc.RegisterName("setWithSet:"), existingCategories)

			allObjects := updatedCategories.Send(objc.RegisterName("allObjects"))
			count := allObjects.Send(objc.RegisterName("count"))

			for i := uintptr(0); i < uintptr(count); i++ {
				cat := allObjects.Send(objc.RegisterName("objectAtIndex:"), i)
				identifier := NSStringToGoString(cat.Send(objc.RegisterName("identifier")))
				if identifier == categoryID {
					updatedCategories.Send(objc.RegisterName("removeObject:"), cat)
					break
				}
			}

			updatedCategories.Send(objc.RegisterName("addObject:"), newCategory)

			unCenterClass := objc.GetClass("UNUserNotificationCenter")
			center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))
			center.Send(objc.RegisterName("setNotificationCategories:"), updatedCategories)

			onCallbackResult(channelID, true, "")
		},
	)
	return block
}

func sendNotificationWithActions(channelID int, opts Options) {
	if !initializeNotificationCenter() {
		onCallbackResult(channelID, false, "sendNotificationWithActions.initializeNotificationCenter: Failed to initialize the notification center")
		return
	}

	content, err := buildNotificationContent(opts.Title, opts.Subtitle, opts.Body, opts.Data)
	if err != nil {
		onCallbackResult(channelID, false, fmt.Sprintf("Content creation failed: %s", err.Error()))
		return
	}

	if opts.CategoryID != "" {
		cOptionCategoryID := GoStringToNSString(opts.CategoryID)
		content.Send(objc.RegisterName("setCategoryIdentifier:"), cOptionCategoryID)
	}

	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	requestClass := objc.GetClass("UNNotificationRequest")
	cOptionID := GoStringToNSString(opts.ID)
	request := objc.ID(requestClass).Send(objc.RegisterName("requestWithIdentifier:content:trigger:"),
		cOptionID, content, 0)

	completionBlock := createSendCompletionBlock(channelID, opts.ID)
	defer completionBlock.Release()

	center.Send(objc.RegisterName("addNotificationRequest:withCompletionHandler:"), request, completionBlock)
}

func sendNotification(channelID int, opts Options) {
	// Debug("sendNotification", "channelID:", channelID, "opts:", opts)
	if !initializeNotificationCenter() {
		onCallbackResult(channelID, false, "sendNotification.initializeNotificationCenter: Failed to initialize the notification center")
		return
	}

	cContent, err := buildNotificationContent(opts.Title, opts.Subtitle, opts.Body, opts.Data)
	if err != nil {
		onCallbackResult(channelID, false, fmt.Sprintf("Content creation failed: %s", err.Error()))
		return
	}

	unCenterClass := objc.GetClass("UNUserNotificationCenter")
	center := objc.ID(unCenterClass).Send(objc.RegisterName("currentNotificationCenter"))

	requestClass := objc.GetClass("UNNotificationRequest")
	cOptionID := GoStringToNSString(opts.ID)
	// Debug("sendNotification", "cOptionID:", cOptionID)
	request := objc.ID(requestClass).Send(objc.RegisterName("requestWithIdentifier:content:trigger:"),
		cOptionID, cContent, 0)

	completionBlock := createSendCompletionBlock(channelID, opts.ID)
	defer completionBlock.Release()
	// Debug("sendNotification", "completionBlock:", completionBlock)

	center.Send(objc.RegisterName("addNotificationRequest:withCompletionHandler:"),
		request, completionBlock)
	// Debug("sendNotification", "end")
}

func buildNotificationContent(title, subtitle, body string, data map[string]any) (objc.ID, error) {
	// Debug("buildNotificationContent", "title:", title, "subtitle:", subtitle, "body:", body, "data:", data)
	contentClass := objc.GetClass("UNMutableNotificationContent")
	content := objc.ID(contentClass).Send(objc.RegisterName("new"))

	cTitle := GoStringToNSString(title)
	cBody := GoStringToNSString(body)
	content.Send(objc.RegisterName("setTitle:"), cTitle)
	content.Send(objc.RegisterName("setBody:"), cBody)

	soundClass := objc.GetClass("UNNotificationSound")
	defaultSound := objc.ID(soundClass).Send(objc.RegisterName("defaultSound"))
	content.Send(objc.RegisterName("setSound:"), defaultSound)

	if subtitle != "" {
		cSubtitle := GoStringToNSString(subtitle)
		content.Send(objc.RegisterName("setSubtitle:"), cSubtitle)
	}

	if len(data) > 0 {
		userInfo := convertGoMapToNSDictionary(data)
		if userInfo != 0 {
			content.Send(objc.RegisterName("setUserInfo:"), userInfo)
		}
	}

	return content, nil
}
