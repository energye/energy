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

// ICefLayout Parent: ICefBaseRefCounted
//
//	A Layout handles the sizing of the children of a Panel according to
//	implementation-specific heuristics. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_layout_capi.h">CEF source file: /include/capi/views/cef_layout_capi.h (cef_layout_t)</a>
type ICefLayout interface {
	ICefBaseRefCounted
	// AsBoxLayout
	//  Returns this Layout as a BoxLayout or NULL if this is not a BoxLayout.
	AsBoxLayout() ICefBoxLayout // function
	// AsFillLayout
	//  Returns this Layout as a FillLayout or NULL if this is not a FillLayout.
	AsFillLayout() ICefFillLayout // function
	// IsValid
	//  Returns true(1) if this Layout is valid.
	IsValid() bool // function
}

// TCefLayout Parent: TCefBaseRefCounted
//
//	A Layout handles the sizing of the children of a Panel according to
//	implementation-specific heuristics. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_layout_capi.h">CEF source file: /include/capi/views/cef_layout_capi.h (cef_layout_t)</a>
type TCefLayout struct {
	TCefBaseRefCounted
}

// LayoutRef -> ICefLayout
var LayoutRef layout

// layout TCefLayout Ref
type layout uintptr

// UnWrap
//
//	Returns a ICefLayout instance using a PCefLayout data pointer.
func (m *layout) UnWrap(data uintptr) ICefLayout {
	var resultCefLayout uintptr
	CEF().SysCallN(1019, uintptr(data), uintptr(unsafePointer(&resultCefLayout)))
	return AsCefLayout(resultCefLayout)
}

func (m *TCefLayout) AsBoxLayout() ICefBoxLayout {
	var resultCefBoxLayout uintptr
	CEF().SysCallN(1016, m.Instance(), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCefLayout) AsFillLayout() ICefFillLayout {
	var resultCefFillLayout uintptr
	CEF().SysCallN(1017, m.Instance(), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}

func (m *TCefLayout) IsValid() bool {
	r1 := CEF().SysCallN(1018, m.Instance())
	return GoBool(r1)
}
