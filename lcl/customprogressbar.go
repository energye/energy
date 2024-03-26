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

// ICustomProgressBar Parent: IWinControl
type ICustomProgressBar interface {
	IWinControl
	Max() int32                                    // property
	SetMax(AValue int32)                           // property
	Min() int32                                    // property
	SetMin(AValue int32)                           // property
	Orientation() TProgressBarOrientation          // property
	SetOrientation(AValue TProgressBarOrientation) // property
	Position() int32                               // property
	SetPosition(AValue int32)                      // property
	Smooth() bool                                  // property
	SetSmooth(AValue bool)                         // property
	Step() int32                                   // property
	SetStep(AValue int32)                          // property
	Style() TProgressBarStyle                      // property
	SetStyle(AValue TProgressBarStyle)             // property
	BarShowText() bool                             // property
	SetBarShowText(AValue bool)                    // property
	StepIt()                                       // procedure
	StepBy(Delta int32)                            // procedure
}

// TCustomProgressBar Parent: TWinControl
type TCustomProgressBar struct {
	TWinControl
}

func NewCustomProgressBar(AOwner IComponent) ICustomProgressBar {
	r1 := LCL().SysCallN(1951, GetObjectUintptr(AOwner))
	return AsCustomProgressBar(r1)
}

func (m *TCustomProgressBar) Max() int32 {
	r1 := LCL().SysCallN(1952, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetMax(AValue int32) {
	LCL().SysCallN(1952, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Min() int32 {
	r1 := LCL().SysCallN(1953, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetMin(AValue int32) {
	LCL().SysCallN(1953, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Orientation() TProgressBarOrientation {
	r1 := LCL().SysCallN(1954, 0, m.Instance(), 0)
	return TProgressBarOrientation(r1)
}

func (m *TCustomProgressBar) SetOrientation(AValue TProgressBarOrientation) {
	LCL().SysCallN(1954, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Position() int32 {
	r1 := LCL().SysCallN(1955, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetPosition(AValue int32) {
	LCL().SysCallN(1955, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Smooth() bool {
	r1 := LCL().SysCallN(1956, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomProgressBar) SetSmooth(AValue bool) {
	LCL().SysCallN(1956, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomProgressBar) Step() int32 {
	r1 := LCL().SysCallN(1957, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetStep(AValue int32) {
	LCL().SysCallN(1957, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Style() TProgressBarStyle {
	r1 := LCL().SysCallN(1960, 0, m.Instance(), 0)
	return TProgressBarStyle(r1)
}

func (m *TCustomProgressBar) SetStyle(AValue TProgressBarStyle) {
	LCL().SysCallN(1960, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) BarShowText() bool {
	r1 := LCL().SysCallN(1949, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomProgressBar) SetBarShowText(AValue bool) {
	LCL().SysCallN(1949, 1, m.Instance(), PascalBool(AValue))
}

func CustomProgressBarClass() TClass {
	ret := LCL().SysCallN(1950)
	return TClass(ret)
}

func (m *TCustomProgressBar) StepIt() {
	LCL().SysCallN(1959, m.Instance())
}

func (m *TCustomProgressBar) StepBy(Delta int32) {
	LCL().SysCallN(1958, m.Instance(), uintptr(Delta))
}
