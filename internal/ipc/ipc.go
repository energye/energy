package ipc

import "github.com/energye/energy/v3/cef/ipc/types"

var (
	listener    *ipcListener
	messageSend map[uint32]IMessageSend
)

type IMessageSend interface {
	MessageSend()
}

type IMessageReceivedDelegate interface {
	Received(windowId uint32, message string)
}

type MessageReceivedDelegate struct {
}

func NewMessageReceivedDelegate() IMessageReceivedDelegate {
	result := &MessageReceivedDelegate{}
	return result
}

func (m *MessageReceivedDelegate) Received(windowId uint32, message string) {

}

type ipcListener struct {
	callbacks map[string]ICallback
}

func init() {
	listener = &ipcListener{
		callbacks: make(map[string]ICallback),
	}
	messageSend = make(map[uint32]IMessageSend)
}

// createCallback
//
//	Create and return a callback function
func createCallback(fn interface{}) ICallback {
	switch fn.(type) {
	case EventCallback:
		return &Callback{callback: fn.(EventCallback)}
	}
	return nil
}

// AddEvent
//
//	Add listening event
//	callback function
//	  1. context 2.argument list
func AddEvent(name string, fn EventCallback) {
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
