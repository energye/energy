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

// ICustomTrackBar Parent: IWinControl
type ICustomTrackBar interface {
	IWinControl
	Frequency() int32                           // property
	SetFrequency(AValue int32)                  // property
	LineSize() int32                            // property
	SetLineSize(AValue int32)                   // property
	Max() int32                                 // property
	SetMax(AValue int32)                        // property
	Min() int32                                 // property
	SetMin(AValue int32)                        // property
	Orientation() TTrackBarOrientation          // property
	SetOrientation(AValue TTrackBarOrientation) // property
	PageSize() int32                            // property
	SetPageSize(AValue int32)                   // property
	Position() int32                            // property
	SetPosition(AValue int32)                   // property
	Reversed() bool                             // property
	SetReversed(AValue bool)                    // property
	ScalePos() TTrackBarScalePos                // property
	SetScalePos(AValue TTrackBarScalePos)       // property
	SelEnd() int32                              // property
	SetSelEnd(AValue int32)                     // property
	SelStart() int32                            // property
	SetSelStart(AValue int32)                   // property
	ShowSelRange() bool                         // property
	SetShowSelRange(AValue bool)                // property
	TickMarks() TTickMark                       // property
	SetTickMarks(AValue TTickMark)              // property
	TickStyle() TTickStyle                      // property
	SetTickStyle(AValue TTickStyle)             // property
	SetTick(Value int32)                        // procedure
	SetParams(APosition, AMin, AMax int32)      // procedure
	SetOnChange(fn TNotifyEvent)                // property event
}

// TCustomTrackBar Parent: TWinControl
type TCustomTrackBar struct {
	TWinControl
	changePtr uintptr
}

func NewCustomTrackBar(AOwner IComponent) ICustomTrackBar {
	r1 := LCL().SysCallN(2358, GetObjectUintptr(AOwner))
	return AsCustomTrackBar(r1)
}

func (m *TCustomTrackBar) Frequency() int32 {
	r1 := LCL().SysCallN(2359, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetFrequency(AValue int32) {
	LCL().SysCallN(2359, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) LineSize() int32 {
	r1 := LCL().SysCallN(2360, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetLineSize(AValue int32) {
	LCL().SysCallN(2360, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Max() int32 {
	r1 := LCL().SysCallN(2361, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMax(AValue int32) {
	LCL().SysCallN(2361, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Min() int32 {
	r1 := LCL().SysCallN(2362, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMin(AValue int32) {
	LCL().SysCallN(2362, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Orientation() TTrackBarOrientation {
	r1 := LCL().SysCallN(2363, 0, m.Instance(), 0)
	return TTrackBarOrientation(r1)
}

func (m *TCustomTrackBar) SetOrientation(AValue TTrackBarOrientation) {
	LCL().SysCallN(2363, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) PageSize() int32 {
	r1 := LCL().SysCallN(2364, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPageSize(AValue int32) {
	LCL().SysCallN(2364, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Position() int32 {
	r1 := LCL().SysCallN(2365, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPosition(AValue int32) {
	LCL().SysCallN(2365, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Reversed() bool {
	r1 := LCL().SysCallN(2366, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetReversed(AValue bool) {
	LCL().SysCallN(2366, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) ScalePos() TTrackBarScalePos {
	r1 := LCL().SysCallN(2367, 0, m.Instance(), 0)
	return TTrackBarScalePos(r1)
}

func (m *TCustomTrackBar) SetScalePos(AValue TTrackBarScalePos) {
	LCL().SysCallN(2367, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelEnd() int32 {
	r1 := LCL().SysCallN(2368, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelEnd(AValue int32) {
	LCL().SysCallN(2368, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelStart() int32 {
	r1 := LCL().SysCallN(2369, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelStart(AValue int32) {
	LCL().SysCallN(2369, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) ShowSelRange() bool {
	r1 := LCL().SysCallN(2373, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetShowSelRange(AValue bool) {
	LCL().SysCallN(2373, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) TickMarks() TTickMark {
	r1 := LCL().SysCallN(2374, 0, m.Instance(), 0)
	return TTickMark(r1)
}

func (m *TCustomTrackBar) SetTickMarks(AValue TTickMark) {
	LCL().SysCallN(2374, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) TickStyle() TTickStyle {
	r1 := LCL().SysCallN(2375, 0, m.Instance(), 0)
	return TTickStyle(r1)
}

func (m *TCustomTrackBar) SetTickStyle(AValue TTickStyle) {
	LCL().SysCallN(2375, 1, m.Instance(), uintptr(AValue))
}

func CustomTrackBarClass() TClass {
	ret := LCL().SysCallN(2357)
	return TClass(ret)
}

func (m *TCustomTrackBar) SetTick(Value int32) {
	LCL().SysCallN(2372, m.Instance(), uintptr(Value))
}

func (m *TCustomTrackBar) SetParams(APosition, AMin, AMax int32) {
	LCL().SysCallN(2371, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func (m *TCustomTrackBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2370, m.Instance(), m.changePtr)
}
