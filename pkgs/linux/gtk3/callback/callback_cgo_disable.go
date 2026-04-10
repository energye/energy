//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package callback

import (
	"github.com/ebitengine/purego"
	"github.com/energye/energy/v3/pkgs/linux"
	"github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

var (
	gSignalConnect           func(object, name, handler, data, destroyData, flags uintptr) uintptr
	gSignalHandlerDisconnect func(widget, handlerId uintptr)
)

func (m *SignalHandlerID) Disconnect() {
	gSignalHandlerDisconnect(m.Widget, uintptr(m.HandlerID))
}

var gobject2_0 *linux.DnyLibrary

func init() {
	gobject2_0 = linux.LibLoad(linux.Libgobject2_0)
	gobject2_0.Table = []*imports.Table{
		imports.NewTable("g_signal_connect_data", 0),
		imports.NewTable("g_signal_handler_disconnect", 0),
	}
	gobject2_0.SetLibClose()
	gobject2_0.MapperIndex()
	purego.RegisterLibFunc(&gSignalConnect, uintptr(gobject2_0.Dll), "g_signal_connect_data")
	purego.RegisterLibFunc(&gSignalHandlerDisconnect, uintptr(gobject2_0.Dll), "g_signal_handler_disconnect")
}

func Connect(widget uintptr, signalName string, fn any, userData uintptr) *SignalHandlerID {
	fnPtr := purego.NewCallback(fn)
	handlerID := gSignalConnect(widget, api.PasStr(signalName), fnPtr, userData, 0, 0)
	return &SignalHandlerID{
		Widget:    widget,
		HandlerID: types.GULong(handlerID),
	}
}
