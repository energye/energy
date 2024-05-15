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

// ICustomCheckBox Parent: IButtonControl
type ICustomCheckBox interface {
	IButtonControl
	Alignment() TLeftRight          // property
	SetAlignment(AValue TLeftRight) // property
	AllowGrayed() bool              // property
	SetAllowGrayed(AValue bool)     // property
	State() TCheckBoxState          // property
	SetState(AValue TCheckBoxState) // property
	ShortCut() TShortCut            // property
	ShortCutKey2() TShortCut        // property
	SetOnChange(fn TNotifyEvent)    // property event
}

// TCustomCheckBox Parent: TButtonControl
type TCustomCheckBox struct {
	TButtonControl
	changePtr uintptr
}

func NewCustomCheckBox(TheOwner IComponent) ICustomCheckBox {
	r1 := LCL().SysCallN(1359, GetObjectUintptr(TheOwner))
	return AsCustomCheckBox(r1)
}

func (m *TCustomCheckBox) Alignment() TLeftRight {
	r1 := LCL().SysCallN(1356, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TCustomCheckBox) SetAlignment(AValue TLeftRight) {
	LCL().SysCallN(1356, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) AllowGrayed() bool {
	r1 := LCL().SysCallN(1357, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckBox) SetAllowGrayed(AValue bool) {
	LCL().SysCallN(1357, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckBox) State() TCheckBoxState {
	r1 := LCL().SysCallN(1363, 0, m.Instance(), 0)
	return TCheckBoxState(r1)
}

func (m *TCustomCheckBox) SetState(AValue TCheckBoxState) {
	LCL().SysCallN(1363, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) ShortCut() TShortCut {
	r1 := LCL().SysCallN(1361, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomCheckBox) ShortCutKey2() TShortCut {
	r1 := LCL().SysCallN(1362, m.Instance())
	return TShortCut(r1)
}

func CustomCheckBoxClass() TClass {
	ret := LCL().SysCallN(1358)
	return TClass(ret)
}

func (m *TCustomCheckBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1360, m.Instance(), m.changePtr)
}
