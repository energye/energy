//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IListBox Parent: ICustomListBox
type IListBox interface {
	ICustomListBox
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TListBox Parent: TCustomListBox
type TListBox struct {
	TCustomListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewListBox(TheOwner IComponent) IListBox {
	r1 := LCL().SysCallN(3331, GetObjectUintptr(TheOwner))
	return AsListBox(r1)
}

func (m *TListBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(3332, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TListBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3332, 1, m.Instance(), uintptr(AValue))
}

func (m *TListBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(3333, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TListBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3333, 1, m.Instance(), uintptr(AValue))
}

func (m *TListBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(3334, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TListBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3334, 1, m.Instance(), uintptr(AValue))
}

func ListBoxClass() TClass {
	ret := LCL().SysCallN(3330)
	return TClass(ret)
}

func (m *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3335, m.Instance(), m.contextPopupPtr)
}

func (m *TListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3336, m.Instance(), m.dragDropPtr)
}

func (m *TListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3337, m.Instance(), m.dragOverPtr)
}

func (m *TListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3338, m.Instance(), m.endDragPtr)
}

func (m *TListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3339, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3340, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3341, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3342, m.Instance(), m.startDragPtr)
}
