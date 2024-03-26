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

// ICustomSpinEdit Parent: ICustomFloatSpinEdit
type ICustomSpinEdit interface {
	ICustomFloatSpinEdit
	ValueForInteger() int32              // property
	SetValueForInteger(AValue int32)     // property
	MinValueForInteger() int32           // property
	SetMinValueForInteger(AValue int32)  // property
	MaxValueForInteger() int32           // property
	SetMaxValueForInteger(AValue int32)  // property
	IncrementForInteger() int32          // property
	SetIncrementForInteger(AValue int32) // property
}

// TCustomSpinEdit Parent: TCustomFloatSpinEdit
type TCustomSpinEdit struct {
	TCustomFloatSpinEdit
}

func NewCustomSpinEdit(TheOwner IComponent) ICustomSpinEdit {
	r1 := LCL().SysCallN(2053, GetObjectUintptr(TheOwner))
	return AsCustomSpinEdit(r1)
}

func (m *TCustomSpinEdit) ValueForInteger() int32 {
	r1 := LCL().SysCallN(2057, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetValueForInteger(AValue int32) {
	LCL().SysCallN(2057, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MinValueForInteger() int32 {
	r1 := LCL().SysCallN(2056, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMinValueForInteger(AValue int32) {
	LCL().SysCallN(2056, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MaxValueForInteger() int32 {
	r1 := LCL().SysCallN(2055, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMaxValueForInteger(AValue int32) {
	LCL().SysCallN(2055, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) IncrementForInteger() int32 {
	r1 := LCL().SysCallN(2054, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetIncrementForInteger(AValue int32) {
	LCL().SysCallN(2054, 1, m.Instance(), uintptr(AValue))
}

func CustomSpinEditClass() TClass {
	ret := LCL().SysCallN(2052)
	return TClass(ret)
}
