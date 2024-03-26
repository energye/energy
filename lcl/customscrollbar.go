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

// ICustomScrollBar Parent: IWinControl
type ICustomScrollBar interface {
	IWinControl
	Kind() TScrollBarKind                             // property
	SetKind(AValue TScrollBarKind)                    // property
	LargeChange() TScrollBarInc                       // property
	SetLargeChange(AValue TScrollBarInc)              // property
	Max() int32                                       // property
	SetMax(AValue int32)                              // property
	Min() int32                                       // property
	SetMin(AValue int32)                              // property
	PageSize() int32                                  // property
	SetPageSize(AValue int32)                         // property
	Position() int32                                  // property
	SetPosition(AValue int32)                         // property
	SmallChange() TScrollBarInc                       // property
	SetSmallChange(AValue TScrollBarInc)              // property
	SetParams(APosition, AMin, AMax, APageSize int32) // procedure
	SetParams1(APosition, AMin, AMax int32)           // procedure
	SetOnChange(fn TNotifyEvent)                      // property event
	SetOnScroll(fn TScrollEvent)                      // property event
}

// TCustomScrollBar Parent: TWinControl
type TCustomScrollBar struct {
	TWinControl
	changePtr uintptr
	scrollPtr uintptr
}

func NewCustomScrollBar(AOwner IComponent) ICustomScrollBar {
	r1 := LCL().SysCallN(2014, GetObjectUintptr(AOwner))
	return AsCustomScrollBar(r1)
}

func (m *TCustomScrollBar) Kind() TScrollBarKind {
	r1 := LCL().SysCallN(2015, 0, m.Instance(), 0)
	return TScrollBarKind(r1)
}

func (m *TCustomScrollBar) SetKind(AValue TScrollBarKind) {
	LCL().SysCallN(2015, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) LargeChange() TScrollBarInc {
	r1 := LCL().SysCallN(2016, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TCustomScrollBar) SetLargeChange(AValue TScrollBarInc) {
	LCL().SysCallN(2016, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Max() int32 {
	r1 := LCL().SysCallN(2017, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetMax(AValue int32) {
	LCL().SysCallN(2017, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Min() int32 {
	r1 := LCL().SysCallN(2018, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetMin(AValue int32) {
	LCL().SysCallN(2018, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) PageSize() int32 {
	r1 := LCL().SysCallN(2019, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetPageSize(AValue int32) {
	LCL().SysCallN(2019, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Position() int32 {
	r1 := LCL().SysCallN(2020, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetPosition(AValue int32) {
	LCL().SysCallN(2020, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) SmallChange() TScrollBarInc {
	r1 := LCL().SysCallN(2025, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TCustomScrollBar) SetSmallChange(AValue TScrollBarInc) {
	LCL().SysCallN(2025, 1, m.Instance(), uintptr(AValue))
}

func CustomScrollBarClass() TClass {
	ret := LCL().SysCallN(2013)
	return TClass(ret)
}

func (m *TCustomScrollBar) SetParams(APosition, AMin, AMax, APageSize int32) {
	LCL().SysCallN(2023, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax), uintptr(APageSize))
}

func (m *TCustomScrollBar) SetParams1(APosition, AMin, AMax int32) {
	LCL().SysCallN(2024, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func (m *TCustomScrollBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2021, m.Instance(), m.changePtr)
}

func (m *TCustomScrollBar) SetOnScroll(fn TScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2022, m.Instance(), m.scrollPtr)
}
