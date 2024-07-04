//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC - Based on pkgs IPC, CEF Internal implementation
// event listeners
// event triggered

package ipc

import (
	"github.com/energye/energy/v3/cef/ipc/types"
)

var (
	listener *ipcListener
)

type IListener interface {
}

type ipcListener struct {
	callbacks map[string]ICallback
}

func init() {
	listener = &ipcListener{
		callbacks: make(map[string]ICallback),
	}
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

// On
//
//	IPC GO Listening for events
func On(name string, fn EventCallback) {
	if name == "" || fn == nil {
		return
	}
	if callback := createCallback(fn); callback != nil {
		listener.addEvent(name, callback)
	}
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	if name == "" {
		return
	}
	delete(listener.callbacks, name)
}

// Emit
//
//	Event that triggers listening
//	default to triggering the main process
//func Emit(name string, argument ...interface{}) bool {
//	if name == "" {
//		return false
//	}
//	// When the window is closed
//	if browser.window == nil || browser.window.IsClosing() {
//		// This window is the first one created and not closed
//		SetProcessMessage(browser.browserWindow.LookForMainWindow())
//	}
//	browser.window.ProcessMessage().EmitRender(0, name, nil, argument...)
//	return true
//}

// addOnEvent
//
//	Add listening event
//	callback function
//	  1. context 2.argument list
func (m *ipcListener) addEvent(name string, fn ICallback) {
	if m == nil || name == "" || fn == nil {
		return
	}
	m.callbacks[name] = fn
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
