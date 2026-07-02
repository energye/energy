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

// List is a representation of GLib's GList.
type List struct {
	data unsafe.Pointer
}

// AsList wraps a raw pointer as an IList.
func AsList(ptr unsafe.Pointer) IList {
	if ptr == nil {
		return nil
	}
	return &List{data: ptr}
}

func (v *List) Instance() uintptr {
	return uintptr(v.data)
}

// Append is a wrapper around g_list_append().
func (v *List) Append(data uintptr) IList {
	gList := glib2_0.SysCall("g_list_append", v.Instance(), data)
	return AsList(unsafe.Pointer(gList))
}

// Prepend is a wrapper around g_list_prepend().
func (v *List) Prepend(data uintptr) IList {
	gList := glib2_0.SysCall("g_list_prepend", v.Instance(), data)
	return AsList(unsafe.Pointer(gList))
}

// Insert is a wrapper around g_list_insert().
func (v *List) Insert(data uintptr, position int) IList {
	gList := glib2_0.SysCall("g_list_insert", v.Instance(), data, uintptr(position))
	return AsList(unsafe.Pointer(gList))
}

// Length is a wrapper around g_list_length().
func (v *List) Length() uint {
	r := glib2_0.SysCall("g_list_length", v.Instance())
	return uint(r)
}

// NthDataRaw is a wrapper around g_list_nth_data().
func (v *List) NthDataRaw(n uint) unsafe.Pointer {
	r := glib2_0.SysCall("g_list_nth_data", v.Instance(), uintptr(n))
	return unsafe.Pointer(r)
}

// Nth is a wrapper around g_list_nth().
func (v *List) Nth(n uint) IList {
	r := glib2_0.SysCall("g_list_nth", v.Instance(), uintptr(n))
	list := AsList(unsafe.Pointer(r))
	return list
}

// NthData acts the same as g_list_nth_data(), but passes
// retrieved value before returning through wrap function, set by DataWrapper().
// If no wrap function is set, it returns raw unsafe.Pointer.
func (v *List) NthData(n uint) any {
	ptr := v.NthDataRaw(n)
	return ptr
}

// Next returns the next element in the list.
func (v *List) Next() IList {
	nextPtr := unsafe.Pointer(v.Instance() + ptrSize)
	return AsList(nextPtr)
}

// Previous returns the previous element in the list.
func (v *List) Previous() IList {
	prevPtr := unsafe.Pointer(v.Instance() + ptrSize*2)
	return AsList(prevPtr)
}

// First is a wrapper around g_list_first().
func (v *List) First() IList {
	r := glib2_0.SysCall("g_list_first", v.Instance())
	return AsList(unsafe.Pointer(r))
}

// Last is a wrapper around g_list_last().
func (v *List) Last() IList {
	r := glib2_0.SysCall("g_list_last", v.Instance())
	return AsList(unsafe.Pointer(r))
}

// Reverse is a wrapper around g_list_reverse().
func (v *List) Reverse() IList {
	r := glib2_0.SysCall("g_list_reverse", v.Instance())
	list := AsList(unsafe.Pointer(r))
	return list
}

// Free is a wrapper around g_list_free().
func (v *List) Free() {
	glib2_0.SysCall("g_list_free", v.Instance())
}
