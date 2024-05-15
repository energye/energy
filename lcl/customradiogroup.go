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

// ICustomRadioGroup Parent: ICustomGroupBox
type ICustomRadioGroup interface {
	ICustomGroupBox
	AutoFill() bool                        // property
	SetAutoFill(AValue bool)               // property
	ItemIndex() int32                      // property
	SetItemIndex(AValue int32)             // property
	Items() IStrings                       // property
	SetItems(AValue IStrings)              // property
	Columns() int32                        // property
	SetColumns(AValue int32)               // property
	ColumnLayout() TColumnLayout           // property
	SetColumnLayout(AValue TColumnLayout)  // property
	CanModify() bool                       // function
	Rows() int32                           // function
	SetOnItemEnter(fn TNotifyEvent)        // property event
	SetOnItemExit(fn TNotifyEvent)         // property event
	SetOnSelectionChanged(fn TNotifyEvent) // property event
}

// TCustomRadioGroup Parent: TCustomGroupBox
type TCustomRadioGroup struct {
	TCustomGroupBox
	itemEnterPtr        uintptr
	itemExitPtr         uintptr
	selectionChangedPtr uintptr
}

func NewCustomRadioGroup(TheOwner IComponent) ICustomRadioGroup {
	r1 := LCL().SysCallN(2156, GetObjectUintptr(TheOwner))
	return AsCustomRadioGroup(r1)
}

func (m *TCustomRadioGroup) AutoFill() bool {
	r1 := LCL().SysCallN(2151, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomRadioGroup) SetAutoFill(AValue bool) {
	LCL().SysCallN(2151, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomRadioGroup) ItemIndex() int32 {
	r1 := LCL().SysCallN(2157, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomRadioGroup) SetItemIndex(AValue int32) {
	LCL().SysCallN(2157, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) Items() IStrings {
	r1 := LCL().SysCallN(2158, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomRadioGroup) SetItems(AValue IStrings) {
	LCL().SysCallN(2158, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomRadioGroup) Columns() int32 {
	r1 := LCL().SysCallN(2155, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomRadioGroup) SetColumns(AValue int32) {
	LCL().SysCallN(2155, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) ColumnLayout() TColumnLayout {
	r1 := LCL().SysCallN(2154, 0, m.Instance(), 0)
	return TColumnLayout(r1)
}

func (m *TCustomRadioGroup) SetColumnLayout(AValue TColumnLayout) {
	LCL().SysCallN(2154, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) CanModify() bool {
	r1 := LCL().SysCallN(2152, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRadioGroup) Rows() int32 {
	r1 := LCL().SysCallN(2159, m.Instance())
	return int32(r1)
}

func CustomRadioGroupClass() TClass {
	ret := LCL().SysCallN(2153)
	return TClass(ret)
}

func (m *TCustomRadioGroup) SetOnItemEnter(fn TNotifyEvent) {
	if m.itemEnterPtr != 0 {
		RemoveEventElement(m.itemEnterPtr)
	}
	m.itemEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2160, m.Instance(), m.itemEnterPtr)
}

func (m *TCustomRadioGroup) SetOnItemExit(fn TNotifyEvent) {
	if m.itemExitPtr != 0 {
		RemoveEventElement(m.itemExitPtr)
	}
	m.itemExitPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2161, m.Instance(), m.itemExitPtr)
}

func (m *TCustomRadioGroup) SetOnSelectionChanged(fn TNotifyEvent) {
	if m.selectionChangedPtr != 0 {
		RemoveEventElement(m.selectionChangedPtr)
	}
	m.selectionChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2162, m.Instance(), m.selectionChangedPtr)
}
