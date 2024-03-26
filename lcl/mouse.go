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
	"unsafe"
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
	r1 := LCL().SysCallN(3710)
	return AsMouse(r1)
}

func (m *TMouse) Capture() HWND {
	r1 := LCL().SysCallN(3708, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMouse) SetCapture(AValue HWND) {
	LCL().SysCallN(3708, 1, m.Instance(), uintptr(AValue))
}

func (m *TMouse) CursorPos() (resultPoint TPoint) {
	LCL().SysCallN(3711, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TMouse) SetCursorPos(AValue *TPoint) {
	LCL().SysCallN(3711, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TMouse) IsDragging() bool {
	r1 := LCL().SysCallN(3714, m.Instance())
	return GoBool(r1)
}

func (m *TMouse) WheelScrollLines() int32 {
	r1 := LCL().SysCallN(3715, m.Instance())
	return int32(r1)
}

func (m *TMouse) DragImmediate() bool {
	r1 := LCL().SysCallN(3712, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMouse) SetDragImmediate(AValue bool) {
	LCL().SysCallN(3712, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMouse) DragThreshold() int32 {
	r1 := LCL().SysCallN(3713, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMouse) SetDragThreshold(AValue int32) {
	LCL().SysCallN(3713, 1, m.Instance(), uintptr(AValue))
}

func MouseClass() TClass {
	ret := LCL().SysCallN(3709)
	return TClass(ret)
}
