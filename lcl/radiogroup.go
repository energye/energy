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

// IRadioGroup Parent: ICustomRadioGroup
type IRadioGroup interface {
	ICustomRadioGroup
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
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
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TRadioGroup Parent: TCustomRadioGroup
type TRadioGroup struct {
	TCustomRadioGroup
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
	startDragPtr      uintptr
}

func NewRadioGroup(TheOwner IComponent) IRadioGroup {
	r1 := LCL().SysCallN(4691, GetObjectUintptr(TheOwner))
	return AsRadioGroup(r1)
}

func (m *TRadioGroup) DragCursor() TCursor {
	r1 := LCL().SysCallN(4692, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRadioGroup) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4692, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioGroup) DragMode() TDragMode {
	r1 := LCL().SysCallN(4693, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRadioGroup) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4693, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioGroup) ParentFont() bool {
	r1 := LCL().SysCallN(4695, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentFont(AValue bool) {
	LCL().SysCallN(4695, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioGroup) ParentColor() bool {
	r1 := LCL().SysCallN(4694, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentColor(AValue bool) {
	LCL().SysCallN(4694, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioGroup) ParentShowHint() bool {
	r1 := LCL().SysCallN(4696, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4696, 1, m.Instance(), PascalBool(AValue))
}

func RadioGroupClass() TClass {
	ret := LCL().SysCallN(4690)
	return TClass(ret)
}

func (m *TRadioGroup) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4697, m.Instance(), m.dblClickPtr)
}

func (m *TRadioGroup) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4698, m.Instance(), m.dragDropPtr)
}

func (m *TRadioGroup) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4699, m.Instance(), m.dragOverPtr)
}

func (m *TRadioGroup) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4700, m.Instance(), m.endDragPtr)
}

func (m *TRadioGroup) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4701, m.Instance(), m.mouseDownPtr)
}

func (m *TRadioGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4702, m.Instance(), m.mouseEnterPtr)
}

func (m *TRadioGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4703, m.Instance(), m.mouseLeavePtr)
}

func (m *TRadioGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4704, m.Instance(), m.mouseMovePtr)
}

func (m *TRadioGroup) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4705, m.Instance(), m.mouseUpPtr)
}

func (m *TRadioGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4706, m.Instance(), m.mouseWheelPtr)
}

func (m *TRadioGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4707, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRadioGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4708, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRadioGroup) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4709, m.Instance(), m.startDragPtr)
}
