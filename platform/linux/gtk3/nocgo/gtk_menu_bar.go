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

// MenuBar is a representation of GTK's GtkMenuBar.
type MenuBar struct {
	MenuShell
}

func AsMenuBar(ptr unsafe.Pointer) IMenuBar {
	if ptr == nil {
		return nil
	}
	m := new(MenuBar)
	m.instance = ptr
	return m
}
