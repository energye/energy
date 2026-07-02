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

// Fixed is a representation of GTK's GtkFixed.
type Fixed struct {
	Container
}

func AsFixed(ptr unsafe.Pointer) IFixed {
	if ptr == nil {
		return nil
	}
	m := new(Fixed)
	m.instance = ptr
	return m
}

// NewFixed is a wrapper around gtk_fixed_new().
func NewFixed() IFixed {
	r := gtk3.SysCall("gtk_fixed_new")
	if r == 0 {
		return nil
	}
	return AsFixed(unsafe.Pointer(r))
}

// Put is a wrapper around gtk_fixed_put().
func (m *Fixed) Put(w IWidget, x, y int) {
	gtk3.SysCall("gtk_fixed_put", m.Instance(), w.Instance(), uintptr(x), uintptr(y))
}

// Move is a wrapper around gtk_fixed_move().
func (m *Fixed) Move(w IWidget, x, y int) {
	gtk3.SysCall("gtk_fixed_move", m.Instance(), w.Instance(), uintptr(x), uintptr(y))
}
