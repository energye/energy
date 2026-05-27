//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package cgo

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include "gtk.go.h"
*/
import "C"

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Settings struct {
	*Object
}

func SettingsGetDefault() *Settings {
	c := C.gtk_settings_get_default()
	if c == nil {
		return nil
	}
	return &Settings{ToGoObject(unsafe.Pointer(c))}
}

func AsSettings(ptr unsafe.Pointer) *Settings {
	if ptr == nil {
		return nil
	}
	return &Settings{ToGoObject(ptr)}
}

func (v *Settings) native() *C.GtkSettings {
	return toGtkSettings(unsafe.Pointer(v.Instance()))
}

func (v *Settings) SetOnThemeChanged(fn types.TThemeChangedEvent) types.ISignalHandlerID {
	signalHandlerID := callback.Connect(v.Instance(), types.EsnNotifyThemeChanged, callback.C_trampoline_3_void, fn, 0)
	return signalHandlerID
}
