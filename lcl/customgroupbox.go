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

// ICustomGroupBox Parent: IWinControl
type ICustomGroupBox interface {
	IWinControl
	ParentBackground() bool          // property
	SetParentBackground(AValue bool) // property
}

// TCustomGroupBox Parent: TWinControl
type TCustomGroupBox struct {
	TWinControl
}

func NewCustomGroupBox(AOwner IComponent) ICustomGroupBox {
	r1 := LCL().SysCallN(1779, GetObjectUintptr(AOwner))
	return AsCustomGroupBox(r1)
}

func (m *TCustomGroupBox) ParentBackground() bool {
	r1 := LCL().SysCallN(1780, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomGroupBox) SetParentBackground(AValue bool) {
	LCL().SysCallN(1780, 1, m.Instance(), PascalBool(AValue))
}

func CustomGroupBoxClass() TClass {
	ret := LCL().SysCallN(1778)
	return TClass(ret)
}
