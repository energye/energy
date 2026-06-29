//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package cef

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/platform/linux/gtk3"
	gtk3Types "github.com/energye/energy/v3/platform/linux/types"
	"github.com/godbus/dbus/v5"
)

var (
	frameWidth  = int32(4)
	frameHeight = int32(4)
	frameCorner = frameWidth + frameHeight
)

func (m *TBrowser) drag(message ipc.ProcessMessage) bool {
	// todo 待实现
	if m.window != nil && m.kind == bkEmbedded {

	} else if m.kind == bkViews {

	}
	return false
}

func (m *TBrowser) resize(ht string) bool {
	// todo 待实现
	if m.window != nil && m.kind == bkEmbedded {

	} else if m.kind == bkViews {

	}
	return false
}

func IsCurrentlyDarkMode() bool {
	conn, err := dbus.SessionBus()
	if err != nil {
		return false
	}
	//defer conn.Close()
	obj := conn.Object("org.freedesktop.portal.Desktop", "/org/freedesktop/portal/desktop")
	call := obj.Call("org.freedesktop.portal.Settings.Read", 0, "org.freedesktop.appearance", "color-scheme")
	if call.Err != nil {
		return false
	}
	var result dbus.Variant
	if err := call.Store(&result); err != nil {
		return false
	}
	innerVariant, ok := result.Value().(dbus.Variant)
	if !ok {
		return false
	}
	colorScheme, ok := innerVariant.Value().(uint32)
	if !ok {
		return false
	}
	return colorScheme == gtk3Types.ColorSchemePreferDark
}

func (m *TViewsBrowser) UpdateTheme() {
	isDark := false
	switch GApplication.Options.Linux.Theme {
	case application.SystemDefault:
		isDark = IsCurrentlyDarkMode()
	case application.Dark:
		isDark = true
	case application.Light:
		isDark = false
	}
	m.doThemeChanged(isDark)
}

func (m *TViewsBrowser) startThemeObserver() {
	settings := gtk3.SettingsGetDefault()
	if settings == nil {
		return
	}
	settings.SetOnThemeChanged(func(sender gtk3Types.PGtkWidget, pspec uintptr, userData gtk3Types.GPointer) {
		m.UpdateTheme()
	})
}
