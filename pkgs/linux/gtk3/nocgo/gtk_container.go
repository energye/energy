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

type Container struct {
	Widget
}

func AsContainer(ptr unsafe.Pointer) IContainer {
	if ptr == nil {
		return nil
	}
	m := new(Container)
	m.instance = ptr
	return m
}

// Add is a wrapper around gtk_container_add().
func (m *Container) Add(w IWidget) {
	gtk3.SysCall("gtk_container_add", m.Instance(), w.Instance())
}

// Remove is a wrapper around gtk_container_remove().
func (m *Container) Remove(w IWidget) {
	gtk3.SysCall("gtk_container_remove", m.Instance(), w.Instance())
}

// CheckResize is a wrapper around gtk_container_check_resize().
func (m *Container) CheckResize() {
	gtk3.SysCall("gtk_container_check_resize", m.Instance())
}

// GetChildren is a wrapper around gtk_container_get_children().
func (m *Container) GetChildren() IList {
	gList := gtk3.SysCall("gtk_container_get_children", m.Instance())
	if gList == 0 {
		return nil
	}
	return AsList(unsafe.Pointer(gList))
}
