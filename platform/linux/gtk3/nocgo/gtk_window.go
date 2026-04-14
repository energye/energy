//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Window struct {
	Bin
}

func AsWindow(ptr unsafe.Pointer) IWindow {
	if ptr == nil {
		return nil
	}
	m := &Window{}
	m.instance = ptr
	return m
}

func (m *Window) GetDefaultSize() (width, height int) {
	gtk3.SysCall("gtk_window_get_default_size", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (m *Window) SetDecorated(setting bool) {
	gtk3.SysCall("gtk_window_set_decorated", m.Instance(), ToCBool(setting))
}

// Maximize is a wrapper around gtk_window_maximize().
func (m *Window) Maximize() {
	gtk3.SysCall("gtk_window_maximize", m.Instance())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (m *Window) Unmaximize() {
	gtk3.SysCall("gtk_window_unmaximize", m.Instance())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (m *Window) Fullscreen() {
	gtk3.SysCall("gtk_window_fullscreen", m.Instance())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (m *Window) Unfullscreen() {
	gtk3.SysCall("gtk_window_unfullscreen", m.Instance())
}

// SetTitle is a wrapper around gtk_window_set_title().
func (m *Window) SetTitle(title string) {
	gtk3.SysCall("gtk_window_set_title", m.Instance(), CStr(title))
}

// GetTitle is a wrapper around gtk_window_get_title().
func (m *Window) GetTitle() string {
	r := gtk3.SysCall("gtk_window_get_title", m.Instance())
	return GoStr(r)
}

func (m *Window) SetOnConfigure(fn TConfigureEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnConfigureEvent, fn, 0)
	return signalHandlerID
}

func (m *Window) SetOnMap(fn TMapEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnMapEvent, fn, 0)
	return signalHandlerID
}

func (m *Window) SetOnDraw(fn TDrawEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDrawEvent, fn, 0)
	return signalHandlerID
}
