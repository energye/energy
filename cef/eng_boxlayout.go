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

// ICefBoxLayout Parent: ICefLayout
//
//	A Layout manager that arranges child views vertically or horizontally in a
//	side-by-side fashion with spacing around and between the child views. The
//	child views are always sized according to their preferred size. If the
//	host's bounds provide insufficient space, child views will be clamped.
//	Excess space will not be distributed. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a cref="uCEFTypes|TCefBoxLayout">Implements TCefBoxLayout</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_box_layout_capi.h">CEF source file: /include/capi/views/cef_box_layout_capi.h (cef_box_layout_t)</a>
type ICefBoxLayout interface {
	ICefLayout
	// SetFlexForView
	//  Set the flex weight for the given |view|. Using the preferred size as the
	//  basis, free space along the main axis is distributed to views in the ratio
	//  of their flex weights. Similarly, if the views will overflow the parent,
	//  space is subtracted in these ratios. A flex of 0 means this view is not
	//  resized. Flex values must not be negative.
	SetFlexForView(view ICefView, flex int32) // procedure
	// ClearFlexForView
	//  Clears the flex for the given |view|, causing it to use the default flex
	//  specified via TCefBoxLayoutSettings.default_flex.
	ClearFlexForView(view ICefView) // procedure
}

// TCefBoxLayout Parent: TCefLayout
//
//	A Layout manager that arranges child views vertically or horizontally in a
//	side-by-side fashion with spacing around and between the child views. The
//	child views are always sized according to their preferred size. If the
//	host's bounds provide insufficient space, child views will be clamped.
//	Excess space will not be distributed. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a cref="uCEFTypes|TCefBoxLayout">Implements TCefBoxLayout</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_box_layout_capi.h">CEF source file: /include/capi/views/cef_box_layout_capi.h (cef_box_layout_t)</a>
type TCefBoxLayout struct {
	TCefLayout
}

// BoxLayoutRef -> ICefBoxLayout
var BoxLayoutRef boxLayout

// boxLayout TCefBoxLayout Ref
type boxLayout uintptr

// UnWrap
//
//	Returns a ICefBoxLayout instance using a PCefBoxLayout data pointer.
func (m *boxLayout) UnWrap(data uintptr) ICefBoxLayout {
	var resultCefBoxLayout uintptr
	CEF().SysCallN(608, uintptr(data), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCefBoxLayout) SetFlexForView(view ICefView, flex int32) {
	CEF().SysCallN(607, m.Instance(), GetObjectUintptr(view), uintptr(flex))
}

func (m *TCefBoxLayout) ClearFlexForView(view ICefView) {
	CEF().SysCallN(606, m.Instance(), GetObjectUintptr(view))
}
