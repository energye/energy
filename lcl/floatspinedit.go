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

// IFloatSpinEdit Parent: ICustomFloatSpinEdit
type IFloatSpinEdit interface {
	ICustomFloatSpinEdit
	AutoSelected() bool                             // property
	SetAutoSelected(AValue bool)                    // property
	AutoSelect() bool                               // property
	SetAutoSelect(AValue bool)                      // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnEditingDone(fn TNotifyEvent)               // property event
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
}

// TFloatSpinEdit Parent: TCustomFloatSpinEdit
type TFloatSpinEdit struct {
	TCustomFloatSpinEdit
	editingDonePtr     uintptr
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
}

func NewFloatSpinEdit(TheOwner IComponent) IFloatSpinEdit {
	r1 := LCL().SysCallN(3021, GetObjectUintptr(TheOwner))
	return AsFloatSpinEdit(r1)
}

func (m *TFloatSpinEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(3019, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(3019, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(3018, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(3018, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentColor() bool {
	r1 := LCL().SysCallN(3022, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(3022, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentFont() bool {
	r1 := LCL().SysCallN(3023, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(3023, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(3024, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3024, 1, m.Instance(), PascalBool(AValue))
}

func FloatSpinEditClass() TClass {
	ret := LCL().SysCallN(3020)
	return TClass(ret)
}

func (m *TFloatSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3025, m.Instance(), m.editingDonePtr)
}

func (m *TFloatSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3026, m.Instance(), m.mouseDownPtr)
}

func (m *TFloatSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3027, m.Instance(), m.mouseEnterPtr)
}

func (m *TFloatSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3028, m.Instance(), m.mouseLeavePtr)
}

func (m *TFloatSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3029, m.Instance(), m.mouseMovePtr)
}

func (m *TFloatSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3030, m.Instance(), m.mouseUpPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3031, m.Instance(), m.mouseWheelPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3032, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3036, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3033, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3034, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3035, m.Instance(), m.mouseWheelRightPtr)
}
