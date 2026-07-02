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
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// StyleContext is a representation of GTK's GtkStyleContext.
type StyleContext struct {
	Object
}

func AsStyleContext(ptr unsafe.Pointer) IStyleContext {
	if ptr == nil {
		return nil
	}
	m := new(StyleContext)
	m.instance = ptr
	return m
}

func (m *StyleContext) AddClass(class_name string) {
	gtk3.SysCall("gtk_style_context_add_class", m.Instance(), CStr(class_name))
}

func (m *StyleContext) RemoveClass(class_name string) {
	gtk3.SysCall("gtk_style_context_remove_class", m.Instance(), CStr(class_name))
}

func (m *StyleContext) HasClass(className string) bool {
	r := gtk3.SysCall("gtk_style_context_has_class", m.Instance(), CStr(className))
	return ToGoBool(r)
}

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (m *StyleContext) AddProvider(provider IStyleProvider, prio uint) {
	gtk3.SysCall("gtk_style_context_add_provider", m.Instance(), provider.Instance(), uintptr(prio))
}

// GetParent is a wrapper around gtk_style_context_get_parent().
func (m *StyleContext) GetParent() IStyleContext {
	r := gtk3.SysCall("gtk_style_context_get_parent", m.Instance())
	if r == 0 {
		return nil
	}
	return AsStyleContext(unsafe.Pointer(r))
}

// GetState is a wrapper around gtk_style_context_get_state().
func (m *StyleContext) GetState() StateFlags {
	r := gtk3.SysCall("gtk_style_context_get_state", m.Instance())
	return StateFlags(r)
}

// Save is a wrapper around gtk_style_context_save().
func (m *StyleContext) Save() {
	gtk3.SysCall("gtk_style_context_save", m.Instance())
}

// Restore is a wrapper around gtk_style_context_restore().
func (m *StyleContext) Restore() {
	gtk3.SysCall("gtk_style_context_restore", m.Instance())
}

// SetParent is a wrapper around gtk_style_context_set_parent().
func (m *StyleContext) SetParent(parent IStyleContext) {
	gtk3.SysCall("gtk_style_context_set_parent", m.Instance(), parent.Instance())
}

// SetScreen is a wrapper around gtk_style_context_set_screen().
func (m *StyleContext) SetScreen(screen IScreen) {
	gtk3.SysCall("gtk_style_context_set_screen", m.Instance(), screen.Instance())
}

// SetState is a wrapper around gtk_style_context_set_state().
func (m *StyleContext) SetState(state StateFlags) {
	gtk3.SysCall("gtk_style_context_set_state", m.Instance(), uintptr(state))
}

// RemoveProvider is a wrapper around gtk_style_context_remove_provider().
func (m *StyleContext) RemoveProvider(provider IStyleProvider) {
	gtk3.SysCall("gtk_style_context_remove_provider", m.Instance(), provider.Instance())
}
