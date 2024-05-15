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

// ICustomStaticText Parent: IWinControl
type ICustomStaticText interface {
	IWinControl
	Alignment() TAlignment                    // property
	SetAlignment(AValue TAlignment)           // property
	BorderStyle() TStaticBorderStyle          // property
	SetBorderStyle(AValue TStaticBorderStyle) // property
	FocusControl() IWinControl                // property
	SetFocusControl(AValue IWinControl)       // property
	ShowAccelChar() bool                      // property
	SetShowAccelChar(AValue bool)             // property
	Transparent() bool                        // property
	SetTransparent(AValue bool)               // property
}

// TCustomStaticText Parent: TWinControl
type TCustomStaticText struct {
	TWinControl
}

func NewCustomStaticText(AOwner IComponent) ICustomStaticText {
	r1 := LCL().SysCallN(2267, GetObjectUintptr(AOwner))
	return AsCustomStaticText(r1)
}

func (m *TCustomStaticText) Alignment() TAlignment {
	r1 := LCL().SysCallN(2264, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomStaticText) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2264, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) BorderStyle() TStaticBorderStyle {
	r1 := LCL().SysCallN(2265, 0, m.Instance(), 0)
	return TStaticBorderStyle(r1)
}

func (m *TCustomStaticText) SetBorderStyle(AValue TStaticBorderStyle) {
	LCL().SysCallN(2265, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) FocusControl() IWinControl {
	r1 := LCL().SysCallN(2268, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomStaticText) SetFocusControl(AValue IWinControl) {
	LCL().SysCallN(2268, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomStaticText) ShowAccelChar() bool {
	r1 := LCL().SysCallN(2269, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(2269, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomStaticText) Transparent() bool {
	r1 := LCL().SysCallN(2270, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetTransparent(AValue bool) {
	LCL().SysCallN(2270, 1, m.Instance(), PascalBool(AValue))
}

func CustomStaticTextClass() TClass {
	ret := LCL().SysCallN(2266)
	return TClass(ret)
}
