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

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	Object
}

func AsCellRendererText(ptr unsafe.Pointer) *CellRendererText {
	if ptr == nil {
		return nil
	}
	m := new(CellRendererText)
	m.instance = ptr
	return m
}

// NewCellRendererText is a wrapper around gtk_cell_renderer_text_new().
func NewCellRendererText() *CellRendererText {
	r := gtk3.SysCall("gtk_cell_renderer_text_new")
	if r == 0 {
		return nil
	}
	return AsCellRendererText(unsafe.Pointer(r))
}
