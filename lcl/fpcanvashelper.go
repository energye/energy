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

// IFPCanvasHelper Parent: IPersistent
type IFPCanvasHelper interface {
	IPersistent
	Allocated() bool                                          // property
	FixedCanvas() bool                                        // property
	Canvas() IFPCustomCanvas                                  // property
	FPColor() (resultFPColor TFPColor)                        // property
	SetFPColor(AValue *TFPColor)                              // property
	DelayAllocate() bool                                      // property
	SetDelayAllocate(AValue bool)                             // property
	AllocateResources(ACanvas IFPCustomCanvas, CanDelay bool) // procedure
	DeallocateResources()                                     // procedure
	SetOnChanging(fn TNotifyEvent)                            // property event
	SetOnChange(fn TNotifyEvent)                              // property event
}

// TFPCanvasHelper Parent: TPersistent
type TFPCanvasHelper struct {
	TPersistent
	changingPtr uintptr
	changePtr   uintptr
}

func NewFPCanvasHelper() IFPCanvasHelper {
	r1 := LCL().SysCallN(2581)
	return AsFPCanvasHelper(r1)
}

func (m *TFPCanvasHelper) Allocated() bool {
	r1 := LCL().SysCallN(2578, m.Instance())
	return GoBool(r1)
}

func (m *TFPCanvasHelper) FixedCanvas() bool {
	r1 := LCL().SysCallN(2585, m.Instance())
	return GoBool(r1)
}

func (m *TFPCanvasHelper) Canvas() IFPCustomCanvas {
	r1 := LCL().SysCallN(2579, m.Instance())
	return AsFPCustomCanvas(r1)
}

func (m *TFPCanvasHelper) FPColor() (resultFPColor TFPColor) {
	r1 := LCL().SysCallN(2584, 0, m.Instance(), 0)
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCanvasHelper) SetFPColor(AValue *TFPColor) {
	LCL().SysCallN(2584, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFPCanvasHelper) DelayAllocate() bool {
	r1 := LCL().SysCallN(2583, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCanvasHelper) SetDelayAllocate(AValue bool) {
	LCL().SysCallN(2583, 1, m.Instance(), PascalBool(AValue))
}

func FPCanvasHelperClass() TClass {
	ret := LCL().SysCallN(2580)
	return TClass(ret)
}

func (m *TFPCanvasHelper) AllocateResources(ACanvas IFPCustomCanvas, CanDelay bool) {
	LCL().SysCallN(2577, m.Instance(), GetObjectUintptr(ACanvas), PascalBool(CanDelay))
}

func (m *TFPCanvasHelper) DeallocateResources() {
	LCL().SysCallN(2582, m.Instance())
}

func (m *TFPCanvasHelper) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2587, m.Instance(), m.changingPtr)
}

func (m *TFPCanvasHelper) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2586, m.Instance(), m.changePtr)
}
