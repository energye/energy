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

// ICefBaseRefCounted
func (m *ICefBaseRefCounted) Wrap(data uintptr) unsafe.Pointer {
	var result uintptr
	imports.Proc(internale_CefBaseRefCounted_Wrap).Call(data, uintptr(unsafe.Pointer(&result)))
	return unsafe.Pointer(result)
}

func (m *ICefBaseRefCounted) Free(data uintptr) {
	imports.Proc(internale_CefBaseRefCounted_Free).Call(uintptr(unsafe.Pointer(&data)))
}
