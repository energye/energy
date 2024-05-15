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

// ICustomFlowPanel Parent: ICustomPanel
type ICustomFlowPanel interface {
	ICustomPanel
	AutoWrap() bool                              // property
	SetAutoWrap(AValue bool)                     // property
	ControlList() IFlowPanelControlList          // property
	SetControlList(AValue IFlowPanelControlList) // property
	FlowStyle() TFlowStyle                       // property
	SetFlowStyle(AValue TFlowStyle)              // property
	FlowLayout() TTextLayout                     // property
	SetFlowLayout(AValue TTextLayout)            // property
}

// TCustomFlowPanel Parent: TCustomPanel
type TCustomFlowPanel struct {
	TCustomPanel
}

func NewCustomFlowPanel(AOwner IComponent) ICustomFlowPanel {
	r1 := LCL().SysCallN(1654, GetObjectUintptr(AOwner))
	return AsCustomFlowPanel(r1)
}

func (m *TCustomFlowPanel) AutoWrap() bool {
	r1 := LCL().SysCallN(1651, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFlowPanel) SetAutoWrap(AValue bool) {
	LCL().SysCallN(1651, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFlowPanel) ControlList() IFlowPanelControlList {
	r1 := LCL().SysCallN(1653, 0, m.Instance(), 0)
	return AsFlowPanelControlList(r1)
}

func (m *TCustomFlowPanel) SetControlList(AValue IFlowPanelControlList) {
	LCL().SysCallN(1653, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomFlowPanel) FlowStyle() TFlowStyle {
	r1 := LCL().SysCallN(1656, 0, m.Instance(), 0)
	return TFlowStyle(r1)
}

func (m *TCustomFlowPanel) SetFlowStyle(AValue TFlowStyle) {
	LCL().SysCallN(1656, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomFlowPanel) FlowLayout() TTextLayout {
	r1 := LCL().SysCallN(1655, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TCustomFlowPanel) SetFlowLayout(AValue TTextLayout) {
	LCL().SysCallN(1655, 1, m.Instance(), uintptr(AValue))
}

func CustomFlowPanelClass() TClass {
	ret := LCL().SysCallN(1652)
	return TClass(ret)
}
