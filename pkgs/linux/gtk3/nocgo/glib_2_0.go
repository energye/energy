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
	"github.com/energye/energy/v3/pkgs/linux"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

//	type GError struct {
//	   GQuark domain;   // uint32
//	   gint code;       // int32
//	   gchar *message;  // pointer (uintptr)
//	}
type GError struct {
	Domain  uint32
	Code    int32
	Message uintptr
}

func GoStr(cStrPtr uintptr) string {
	if cStrPtr == 0 {
		return ""
	}
	strLen := int(glib2_0.SysCall("strlen", cStrPtr))
	if strLen == 0 {
		return ""
	}
	return string((*[1 << 30]byte)(unsafe.Pointer(cStrPtr))[:strLen:strLen])
	//p := (*byte)(unsafe.Pointer(cStrPtr))
	//return unsafe.String(p, strLen)
}

// CStr string 转 底层字符串指针
func CStr(str string) uintptr {
	return api.PasStr(str)
}

func GErrorFree(ptr uintptr) {
	glib2_0.SysCall("g_error_free", ptr)
}

func GFree(ptr uintptr) {
	glib2_0.SysCall("g_free", ptr)
}

func GMalloc(size uintptr) uintptr {
	r := glib2_0.SysCall("g_malloc", size)
	return r
}

// makeCStringArray 构建char**，返回 base 指针和每个字符串指针列表
func makeCStringArray(strs []string) (base uintptr, items []uintptr) {
	count := len(strs)
	arraySize := uintptr(count+1) * ptrSize
	base = GMalloc(arraySize)
	if base == 0 {
		return 0, nil
	}
	items = make([]uintptr, count)
	for i, s := range strs {
		cstr := CStr(s)
		items[i] = cstr
		ptrAddr := base + uintptr(i)*ptrSize
		*(*unsafe.Pointer)(unsafe.Pointer(ptrAddr)) = unsafe.Pointer(cstr)
	}

	nullAddr := base + uintptr(count)*ptrSize
	*(*unsafe.Pointer)(unsafe.Pointer(nullAddr)) = nil
	return base, items
}

// freeCStringArray 释放 base 和所有字符串
func freeCStringArray(base uintptr, items []uintptr) {
	for _, p := range items {
		if p != 0 {
			GFree(p)
		}
	}
	if base != 0 {
		GFree(base)
	}
}

func toGoStringArray(cArrPtr uintptr) []string {
	if cArrPtr == 0 {
		return nil
	}
	var array []string
	for i := 0; ; i++ {
		ptr := *(*uintptr)(unsafe.Pointer(cArrPtr + uintptr(i)*ptrSize))
		if ptr == 0 {
			break
		}
		s := GoStr(ptr)
		array = append(array, s)
	}
	return array
}

var glib2_0 *linux.DnyLibrary

func init() {
	glib2_0 = linux.LibLoad(linux.Libglib2_0)
	glib2_0.Table = []*imports.Table{
		imports.NewTable("strlen", 0),
		imports.NewTable("g_error_free", 0),
		imports.NewTable("g_free", 0),
		imports.NewTable("g_malloc", 0),
		// list
		imports.NewTable("g_list_append", 0),
		imports.NewTable("g_list_prepend", 0),
		imports.NewTable("g_list_insert", 0),
		imports.NewTable("g_list_length", 0),
		imports.NewTable("g_list_nth_data", 0),
		imports.NewTable("g_list_first", 0),
		imports.NewTable("g_list_last", 0),
		imports.NewTable("g_list_free", 0),
	}
	glib2_0.SetLibClose()
	glib2_0.MapperIndex()
}
