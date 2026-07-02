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

import "unsafe"

// Statusbar is a representation of GTK's GtkStatusbar.
type Statusbar struct {
	Box
}

func AsStatusbar(ptr unsafe.Pointer) *Statusbar {
	if ptr == nil {
		return nil
	}
	m := new(Statusbar)
	m.instance = ptr
	return m
}

// NewStatusbar is a wrapper around gtk_statusbar_new().
func NewStatusbar() *Statusbar {
	r := gtk3.SysCall("gtk_statusbar_new")
	if r == 0 {
		return nil
	}
	return AsStatusbar(unsafe.Pointer(r))
}

// GetContextId is a wrapper around gtk_statusbar_get_context_id().
func (m *Statusbar) GetContextId(contextDescription string) uint {
	r := gtk3.SysCall("gtk_statusbar_get_context_id", m.Instance(), CStr(contextDescription))
	return uint(r)
}

// Push is a wrapper around gtk_statusbar_push().
func (m *Statusbar) Push(contextId uint, text string) uint {
	r := gtk3.SysCall("gtk_statusbar_push", m.Instance(), uintptr(contextId), CStr(text))
	return uint(r)
}

// Pop is a wrapper around gtk_statusbar_pop().
func (m *Statusbar) Pop(contextId uint) {
	gtk3.SysCall("gtk_statusbar_pop", m.Instance(), uintptr(contextId))
}

// RemoveAll is a wrapper around gtk_statusbar_remove_all().
func (m *Statusbar) RemoveAll(contextId uint) {
	gtk3.SysCall("gtk_statusbar_remove_all", m.Instance(), uintptr(contextId))
}
