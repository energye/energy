//go:build windows

package notification

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	. "github.com/energye/energy/v3/platform/notification/types"
	"github.com/energye/energy/v3/platform/win32/go-toast"
	"github.com/energye/energy/v3/platform/win32/go-toast/wintoast"
	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

var (
	once          sync.Once
	gNotification *Notification
)

const (
	ToastRegistryPath                  = `Software\Classes\AppUserModelId\`
	ToastRegistryGuidKey               = "CustomActivator"
	NotificationCategoriesRegistryPath = `SOFTWARE\%s\NotificationCategories`
	NotificationCategoriesRegistryKey  = "Categories"
)

// NotificationPayload combines the action ID and user data into a single structure
type NotificationPayload struct {
	Action  string  `json:"action"`
	Options Options `json:"payload,omitempty"`
}

// Notification implements INotification interface for Windows
type Notification struct {
	categories                 map[string]Category
	categoriesLock             sync.RWMutex
	appName                    string
	appGUID                    string
	iconPath                   string
	exePath                    string
	notificationResultCallback TNotificationResponseEvent
	callbackLock               sync.RWMutex
}

// New creates a new Notification instance
func New() INotification {
	once.Do(func() {
		gNotification = &Notification{
			categories: make(map[string]Category),
		}
		// 自动初始化
		if err := gNotification.Initialize(); err != nil {
			fmt.Printf("[Energy] Notification service initialization warning: %v\n", err)
		}
	})
	return gNotification
}

// Initialize sets up the notification service
func (n *Notification) Initialize() error {
	n.categoriesLock.Lock()
	defer n.categoriesLock.Unlock()

	n.appName = "TestAPPName" //TODO 需要动态配置

	guid, err := n.getGUID()
	if err != nil {
		return fmt.Errorf("failed to get GUID: %w", err)
	}
	n.appGUID = guid

	n.iconPath = filepath.Join(os.TempDir(), n.appName+n.appGUID+".png")

	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	n.exePath = exe

	// Create the registry key for the toast activator
	key, _, err := registry.CreateKey(registry.CURRENT_USER,
		`Software\Classes\CLSID\`+n.appGUID+`\LocalServer32`, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to create CLSID key: %w", err)
	}

	if err := key.SetStringValue("", fmt.Sprintf("\"%s\" %%1", n.exePath)); err != nil {
		key.Close()
		return fmt.Errorf("failed to set CLSID server path: %w", err)
	}
	key.Close()

	if err := toast.SetAppData(toast.AppData{
		AppID:         n.appName,
		GUID:          guid,
		IconPath:      n.iconPath,
		ActivationExe: n.exePath,
	}); err != nil {
		return fmt.Errorf("failed to set app data: %w", err)
	}

	toast.SetActivationCallback(func(args string, data []toast.UserData) {
		result := Result{}
		actionIdentifier, options, err := parseNotificationResponse(args)
		if err != nil {
			result.Error = err
		} else {
			response := Response{
				ID:               options.ID,
				ActionIdentifier: actionIdentifier,
				Title:            options.Title,
				Subtitle:         options.Subtitle,
				Body:             options.Body,
				CategoryID:       options.CategoryID,
				UserInfo:         options.Data,
			}
			if userText, found := n.getUserText(data); found {
				response.UserText = userText
			}
			result.Response = response
		}
		n.handleNotificationResult(result)
	})

	if err := wintoast.RegisterClassFactory(wintoast.ClassFactory); err != nil {
		return fmt.Errorf("RegisterClassFactory failed: %w", err)
	}
	if err := n.loadCategoriesFromRegistry(); err != nil {
		return fmt.Errorf("failed to load categories: %w", err)
	}

	return nil
}

// RequestNotificationAuthorization is a Windows stub that always returns true, nil.
// (user authorization is macOS-specific)
func (n *Notification) RequestNotificationAuthorization() (bool, error) {
	return true, nil
}

// CheckNotificationAuthorization is a Windows stub that always returns true.
// (user authorization is macOS-specific)
func (n *Notification) CheckNotificationAuthorization() (bool, error) {
	return true, nil
}

// SendNotification sends a basic notification with a name, title, and body.
func (n *Notification) SendNotification(options Options) error {
	if err := validateNotificationOptions(options); err != nil {
		return err
	}
	if err := n.saveIconToDir(); err != nil {
		fmt.Printf("Error saving icon: %v\n", err)
	}
	notif := toast.Notification{
		Title:               options.Title,
		Body:                options.Body,
		ActivationType:      toast.Foreground,
		ActivationArguments: DefaultActionIdentifier,
	}

	encodedPayload, err := n.encodePayload(DefaultActionIdentifier, options)
	if err != nil {
		return fmt.Errorf("failed to encode notification payload: %w", err)
	}
	notif.ActivationArguments = encodedPayload

	return notif.Push()
}

// SendNotificationWithActions sends a notification with additional actions and inputs.
func (n *Notification) SendNotificationWithActions(options Options) error {
	if err := validateNotificationOptions(options); err != nil {
		return err
	}

	if err := n.saveIconToDir(); err != nil {
		fmt.Printf("Error saving icon: %v\n", err)
	}

	n.categoriesLock.RLock()
	nCategory, categoryExists := n.categories[options.CategoryID]
	n.categoriesLock.RUnlock()

	if options.CategoryID == "" || !categoryExists {
		fmt.Printf("Category '%s' not found, sending basic notification without actions\n", options.CategoryID)
	}

	notif := toast.Notification{
		Title:               options.Title,
		Body:                options.Body,
		ActivationType:      toast.Foreground,
		ActivationArguments: DefaultActionIdentifier,
	}

	for _, action := range nCategory.Actions {
		notif.Actions = append(notif.Actions, toast.Action{
			Content:   action.Title,
			Arguments: action.ID,
		})
	}

	if nCategory.HasReplyField {
		notif.Inputs = append(notif.Inputs, toast.Input{
			ID:          "userText",
			Placeholder: nCategory.ReplyPlaceholder,
		})

		notif.Actions = append(notif.Actions, toast.Action{
			Content:   nCategory.ReplyButtonTitle,
			Arguments: "TEXT_REPLY",
			InputID:   "userText",
		})
	}

	encodedPayload, err := n.encodePayload(notif.ActivationArguments, options)
	if err != nil {
		return fmt.Errorf("failed to encode notification payload: %w", err)
	}
	notif.ActivationArguments = encodedPayload

	for index := range notif.Actions {
		encodedPayload, err := n.encodePayload(notif.Actions[index].Arguments, options)
		if err != nil {
			return fmt.Errorf("failed to encode notification payload: %w", err)
		}
		notif.Actions[index].Arguments = encodedPayload
	}

	return notif.Push()
}

// RegisterNotificationCategory registers a new NotificationCategory.
func (n *Notification) RegisterNotificationCategory(category Category) error {
	n.categoriesLock.Lock()
	defer n.categoriesLock.Unlock()

	n.categories[category.ID] = Category{
		ID:               category.ID,
		Actions:          category.Actions,
		HasReplyField:    category.HasReplyField,
		ReplyPlaceholder: category.ReplyPlaceholder,
		ReplyButtonTitle: category.ReplyButtonTitle,
	}

	return n.saveCategoriesToRegistry()
}

// RemoveNotificationCategory removes a previously registered NotificationCategory.
func (n *Notification) RemoveNotificationCategory(categoryID string) error {
	n.categoriesLock.Lock()
	defer n.categoriesLock.Unlock()

	delete(n.categories, categoryID)

	return n.saveCategoriesToRegistry()
}

// RemoveAllPendingNotifications is a Windows stub that always returns nil.
func (n *Notification) RemoveAllPendingNotifications() error {
	return nil
}

// RemovePendingNotification is a Windows stub that always returns nil.
func (n *Notification) RemovePendingNotification(_ string) error {
	return nil
}

// RemoveAllDeliveredNotifications is a Windows stub that always returns nil.
func (n *Notification) RemoveAllDeliveredNotifications() error {
	return nil
}

// RemoveDeliveredNotification is a Windows stub that always returns nil.
func (n *Notification) RemoveDeliveredNotification(_ string) error {
	return nil
}

// RemoveNotification is a Windows stub that always returns nil.
func (n *Notification) RemoveNotification(_ string) error {
	return nil
}

// SetOnNotificationResponse registers notification response callback
func (n *Notification) SetOnNotificationResponse(callback TNotificationResponseEvent) {
	n.callbackLock.Lock()
	defer n.callbackLock.Unlock()
	n.notificationResultCallback = callback
}

// handleNotificationResult processes notification result
func (n *Notification) handleNotificationResult(result Result) {
	n.callbackLock.RLock()
	callback := n.notificationResultCallback
	n.callbackLock.RUnlock()

	if callback != nil {
		callback(result)
	}
}

// encodePayload combines an action ID and user data into a single encoded string
func (n *Notification) encodePayload(actionID string, options Options) (string, error) {
	payload := NotificationPayload{
		Action:  actionID,
		Options: options,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return actionID, err
	}
	encodedPayload := base64.StdEncoding.EncodeToString(jsonData)
	return encodedPayload, nil
}

// decodePayload extracts the action ID and user data from an encoded payload
func decodePayload(encodedString string) (string, Options, error) {
	jsonData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return encodedString, Options{}, fmt.Errorf("failed to decode base64 payload: %w", err)
	}

	var payload NotificationPayload
	if err := json.Unmarshal(jsonData, &payload); err != nil {
		return encodedString, Options{}, fmt.Errorf("failed to unmarshal notification payload: %w", err)
	}

	return payload.Action, payload.Options, nil
}

// parseNotificationResponse parses notification response
func parseNotificationResponse(response string) (action string, options Options, err error) {
	actionID, options, err := decodePayload(response)

	if err != nil {
		fmt.Printf("Warning: Failed to decode notification response: %v\n", err)
		return response, Options{}, err
	}

	return actionID, options, nil
}

func (n *Notification) saveIconToDir() error {
	// 注意: 这里需要使用 w32 包获取图标，但原代码中引用的函数可能不存在
	// 使用伪代码标记，实际实现需要根据可用的 API 调整
	// icon, err := application.NewIconFromResource(w32.GetModuleHandle(""), uint16(3))
	// if err != nil {
	// 	return fmt.Errorf("failed to retrieve application icon: %w", err)
	// }
	// return w32.SaveHIconAsPNG(icon, n.iconPath)
	return nil
}

func (n *Notification) saveCategoriesToRegistry() error {
	registryPath := fmt.Sprintf(NotificationCategoriesRegistryPath, n.appName)

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		registryPath,
		registry.ALL_ACCESS,
	)
	if err != nil {
		return err
	}
	defer key.Close()

	data, err := json.Marshal(n.categories)
	if err != nil {
		return err
	}

	return key.SetStringValue(NotificationCategoriesRegistryKey, string(data))
}

func (n *Notification) loadCategoriesFromRegistry() error {
	registryPath := fmt.Sprintf(NotificationCategoriesRegistryPath, n.appName)

	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		registryPath,
		registry.QUERY_VALUE,
	)
	if err != nil {
		if err == registry.ErrNotExist {
			return nil
		}
		return fmt.Errorf("failed to open registry key: %w", err)
	}
	defer key.Close()

	data, _, err := key.GetStringValue(NotificationCategoriesRegistryKey)
	if err != nil {
		if err == registry.ErrNotExist {
			return nil
		}
		return fmt.Errorf("failed to read categories from registry: %w", err)
	}

	categories := make(map[string]Category)
	if err := json.Unmarshal([]byte(data), &categories); err != nil {
		return fmt.Errorf("failed to parse notification categories from registry: %w", err)
	}

	n.categories = categories

	return nil
}

func (n *Notification) getUserText(data []toast.UserData) (string, bool) {
	for _, d := range data {
		if d.Key == "userText" {
			return d.Value, true
		}
	}
	return "", false
}

func (n *Notification) getGUID() (string, error) {
	keyPath := ToastRegistryPath + n.appName

	k, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.QUERY_VALUE)
	if err == nil {
		guid, _, err := k.GetStringValue(ToastRegistryGuidKey)
		k.Close()
		if err == nil && guid != "" {
			return guid, nil
		}
	}
	guid := generateGUID()

	k, _, err = registry.CreateKey(registry.CURRENT_USER, keyPath, registry.WRITE)
	if err != nil {
		return "", fmt.Errorf("failed to create registry key: %w", err)
	}
	defer k.Close()

	if err := k.SetStringValue(ToastRegistryGuidKey, guid); err != nil {
		return "", fmt.Errorf("failed to write GUID to registry: %w", err)
	}

	return guid, nil
}

func generateGUID() string {
	guid := uuid.New()
	return fmt.Sprintf("{%s}", guid.String())
}

// validateNotificationOptions validates notification options
func validateNotificationOptions(options Options) error {
	if options.ID == "" {
		return fmt.Errorf("notification ID cannot be empty")
	}
	if options.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	return nil
}
