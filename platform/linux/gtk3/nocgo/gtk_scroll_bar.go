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

// Scrollbar is a representation of GTK's GtkScrollbar.
type Scrollbar struct {
	Range
}

func AsScrollbar(ptr unsafe.Pointer) IScrollbar {
	if ptr == nil {
		return nil
	}
	m := new(Scrollbar)
	m.instance = ptr
	return m
}

// NewScrollbar is a wrapper around gtk_scrollbar_new().
func NewScrollbar(orientation Orientation, adjustment IAdjustment) IScrollbar {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	r := gtk3.SysCall("gtk_scrollbar_new", uintptr(orientation), adj)
	if r == 0 {
		return nil
	}
	return AsScrollbar(unsafe.Pointer(r))
}
