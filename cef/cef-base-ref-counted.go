//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"unsafe"
)

type ICefBaseRefCounted interface {
	Wrap(data uintptr) unsafe.Pointer
	Free(data uintptr)
}

// Wrap 指针引用包裹
func (m *TCefBaseRefCounted) Wrap(data uintptr) unsafe.Pointer {
	var result uintptr
	imports.Proc(internale_CefBaseRefCounted_Wrap).Call(data, uintptr(unsafe.Pointer(&result)))
	return unsafe.Pointer(result)
}

// Free 释放底层指针
func (m *TCefBaseRefCounted) Free(data uintptr) {
	imports.Proc(internale_CefBaseRefCounted_Free).Call(uintptr(unsafe.Pointer(&data)))
	m.instance = nil
}
