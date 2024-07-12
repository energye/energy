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
	"encoding/json"
	"github.com/energye/energy/v3/ipc/callback"
	"sync"
)

var (
	listener           *ipcListener
	listenerLock       sync.RWMutex
	processMessage     map[uint32]IProcessMessage
	processMessageLock sync.RWMutex
	mainWindowId       uint32 // ipc message default window flag
)

type IProcessMessage interface {
	WindowId() uint32
	SendMessage(data []byte)
}

type IMessageReceivedDelegate interface {
	Received(windowId uint32, message string)
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

func (m *MessageReceivedDelegate) Received(windowId uint32, messageData string) {
	var message ProcessMessage
	err := json.Unmarshal([]byte(messageData), &message)
	process := GetProcessMessage(windowId)
	if err != nil {
		//sendError := &ProcessMessage{}
		//process.SendMessage()
	} else {
		// call go ipc callback
		if message.Name == "" {
			return
		}
		listenerLock.Lock()
		fn, ok := listener.onCallbacks[message.Name]
		listenerLock.Unlock()
		var result interface{}
		if ok {
			ctx := callback.NewContext(windowId, message.Data)
			result = fn.Invoke(ctx)
		}
		//
		if message.Id != 0 {
			message.Name = ""
			message.Data = result
			tmpMsg, err := message.ToJSON()
			if err != nil {
				// Log ???
			}
			process.SendMessage(tmpMsg)
		}
	}
}

func init() {
	listener = &ipcListener{
		onCallbacks:   make(map[string]callback.ICallback),
		emitCallbacks: make(map[uint32]callback.ICallback),
	}
	processMessage = make(map[uint32]IProcessMessage)
}

// SetMainWindowId For IPC messages
func SetMainWindowId(windowId uint32) {
	mainWindowId = windowId
}

// RegisterProcessMessage reegister process message object
func RegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	processMessage[window.WindowId()] = window
	processMessageLock.Unlock()
}

// UnRegisterProcessMessage cancel process message object
func UnRegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	delete(processMessage, window.WindowId())
	processMessageLock.Unlock()
}

// GetProcessMessage return process message object
func GetProcessMessage(windowId uint32) IProcessMessage {
	processMessageLock.Lock()
	defer processMessageLock.Unlock()
	return processMessage[windowId]
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
func Emit(windowId uint32, name string, arguments ...interface{}) {
	if windowId == 0 {
		windowId = mainWindowId
	}
	processMessageLock.Lock()
	emitMsg, ok := processMessage[windowId]
	processMessageLock.Unlock()
	if ok {
		var resultCallbackId uint32
		if len(arguments) > 0 {
			argsl := arguments[len(arguments)-1]
			switch argsl.(type) {
			case callback.EventCallback:
				resultCallbackId = NextExecutionID()
				listener.emitCallbacks[resultCallbackId] = callback.New(argsl.(callback.EventCallback))
				arguments = arguments[:len(arguments)-1]
			}
		}
		message := &ProcessMessage{
			Name: name,
			Data: arguments,
			Id:   resultCallbackId,
		}
		data, err := message.ToJSON()
		if err != nil {
			// Log ???
		}
		emitMsg.SendMessage(data)
	}
}

// Add emit callback function
func (m *ipcListener) emitCallback(fn interface{}) int32 {
	if m == nil || fn == nil {
		return 0
	}
	//if callbackFN := createCallback(fn); callbackFN != nil {
	//	if m.emitCallbackMessageId == -1 {
	//		m.emitCallbackMessageId = 1
	//	} else {
	//		m.emitCallbackMessageId++
	//	}
	//	m.emitCallback[m.emitCallbackMessageId] = callbackFN
	//	return m.emitCallbackMessageId
	//}
	return 0
}
