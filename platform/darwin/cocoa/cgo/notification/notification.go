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

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa -framework UserNotifications

#include "notification.h"
*/
import "C"

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/energye/energy/v3/platform/notification/types"
	"sync"
	"time"
	"unsafe"
)

var (
	once          sync.Once
	gNotification *Notification
)

// Apple system constants
const (
	appleDefaultActionID = "com.apple.UNNotificationDefaultActionIdentifier"
	authorizationTimeout = 180 * time.Second // 3 minutes for user interaction
	operationTimeout     = 5 * time.Second   // Standard operation timeout
)

type Notification struct {
	notificationResultCallback TNotificationResponseEvent
	callbackLock               sync.RWMutex
	callbackChannels           map[int]chan callbackResult
	mu                         sync.Mutex
	nextID                     int
	isInitialized              bool
}

// callbackResult carries async operation results from C to Go
type callbackResult struct {
	ok  bool
	err error
}

// New 创建通知服务
func New() INotification {
	once.Do(func() {
		gNotification = &Notification{
			callbackChannels: make(map[int]chan callbackResult),
			nextID:           0,
		}
		// 自动初始化
		if err := gNotification.Initialize(); err != nil {
			fmt.Printf("[Energy] Notification service initialization warning: %v\n", err)
		}
	})
	return gNotification
}

// Initialize sets up the notification service
func (m *Notification) Initialize() error {
	if m.isInitialized {
		return nil
	}

	if !bool(C.isNotificationAvailable()) {
		return fmt.Errorf("notifications unavailable: requires macOS 10.15 or later")
	}

	if !bool(C.checkBundleIdentifier()) {
		return fmt.Errorf("missing bundle identifier: package app as .app bundle with Info.plist")
	}

	if !bool(C.initializeNotificationCenter()) {
		return fmt.Errorf("failed to initialize notification delegate")
	}

	m.isInitialized = true
	return nil
}

// RequestNotificationAuthorization prompts user for notification permissions
func (m *Notification) RequestNotificationAuthorization() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authorizationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.requestNotificationAuthorization(C.int(id))

	select {
	case result := <-ch:
		return result.ok, result.err
	case <-ctx.Done():
		return false, fmt.Errorf("authorization request timeout: %w", ctx.Err())
	}
}

// CheckNotificationAuthorization queries current permission status
func (m *Notification) CheckNotificationAuthorization() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.checkNotificationAuthorization(C.int(id))

	select {
	case result := <-ch:
		return result.ok, result.err
	case <-ctx.Done():
		return false, fmt.Errorf("authorization check timeout: %w", ctx.Err())
	}
}

// SendNotification delivers a basic notification
func (m *Notification) SendNotification(opts Options) error {
	if err := validateNotificationOptions(opts); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	// Prepare C strings
	cID := C.CString(opts.ID)
	cTitle := C.CString(opts.Title)
	cSubtitle := C.CString(opts.Subtitle)
	cBody := C.CString(opts.Body)
	defer func() {
		C.free(unsafe.Pointer(cID))
		C.free(unsafe.Pointer(cTitle))
		C.free(unsafe.Pointer(cSubtitle))
		C.free(unsafe.Pointer(cBody))
	}()

	// Marshal custom data
	var cDataJSON *C.char
	if opts.Data != nil {
		jsonBytes, marshalErr := json.Marshal(opts.Data)
		if marshalErr != nil {
			return fmt.Errorf("data serialization failed: %w", marshalErr)
		}
		cDataJSON = C.CString(string(jsonBytes))
		defer C.free(unsafe.Pointer(cDataJSON))
	}

	// Execute send
	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.sendNotification(C.int(id), cID, cTitle, cSubtitle, cBody, cDataJSON)

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

// SendNotificationWithActions delivers a notification with interactive actions
func (m *Notification) SendNotificationWithActions(opts Options) error {
	if err := validateNotificationOptions(opts); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	// Prepare C strings
	cID := C.CString(opts.ID)
	cTitle := C.CString(opts.Title)
	cSubtitle := C.CString(opts.Subtitle)
	cBody := C.CString(opts.Body)
	cCategoryID := C.CString(opts.CategoryID)
	defer func() {
		C.free(unsafe.Pointer(cID))
		C.free(unsafe.Pointer(cTitle))
		C.free(unsafe.Pointer(cSubtitle))
		C.free(unsafe.Pointer(cBody))
		C.free(unsafe.Pointer(cCategoryID))
	}()

	// Marshal custom data
	var cDataJSON *C.char
	if opts.Data != nil {
		jsonBytes, marshalErr := json.Marshal(opts.Data)
		if marshalErr != nil {
			return fmt.Errorf("data serialization failed: %w", marshalErr)
		}
		cDataJSON = C.CString(string(jsonBytes))
		defer C.free(unsafe.Pointer(cDataJSON))
	}

	// Execute send
	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.sendNotificationWithActions(C.int(id), cID, cTitle, cSubtitle, cBody, cCategoryID, cDataJSON)

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

// RegisterNotificationCategory registers a notification category with actions
func (m *Notification) RegisterNotificationCategory(category Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	// Serialize actions
	actionsJSON, jsonErr := json.Marshal(category.Actions)
	if jsonErr != nil {
		return fmt.Errorf("actions serialization failed: %w", jsonErr)
	}

	cCategoryID := C.CString(category.ID)
	cActionsJSON := C.CString(string(actionsJSON))
	defer func() {
		C.free(unsafe.Pointer(cCategoryID))
		C.free(unsafe.Pointer(cActionsJSON))
	}()

	// Prepare text input parameters
	var cPlaceholder, cButtonTitle *C.char
	if category.HasReplyField {
		cPlaceholder = C.CString(category.ReplyPlaceholder)
		cButtonTitle = C.CString(category.ReplyButtonTitle)
		defer func() {
			C.free(unsafe.Pointer(cPlaceholder))
			C.free(unsafe.Pointer(cButtonTitle))
		}()
	}

	// Execute registration
	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.registerNotificationCategory(
		C.int(id),
		cCategoryID,
		cActionsJSON,
		C.bool(category.HasReplyField),
		cPlaceholder,
		cButtonTitle,
	)

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

// RemoveNotificationCategory unregisters a notification category
func (m *Notification) RemoveNotificationCategory(categoryID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	cCategoryID := C.CString(categoryID)
	defer C.free(unsafe.Pointer(cCategoryID))

	id, ch := m.allocateChannel()
	defer m.releaseChannel(id)

	C.removeNotificationCategory(C.int(id), cCategoryID)

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

// RemoveAllPendingNotifications cancels all scheduled notifications
func (m *Notification) RemoveAllPendingNotifications() error {
	C.removeAllPendingNotifications()
	return nil
}

// RemovePendingNotification cancels a specific scheduled notification
func (m *Notification) RemovePendingNotification(identifier string) error {
	cID := C.CString(identifier)
	defer C.free(unsafe.Pointer(cID))

	C.removePendingNotification(cID)
	return nil
}

// RemoveAllDeliveredNotifications clears all displayed notifications
func (m *Notification) RemoveAllDeliveredNotifications() error {
	C.removeAllDeliveredNotifications()
	return nil
}

// RemoveDeliveredNotification removes a specific displayed notification
func (m *Notification) RemoveDeliveredNotification(identifier string) error {
	cID := C.CString(identifier)
	defer C.free(unsafe.Pointer(cID))

	C.removeDeliveredNotification(cID)
	return nil
}

// RemoveNotification stub for cross-platform compatibility
func (m *Notification) RemoveNotification(_ string) error {
	return nil
}

// ----------------------------------------------------------------------------
// CGO Callback Handlers
// ----------------------------------------------------------------------------

// captureResult receives async operation results from Objective-C
//
//export captureResult
func captureResult(channelID C.int, success C.bool, errorMsg *C.char) {
	if gNotification == nil {
		return
	}
	ch, exists := gNotification.GetChannel(int(channelID))
	if !exists {
		return
	}
	var opErr error
	if errorMsg != nil {
		opErr = fmt.Errorf("%s", C.GoString(errorMsg))
	}
	ch <- callbackResult{
		ok:  bool(success),
		err: opErr,
	}
	close(ch)
}

// didReceiveNotificationResponse handles user interactions with notifications
//
//export didReceiveNotificationResponse
func didReceiveNotificationResponse(jsonPayload *C.char, errStr *C.char) {
	result := Result{}

	// Handle error case
	if errStr != nil {
		result.Error = fmt.Errorf("response error: %s", C.GoString(errStr))
		dispatchResult(result)
		return
	}

	// Validate payload
	if jsonPayload == nil {
		result.Error = fmt.Errorf("nil response payload")
		dispatchResult(result)
		return
	}

	// Parse JSON response
	payload := C.GoString(jsonPayload)
	var response Response
	if parseErr := json.Unmarshal([]byte(payload), &response); parseErr != nil {
		result.Error = fmt.Errorf("response parsing failed: %w", parseErr)
		dispatchResult(result)
		return
	}

	// Normalize action identifier
	if response.ActionIdentifier == appleDefaultActionID {
		response.ActionIdentifier = DefaultActionIdentifier
	}

	result.Response = response
	dispatchResult(result)
}

// dispatchResult sends notification result to registered handler
func dispatchResult(result Result) {
	if gNotification != nil {
		gNotification.handleNotificationResult(result)
	}
}

// allocateChannel creates a new callback channel and returns its ID
func (m *Notification) allocateChannel() (int, chan callbackResult) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := m.nextID
	m.nextID++

	ch := make(chan callbackResult, 1)
	m.callbackChannels[id] = ch

	return id, ch
}

// GetChannel retrieves and removes a callback channel (implements ChannelHandler)
func (m *Notification) GetChannel(id int) (chan callbackResult, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ch, exists := m.callbackChannels[id]
	if exists {
		delete(m.callbackChannels, id)
	}
	return ch, exists
}

// releaseChannel cleans up a callback channel on timeout
func (m *Notification) releaseChannel(id int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ch, exists := m.callbackChannels[id]; exists {
		delete(m.callbackChannels, id)
		close(ch)
	}
}

// SetOnNotificationResponse 注册通知响应回调
func (m *Notification) SetOnNotificationResponse(callback TNotificationResponseEvent) {
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.notificationResultCallback = callback
}

// handleNotificationResult 处理通知结果
func (m *Notification) handleNotificationResult(result Result) {
	m.callbackLock.RLock()
	callback := m.notificationResultCallback
	m.callbackLock.RUnlock()

	if callback != nil {
		callback(result)
	}
}

// validateNotificationOptions 验证通知选项
func validateNotificationOptions(options Options) error {
	if options.ID == "" {
		return fmt.Errorf("notification ID cannot be empty")
	}
	if options.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	return nil
}
