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

// ICustomMemo Parent: ICustomEdit
type ICustomMemo interface {
	ICustomEdit
	Lines() IStrings                        // property
	SetLines(AValue IStrings)               // property
	HorzScrollBar() IMemoScrollBar          // property
	SetHorzScrollBar(AValue IMemoScrollBar) // property
	VertScrollBar() IMemoScrollBar          // property
	SetVertScrollBar(AValue IMemoScrollBar) // property
	ScrollBars() TScrollStyle               // property
	SetScrollBars(AValue TScrollStyle)      // property
	WantReturns() bool                      // property
	SetWantReturns(AValue bool)             // property
	WantTabs() bool                         // property
	SetWantTabs(AValue bool)                // property
	WordWrap() bool                         // property
	SetWordWrap(AValue bool)                // property
	Append(AValue string)                   // procedure
}

// TCustomMemo Parent: TCustomEdit
type TCustomMemo struct {
	TCustomEdit
}

func NewCustomMemo(AOwner IComponent) ICustomMemo {
	r1 := LCL().SysCallN(2094, GetObjectUintptr(AOwner))
	return AsCustomMemo(r1)
}

func (m *TCustomMemo) Lines() IStrings {
	r1 := LCL().SysCallN(2096, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomMemo) SetLines(AValue IStrings) {
	LCL().SysCallN(2096, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) HorzScrollBar() IMemoScrollBar {
	r1 := LCL().SysCallN(2095, 0, m.Instance(), 0)
	return AsMemoScrollBar(r1)
}

func (m *TCustomMemo) SetHorzScrollBar(AValue IMemoScrollBar) {
	LCL().SysCallN(2095, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) VertScrollBar() IMemoScrollBar {
	r1 := LCL().SysCallN(2098, 0, m.Instance(), 0)
	return AsMemoScrollBar(r1)
}

func (m *TCustomMemo) SetVertScrollBar(AValue IMemoScrollBar) {
	LCL().SysCallN(2098, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(2097, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomMemo) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(2097, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomMemo) WantReturns() bool {
	r1 := LCL().SysCallN(2099, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWantReturns(AValue bool) {
	LCL().SysCallN(2099, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMemo) WantTabs() bool {
	r1 := LCL().SysCallN(2100, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWantTabs(AValue bool) {
	LCL().SysCallN(2100, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMemo) WordWrap() bool {
	r1 := LCL().SysCallN(2101, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWordWrap(AValue bool) {
	LCL().SysCallN(2101, 1, m.Instance(), PascalBool(AValue))
}

func CustomMemoClass() TClass {
	ret := LCL().SysCallN(2093)
	return TClass(ret)
}

func (m *TCustomMemo) Append(AValue string) {
	LCL().SysCallN(2092, m.Instance(), PascalStr(AValue))
}
