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

// IColorListBox Parent: ICustomColorListBox
type IColorListBox interface {
	ICustomColorListBox
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

// TColorListBox Parent: TCustomColorListBox
type TColorListBox struct {
	TCustomColorListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewColorListBox(AOwner IComponent) IColorListBox {
	r1 := LCL().SysCallN(768, GetObjectUintptr(AOwner))
	return AsColorListBox(r1)
}

func (m *TColorListBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(769, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TColorListBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(769, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorListBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(770, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TColorListBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(770, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorListBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(771, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TColorListBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(771, 1, m.Instance(), uintptr(AValue))
}

func ColorListBoxClass() TClass {
	ret := LCL().SysCallN(767)
	return TClass(ret)
}

func (m *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(772, m.Instance(), m.contextPopupPtr)
}

func (m *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(773, m.Instance(), m.dragDropPtr)
}

func (m *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(774, m.Instance(), m.dragOverPtr)
}

func (m *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(775, m.Instance(), m.endDragPtr)
}

func (m *TColorListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(776, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TColorListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(777, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TColorListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(778, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TColorListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(779, m.Instance(), m.startDragPtr)
}
