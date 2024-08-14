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

type OptionsEvent struct {
	BrowserId uint32
	Name      string
	Data      interface{}
	Callback  callback.EventCallback
}

// On
//
//	IPC GO Listening for events
func On(name string, fn callback.EventCallback) {
	ipc.On(name, fn)
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
func Emit(name string, arguments ...interface{}) {
	if name == "" {
		return
	}
	ipc.Emit(0, name, arguments...)
}

// EmitOptions
//
//	Event that triggers listening
//	default: 0 to triggering the main browser
func EmitOptions(options *OptionsEvent) {
	if options == nil || options.Name == "" {
		return
	}
	ipc.Emit(options.BrowserId, options.Name, options.Data, options.Callback)
}
