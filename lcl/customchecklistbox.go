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

// ICustomCheckListBox Parent: ICustomListBox
type ICustomCheckListBox interface {
	ICustomListBox
	AllowGrayed() bool                                                      // property
	SetAllowGrayed(AValue bool)                                             // property
	Checked(AIndex int32) bool                                              // property
	SetChecked(AIndex int32, AValue bool)                                   // property
	Header(AIndex int32) bool                                               // property
	SetHeader(AIndex int32, AValue bool)                                    // property
	HeaderBackgroundColor() TColor                                          // property
	SetHeaderBackgroundColor(AValue TColor)                                 // property
	HeaderColor() TColor                                                    // property
	SetHeaderColor(AValue TColor)                                           // property
	ItemEnabled(AIndex int32) bool                                          // property
	SetItemEnabled(AIndex int32, AValue bool)                               // property
	State(AIndex int32) TCheckBoxState                                      // property
	SetState(AIndex int32, AValue TCheckBoxState)                           // property
	CalculateStandardItemHeight() int32                                     // function
	Toggle(AIndex int32)                                                    // procedure
	CheckAll(AState TCheckBoxState, aAllowGrayed bool, aAllowDisabled bool) // procedure
	Exchange(AIndex1, AIndex2 int32)                                        // procedure
	SetOnClickCheck(fn TNotifyEvent)                                        // property event
}

// TCustomCheckListBox Parent: TCustomListBox
type TCustomCheckListBox struct {
	TCustomListBox
	clickCheckPtr uintptr
}

func NewCustomCheckListBox(AOwner IComponent) ICustomCheckListBox {
	r1 := LCL().SysCallN(1202, GetObjectUintptr(AOwner))
	return AsCustomCheckListBox(r1)
}

func (m *TCustomCheckListBox) AllowGrayed() bool {
	r1 := LCL().SysCallN(1197, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetAllowGrayed(AValue bool) {
	LCL().SysCallN(1197, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckListBox) Checked(AIndex int32) bool {
	r1 := LCL().SysCallN(1200, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetChecked(AIndex int32, AValue bool) {
	LCL().SysCallN(1200, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) Header(AIndex int32) bool {
	r1 := LCL().SysCallN(1204, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetHeader(AIndex int32, AValue bool) {
	LCL().SysCallN(1204, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) HeaderBackgroundColor() TColor {
	r1 := LCL().SysCallN(1205, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomCheckListBox) SetHeaderBackgroundColor(AValue TColor) {
	LCL().SysCallN(1205, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckListBox) HeaderColor() TColor {
	r1 := LCL().SysCallN(1206, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomCheckListBox) SetHeaderColor(AValue TColor) {
	LCL().SysCallN(1206, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckListBox) ItemEnabled(AIndex int32) bool {
	r1 := LCL().SysCallN(1207, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetItemEnabled(AIndex int32, AValue bool) {
	LCL().SysCallN(1207, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) State(AIndex int32) TCheckBoxState {
	r1 := LCL().SysCallN(1209, 0, m.Instance(), uintptr(AIndex))
	return TCheckBoxState(r1)
}

func (m *TCustomCheckListBox) SetState(AIndex int32, AValue TCheckBoxState) {
	LCL().SysCallN(1209, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func (m *TCustomCheckListBox) CalculateStandardItemHeight() int32 {
	r1 := LCL().SysCallN(1198, m.Instance())
	return int32(r1)
}

func CustomCheckListBoxClass() TClass {
	ret := LCL().SysCallN(1201)
	return TClass(ret)
}

func (m *TCustomCheckListBox) Toggle(AIndex int32) {
	LCL().SysCallN(1210, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckListBox) CheckAll(AState TCheckBoxState, aAllowGrayed bool, aAllowDisabled bool) {
	LCL().SysCallN(1199, m.Instance(), uintptr(AState), PascalBool(aAllowGrayed), PascalBool(aAllowDisabled))
}

func (m *TCustomCheckListBox) Exchange(AIndex1, AIndex2 int32) {
	LCL().SysCallN(1203, m.Instance(), uintptr(AIndex1), uintptr(AIndex2))
}

func (m *TCustomCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
	if m.clickCheckPtr != 0 {
		RemoveEventElement(m.clickCheckPtr)
	}
	m.clickCheckPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1208, m.Instance(), m.clickCheckPtr)
}
