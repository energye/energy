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
	r1 := LCL().SysCallN(1169, GetObjectUintptr(TheOwner))
	return AsCustomCheckBox(r1)
}

func (m *TCustomCheckBox) Alignment() TLeftRight {
	r1 := LCL().SysCallN(1166, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TCustomCheckBox) SetAlignment(AValue TLeftRight) {
	LCL().SysCallN(1166, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) AllowGrayed() bool {
	r1 := LCL().SysCallN(1167, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckBox) SetAllowGrayed(AValue bool) {
	LCL().SysCallN(1167, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckBox) State() TCheckBoxState {
	r1 := LCL().SysCallN(1173, 0, m.Instance(), 0)
	return TCheckBoxState(r1)
}

func (m *TCustomCheckBox) SetState(AValue TCheckBoxState) {
	LCL().SysCallN(1173, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) ShortCut() TShortCut {
	r1 := LCL().SysCallN(1171, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomCheckBox) ShortCutKey2() TShortCut {
	r1 := LCL().SysCallN(1172, m.Instance())
	return TShortCut(r1)
}

func CustomCheckBoxClass() TClass {
	ret := LCL().SysCallN(1168)
	return TClass(ret)
}

func (m *TCustomCheckBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1170, m.Instance(), m.changePtr)
}
