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

// IButton Parent: ICustomButton
type IButton interface {
	ICustomButton
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
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

// TButton Parent: TCustomButton
type TButton struct {
	TCustomButton
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
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

func NewButton(TheOwner IComponent) IButton {
	r1 := LCL().SysCallN(290, GetObjectUintptr(TheOwner))
	return AsButton(r1)
}

func (m *TButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(291, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(291, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(292, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(292, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(293, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(293, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) ParentFont() bool {
	r1 := LCL().SysCallN(294, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TButton) SetParentFont(AValue bool) {
	LCL().SysCallN(294, 1, m.Instance(), PascalBool(AValue))
}

func (m *TButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(295, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(295, 1, m.Instance(), PascalBool(AValue))
}

func ButtonClass() TClass {
	ret := LCL().SysCallN(289)
	return TClass(ret)
}

func (m *TButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(296, m.Instance(), m.contextPopupPtr)
}

func (m *TButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(297, m.Instance(), m.dragDropPtr)
}

func (m *TButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(298, m.Instance(), m.dragOverPtr)
}

func (m *TButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(299, m.Instance(), m.endDragPtr)
}

func (m *TButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(300, m.Instance(), m.mouseDownPtr)
}

func (m *TButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(301, m.Instance(), m.mouseEnterPtr)
}

func (m *TButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(302, m.Instance(), m.mouseLeavePtr)
}

func (m *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(303, m.Instance(), m.mouseMovePtr)
}

func (m *TButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(304, m.Instance(), m.mouseUpPtr)
}

func (m *TButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(305, m.Instance(), m.mouseWheelPtr)
}

func (m *TButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(306, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(307, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(308, m.Instance(), m.startDragPtr)
}
