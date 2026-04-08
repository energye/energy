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
	"github.com/energye/energy/v3/pkgs/linux"
	"github.com/energye/lcl/api/imports"
	"reflect"
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

func ucharString(guchar uintptr) string {
	// Seek and find the string length.
	var strlen int
	for ptr := guchar; ptr != 0; ptr = nextguchar(ptr) {
		strlen++
	}
	// Array of unsigned char means GoString is unavailable, so maybe this is
	// fine.
	var data []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Len = strlen
	sliceHeader.Cap = strlen
	sliceHeader.Data = guchar

	// Return a copy of the string.
	return string(data)
}

func nextguchar(guchar uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(guchar + 1))
}

var gobject2_0 *linux.DnyLibrary

func init() {
	gobject2_0 = linux.LibLoad(linux.Libgobject2_0)
	gobject2_0.Table = []*imports.Table{
		imports.NewTable("g_object_ref", 0),
		imports.NewTable("g_object_unref", 0),
	}
	gobject2_0.SetLibClose()
	gobject2_0.MapperIndex()
}
