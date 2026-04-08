//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package nocgo

import (
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

func GErrorFree(errPtr uintptr) {
	glib2_0.SysCall("g_error_free", errPtr)
}

var glib2_0 *dnyLibrary

func init() {
	glib2_0 = libLoad(libglib2_0)
	setLibClose(glib2_0)
	glib2_0.Table = []*imports.Table{
		imports.NewTable("strlen", 0),
		imports.NewTable("g_error_free", 0),
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
	glib2_0.mapperIndex()
}
