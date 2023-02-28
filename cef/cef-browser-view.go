//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Browser View
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

func (m *ICefBrowserView) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBrowserView) Browser() *ICefBrowser {
	var result uintptr
	imports.Proc(internale_CefBrowserView_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowser{instance: unsafe.Pointer(result)}
}

func (m *ICefBrowserView) ChromeToolbar() *ICefView {
	var result uintptr
	imports.Proc(internale_CefBrowserView_ChromeToolbar).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefView{instance: unsafe.Pointer(result)}
}

func (m *ICefBrowserView) SetPreferAccelerators(preferAccelerators bool) {
	imports.Proc(internale_CefBrowserView_SetPreferAccelerators).Call(m.Instance(), api.PascalBool(preferAccelerators))
}
