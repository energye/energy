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

// ICustomDateTimePicker Parent: ICustomControl
type ICustomDateTimePicker interface {
	ICustomControl
	DateIsNull() bool             // function
	SelectDate()                  // procedure
	SelectTime()                  // procedure
	SendExternalKey(aKey Char)    // procedure
	SendExternalKeyCode(Key Word) // procedure
	Paint()                       // procedure
}

// TCustomDateTimePicker Parent: TCustomControl
type TCustomDateTimePicker struct {
	TCustomControl
}

func NewCustomDateTimePicker(AOwner IComponent) ICustomDateTimePicker {
	r1 := LCL().SysCallN(1496, GetObjectUintptr(AOwner))
	return AsCustomDateTimePicker(r1)
}

func (m *TCustomDateTimePicker) DateIsNull() bool {
	r1 := LCL().SysCallN(1497, m.Instance())
	return GoBool(r1)
}

func CustomDateTimePickerClass() TClass {
	ret := LCL().SysCallN(1495)
	return TClass(ret)
}

func (m *TCustomDateTimePicker) SelectDate() {
	LCL().SysCallN(1499, m.Instance())
}

func (m *TCustomDateTimePicker) SelectTime() {
	LCL().SysCallN(1500, m.Instance())
}

func (m *TCustomDateTimePicker) SendExternalKey(aKey Char) {
	LCL().SysCallN(1501, m.Instance(), uintptr(aKey))
}

func (m *TCustomDateTimePicker) SendExternalKeyCode(Key Word) {
	LCL().SysCallN(1502, m.Instance(), uintptr(Key))
}

func (m *TCustomDateTimePicker) Paint() {
	LCL().SysCallN(1498, m.Instance())
}
