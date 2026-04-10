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
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// Layout is a representation of GTK's GtkLayout.
type Layout struct {
	Container
}

func AsLayout(ptr unsafe.Pointer) ILayout {
	if ptr == nil {
		return nil
	}
	m := new(Layout)
	m.instance = ptr
	return m
}

// Layout.Put is a wrapper around gtk_layout_put().
func (m *Layout) Put(w IWidget, x, y int) {
	gtk3.SysCall("gtk_layout_put", m.Instance(), w.Instance(), uintptr(x), uintptr(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (m *Layout) Move(w IWidget, x, y int) {
	gtk3.SysCall("gtk_layout_move", m.Instance(), w.Instance(), uintptr(x), uintptr(y))
}

// Layout.SetSize is a wrapper around gtk_layout_set_size
func (m *Layout) SetSize(width, height uint) {
	gtk3.SysCall("gtk_layout_set_size", m.Instance(), uintptr(width), uintptr(height))
}

// Layout.GetSize is a wrapper around gtk_layout_get_size
func (m *Layout) GetSize() (width, height uint) {
	gtk3.SysCall("gtk_layout_get_size", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}
