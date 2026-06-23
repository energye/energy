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

// Iconify is a wrapper around gtk_window_iconify().
func (m *Window) Iconify() {
	gtk3.SysCall("gtk_window_iconify", m.Instance())
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (m *Window) Deiconify() {
	gtk3.SysCall("gtk_window_deiconify", m.Instance())
}

// Stick is a wrapper around gtk_window_stick().
func (m *Window) Stick() {
	gtk3.SysCall("gtk_window_stick", m.Instance())
}

// Unstick is a wrapper around gtk_window_unstick().
func (m *Window) Unstick() {
	gtk3.SysCall("gtk_window_unstick", m.Instance())
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

// BeginResizeDrag is a wrapper around gtk_window_begin_resize_drag().
func (m *Window) BeginResizeDrag(edge WindowEdge, button ButtonType, rootX, rootY int, timestamp uint32) {
	gtk3.SysCall("gtk_window_begin_resize_drag", m.Instance(), uintptr(edge), uintptr(button), uintptr(rootX), uintptr(rootY), uintptr(timestamp))
}

// BeginMoveDrag is a wrapper around gtk_window_begin_move_drag().
func (m *Window) BeginMoveDrag(button ButtonType, rootX, rootY int, timestamp uint32) {
	gtk3.SysCall("gtk_window_begin_resize_drag", m.Instance(), uintptr(button), uintptr(rootX), uintptr(rootY), uintptr(timestamp))
}

func (m *Window) SetOnConfigure(fn TConfigureEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnConfigureEvent, callback.C_trampoline_2_void, fn, 0)
	return signalHandlerID
}

func (m *Window) SetOnMap(fn TMapEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnMapEvent, callback.C_trampoline_2_void, fn, 0)
	return signalHandlerID
}

func (m *Window) SetOnDraw(fn TDrawEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDrawEvent, callback.C_trampoline_3_gboolean, fn, 0)
	return signalHandlerID
}
