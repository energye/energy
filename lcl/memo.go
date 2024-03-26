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

// IMemo Parent: ICustomMemo
type IMemo interface {
	ICustomMemo
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEditingDone(fn TNotifyEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TMemo Parent: TCustomMemo
type TMemo struct {
	TCustomMemo
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	editingDonePtr     uintptr
	endDragPtr         uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewMemo(AOwner IComponent) IMemo {
	r1 := LCL().SysCallN(3564, GetObjectUintptr(AOwner))
	return AsMemo(r1)
}

func (m *TMemo) DragCursor() TCursor {
	r1 := LCL().SysCallN(3565, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TMemo) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3565, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) DragKind() TDragKind {
	r1 := LCL().SysCallN(3566, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TMemo) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3566, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) DragMode() TDragMode {
	r1 := LCL().SysCallN(3567, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TMemo) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3567, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) ParentColor() bool {
	r1 := LCL().SysCallN(3568, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentColor(AValue bool) {
	LCL().SysCallN(3568, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMemo) ParentFont() bool {
	r1 := LCL().SysCallN(3569, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentFont(AValue bool) {
	LCL().SysCallN(3569, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMemo) ParentShowHint() bool {
	r1 := LCL().SysCallN(3570, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3570, 1, m.Instance(), PascalBool(AValue))
}

func MemoClass() TClass {
	ret := LCL().SysCallN(3563)
	return TClass(ret)
}

func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3571, m.Instance(), m.contextPopupPtr)
}

func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3572, m.Instance(), m.dblClickPtr)
}

func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3573, m.Instance(), m.dragDropPtr)
}

func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3574, m.Instance(), m.dragOverPtr)
}

func (m *TMemo) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3575, m.Instance(), m.editingDonePtr)
}

func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3576, m.Instance(), m.endDragPtr)
}

func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3577, m.Instance(), m.mouseDownPtr)
}

func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3578, m.Instance(), m.mouseEnterPtr)
}

func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3579, m.Instance(), m.mouseLeavePtr)
}

func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3580, m.Instance(), m.mouseMovePtr)
}

func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3581, m.Instance(), m.mouseUpPtr)
}

func (m *TMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3582, m.Instance(), m.mouseWheelPtr)
}

func (m *TMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3583, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3587, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TMemo) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3584, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TMemo) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3585, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TMemo) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3586, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TMemo) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3588, m.Instance(), m.startDragPtr)
}
