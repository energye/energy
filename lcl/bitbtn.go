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

// IBitBtn Parent: ICustomBitBtn
type IBitBtn interface {
	ICustomBitBtn
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

// TBitBtn Parent: TCustomBitBtn
type TBitBtn struct {
	TCustomBitBtn
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

func NewBitBtn(TheOwner IComponent) IBitBtn {
	r1 := LCL().SysCallN(424, GetObjectUintptr(TheOwner))
	return AsBitBtn(r1)
}

func (m *TBitBtn) DragCursor() TCursor {
	r1 := LCL().SysCallN(425, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TBitBtn) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(425, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) DragKind() TDragKind {
	r1 := LCL().SysCallN(426, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TBitBtn) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(426, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) DragMode() TDragMode {
	r1 := LCL().SysCallN(427, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TBitBtn) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(427, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) ParentFont() bool {
	r1 := LCL().SysCallN(428, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBitBtn) SetParentFont(AValue bool) {
	LCL().SysCallN(428, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBitBtn) ParentShowHint() bool {
	r1 := LCL().SysCallN(429, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBitBtn) SetParentShowHint(AValue bool) {
	LCL().SysCallN(429, 1, m.Instance(), PascalBool(AValue))
}

func BitBtnClass() TClass {
	ret := LCL().SysCallN(423)
	return TClass(ret)
}

func (m *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(430, m.Instance(), m.contextPopupPtr)
}

func (m *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(431, m.Instance(), m.dragDropPtr)
}

func (m *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(432, m.Instance(), m.dragOverPtr)
}

func (m *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(433, m.Instance(), m.endDragPtr)
}

func (m *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(434, m.Instance(), m.mouseDownPtr)
}

func (m *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(435, m.Instance(), m.mouseEnterPtr)
}

func (m *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(436, m.Instance(), m.mouseLeavePtr)
}

func (m *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(437, m.Instance(), m.mouseMovePtr)
}

func (m *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(438, m.Instance(), m.mouseUpPtr)
}

func (m *TBitBtn) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(439, m.Instance(), m.mouseWheelPtr)
}

func (m *TBitBtn) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(440, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBitBtn) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(441, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBitBtn) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(442, m.Instance(), m.startDragPtr)
}
