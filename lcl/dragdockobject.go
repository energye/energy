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

// IDragDockObject Parent: IDragObject
type IDragDockObject interface {
	IDragObject
	DockOffset() (resultPoint TPoint)  // property
	SetDockOffset(AValue *TPoint)      // property
	DockRect() (resultRect TRect)      // property
	SetDockRect(AValue *TRect)         // property
	DropAlign() TAlign                 // property
	SetDropAlign(AValue TAlign)        // property
	DropOnControl() IControl           // property
	SetDropOnControl(AValue IControl)  // property
	Floating() bool                    // property
	SetFloating(AValue bool)           // property
	IncreaseDockArea() bool            // property
	EraseDockRect() (resultRect TRect) // property
	SetEraseDockRect(AValue *TRect)    // property
}

// TDragDockObject Parent: TDragObject
type TDragDockObject struct {
	TDragObject
}

func NewDragDockObject(AControl IControl) IDragDockObject {
	r1 := LCL().SysCallN(2438, GetObjectUintptr(AControl))
	return AsDragDockObject(r1)
}

func (m *TDragDockObject) DockOffset() (resultPoint TPoint) {
	LCL().SysCallN(2439, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDragDockObject) SetDockOffset(AValue *TPoint) {
	LCL().SysCallN(2439, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragDockObject) DockRect() (resultRect TRect) {
	LCL().SysCallN(2440, 0, m.Instance(), uintptr(unsafe.Pointer(&resultRect)), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TDragDockObject) SetDockRect(AValue *TRect) {
	LCL().SysCallN(2440, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragDockObject) DropAlign() TAlign {
	r1 := LCL().SysCallN(2441, 0, m.Instance(), 0)
	return TAlign(r1)
}

func (m *TDragDockObject) SetDropAlign(AValue TAlign) {
	LCL().SysCallN(2441, 1, m.Instance(), uintptr(AValue))
}

func (m *TDragDockObject) DropOnControl() IControl {
	r1 := LCL().SysCallN(2442, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragDockObject) SetDropOnControl(AValue IControl) {
	LCL().SysCallN(2442, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragDockObject) Floating() bool {
	r1 := LCL().SysCallN(2444, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDragDockObject) SetFloating(AValue bool) {
	LCL().SysCallN(2444, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDragDockObject) IncreaseDockArea() bool {
	r1 := LCL().SysCallN(2445, m.Instance())
	return GoBool(r1)
}

func (m *TDragDockObject) EraseDockRect() (resultRect TRect) {
	LCL().SysCallN(2443, 0, m.Instance(), uintptr(unsafe.Pointer(&resultRect)), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TDragDockObject) SetEraseDockRect(AValue *TRect) {
	LCL().SysCallN(2443, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func DragDockObjectClass() TClass {
	ret := LCL().SysCallN(2437)
	return TClass(ret)
}
