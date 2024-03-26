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

// ISpinEdit Parent: ICustomSpinEdit
type ISpinEdit interface {
	ICustomSpinEdit
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

// TSpinEdit Parent: TCustomSpinEdit
type TSpinEdit struct {
	TCustomSpinEdit
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

func NewSpinEdit(TheOwner IComponent) ISpinEdit {
	r1 := LCL().SysCallN(4402, GetObjectUintptr(TheOwner))
	return AsSpinEdit(r1)
}

func (m *TSpinEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(4400, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(4400, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(4399, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(4399, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentColor() bool {
	r1 := LCL().SysCallN(4403, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(4403, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentFont() bool {
	r1 := LCL().SysCallN(4404, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(4404, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(4405, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4405, 1, m.Instance(), PascalBool(AValue))
}

func SpinEditClass() TClass {
	ret := LCL().SysCallN(4401)
	return TClass(ret)
}

func (m *TSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4406, m.Instance(), m.editingDonePtr)
}

func (m *TSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4407, m.Instance(), m.mouseDownPtr)
}

func (m *TSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4408, m.Instance(), m.mouseEnterPtr)
}

func (m *TSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4409, m.Instance(), m.mouseLeavePtr)
}

func (m *TSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4410, m.Instance(), m.mouseMovePtr)
}

func (m *TSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4411, m.Instance(), m.mouseUpPtr)
}

func (m *TSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4412, m.Instance(), m.mouseWheelPtr)
}

func (m *TSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4413, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4417, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4414, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4415, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4416, m.Instance(), m.mouseWheelRightPtr)
}
