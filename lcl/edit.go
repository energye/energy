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

// IEdit Parent: ICustomEdit
type IEdit interface {
	ICustomEdit
	AutoSelected() bool                            // property
	SetAutoSelected(AValue bool)                   // property
	AutoSelect() bool                              // property
	SetAutoSelect(AValue bool)                     // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TEdit Parent: TCustomEdit
type TEdit struct {
	TCustomEdit
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewEdit(AOwner IComponent) IEdit {
	r1 := LCL().SysCallN(2793, GetObjectUintptr(AOwner))
	return AsEdit(r1)
}

func (m *TEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(2791, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(2791, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(2790, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(2790, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(2794, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2794, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) DragKind() TDragKind {
	r1 := LCL().SysCallN(2795, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TEdit) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(2795, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(2796, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2796, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) ParentColor() bool {
	r1 := LCL().SysCallN(2797, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(2797, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) ParentFont() bool {
	r1 := LCL().SysCallN(2798, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(2798, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(2799, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2799, 1, m.Instance(), PascalBool(AValue))
}

func EditClass() TClass {
	ret := LCL().SysCallN(2792)
	return TClass(ret)
}

func (m *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2800, m.Instance(), m.contextPopupPtr)
}

func (m *TEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2801, m.Instance(), m.dblClickPtr)
}

func (m *TEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2802, m.Instance(), m.dragDropPtr)
}

func (m *TEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2803, m.Instance(), m.dragOverPtr)
}

func (m *TEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2804, m.Instance(), m.editingDonePtr)
}

func (m *TEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2805, m.Instance(), m.endDragPtr)
}

func (m *TEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2806, m.Instance(), m.mouseDownPtr)
}

func (m *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2807, m.Instance(), m.mouseEnterPtr)
}

func (m *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2808, m.Instance(), m.mouseLeavePtr)
}

func (m *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2809, m.Instance(), m.mouseMovePtr)
}

func (m *TEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2810, m.Instance(), m.mouseUpPtr)
}

func (m *TEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2811, m.Instance(), m.mouseWheelPtr)
}

func (m *TEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2812, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2813, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2814, m.Instance(), m.startDragPtr)
}
