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

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

func AsMenuShell(ptr unsafe.Pointer) IMenuShell {
	if ptr == nil {
		return nil
	}
	m := new(MenuShell)
	m.instance = ptr
	return m
}
