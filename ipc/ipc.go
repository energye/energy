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
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/ipc/callback"
)

// On
//
//	IPC GO Listening for events
func On(name string, fn callback.EventCallback) {
	ipc.AddEvent(name, fn)
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	ipc.RemoveOn(name)
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
