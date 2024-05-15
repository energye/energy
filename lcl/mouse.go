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

// IMouse Parent: IObject
type IMouse interface {
	IObject
	Capture() HWND                   // property
	SetCapture(AValue HWND)          // property
	CursorPos() (resultPoint TPoint) // property
	SetCursorPos(AValue *TPoint)     // property
	IsDragging() bool                // property
	WheelScrollLines() int32         // property
	DragImmediate() bool             // property
	SetDragImmediate(AValue bool)    // property
	DragThreshold() int32            // property
	SetDragThreshold(AValue int32)   // property
}

// TMouse Parent: TObject
type TMouse struct {
	TObject
}

func NewMouse() IMouse {
	r1 := LCL().SysCallN(4352)
	return AsMouse(r1)
}

func (m *TMouse) Capture() HWND {
	r1 := LCL().SysCallN(4350, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMouse) SetCapture(AValue HWND) {
	LCL().SysCallN(4350, 1, m.Instance(), uintptr(AValue))
}

func (m *TMouse) CursorPos() (resultPoint TPoint) {
	LCL().SysCallN(4353, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TMouse) SetCursorPos(AValue *TPoint) {
	LCL().SysCallN(4353, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TMouse) IsDragging() bool {
	r1 := LCL().SysCallN(4356, m.Instance())
	return GoBool(r1)
}

func (m *TMouse) WheelScrollLines() int32 {
	r1 := LCL().SysCallN(4357, m.Instance())
	return int32(r1)
}

func (m *TMouse) DragImmediate() bool {
	r1 := LCL().SysCallN(4354, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMouse) SetDragImmediate(AValue bool) {
	LCL().SysCallN(4354, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMouse) DragThreshold() int32 {
	r1 := LCL().SysCallN(4355, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMouse) SetDragThreshold(AValue int32) {
	LCL().SysCallN(4355, 1, m.Instance(), uintptr(AValue))
}

func MouseClass() TClass {
	ret := LCL().SysCallN(4351)
	return TClass(ret)
}
