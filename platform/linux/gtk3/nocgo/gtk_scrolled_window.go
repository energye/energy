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

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

func AsScrolledWindow(ptr unsafe.Pointer) IScrolledWindow {
	if ptr == nil {
		return nil
	}
	m := new(ScrolledWindow)
	m.instance = ptr
	return m
}
