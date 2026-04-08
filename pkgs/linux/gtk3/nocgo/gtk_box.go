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
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

func AsBox(ptr unsafe.Pointer) IBox {
	if ptr == nil {
		return nil
	}
	m := new(Box)
	m.instance = ptr
	return m
}

// PackStart is a wrapper around gtk_box_pack_start().
func (m *Box) PackStart(child IWidget, expand, fill bool, padding uint) {
	gtk3.SysCall("gtk_box_pack_start", m.Instance(), child.Instance(), ToCBool(expand), ToCBool(fill), uintptr(padding))
}

// PackEnd is a wrapper around gtk_box_pack_end().
func (m *Box) PackEnd(child IWidget, expand, fill bool, padding uint) {
	gtk3.SysCall("gtk_box_pack_end", m.Instance(), child.Instance(), ToCBool(expand), ToCBool(fill), uintptr(padding))
}
