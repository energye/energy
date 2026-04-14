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

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (m *StyleContext) AddProvider(provider IStyleProvider, prio uint) {
	gtk3.SysCall("gtk_style_context_add_provider", m.Instance(), provider.Instance(), uintptr(prio))
}
