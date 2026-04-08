// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package nocgo

import (
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

type Bin struct {
	Container
}

type Object struct {
	instance unsafe.Pointer
}

func (m *Object) Instance() uintptr {
	return uintptr(m.instance)
}

// Ref is a wrapper around g_object_ref().
func (m *Object) Ref() {
	gobject2_0.SysCall("g_object_ref", m.Instance())
}

// Unref is a wrapper around g_object_unref().
func (m *Object) Unref() {
	gobject2_0.SysCall("g_object_unref", m.Instance())
}

var gobject2_0 *dnyLibrary

func init() {
	gobject2_0 = libLoad(libgobject2_0)
	setLibClose(gobject2_0)
	gobject2_0.Table = []*imports.Table{
		imports.NewTable("g_object_ref", 0),
		imports.NewTable("g_object_unref", 0),
	}
	gobject2_0.mapperIndex()
}
