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
		"EnergyNotificationDelegate",
		objc.GetClass("NSObject"),
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

func buildNotificationResponsePayload(response objc.ID) map[string]interface{} {
	payload := make(map[string]interface{})

	notification := response.Send(objc.RegisterName("notification"))
	request := notification.Send(objc.RegisterName("request"))
	content := request.Send(objc.RegisterName("content"))

	identifier := getStringFromObjC(request.Send(objc.RegisterName("identifier")))
	actionIdentifier := getStringFromObjC(response.Send(objc.RegisterName("actionIdentifier")))
	title := getStringFromObjC(content.Send(objc.RegisterName("title")))
	body := getStringFromObjC(content.Send(objc.RegisterName("body")))

	payload["id"] = identifier
	payload["actionIdentifier"] = actionIdentifier
	payload["title"] = title
	payload["body"] = body

	categoryIdentifier := getStringFromObjC(content.Send(objc.RegisterName("categoryIdentifier")))
	if categoryIdentifier != "" {
		payload["categoryIdentifier"] = categoryIdentifier
	}

	subtitle := getStringFromObjC(content.Send(objc.RegisterName("subtitle")))
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
		userText := getStringFromObjC(response.Send(objc.RegisterName("userText")))
		if userText != "" {
			payload["userText"] = userText
		}
	}

	return payload
}

func getStringFromObjC(nsString objc.ID) string {
	if nsString == 0 {
		return ""
	}
	typeStr := nsString.Send(objc.RegisterName("UTF8String"))
	if typeStr == 0 {
		return ""
	}
	utf8Length := nsString.Send(objc.RegisterName("lengthOfBytesUsingEncoding:"), 4)
	if utf8Length == 0 {
		return ""
	}
	value := unsafe.String((*byte)(unsafe.Pointer(typeStr)), uintptr(utf8Length))
	return value
}

func convertNSDictionaryToGoMap(dict objc.ID) map[string]interface{} {
	result := make(map[string]interface{})

	allKeys := dict.Send(objc.RegisterName("allKeys"))
	count := allKeys.Send(objc.RegisterName("count"))

	for i := uintptr(0); i < uintptr(count); i++ {
		key := allKeys.Send(objc.RegisterName("objectAtIndex:"), i)
		value := dict.Send(objc.RegisterName("objectForKey:"), key)
		keyStr := getStringFromObjC(key)
		switch {
		case value.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSString")) != 0:
			result[keyStr] = getStringFromObjC(value)
		case value.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSNumber")) != 0:
			objCType := getStringFromObjC(value.Send(objc.RegisterName("objCType")))
			if objCType == "c" || objCType == "B" {
				result[keyStr] = value.Send(objc.RegisterName("boolValue")) != 0
			} else if objCType == "i" || objCType == "l" || objCType == "q" {
				result[keyStr] = value.Send(objc.RegisterName("intValue"))
			} else if objCType == "f" || objCType == "d" {
				result[keyStr] = value.Send(objc.RegisterName("doubleValue"))
			}
		case value.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSArray")) != 0:
			arr := convertNSArrayToGoSlice(value)
			result[keyStr] = arr
		case value.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSDictionary")) != 0:
			nestedMap := convertNSDictionaryToGoMap(value)
			result[keyStr] = nestedMap
		default:
			result[keyStr] = getStringFromObjC(value)
		}
	}

	return result
}

func convertNSArrayToGoSlice(arr objc.ID) []interface{} {
	count := arr.Send(objc.RegisterName("count"))
	result := make([]interface{}, 0, count)
	for i := uintptr(0); i < uintptr(count); i++ {
		item := arr.Send(objc.RegisterName("objectAtIndex:"), i)
		switch {
		case item.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSString")) != 0:
			result = append(result, getStringFromObjC(item))
		case item.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSNumber")) != 0:
			objCType := getStringFromObjC(item.Send(objc.RegisterName("objCType")))
			if objCType == "c" || objCType == "B" {
				result = append(result, item.Send(objc.RegisterName("boolValue")) != 0)
			} else if objCType == "i" || objCType == "l" || objCType == "q" {
				result = append(result, item.Send(objc.RegisterName("intValue")))
			} else if objCType == "f" || objCType == "d" {
				result = append(result, item.Send(objc.RegisterName("doubleValue")))
			}
		case item.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSArray")) != 0:
			result = append(result, convertNSArrayToGoSlice(item))
		case item.Send(objc.RegisterName("isKindOfClass:"), objc.GetClass("NSDictionary")) != 0:
			result = append(result, convertNSDictionaryToGoMap(item))
		default:
			result = append(result, getStringFromObjC(item))
		}
	}

	return result
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
	//Debug("bundle identifier:", getStringFromObjC(bundleID))
	return bundleID != 0 && getStringFromObjC(bundleID) != ""
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
			errMsg = getStringFromObjC(localizedDescription)
		}
		// Debug("AuthorizationCompletionBlock", "channelID:", channelID, "granted:", granted, "errMsg:", errMsg)
		onCallbackResult(channelID, granted, errMsg)
	})
	return block
}

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

func buildNotificationContent(title, subtitle, body string, data map[string]interface{}) (objc.ID, error) {
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

func convertGoMapToNSDictionary(data map[string]interface{}) objc.ID {
	dictClass := objc.GetClass("NSMutableDictionary")
	dict := objc.ID(dictClass).Send(objc.RegisterName("dictionary"))

	numClass := objc.GetClass("NSNumber")
	for key, value := range data {
		nsKey := GoStringToNSString(key)
		var nsValue objc.ID
		switch v := value.(type) {
		case string:
			nsValue = GoStringToNSString(v)
		case bool:
			if v {
				nsValue = objc.ID(numClass).Send(objc.RegisterName("numberWithBool:"), 1)
			} else {
				nsValue = objc.ID(numClass).Send(objc.RegisterName("numberWithBool:"), 0)
			}
		case int:
			nsValue = objc.ID(numClass).Send(objc.RegisterName("numberWithInt:"), v)
		case int64:
			nsValue = objc.ID(numClass).Send(objc.RegisterName("numberWithLongLong:"), v)
		case float64:
			nsValue = objc.ID(numClass).Send(objc.RegisterName("numberWithDouble:"), v)
		case map[string]interface{}:
			nsValue = convertGoMapToNSDictionary(v)
		default:
			continue
		}

		if nsValue != 0 {
			dict.Send(objc.RegisterName("setObject:forKey:"), nsValue, nsKey)
		}
	}

	return dict
}

func createSendCompletionBlock(channelID int, identifier string) objc.Block {
	block := objc.NewBlock(func(self objc.Block, errorPtr unsafe.Pointer) {
		// Debug("SendCompletionBlock", "self:", self, "error:", uintptr(errorPtr))
		if errorPtr != nil {
			errorObj := objc.ID(errorPtr)
			localizedDescription := errorObj.Send(objc.RegisterName("localizedDescription"))
			errMsg := getStringFromObjC(localizedDescription)
			onCallbackResult(channelID, false, fmt.Sprintf("Send failed: %s", errMsg))
		} else {
			onCallbackResult(channelID, true, "")
		}
	})
	return block
}

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
	arrayClass := objc.GetClass("NSMutableArray")
	actionsArray := objc.ID(arrayClass).Send(objc.RegisterName("array"))

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
				const UNNotificationActionOptionDestructive = 1 << 0
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
				identifier := getStringFromObjC(cat.Send(objc.RegisterName("identifier")))
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
				identifier := getStringFromObjC(cat.Send(objc.RegisterName("identifier")))
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

func validateNotificationOptions(options Options) error {
	if options.ID == "" {
		return fmt.Errorf("notification ID cannot be empty")
	}
	if options.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	return nil
}

func GoStringToNSString(str string) objc.ID {
	return objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"), str)
}

func NSStringToGoString(nsString objc.ID) string {
	typeStr := nsString.Send(objc.RegisterName("UTF8String"))
	utf8Length := nsString.Send(objc.RegisterName("lengthOfBytesUsingEncoding:"), 4)
	value := unsafe.String((*byte)(unsafe.Pointer(typeStr)), uintptr(utf8Length))
	return value
}
