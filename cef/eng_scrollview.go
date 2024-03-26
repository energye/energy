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

// ICefScrollView Parent: ICefView
//
//	A ScrollView will show horizontal and/or vertical scrollbars when necessary
//	based on the size of the attached content view. Methods must be called on
//	the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_scroll_view_capi.h">CEF source file: /include/capi/views/cef_scroll_view_capi.h (cef_scroll_view_t)</a>
type ICefScrollView interface {
	ICefView
	// GetContentView
	//  Returns the content View.
	GetContentView() ICefView // function
	// GetVisibleContentRect
	//  Returns the visible region of the content View.
	GetVisibleContentRect() (resultCefRect TCefRect) // function
	// HasHorizontalScrollbar
	//  Returns true(1) if the horizontal scrollbar is currently showing.
	HasHorizontalScrollbar() bool // function
	// GetHorizontalScrollbarHeight
	//  Returns the height of the horizontal scrollbar.
	GetHorizontalScrollbarHeight() int32 // function
	// HasVerticalScrollbar
	//  Returns true(1) if the vertical scrollbar is currently showing.
	HasVerticalScrollbar() bool // function
	// GetVerticalScrollbarWidth
	//  Returns the width of the vertical scrollbar.
	GetVerticalScrollbarWidth() int32 // function
	// SetContentView
	//  Set the content View. The content View must have a specified size(e.g.
	//  via ICefView.SetBounds or ICefViewDelegate.GetPreferredSize).
	SetContentView(view ICefView) // procedure
}

// TCefScrollView Parent: TCefView
//
//	A ScrollView will show horizontal and/or vertical scrollbars when necessary
//	based on the size of the attached content view. Methods must be called on
//	the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_scroll_view_capi.h">CEF source file: /include/capi/views/cef_scroll_view_capi.h (cef_scroll_view_t)</a>
type TCefScrollView struct {
	TCefView
}

// ScrollViewRef -> ICefScrollView
var ScrollViewRef scrollView

// scrollView TCefScrollView Ref
type scrollView uintptr

// UnWrap
//
//	Returns a ICefScrollView instance using a PCefScrollView data pointer.
func (m *scrollView) UnWrap(data uintptr) ICefScrollView {
	var resultCefScrollView uintptr
	CEF().SysCallN(1353, uintptr(data), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

// CreateScrollView
//
//	Create a new ScrollView.
func (m *scrollView) CreateScrollView(delegate ICefViewDelegate) ICefScrollView {
	var resultCefScrollView uintptr
	CEF().SysCallN(1345, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

func (m *TCefScrollView) GetContentView() ICefView {
	var resultCefView uintptr
	CEF().SysCallN(1346, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefScrollView) GetVisibleContentRect() (resultCefRect TCefRect) {
	CEF().SysCallN(1349, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefScrollView) HasHorizontalScrollbar() bool {
	r1 := CEF().SysCallN(1350, m.Instance())
	return GoBool(r1)
}

func (m *TCefScrollView) GetHorizontalScrollbarHeight() int32 {
	r1 := CEF().SysCallN(1347, m.Instance())
	return int32(r1)
}

func (m *TCefScrollView) HasVerticalScrollbar() bool {
	r1 := CEF().SysCallN(1351, m.Instance())
	return GoBool(r1)
}

func (m *TCefScrollView) GetVerticalScrollbarWidth() int32 {
	r1 := CEF().SysCallN(1348, m.Instance())
	return int32(r1)
}

func (m *TCefScrollView) SetContentView(view ICefView) {
	CEF().SysCallN(1352, m.Instance(), GetObjectUintptr(view))
}
