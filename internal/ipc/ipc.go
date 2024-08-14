// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package ipc

import (
	"github.com/energye/energy/v3/ipc/callback"
	"sync"
)

var (
	listener           *ipcListener
	listenerLock       sync.RWMutex
	processMessage     map[uint32]IProcessMessage
	processMessageLock sync.RWMutex
	defaultBrowserId   uint32 // ipc message default browser id
)

type IProcessMessage interface {
	BrowserId() uint32
	SendMessage(payload []byte)
}

// IMessageReceivedDelegate Process message agent processor
type IMessageReceivedDelegate interface {
	// Received message
	//  windowId: The currently received browser window id
	//  message: Process message JSON string
	Received(windowId uint32, message *ProcessMessage) bool
}

type MessageReceivedDelegate struct {
}

type ipcListener struct {
	onCallbacks   map[string]callback.ICallback
	emitCallbacks map[uint32]callback.ICallback
}

// NewMessageReceivedDelegate create MessageReceivedDelegate
func NewMessageReceivedDelegate() IMessageReceivedDelegate {
	result := &MessageReceivedDelegate{}
	return result
}

func (m *MessageReceivedDelegate) Received(windowId uint32, message *ProcessMessage) bool {
	// Log ???
	//sendError := &ProcessMessage{}
	//process.SendMessage()
	// call go ipc callback
	if message.Name == "" {
		// Log ???
		return false
	}
	switch message.Type {
	case MT_EVENT_JS_EMIT: // js ipc.emit
		m.handlerJSEMIT(windowId, message)
	case MT_EVENT_GO_EMIT_CALLBACK: // go ipc.emit - callback function
		m.handlerGOEMITCallback(windowId, message)
	default:
		return false
	}
	return true
}

// go ipc.emit - callback function
func (m *MessageReceivedDelegate) handlerGOEMITCallback(browserId uint32, message *ProcessMessage) {
	if callbackFunc, ok := listener.emitCallbacks[message.Id]; ok {
		ctx := callback.NewContext(browserId, message.Data)
		callbackFunc.Invoke(ctx)
	}
}

// js ipc.emit
func (m *MessageReceivedDelegate) handlerJSEMIT(browserId uint32, message *ProcessMessage) {
	listenerLock.Lock()
	fn, ok := listener.onCallbacks[message.Name]
	listenerLock.Unlock()
	var result interface{}
	if ok {
		ctx := callback.NewContext(browserId, message.Data)
		// call ipc.on callback function
		result = fn.Invoke(ctx)
	}
	// not 0 js has callback function
	if message.Id != 0 {
		message.Type = MT_EVENT_JS_EMIT_CALLBACK
		message.Data = result
		payload, err := message.ToJSON()
		if err != nil {
			// Log ???
		}
		process := GetProcessMessage(browserId)
		if process == nil {
			// Log ???
			return
		}
		process.SendMessage(payload)
	}
}

func init() {
	listener = &ipcListener{
		onCallbacks:   make(map[string]callback.ICallback),
		emitCallbacks: make(map[uint32]callback.ICallback),
	}
	processMessage = make(map[uint32]IProcessMessage)
}

// SetMainDefaultBrowserId For IPC messages
func SetMainDefaultBrowserId(browserId uint32) {
	defaultBrowserId = browserId
}

// RegisterProcessMessage register process message object
func RegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	processMessage[window.BrowserId()] = window
	processMessageLock.Unlock()
}

// UnRegisterProcessMessage cancel process message object
func UnRegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	delete(processMessage, window.BrowserId())
	processMessageLock.Unlock()
}

// GetProcessMessage return process message object
func GetProcessMessage(browserId uint32) IProcessMessage {
	processMessageLock.Lock()
	defer processMessageLock.Unlock()
	return processMessage[browserId]
}

// createCallback
//
//	Create and return a callback function
func createCallback(fn interface{}) callback.ICallback {
	switch fn.(type) {
	case callback.EventCallback:
		return callback.New(fn.(callback.EventCallback))
	}
	return nil
}

// On
//
//	Add listening event
//	callback function
//	  1. context 2.argument list
func On(name string, fn callback.EventCallback) {
	if listener == nil || name == "" || fn == nil {
		return
	}
	if newCallback := createCallback(fn); newCallback != nil {
		listenerLock.Lock()
		listener.onCallbacks[name] = newCallback
		listenerLock.Unlock()
	}
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	listenerLock.Lock()
	delete(listener.onCallbacks, name)
	listenerLock.Unlock()
}

// Emit Sends an ipc message to the specified window
//
//	windowId: target window, default 0 = mainWindow, otherwise, it must be a valid window identifier
//	name: emit message name
//	arguments: emit message data
func Emit(browserId uint32, name string, arguments ...interface{}) {
	if browserId == 0 {
		browserId = defaultBrowserId // default main browser window id
	}
	processMessageLock.Lock()
	emitMsg, ok := processMessage[browserId]
	processMessageLock.Unlock()
	if ok {
		var executionID uint32
		if len(arguments) > 0 {
			// Check if the last parameter is a callback function
			callbackFunction := arguments[len(arguments)-1]
			switch callbackFunction.(type) {
			case callback.EventCallback:
				executionID = NextExecutionID()
				listener.emitCallbacks[executionID] = callback.New(callbackFunction.(callback.EventCallback))
				arguments = arguments[:len(arguments)-1]
			}
		}
		message := &ProcessMessage{
			Type: MT_EVENT_GO_EMIT,
			Name: name,
			Data: arguments,
			Id:   executionID,
		}
		payload, err := message.ToJSON()
		if err != nil {
			// Log ???
		}
		emitMsg.SendMessage(payload)
	}
}
