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

// IBevel Parent: IGraphicControl
type IBevel interface {
	IGraphicControl
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Shape() TBevelShape                            // property
	SetShape(AValue TBevelShape)                   // property
	Style() TBevelStyle                            // property
	SetStyle(AValue TBevelStyle)                   // property
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

// TBevel Parent: TGraphicControl
type TBevel struct {
	TGraphicControl
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	paintPtr          uintptr
}

func NewBevel(AOwner IComponent) IBevel {
	r1 := LCL().SysCallN(410, GetObjectUintptr(AOwner))
	return AsBevel(r1)
}

func (m *TBevel) ParentShowHint() bool {
	r1 := LCL().SysCallN(411, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBevel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(411, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBevel) Shape() TBevelShape {
	r1 := LCL().SysCallN(421, 0, m.Instance(), 0)
	return TBevelShape(r1)
}

func (m *TBevel) SetShape(AValue TBevelShape) {
	LCL().SysCallN(421, 1, m.Instance(), uintptr(AValue))
}

func (m *TBevel) Style() TBevelStyle {
	r1 := LCL().SysCallN(422, 0, m.Instance(), 0)
	return TBevelStyle(r1)
}

func (m *TBevel) SetStyle(AValue TBevelStyle) {
	LCL().SysCallN(422, 1, m.Instance(), uintptr(AValue))
}

func BevelClass() TClass {
	ret := LCL().SysCallN(409)
	return TClass(ret)
}

func (m *TBevel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(412, m.Instance(), m.mouseDownPtr)
}

func (m *TBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(413, m.Instance(), m.mouseEnterPtr)
}

func (m *TBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(414, m.Instance(), m.mouseLeavePtr)
}

func (m *TBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(415, m.Instance(), m.mouseMovePtr)
}

func (m *TBevel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(416, m.Instance(), m.mouseUpPtr)
}

func (m *TBevel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(417, m.Instance(), m.mouseWheelPtr)
}

func (m *TBevel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(418, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBevel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(419, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBevel) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(420, m.Instance(), m.paintPtr)
}
