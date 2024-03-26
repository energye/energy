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

// IDragImageListResolution Parent: ICustomImageListResolution
type IDragImageListResolution interface {
	ICustomImageListResolution
	DragHotspot() (resultPoint TPoint)           // property
	SetDragHotspot(AValue *TPoint)               // property
	Dragging() bool                              // property
	BeginDrag(Window HWND, X, Y int32) bool      // function
	DragLock(Window HWND, XPos, YPos int32) bool // function
	DragMove(X, Y int32) bool                    // function
	EndDrag() bool                               // function
	DragUnlock()                                 // procedure
	HideDragImage()                              // procedure
	ShowDragImage()                              // procedure
}

// TDragImageListResolution Parent: TCustomImageListResolution
type TDragImageListResolution struct {
	TCustomImageListResolution
}

func NewDragImageListResolution(TheOwner IComponent) IDragImageListResolution {
	r1 := LCL().SysCallN(2448, GetObjectUintptr(TheOwner))
	return AsDragImageListResolution(r1)
}

func (m *TDragImageListResolution) DragHotspot() (resultPoint TPoint) {
	LCL().SysCallN(2449, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDragImageListResolution) SetDragHotspot(AValue *TPoint) {
	LCL().SysCallN(2449, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragImageListResolution) Dragging() bool {
	r1 := LCL().SysCallN(2453, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageListResolution) BeginDrag(Window HWND, X, Y int32) bool {
	r1 := LCL().SysCallN(2446, m.Instance(), uintptr(Window), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageListResolution) DragLock(Window HWND, XPos, YPos int32) bool {
	r1 := LCL().SysCallN(2450, m.Instance(), uintptr(Window), uintptr(XPos), uintptr(YPos))
	return GoBool(r1)
}

func (m *TDragImageListResolution) DragMove(X, Y int32) bool {
	r1 := LCL().SysCallN(2451, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageListResolution) EndDrag() bool {
	r1 := LCL().SysCallN(2454, m.Instance())
	return GoBool(r1)
}

func DragImageListResolutionClass() TClass {
	ret := LCL().SysCallN(2447)
	return TClass(ret)
}

func (m *TDragImageListResolution) DragUnlock() {
	LCL().SysCallN(2452, m.Instance())
}

func (m *TDragImageListResolution) HideDragImage() {
	LCL().SysCallN(2455, m.Instance())
}

func (m *TDragImageListResolution) ShowDragImage() {
	LCL().SysCallN(2456, m.Instance())
}
