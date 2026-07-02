//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
// Licensed under Apache License Version 2.0, January 2004
// https://www.apache.org/licenses/LICENSE-2.0
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Separator is a representation of GTK's GtkSeparator.
type Separator struct {
	Widget
}

func AsSeparator(ptr unsafe.Pointer) *Separator {
	if ptr == nil {
		return nil
	}
	m := new(Separator)
	m.instance = ptr
	return m
}

// NewSeparator is a wrapper around gtk_separator_new().
func NewSeparator(orientation Orientation) *Separator {
	r := gtk3.SysCall("gtk_separator_new", uintptr(orientation))
	if r == 0 {
		return nil
	}
	return AsSeparator(unsafe.Pointer(r))
}
