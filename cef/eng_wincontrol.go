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

// ICEFWinControl Parent: IWinControl
//
//	Custom TWinControl used by CEF browsers.
type ICEFWinControl interface {
	IWinControl
	// ChildWindowHandle
	//  Handle of the first child window created by the browser.
	ChildWindowHandle() THandle   // property
	DragKind() TDragKind          // property
	SetDragKind(AValue TDragKind) // property
	DragCursor() TCursor          // property
	SetDragCursor(AValue TCursor) // property
	DragMode() TDragMode          // property
	SetDragMode(AValue TDragMode) // property
	// TakeSnapshot
	//  Take a snapshot of the browser contents into aBitmap. This function only works in Windows without hardware acceleration.
	TakeSnapshot(aBitmap *IBitmap) bool // function
	// DestroyChildWindow
	//  Destroy the child windows created by the browser.
	DestroyChildWindow() bool // function
	// CreateHandle
	//  Exposes the CreateHandle procedure to create the Handle at any time.
	CreateHandle() // procedure
	// InvalidateChildren
	//  Invalidate the child windows created by the browser.
	InvalidateChildren() // procedure
	// UpdateSize
	//  Updates the size of the child windows created by the browser.
	UpdateSize()                  // procedure
	SetOnDragDrop(fn TDragDrop)   // property event
	SetOnDragOver(fn TDragOver)   // property event
	SetOnStartDrag(fn TStartDrag) // property event
	SetOnEndDrag(fn TEndDrag)     // property event
}

// TCEFWinControl Parent: TWinControl
//
//	Custom TWinControl used by CEF browsers.
type TCEFWinControl struct {
	TWinControl
	dragDropPtr  uintptr
	dragOverPtr  uintptr
	startDragPtr uintptr
	endDragPtr   uintptr
}

func NewCEFWinControl(theOwner IComponent) ICEFWinControl {
	r1 := CEF().SysCallN(334, GetObjectUintptr(theOwner))
	return AsCEFWinControl(r1)
}

func (m *TCEFWinControl) ChildWindowHandle() THandle {
	r1 := CEF().SysCallN(332, m.Instance())
	return THandle(r1)
}

func (m *TCEFWinControl) DragKind() TDragKind {
	r1 := CEF().SysCallN(338, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCEFWinControl) SetDragKind(AValue TDragKind) {
	CEF().SysCallN(338, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) DragCursor() TCursor {
	r1 := CEF().SysCallN(337, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCEFWinControl) SetDragCursor(AValue TCursor) {
	CEF().SysCallN(337, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) DragMode() TDragMode {
	r1 := CEF().SysCallN(339, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCEFWinControl) SetDragMode(AValue TDragMode) {
	CEF().SysCallN(339, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) TakeSnapshot(aBitmap *IBitmap) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(345, m.Instance(), uintptr(unsafePointer(&result0)))
	*aBitmap = AsBitmap(result0)
	return GoBool(r1)
}

func (m *TCEFWinControl) DestroyChildWindow() bool {
	r1 := CEF().SysCallN(336, m.Instance())
	return GoBool(r1)
}

func CEFWinControlClass() TClass {
	ret := CEF().SysCallN(333)
	return TClass(ret)
}

func (m *TCEFWinControl) CreateHandle() {
	CEF().SysCallN(335, m.Instance())
}

func (m *TCEFWinControl) InvalidateChildren() {
	CEF().SysCallN(340, m.Instance())
}

func (m *TCEFWinControl) UpdateSize() {
	CEF().SysCallN(347, m.Instance())
}

func (m *TCEFWinControl) SetOnDragDrop(fn TDragDrop) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(341, m.Instance(), m.dragDropPtr)
}

func (m *TCEFWinControl) SetOnDragOver(fn TDragOver) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(342, m.Instance(), m.dragOverPtr)
}

func (m *TCEFWinControl) SetOnStartDrag(fn TStartDrag) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(344, m.Instance(), m.startDragPtr)
}

func (m *TCEFWinControl) SetOnEndDrag(fn TEndDrag) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(343, m.Instance(), m.endDragPtr)
}
