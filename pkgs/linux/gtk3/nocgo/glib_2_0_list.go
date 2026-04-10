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

type List struct {
	data unsafe.Pointer
}

func AsList(ptr unsafe.Pointer) IList {
	if ptr == nil {
		return nil
	}
	return &List{data: ptr}
}

func (m *List) Instance() uintptr {
	return uintptr(m.data)
}

// Append is a wrapper around g_list_append().
func (m *List) Append(data uintptr) IList {
	gList := glib2_0.SysCall("gtk_container_add", m.Instance(), data)
	return AsList(unsafe.Pointer(gList))
}

// Prepend is a wrapper around g_list_prepend().
func (m *List) Prepend(data uintptr) IList {
	gList := glib2_0.SysCall("g_list_prepend", m.Instance(), data)
	return AsList(unsafe.Pointer(gList))
}

// Insert is a wrapper around g_list_insert().
func (m *List) Insert(data uintptr, position int) IList {
	gList := glib2_0.SysCall("g_list_insert", m.Instance(), data, uintptr(position))
	return AsList(unsafe.Pointer(gList))
}

// Length is a wrapper around g_list_length().
func (m *List) Length() uint {
	r := glib2_0.SysCall("g_list_length", m.Instance())
	return uint(r)
}

// NthDataRaw is a wrapper around g_list_nth_data().
func (m *List) NthDataRaw(n uint) unsafe.Pointer {
	r := glib2_0.SysCall("g_list_nth_data", m.Instance(), uintptr(n))
	return unsafe.Pointer(r)
}

// Next is a wrapper around the next struct field
func (m *List) Next() IList {
	nextPtr := unsafe.Pointer(m.Instance() + ptrSize)
	return AsList(nextPtr)
}

// Previous is a wrapper around the prev struct field
func (m *List) Previous() IList {
	prevPtr := unsafe.Pointer(m.Instance() + ptrSize*2)
	return AsList(prevPtr)
}

// First is a wrapper around g_list_first().
func (m *List) First() IList {
	gList := glib2_0.SysCall("g_list_first", m.Instance())
	return AsList(unsafe.Pointer(gList))
}

// Last is a wrapper around g_list_last().
func (m *List) Last() IList {
	gList := glib2_0.SysCall("g_list_last", m.Instance())
	return AsList(unsafe.Pointer(gList))
}

// Free is a wrapper around g_list_free().
func (m *List) Free() {
	glib2_0.SysCall("g_list_free", m.Instance())
}
