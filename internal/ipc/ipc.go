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
	"fmt"
	"github.com/energye/energy/v3/ipc/callback"
	"sync"
)

var (
	listener           *ipcListener
	listenerLock       sync.RWMutex
	processMessage     map[uint32]IProcessMessage
	processMessageLock sync.RWMutex
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
	callbacks map[string]callback.ICallback
}

// NewMessageReceivedDelegate create MessageReceivedDelegate
func NewMessageReceivedDelegate() IMessageReceivedDelegate {
	result := &MessageReceivedDelegate{}
	return result
}

func (m *MessageReceivedDelegate) Received(windowId uint32, messageData string) {
	fmt.Println("MessageReceivedDelegate windowId:", windowId, "messageData:", messageData)
	var message ProcessMessage
	err := json.Unmarshal([]byte(messageData), &message)
	fmt.Printf("\t%+v:\n", message)
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
		fn, ok := listener.callbacks[message.Name]
		listenerLock.Unlock()
		if ok {
			ctx := callback.NewContext(windowId, message.Data)
			fn.Invoke(ctx)
		}
		//
		if message.Id != 0 {
			process.SendMessage([]byte(""))
		}
	}
}

func init() {
	listener = &ipcListener{
		callbacks: make(map[string]callback.ICallback),
	}
	processMessage = make(map[uint32]IProcessMessage)
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
		listenerLock.Lock()
		listener.callbacks[name] = newCallback
		listenerLock.Unlock()
	}
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	listenerLock.Lock()
	delete(listener.callbacks, name)
	listenerLock.Unlock()
}

// emitOnEvent
//
//	Trigger listening event
func (m *ipcListener) emitEvent(name string, data interface{}) {

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
