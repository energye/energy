//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package win32

import (
	"unicode/utf16"
	"unsafe"
)

func UTF16PtrToString(cStr *uint16) string {
	if cStr != nil {
		us := make([]uint16, 0, 256)
		for ptr := uintptr(unsafe.Pointer(cStr)); ; ptr += 2 {
			u := *(*uint16)(unsafe.Pointer(ptr))
			if u == 0 { // null end
				return string(utf16.Decode(us))
			}
			us = append(us, u)
		}
	}
	return ""
}
