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

// ICheckListBox Parent: ICustomCheckListBox
type ICheckListBox interface {
	ICustomCheckListBox
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
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

// TCheckListBox Parent: TCustomCheckListBox
type TCheckListBox struct {
	TCustomCheckListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewCheckListBox(AOwner IComponent) ICheckListBox {
	r1 := LCL().SysCallN(456, GetObjectUintptr(AOwner))
	return AsCheckListBox(r1)
}

func (m *TCheckListBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(457, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckListBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(457, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckListBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(458, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckListBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(458, 1, m.Instance(), uintptr(AValue))
}

func CheckListBoxClass() TClass {
	ret := LCL().SysCallN(455)
	return TClass(ret)
}

func (m *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(459, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(460, m.Instance(), m.dragDropPtr)
}

func (m *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(461, m.Instance(), m.dragOverPtr)
}

func (m *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(462, m.Instance(), m.endDragPtr)
}

func (m *TCheckListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(463, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TCheckListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(464, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TCheckListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(465, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TCheckListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(466, m.Instance(), m.startDragPtr)
}
