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

// IDragImageList Parent: ICustomImageList
type IDragImageList interface {
	ICustomImageList
	DragCursor() TCursor                                                             // property
	SetDragCursor(AValue TCursor)                                                    // property
	DragHotspot() (resultPoint TPoint)                                               // property
	SetDragHotspot(AValue *TPoint)                                                   // property
	Dragging() bool                                                                  // property
	DraggingResolution() IDragImageListResolution                                    // property
	ResolutionForDragImageListResolution(AImageWidth int32) IDragImageListResolution // property
	BeginDrag(Window HWND, X, Y int32) bool                                          // function
	DragLock(Window HWND, XPos, YPos int32) bool                                     // function
	DragMove(X, Y int32) bool                                                        // function
	EndDrag() bool                                                                   // function
	SetDragImage(Index, HotSpotX, HotSpotY int32) bool                               // function
	DragUnlock()                                                                     // procedure
	HideDragImage()                                                                  // procedure
	ShowDragImage()                                                                  // procedure
}

// TDragImageList Parent: TCustomImageList
type TDragImageList struct {
	TCustomImageList
}

func NewDragImageList(AOwner IComponent) IDragImageList {
	r1 := LCL().SysCallN(2459, GetObjectUintptr(AOwner))
	return AsDragImageList(r1)
}

func (m *TDragImageList) DragCursor() TCursor {
	r1 := LCL().SysCallN(2460, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDragImageList) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2460, 1, m.Instance(), uintptr(AValue))
}

func (m *TDragImageList) DragHotspot() (resultPoint TPoint) {
	LCL().SysCallN(2461, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDragImageList) SetDragHotspot(AValue *TPoint) {
	LCL().SysCallN(2461, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDragImageList) Dragging() bool {
	r1 := LCL().SysCallN(2465, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) DraggingResolution() IDragImageListResolution {
	r1 := LCL().SysCallN(2466, m.Instance())
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) ResolutionForDragImageListResolution(AImageWidth int32) IDragImageListResolution {
	r1 := LCL().SysCallN(2469, m.Instance(), uintptr(AImageWidth))
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) BeginDrag(Window HWND, X, Y int32) bool {
	r1 := LCL().SysCallN(2457, m.Instance(), uintptr(Window), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) DragLock(Window HWND, XPos, YPos int32) bool {
	r1 := LCL().SysCallN(2462, m.Instance(), uintptr(Window), uintptr(XPos), uintptr(YPos))
	return GoBool(r1)
}

func (m *TDragImageList) DragMove(X, Y int32) bool {
	r1 := LCL().SysCallN(2463, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) EndDrag() bool {
	r1 := LCL().SysCallN(2467, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) SetDragImage(Index, HotSpotX, HotSpotY int32) bool {
	r1 := LCL().SysCallN(2470, m.Instance(), uintptr(Index), uintptr(HotSpotX), uintptr(HotSpotY))
	return GoBool(r1)
}

func DragImageListClass() TClass {
	ret := LCL().SysCallN(2458)
	return TClass(ret)
}

func (m *TDragImageList) DragUnlock() {
	LCL().SysCallN(2464, m.Instance())
}

func (m *TDragImageList) HideDragImage() {
	LCL().SysCallN(2468, m.Instance())
}

func (m *TDragImageList) ShowDragImage() {
	LCL().SysCallN(2471, m.Instance())
}
