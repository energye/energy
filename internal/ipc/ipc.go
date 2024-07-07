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
	"github.com/energye/energy/v3/cef/ipc/types"
	"github.com/energye/energy/v3/ipc/callback"
	"sync"
)

var (
	listener           *ipcListener
	processMessage     map[uint32]IProcessMessage
	processMessageLock sync.RWMutex
)

type IProcessMessage interface {
	WindowId() uint32
	SendMessage()
}

type IMessageReceivedDelegate interface {
	Received(windowId uint32, message string)
}

type MessageReceivedDelegate struct {
}

type ipcListener struct {
	callbacks map[string]callback.ICallback
}

// NewMessageReceivedDelegate create MessageReceivedDelegate
func NewMessageReceivedDelegate() IMessageReceivedDelegate {
	result := &MessageReceivedDelegate{}
	return result
}

func (m *MessageReceivedDelegate) Received(windowId uint32, messageData string) {
	var message ProcessMessage
	err := json.Unmarshal([]byte(messageData), &message)
	if err != nil {

	} else {

	}
}

func init() {
	listener = &ipcListener{
		callbacks: make(map[string]callback.ICallback),
	}
	processMessage = make(map[uint32]IProcessMessage)
}

// RegisterProcessMessage process message
func RegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	processMessage[window.WindowId()] = window
	processMessageLock.Unlock()
}

// UnRegisterProcessMessage cancel process message
func UnRegisterProcessMessage(window IProcessMessage) {
	processMessageLock.Lock()
	delete(processMessage, window.WindowId())
	processMessageLock.Unlock()
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

// AddEvent
//
//	Add listening event
//	callback function
//	  1. context 2.argument list
func AddEvent(name string, fn callback.EventCallback) {
	if listener == nil || name == "" || fn == nil {
		return
	}
	if newCallback := createCallback(fn); newCallback != nil {
		listener.callbacks[name] = newCallback
	}
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	delete(listener.callbacks, name)
}

// emitOnEvent
//
//	Trigger listening event
func (m *ipcListener) emitEvent(name string, argumentList types.IArrayValue) {
	if m == nil || name == "" || argumentList == nil {
		return
	}
}

// addOnEvent
//
//	Add emit callback function
func (m *ipcListener) addEmitCallback(fn interface{}) int32 {
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
