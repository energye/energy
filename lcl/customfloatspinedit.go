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

// ICustomFloatSpinEdit Parent: ICustomEdit
type ICustomFloatSpinEdit interface {
	ICustomEdit
	DecimalPlaces() int32                                  // property
	SetDecimalPlaces(AValue int32)                         // property
	EditorEnabled() bool                                   // property
	SetEditorEnabled(AValue bool)                          // property
	Increment() (resultDouble float64)                     // property
	SetIncrement(AValue float64)                           // property
	MinValue() (resultDouble float64)                      // property
	SetMinValue(AValue float64)                            // property
	MaxValue() (resultDouble float64)                      // property
	SetMaxValue(AValue float64)                            // property
	Value() (resultDouble float64)                         // property
	SetValue(AValue float64)                               // property
	ValueEmpty() bool                                      // property
	SetValueEmpty(AValue bool)                             // property
	GetLimitedValue(AValue float64) (resultDouble float64) // function
	ValueToStr(AValue float64) string                      // function
	StrToValue(S string) (resultDouble float64)            // function
}

// TCustomFloatSpinEdit Parent: TCustomEdit
type TCustomFloatSpinEdit struct {
	TCustomEdit
}

func NewCustomFloatSpinEdit(TheOwner IComponent) ICustomFloatSpinEdit {
	r1 := LCL().SysCallN(1450, GetObjectUintptr(TheOwner))
	return AsCustomFloatSpinEdit(r1)
}

func (m *TCustomFloatSpinEdit) DecimalPlaces() int32 {
	r1 := LCL().SysCallN(1451, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomFloatSpinEdit) SetDecimalPlaces(AValue int32) {
	LCL().SysCallN(1451, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomFloatSpinEdit) EditorEnabled() bool {
	r1 := LCL().SysCallN(1452, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFloatSpinEdit) SetEditorEnabled(AValue bool) {
	LCL().SysCallN(1452, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFloatSpinEdit) Increment() (resultDouble float64) {
	LCL().SysCallN(1454, 0, m.Instance(), uintptr(unsafe.Pointer(&resultDouble)), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetIncrement(AValue float64) {
	LCL().SysCallN(1454, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) MinValue() (resultDouble float64) {
	LCL().SysCallN(1456, 0, m.Instance(), uintptr(unsafe.Pointer(&resultDouble)), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetMinValue(AValue float64) {
	LCL().SysCallN(1456, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) MaxValue() (resultDouble float64) {
	LCL().SysCallN(1455, 0, m.Instance(), uintptr(unsafe.Pointer(&resultDouble)), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetMaxValue(AValue float64) {
	LCL().SysCallN(1455, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) Value() (resultDouble float64) {
	LCL().SysCallN(1458, 0, m.Instance(), uintptr(unsafe.Pointer(&resultDouble)), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetValue(AValue float64) {
	LCL().SysCallN(1458, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) ValueEmpty() bool {
	r1 := LCL().SysCallN(1459, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFloatSpinEdit) SetValueEmpty(AValue bool) {
	LCL().SysCallN(1459, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFloatSpinEdit) GetLimitedValue(AValue float64) (resultDouble float64) {
	LCL().SysCallN(1453, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) ValueToStr(AValue float64) string {
	r1 := LCL().SysCallN(1460, m.Instance(), uintptr(unsafe.Pointer(&AValue)))
	return GoStr(r1)
}

func (m *TCustomFloatSpinEdit) StrToValue(S string) (resultDouble float64) {
	LCL().SysCallN(1457, m.Instance(), PascalStr(S), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func CustomFloatSpinEditClass() TClass {
	ret := LCL().SysCallN(1449)
	return TClass(ret)
}
