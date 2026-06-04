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

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Settings struct {
	*Object
}

func SettingsGetDefault() *Settings {
	r := gtk3.SysCall("gtk_settings_get_default")
	if r == 0 {
		return nil
	}
	return &Settings{&Object{instance: unsafe.Pointer(r)}}
}

func AsSettings(ptr unsafe.Pointer) *Settings {
	if ptr == nil {
		return nil
	}
	return &Settings{&Object{instance: ptr}}
}

func (v *Settings) SetOnThemeChanged(fn types.TThemeChangedEvent) types.ISignalHandlerID {
	signalHandlerID := callback.Connect(v.Instance(), types.EsnNotifyThemeChanged, "", fn, 0)
	return signalHandlerID
}
