//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefFillLayout Parent: ICefLayout
//
//	A simple Layout that causes the associated Panel's one child to be sized to
//	match the bounds of its parent. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_fill_layout_capi.h">CEF source file: /include/capi/views/cef_fill_layout_capi.h (cef_fill_layout_t)</a>
type ICefFillLayout interface {
	ICefLayout
}

// TCefFillLayout Parent: TCefLayout
//
//	A simple Layout that causes the associated Panel's one child to be sized to
//	match the bounds of its parent. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_fill_layout_capi.h">CEF source file: /include/capi/views/cef_fill_layout_capi.h (cef_fill_layout_t)</a>
type TCefFillLayout struct {
	TCefLayout
}

// FillLayoutRef -> ICefFillLayout
var FillLayoutRef fillLayout

// fillLayout TCefFillLayout Ref
type fillLayout uintptr

// UnWrap
//
//	Returns a ICefFillLayout instance using a PCefFillLayout data pointer.
func (m *fillLayout) UnWrap(data uintptr) ICefFillLayout {
	var resultCefFillLayout uintptr
	CEF().SysCallN(954, uintptr(data), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}
