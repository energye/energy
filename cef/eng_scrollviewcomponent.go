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

// ICEFScrollViewComponent Parent: ICEFViewComponent
type ICEFScrollViewComponent interface {
	ICEFViewComponent
	// ContentView
	//  Get and set the content View. The content View must have a specified size(e.g.
	//  via ICefView.SetBounds or ICefViewDelegate.GetPreferredSize).
	ContentView() ICefView // property
	// SetContentView Set ContentView
	SetContentView(AValue ICefView) // property
	// VisibleContentRect
	//  Returns the visible region of the content View.
	VisibleContentRect() (resultCefRect TCefRect) // property
	// HorizontalScrollbarHeight
	//  Returns the height of the horizontal scrollbar.
	HorizontalScrollbarHeight() int32 // property
	// VerticalScrollbarWidth
	//  Returns the width of the vertical scrollbar.
	VerticalScrollbarWidth() int32 // property
	// HasHorizontalScrollbar
	//  Returns true(1) if the horizontal scrollbar is currently showing.
	HasHorizontalScrollbar() bool // property
	// HasVerticalScrollbar
	//  Returns true(1) if the vertical scrollbar is currently showing.
	HasVerticalScrollbar() bool // property
	// CreateScrollView
	//  Create a new ScrollView.
	CreateScrollView() // procedure
}

// TCEFScrollViewComponent Parent: TCEFViewComponent
type TCEFScrollViewComponent struct {
	TCEFViewComponent
}

func NewCEFScrollViewComponent(aOwner IComponent) ICEFScrollViewComponent {
	r1 := CEF().SysCallN(183, GetObjectUintptr(aOwner))
	return AsCEFScrollViewComponent(r1)
}

func (m *TCEFScrollViewComponent) ContentView() ICefView {
	var resultCefView uintptr
	CEF().SysCallN(182, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFScrollViewComponent) SetContentView(AValue ICefView) {
	CEF().SysCallN(182, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFScrollViewComponent) VisibleContentRect() (resultCefRect TCefRect) {
	CEF().SysCallN(189, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCEFScrollViewComponent) HorizontalScrollbarHeight() int32 {
	r1 := CEF().SysCallN(187, m.Instance())
	return int32(r1)
}

func (m *TCEFScrollViewComponent) VerticalScrollbarWidth() int32 {
	r1 := CEF().SysCallN(188, m.Instance())
	return int32(r1)
}

func (m *TCEFScrollViewComponent) HasHorizontalScrollbar() bool {
	r1 := CEF().SysCallN(185, m.Instance())
	return GoBool(r1)
}

func (m *TCEFScrollViewComponent) HasVerticalScrollbar() bool {
	r1 := CEF().SysCallN(186, m.Instance())
	return GoBool(r1)
}

func CEFScrollViewComponentClass() TClass {
	ret := CEF().SysCallN(181)
	return TClass(ret)
}

func (m *TCEFScrollViewComponent) CreateScrollView() {
	CEF().SysCallN(184, m.Instance())
}
