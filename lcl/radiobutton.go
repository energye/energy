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

// IRadioButton Parent: ICustomCheckBox
type IRadioButton interface {
	ICustomCheckBox
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
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

// TRadioButton Parent: TCustomCheckBox
type TRadioButton struct {
	TCustomCheckBox
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

func NewRadioButton(TheOwner IComponent) IRadioButton {
	r1 := LCL().SysCallN(4670, GetObjectUintptr(TheOwner))
	return AsRadioButton(r1)
}

func (m *TRadioButton) Checked() bool {
	r1 := LCL().SysCallN(4668, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetChecked(AValue bool) {
	LCL().SysCallN(4668, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(4671, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRadioButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4671, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(4672, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TRadioButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4672, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(4673, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRadioButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4673, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) ParentColor() bool {
	r1 := LCL().SysCallN(4674, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentColor(AValue bool) {
	LCL().SysCallN(4674, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentFont() bool {
	r1 := LCL().SysCallN(4675, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentFont(AValue bool) {
	LCL().SysCallN(4675, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(4676, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4676, 1, m.Instance(), PascalBool(AValue))
}

func RadioButtonClass() TClass {
	ret := LCL().SysCallN(4669)
	return TClass(ret)
}

func (m *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4677, m.Instance(), m.contextPopupPtr)
}

func (m *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4678, m.Instance(), m.dragDropPtr)
}

func (m *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4679, m.Instance(), m.dragOverPtr)
}

func (m *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4680, m.Instance(), m.endDragPtr)
}

func (m *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4681, m.Instance(), m.mouseDownPtr)
}

func (m *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4682, m.Instance(), m.mouseEnterPtr)
}

func (m *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4683, m.Instance(), m.mouseLeavePtr)
}

func (m *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4684, m.Instance(), m.mouseMovePtr)
}

func (m *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4685, m.Instance(), m.mouseUpPtr)
}

func (m *TRadioButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4686, m.Instance(), m.mouseWheelPtr)
}

func (m *TRadioButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4687, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRadioButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4688, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRadioButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4689, m.Instance(), m.startDragPtr)
}
