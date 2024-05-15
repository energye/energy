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

// ISpeedButton Parent: ICustomSpeedButton
type ISpeedButton interface {
	ICustomSpeedButton
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
	SetOnDblClick(fn TNotifyEvent)                 // property event
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
	SetOnPaint(fn TNotifyEvent)                    // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TSpeedButton Parent: TCustomSpeedButton
type TSpeedButton struct {
	TCustomSpeedButton
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
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
	paintPtr          uintptr
	startDragPtr      uintptr
}

func NewSpeedButton(AOwner IComponent) ISpeedButton {
	r1 := LCL().SysCallN(5027, GetObjectUintptr(AOwner))
	return AsSpeedButton(r1)
}

func (m *TSpeedButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(5028, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TSpeedButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5028, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(5029, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TSpeedButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5029, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(5030, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TSpeedButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5030, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) ParentFont() bool {
	r1 := LCL().SysCallN(5031, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpeedButton) SetParentFont(AValue bool) {
	LCL().SysCallN(5031, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpeedButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(5032, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpeedButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5032, 1, m.Instance(), PascalBool(AValue))
}

func SpeedButtonClass() TClass {
	ret := LCL().SysCallN(5026)
	return TClass(ret)
}

func (m *TSpeedButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5033, m.Instance(), m.contextPopupPtr)
}

func (m *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5034, m.Instance(), m.dblClickPtr)
}

func (m *TSpeedButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5035, m.Instance(), m.dragDropPtr)
}

func (m *TSpeedButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5036, m.Instance(), m.dragOverPtr)
}

func (m *TSpeedButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5037, m.Instance(), m.endDragPtr)
}

func (m *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5038, m.Instance(), m.mouseDownPtr)
}

func (m *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5039, m.Instance(), m.mouseEnterPtr)
}

func (m *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5040, m.Instance(), m.mouseLeavePtr)
}

func (m *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5041, m.Instance(), m.mouseMovePtr)
}

func (m *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5042, m.Instance(), m.mouseUpPtr)
}

func (m *TSpeedButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5043, m.Instance(), m.mouseWheelPtr)
}

func (m *TSpeedButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5044, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSpeedButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5045, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSpeedButton) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5046, m.Instance(), m.paintPtr)
}

func (m *TSpeedButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5047, m.Instance(), m.startDragPtr)
}
