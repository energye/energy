package ipc

import (
	"encoding/json"
	"github.com/energye/energy/v3/cef/ipc/types"
	"github.com/energye/energy/v3/ipc/callback"
	"sync"
)

var (
	listener        *ipcListener
	messageSend     map[uint32]IMessageSend
	messageSendLock sync.RWMutex
)

type IMessageSend interface {
	WindowId() uint32
	MessageSend()
}

type IMessageReceivedDelegate interface {
	Received(windowId uint32, message string)
}

type MessageReceivedDelegate struct {
}

type ipcListener struct {
	callbacks map[string]callback.ICallback
}

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
	messageSend = make(map[uint32]IMessageSend)
}

// RegisterMessageSend process message send
func RegisterMessageSend(window IMessageSend) {
	messageSendLock.Lock()
	messageSend[window.WindowId()] = window
	messageSendLock.Unlock()
}

func UnRegisterMessageSend(window IMessageSend) {
	messageSendLock.Lock()
	delete(messageSend, window.WindowId())
	messageSendLock.Unlock()
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
	if callback := createCallback(fn); callback != nil {
		listener.callbacks[name] = callback
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
