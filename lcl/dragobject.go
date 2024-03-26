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

// IDragObject Parent: IObject
type IDragObject interface {
	IObject
	AlwaysShowDragImages() bool          // property
	SetAlwaysShowDragImages(AValue bool) // property
	AutoCreated() bool                   // property
	AutoFree() bool                      // property
	Control() IControl                   // property
	SetControl(AValue IControl)          // property
	DragPos() (resultPoint TPoint)       // property
	SetDragPos(AValue *TPoint)           // property
	DragTarget() IControl                // property
	SetDragTarget(AValue IControl)       // property
	DragTargetPos() (resultPoint TPoint) // property
	SetDragTargetPos(AValue *TPoint)     // property
	Dropped() bool                       // property
	HideDragImage()                      // procedure
	ShowDragImage()                      // procedure
}

// TDragObject Parent: TObject
type TDragObject struct {
	TObject
}

func NewDragObject(AControl IControl) IDragObject {
	r1 := LCL().SysCallN(2478, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func NewDragObjectAuto(AControl IControl) IDragObject {
	r1 := LCL().SysCallN(2473, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func (m *TDragObject) AlwaysShowDragImages() bool {
	r1 := LCL().SysCallN(2472, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDragObject) SetAlwaysShowDragImages(AValue bool) {
	LCL().SysCallN(2472, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDragObject) AutoCreated() bool {
	r1 := LCL().SysCallN(2474, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) AutoFree() bool {
	r1 := LCL().SysCallN(2475, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) Control() IControl {
	r1 := LCL().SysCallN(2477, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetControl(AValue IControl) {
	LCL().SysCallN(2477, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragPos() (resultPoint TPoint) {
	LCL().SysCallN(2479, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragPos(AValue *TPoint) {
	LCL().SysCallN(2479, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragObject) DragTarget() IControl {
	r1 := LCL().SysCallN(2480, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetDragTarget(AValue IControl) {
	LCL().SysCallN(2480, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragTargetPos() (resultPoint TPoint) {
	LCL().SysCallN(2481, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragTargetPos(AValue *TPoint) {
	LCL().SysCallN(2481, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragObject) Dropped() bool {
	r1 := LCL().SysCallN(2482, m.Instance())
	return GoBool(r1)
}

func DragObjectClass() TClass {
	ret := LCL().SysCallN(2476)
	return TClass(ret)
}

func (m *TDragObject) HideDragImage() {
	LCL().SysCallN(2483, m.Instance())
}

func (m *TDragObject) ShowDragImage() {
	LCL().SysCallN(2484, m.Instance())
}
