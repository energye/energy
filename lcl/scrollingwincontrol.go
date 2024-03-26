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

// IScrollingWinControl Parent: ICustomControl
type IScrollingWinControl interface {
	ICustomControl
	HorzScrollBar() IControlScrollBar          // property
	SetHorzScrollBar(AValue IControlScrollBar) // property
	VertScrollBar() IControlScrollBar          // property
	SetVertScrollBar(AValue IControlScrollBar) // property
	UpdateScrollbars()                         // procedure
	ScrollInView(AControl IControl)            // procedure
}

// TScrollingWinControl Parent: TCustomControl
type TScrollingWinControl struct {
	TCustomControl
}

func NewScrollingWinControl(TheOwner IComponent) IScrollingWinControl {
	r1 := LCL().SysCallN(4315, GetObjectUintptr(TheOwner))
	return AsScrollingWinControl(r1)
}

func (m *TScrollingWinControl) HorzScrollBar() IControlScrollBar {
	r1 := LCL().SysCallN(4316, 0, m.Instance(), 0)
	return AsControlScrollBar(r1)
}

func (m *TScrollingWinControl) SetHorzScrollBar(AValue IControlScrollBar) {
	LCL().SysCallN(4316, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScrollingWinControl) VertScrollBar() IControlScrollBar {
	r1 := LCL().SysCallN(4319, 0, m.Instance(), 0)
	return AsControlScrollBar(r1)
}

func (m *TScrollingWinControl) SetVertScrollBar(AValue IControlScrollBar) {
	LCL().SysCallN(4319, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ScrollingWinControlClass() TClass {
	ret := LCL().SysCallN(4314)
	return TClass(ret)
}

func (m *TScrollingWinControl) UpdateScrollbars() {
	LCL().SysCallN(4318, m.Instance())
}

func (m *TScrollingWinControl) ScrollInView(AControl IControl) {
	LCL().SysCallN(4317, m.Instance(), GetObjectUintptr(AControl))
}
