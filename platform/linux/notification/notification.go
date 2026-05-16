// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0

//go:build linux

package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v3/application/pack"
	. "github.com/energye/energy/v3/platform/notification/types"
	"github.com/godbus/dbus/v5"
	"os"
	"path/filepath"
	"sync"
)

type Notification struct {
	conn              *dbus.Conn
	categories        map[string]Category
	categoriesLock    sync.RWMutex
	notifications     map[uint32]*notificationData
	notificationsLock sync.RWMutex
	appName           string
	cancel            context.CancelFunc
	callback          TNotificationResponseEvent
	callbackLock      sync.RWMutex
}

type notificationData struct {
	ID         string
	Title      string
	Subtitle   string
	Body       string
	CategoryID string
	Data       map[string]interface{}
	DBusID     uint32
	ActionMap  map[string]string
}

const (
	dbusNotificationInterface = "org.freedesktop.Notifications"
	dbusNotificationPath      = "/org/freedesktop/Notifications"
)

var (
	once          sync.Once
	gNotification INotification
)

// New creates a new Notification instance
func New() INotification {
	once.Do(func() {
		impl := &Notification{
			categories:    make(map[string]Category),
			notifications: make(map[uint32]*notificationData),
		}
		gNotification = impl

		if err := impl.Initialize(); err != nil {
			fmt.Printf("[Energy] Notification service initialization warning: %v\n", err)
		}
	})
	return gNotification
}

// Initialize sets up the notification service
func (m *Notification) Initialize() error {
	name := pack.Info.Name
	if name == "" {
		name = "ENERGY APP"
	}
	m.appName = name

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to session bus: %w", err)
	}
	m.conn = conn

	if err := m.loadCategories(); err != nil {
		fmt.Printf("Failed to load notification categories: %v\n", err)
	}

	var signalCtx context.Context
	signalCtx, m.cancel = context.WithCancel(context.Background())

	if err := m.setupSignalHandling(signalCtx); err != nil {
		return fmt.Errorf("failed to set up notification signal handling: %w", err)
	}

	return nil
}

// RequestNotificationAuthorization is a Linux stub that always returns true, nil.
// (authorization is macOS-specific)
func (m *Notification) RequestNotificationAuthorization() (bool, error) {
	return true, nil
}

// CheckNotificationAuthorization is a Linux stub that always returns true.
// (authorization is macOS-specific)
func (m *Notification) CheckNotificationAuthorization() (bool, error) {
	return true, nil
}

// SendNotification sends a basic notification with a unique identifier, title, subtitle, and body.
func (m *Notification) SendNotification(options Options) error {
	if err := validateNotificationOptions(options); err != nil {
		return err
	}

	hints := map[string]dbus.Variant{}

	body := options.Body
	if options.Subtitle != "" {
		body = options.Subtitle + "\n" + body
	}

	defaultActionID := "default"
	actions := []string{defaultActionID, "Default"}

	actionMap := map[string]string{
		defaultActionID: DefaultActionIdentifier,
	}

	hints["x-notification-id"] = dbus.MakeVariant(options.ID)

	if options.Data != nil {
		userData, err := json.Marshal(options.Data)
		if err == nil {
			hints["x-user-data"] = dbus.MakeVariant(string(userData))
		}
	}

	obj := m.conn.Object(dbusNotificationInterface, dbusNotificationPath)
	call := obj.Call(dbusNotificationInterface+".Notify", 0, m.appName, uint32(0), "",
		options.Title, body, actions, hints, int32(-1))

	if call.Err != nil {
		return fmt.Errorf("failed to send notification: %w", call.Err)
	}

	var dbusID uint32
	if err := call.Store(&dbusID); err != nil {
		return fmt.Errorf("failed to store notification ID: %w", err)
	}

	notification := &notificationData{
		ID:        options.ID,
		Title:     options.Title,
		Subtitle:  options.Subtitle,
		Body:      options.Body,
		Data:      options.Data,
		DBusID:    dbusID,
		ActionMap: actionMap,
	}

	m.notificationsLock.Lock()
	m.notifications[dbusID] = notification
	m.notificationsLock.Unlock()

	return nil
}

// SendNotificationWithActions sends a notification with additional actions.
func (m *Notification) SendNotificationWithActions(options Options) error {
	if err := validateNotificationOptions(options); err != nil {
		return err
	}

	m.categoriesLock.RLock()
	category, exists := m.categories[options.CategoryID]
	m.categoriesLock.RUnlock()

	if options.CategoryID == "" || !exists {
		return m.SendNotification(options)
	}

	body := options.Body
	if options.Subtitle != "" {
		body = options.Subtitle + "\n" + body
	}

	var actions []string
	actionMap := make(map[string]string)

	defaultActionID := "default"
	actions = append(actions, defaultActionID, "Default")
	actionMap[defaultActionID] = DefaultActionIdentifier

	for _, action := range category.Actions {
		actions = append(actions, action.ID, action.Title)
		actionMap[action.ID] = action.ID
	}

	hints := map[string]dbus.Variant{}

	hints["x-notification-id"] = dbus.MakeVariant(options.ID)

	hints["x-category-id"] = dbus.MakeVariant(options.CategoryID)

	if options.Data != nil {
		userData, err := json.Marshal(options.Data)
		if err == nil {
			hints["x-user-data"] = dbus.MakeVariant(string(userData))
		}
	}

	obj := m.conn.Object(dbusNotificationInterface, dbusNotificationPath)
	call := obj.Call(dbusNotificationInterface+".Notify", 0, m.appName, uint32(0), "",
		options.Title, body, actions, hints, int32(-1))

	if call.Err != nil {
		return fmt.Errorf("failed to send notification: %w", call.Err)
	}

	var dbusID uint32
	if err := call.Store(&dbusID); err != nil {
		return fmt.Errorf("failed to store notification ID: %w", err)
	}

	notification := &notificationData{
		ID:         options.ID,
		Title:      options.Title,
		Subtitle:   options.Subtitle,
		Body:       options.Body,
		CategoryID: options.CategoryID,
		Data:       options.Data,
		DBusID:     dbusID,
		ActionMap:  actionMap,
	}

	m.notificationsLock.Lock()
	m.notifications[dbusID] = notification
	m.notificationsLock.Unlock()

	return nil
}

// RegisterNotificationCategory registers a new NotificationCategory to be used with SendNotificationWithActions.
func (m *Notification) RegisterNotificationCategory(category Category) error {
	m.categoriesLock.Lock()
	defer m.categoriesLock.Unlock()

	m.categories[category.ID] = category

	if err := m.saveCategories(); err != nil {
		fmt.Printf("Failed to save notification categories: %v\n", err)
	}

	return nil
}

// RemoveNotificationCategory removes a previously registered NotificationCategory.
func (m *Notification) RemoveNotificationCategory(categoryId string) error {
	m.categoriesLock.Lock()
	defer m.categoriesLock.Unlock()

	delete(m.categories, categoryId)

	if err := m.saveCategories(); err != nil {
		fmt.Printf("Failed to save notification categories: %v\n", err)
	}

	return nil
}

// RemoveAllPendingNotifications attempts to remove all active notifications.
func (m *Notification) RemoveAllPendingNotifications() error {
	m.notificationsLock.Lock()
	dbusIDs := make([]uint32, 0, len(m.notifications))
	for id := range m.notifications {
		dbusIDs = append(dbusIDs, id)
	}
	m.notificationsLock.Unlock()

	for _, id := range dbusIDs {
		m.closeNotification(id)
	}

	return nil
}

// RemovePendingNotification removes a pending notification.
func (m *Notification) RemovePendingNotification(identifier string) error {
	var dbusID uint32
	found := false

	m.notificationsLock.Lock()
	for id, notif := range m.notifications {
		if notif.ID == identifier {
			dbusID = id
			found = true
			break
		}
	}
	m.notificationsLock.Unlock()

	if !found {
		return nil
	}

	return m.closeNotification(dbusID)
}

// RemoveAllDeliveredNotifications functionally equivalent to RemoveAllPendingNotification on Linux.
func (m *Notification) RemoveAllDeliveredNotifications() error {
	return m.RemoveAllPendingNotifications()
}

// RemoveDeliveredNotification functionally equivalent RemovePendingNotification on Linux.
func (m *Notification) RemoveDeliveredNotification(identifier string) error {
	return m.RemovePendingNotification(identifier)
}

// RemoveNotification removes a notification by identifier.
func (m *Notification) RemoveNotification(identifier string) error {
	return m.RemovePendingNotification(identifier)
}

// SetOnNotificationResponse registers notification response callback
func (m *Notification) SetOnNotificationResponse(callback TNotificationResponseEvent) {
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.callback = callback
}

// handleNotificationResult processes notification result
func (m *Notification) handleNotificationResult(result Result) {
	m.callbackLock.RLock()
	callback := m.callback
	m.callbackLock.RUnlock()

	if callback != nil {
		callback(result)
	}
}

// Helper method to close a notification.
func (m *Notification) closeNotification(id uint32) error {
	obj := m.conn.Object(dbusNotificationInterface, dbusNotificationPath)
	call := obj.Call(dbusNotificationInterface+".CloseNotification", 0, id)

	if call.Err != nil {
		return fmt.Errorf("failed to close notification: %w", call.Err)
	}

	m.notificationsLock.Lock()
	delete(m.notifications, id)
	m.notificationsLock.Unlock()

	return nil
}

func (m *Notification) getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config directory: %w", err)
	}

	appConfigDir := filepath.Join(configDir, m.appName)
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create app config directory: %w", err)
	}

	return appConfigDir, nil
}

// Save notification categories.
func (m *Notification) saveCategories() error {
	configDir, err := m.getConfigDir()
	if err != nil {
		return err
	}

	categoriesFile := filepath.Join(configDir, "notification-categories.json")

	categoriesData, err := json.MarshalIndent(m.categories, "", "  ")

	if err != nil {
		return fmt.Errorf("failed to marshal notification categories: %w", err)
	}

	if err := os.WriteFile(categoriesFile, categoriesData, 0644); err != nil {
		return fmt.Errorf("failed to write-notification categories to disk: %w", err)
	}

	return nil
}

// Load notification categories.
func (m *Notification) loadCategories() error {
	configDir, err := m.getConfigDir()
	if err != nil {
		return err
	}

	categoriesFile := filepath.Join(configDir, "notification-categories.json")

	if _, err := os.Stat(categoriesFile); os.IsNotExist(err) {
		return nil
	}

	categoriesData, err := os.ReadFile(categoriesFile)
	if err != nil {
		return fmt.Errorf("failed to read notification categories from disk: %w", err)
	}

	categories := make(map[string]Category)
	if err := json.Unmarshal(categoriesData, &categories); err != nil {
		return fmt.Errorf("failed to unmarshal notification categories: %w", err)
	}

	m.categoriesLock.Lock()
	m.categories = categories
	m.categoriesLock.Unlock()

	return nil
}

// Setup signal handling for notification actions.
func (m *Notification) setupSignalHandling(ctx context.Context) error {
	if err := m.conn.AddMatchSignal(dbus.WithMatchInterface(dbusNotificationInterface),
		dbus.WithMatchMember("ActionInvoked")); err != nil {
		return err
	}

	if err := m.conn.AddMatchSignal(dbus.WithMatchInterface(dbusNotificationInterface),
		dbus.WithMatchMember("NotificationClosed")); err != nil {
		return err
	}

	c := make(chan *dbus.Signal, 10)
	m.conn.Signal(c)

	go m.handleSignals(ctx, c)

	return nil
}

// Handle incoming D-Bus signals.
func (m *Notification) handleSignals(ctx context.Context, c chan *dbus.Signal) {
	for {
		select {
		case <-ctx.Done():
			return
		case signal, ok := <-c:
			if !ok {
				return
			}
			switch signal.Name {
			case dbusNotificationInterface + ".ActionInvoked":
				m.handleActionInvoked(signal)
			case dbusNotificationInterface + ".NotificationClosed":
				m.handleNotificationClosed(signal)
			}
		}
	}
}

// Handle ActionInvoked signal.
func (m *Notification) handleActionInvoked(signal *dbus.Signal) {
	if len(signal.Body) < 2 {
		return
	}

	dbusID, ok := signal.Body[0].(uint32)
	if !ok {
		return
	}

	actionID, ok := signal.Body[1].(string)
	if !ok {
		return
	}

	m.notificationsLock.Lock()
	notification, exists := m.notifications[dbusID]
	if exists {
		delete(m.notifications, dbusID)
	}
	m.notificationsLock.Unlock()

	if !exists {
		return
	}

	appActionID, ok := notification.ActionMap[actionID]
	if !ok {
		appActionID = actionID
	}

	response := Response{
		ID:               notification.ID,
		ActionIdentifier: appActionID,
		Title:            notification.Title,
		Subtitle:         notification.Subtitle,
		Body:             notification.Body,
		CategoryID:       notification.CategoryID,
		UserInfo:         notification.Data,
	}

	result := Result{
		Response: response,
	}

	m.handleNotificationResult(result)
}

// Handle NotificationClosed signal.
// Reason codes:
// 1 - expired timeout
// 2 - dismissed by user (click on X)
// 3 - closed by CloseNotification call
// 4 - undefined/reserved
func (m *Notification) handleNotificationClosed(signal *dbus.Signal) {
	if len(signal.Body) < 2 {
		return
	}

	dbusID, ok := signal.Body[0].(uint32)
	if !ok {
		return
	}

	reason, ok := signal.Body[1].(uint32)
	if !ok {
		reason = 0
	}

	m.notificationsLock.Lock()
	notification, exists := m.notifications[dbusID]
	if exists {
		delete(m.notifications, dbusID)
	}
	m.notificationsLock.Unlock()

	if !exists {
		return
	}

	if reason == 2 {
		response := Response{
			ID:               notification.ID,
			ActionIdentifier: DefaultActionIdentifier,
			Title:            notification.Title,
			Subtitle:         notification.Subtitle,
			Body:             notification.Body,
			CategoryID:       notification.CategoryID,
			UserInfo:         notification.Data,
		}

		result := Result{
			Response: response,
		}

		m.handleNotificationResult(result)
	}
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
