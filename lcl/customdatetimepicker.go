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
	r1 := LCL().SysCallN(1306, GetObjectUintptr(AOwner))
	return AsCustomDateTimePicker(r1)
}

func (m *TCustomDateTimePicker) DateIsNull() bool {
	r1 := LCL().SysCallN(1307, m.Instance())
	return GoBool(r1)
}

func CustomDateTimePickerClass() TClass {
	ret := LCL().SysCallN(1305)
	return TClass(ret)
}

func (m *TCustomDateTimePicker) SelectDate() {
	LCL().SysCallN(1309, m.Instance())
}

func (m *TCustomDateTimePicker) SelectTime() {
	LCL().SysCallN(1310, m.Instance())
}

func (m *TCustomDateTimePicker) SendExternalKey(aKey Char) {
	LCL().SysCallN(1311, m.Instance(), uintptr(aKey))
}

func (m *TCustomDateTimePicker) SendExternalKeyCode(Key Word) {
	LCL().SysCallN(1312, m.Instance(), uintptr(Key))
}

func (m *TCustomDateTimePicker) Paint() {
	LCL().SysCallN(1308, m.Instance())
}
