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

// ICustomCheckGroup Parent: ICustomGroupBox
type ICustomCheckGroup interface {
	ICustomGroupBox
	AutoFill() bool                           // property
	SetAutoFill(AValue bool)                  // property
	Items() IStrings                          // property
	SetItems(AValue IStrings)                 // property
	Checked(Index int32) bool                 // property
	SetChecked(Index int32, AValue bool)      // property
	CheckEnabled(Index int32) bool            // property
	SetCheckEnabled(Index int32, AValue bool) // property
	Columns() int32                           // property
	SetColumns(AValue int32)                  // property
	ColumnLayout() TColumnLayout              // property
	SetColumnLayout(AValue TColumnLayout)     // property
	Rows() int32                              // function
}

// TCustomCheckGroup Parent: TCustomGroupBox
type TCustomCheckGroup struct {
	TCustomGroupBox
}

func NewCustomCheckGroup(TheOwner IComponent) ICustomCheckGroup {
	r1 := LCL().SysCallN(1384, GetObjectUintptr(TheOwner))
	return AsCustomCheckGroup(r1)
}

func (m *TCustomCheckGroup) AutoFill() bool {
	r1 := LCL().SysCallN(1378, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetAutoFill(AValue bool) {
	LCL().SysCallN(1378, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckGroup) Items() IStrings {
	r1 := LCL().SysCallN(1385, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomCheckGroup) SetItems(AValue IStrings) {
	LCL().SysCallN(1385, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCheckGroup) Checked(Index int32) bool {
	r1 := LCL().SysCallN(1380, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetChecked(Index int32, AValue bool) {
	LCL().SysCallN(1380, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomCheckGroup) CheckEnabled(Index int32) bool {
	r1 := LCL().SysCallN(1379, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetCheckEnabled(Index int32, AValue bool) {
	LCL().SysCallN(1379, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomCheckGroup) Columns() int32 {
	r1 := LCL().SysCallN(1383, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCheckGroup) SetColumns(AValue int32) {
	LCL().SysCallN(1383, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckGroup) ColumnLayout() TColumnLayout {
	r1 := LCL().SysCallN(1382, 0, m.Instance(), 0)
	return TColumnLayout(r1)
}

func (m *TCustomCheckGroup) SetColumnLayout(AValue TColumnLayout) {
	LCL().SysCallN(1382, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckGroup) Rows() int32 {
	r1 := LCL().SysCallN(1386, m.Instance())
	return int32(r1)
}

func CustomCheckGroupClass() TClass {
	ret := LCL().SysCallN(1381)
	return TClass(ret)
}
