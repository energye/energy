//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IWVWinControl Parent: IWinControl
//
//	Parent control that will be used by TWVWindowParent in VCL and LCL applications to show the web contents.
type IWVWinControl interface {
	IWinControl
	// ChildWindowHandle
	//  Handle of the first child control created by the browser.
	ChildWindowHandle() THandle   // property
	DragKind() TDragKind          // property
	SetDragKind(AValue TDragKind) // property
	DragCursor() TCursor          // property
	SetDragCursor(AValue TCursor) // property
	DragMode() TDragMode          // property
	SetDragMode(AValue TDragMode) // property
	// CreateHandle
	//  Creates underlying screen object.
	CreateHandle() // procedure
	// InvalidateChildren
	//  Invalidates all child controls created by the browser.
	InvalidateChildren() // procedure
	// UpdateSize
	//  Updates the size of the child controls created by the browser.
	UpdateSize()                       // procedure
	SetOnDragDrop(fn TDragDropEvent)   // property event
	SetOnDragOver(fn TDragOverEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent) // property event
	SetOnEndDrag(fn TEndDragEvent)     // property event
}

// TWVWinControl Parent: TWinControl
//
//	Parent control that will be used by TWVWindowParent in VCL and LCL applications to show the web contents.
type TWVWinControl struct {
	TWinControl
	dragDropPtr  uintptr
	dragOverPtr  uintptr
	startDragPtr uintptr
	endDragPtr   uintptr
}

func NewWVWinControl(TheOwner IComponent) IWVWinControl {
	r1 := WV().SysCallN(1127, GetObjectUintptr(TheOwner))
	return AsWVWinControl(r1)
}

func (m *TWVWinControl) ChildWindowHandle() THandle {
	r1 := WV().SysCallN(1125, m.Instance())
	return THandle(r1)
}

func (m *TWVWinControl) DragKind() TDragKind {
	r1 := WV().SysCallN(1130, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TWVWinControl) SetDragKind(AValue TDragKind) {
	WV().SysCallN(1130, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVWinControl) DragCursor() TCursor {
	r1 := WV().SysCallN(1129, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TWVWinControl) SetDragCursor(AValue TCursor) {
	WV().SysCallN(1129, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVWinControl) DragMode() TDragMode {
	r1 := WV().SysCallN(1131, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TWVWinControl) SetDragMode(AValue TDragMode) {
	WV().SysCallN(1131, 1, m.Instance(), uintptr(AValue))
}

func WVWinControlClass() TClass {
	ret := WV().SysCallN(1126)
	return TClass(ret)
}

func (m *TWVWinControl) CreateHandle() {
	WV().SysCallN(1128, m.Instance())
}

func (m *TWVWinControl) InvalidateChildren() {
	WV().SysCallN(1132, m.Instance())
}

func (m *TWVWinControl) UpdateSize() {
	WV().SysCallN(1137, m.Instance())
}

func (m *TWVWinControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1133, m.Instance(), m.dragDropPtr)
}

func (m *TWVWinControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1134, m.Instance(), m.dragOverPtr)
}

func (m *TWVWinControl) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1136, m.Instance(), m.startDragPtr)
}

func (m *TWVWinControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1135, m.Instance(), m.endDragPtr)
}
