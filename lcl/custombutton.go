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

// ICustomButton Parent: IButtonControl
type ICustomButton interface {
	IButtonControl
	Active() bool                       // property
	Default() bool                      // property
	SetDefault(AValue bool)             // property
	ModalResult() TModalResult          // property
	SetModalResult(AValue TModalResult) // property
	ShortCut() TShortCut                // property
	ShortCutKey2() TShortCut            // property
	Cancel() bool                       // property
	SetCancel(AValue bool)              // property
	Click()                             // procedure
}

// TCustomButton Parent: TButtonControl
type TCustomButton struct {
	TButtonControl
}

func NewCustomButton(TheOwner IComponent) ICustomButton {
	r1 := LCL().SysCallN(1337, GetObjectUintptr(TheOwner))
	return AsCustomButton(r1)
}

func (m *TCustomButton) Active() bool {
	r1 := LCL().SysCallN(1333, m.Instance())
	return GoBool(r1)
}

func (m *TCustomButton) Default() bool {
	r1 := LCL().SysCallN(1338, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomButton) SetDefault(AValue bool) {
	LCL().SysCallN(1338, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomButton) ModalResult() TModalResult {
	r1 := LCL().SysCallN(1339, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomButton) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(1339, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomButton) ShortCut() TShortCut {
	r1 := LCL().SysCallN(1340, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomButton) ShortCutKey2() TShortCut {
	r1 := LCL().SysCallN(1341, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomButton) Cancel() bool {
	r1 := LCL().SysCallN(1334, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomButton) SetCancel(AValue bool) {
	LCL().SysCallN(1334, 1, m.Instance(), PascalBool(AValue))
}

func CustomButtonClass() TClass {
	ret := LCL().SysCallN(1335)
	return TClass(ret)
}

func (m *TCustomButton) Click() {
	LCL().SysCallN(1336, m.Instance())
}
