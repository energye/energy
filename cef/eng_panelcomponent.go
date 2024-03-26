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

// ICEFPanelComponent Parent: ICEFViewComponent
type ICEFPanelComponent interface {
	ICEFViewComponent
	// AsWindow
	//  Returns this Panel as a Window or NULL if this is not a Window.
	AsWindow() ICefWindow // property
	// ChildViewCount
	//  Returns the number of child Views.
	ChildViewCount() NativeUInt // property
	// SetToFillLayout
	//  Set this Panel's Layout to FillLayout and return the FillLayout object.
	SetToFillLayout() ICefFillLayout // function
	// SetToBoxLayout
	//  Set this Panel's Layout to BoxLayout and return the BoxLayout object.
	SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout // function
	// GetLayout
	//  Get the Layout.
	GetLayout() ICefLayout // function
	// GetChildViewAt
	//  Returns the child View at the specified |index|.
	GetChildViewAt(index int32) ICefView // function
	// CreatePanel
	//  Create a new Panel.
	CreatePanel() // procedure
	// Layout
	//  Lay out the child Views(set their bounds based on sizing heuristics
	//  specific to the current Layout).
	Layout() // procedure
	// AddChildView
	//  Add a child View.
	AddChildView(view ICefView) // procedure
	// AddChildViewAt
	//  Add a child View at the specified |index|. If |index| matches the result
	//  of GetChildCount() then the View will be added at the end.
	AddChildViewAt(view ICefView, index int32) // procedure
	// ReorderChildView
	//  Move the child View to the specified |index|. A negative value for |index|
	//  will move the View to the end.
	ReorderChildView(view ICefView, index int32) // procedure
	// RemoveChildView
	//  Remove a child View. The View can then be added to another Panel.
	RemoveChildView(view ICefView) // procedure
	// RemoveAllChildViews
	//  Remove all child Views. The removed Views will be deleted if the client
	//  holds no references to them.
	RemoveAllChildViews() // procedure
}

// TCEFPanelComponent Parent: TCEFViewComponent
type TCEFPanelComponent struct {
	TCEFViewComponent
}

func NewCEFPanelComponent(aOwner IComponent) ICEFPanelComponent {
	r1 := CEF().SysCallN(171, GetObjectUintptr(aOwner))
	return AsCEFPanelComponent(r1)
}

func (m *TCEFPanelComponent) AsWindow() ICefWindow {
	var resultCefWindow uintptr
	CEF().SysCallN(168, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCEFPanelComponent) ChildViewCount() NativeUInt {
	r1 := CEF().SysCallN(169, m.Instance())
	return NativeUInt(r1)
}

func (m *TCEFPanelComponent) SetToFillLayout() ICefFillLayout {
	var resultCefFillLayout uintptr
	CEF().SysCallN(180, m.Instance(), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}

func (m *TCEFPanelComponent) SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout {
	var resultCefBoxLayout uintptr
	CEF().SysCallN(179, m.Instance(), uintptr(unsafePointer(settings)), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCEFPanelComponent) GetLayout() ICefLayout {
	var resultCefLayout uintptr
	CEF().SysCallN(174, m.Instance(), uintptr(unsafePointer(&resultCefLayout)))
	return AsCefLayout(resultCefLayout)
}

func (m *TCEFPanelComponent) GetChildViewAt(index int32) ICefView {
	var resultCefView uintptr
	CEF().SysCallN(173, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func CEFPanelComponentClass() TClass {
	ret := CEF().SysCallN(170)
	return TClass(ret)
}

func (m *TCEFPanelComponent) CreatePanel() {
	CEF().SysCallN(172, m.Instance())
}

func (m *TCEFPanelComponent) Layout() {
	CEF().SysCallN(175, m.Instance())
}

func (m *TCEFPanelComponent) AddChildView(view ICefView) {
	CEF().SysCallN(166, m.Instance(), GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) AddChildViewAt(view ICefView, index int32) {
	CEF().SysCallN(167, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) ReorderChildView(view ICefView, index int32) {
	CEF().SysCallN(178, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) RemoveChildView(view ICefView) {
	CEF().SysCallN(177, m.Instance(), GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) RemoveAllChildViews() {
	CEF().SysCallN(176, m.Instance())
}
