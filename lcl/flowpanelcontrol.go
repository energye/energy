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

// IFlowPanelControl Parent: ICollectionItem
type IFlowPanelControl interface {
	ICollectionItem
	Control() IControl              // property
	SetControl(AValue IControl)     // property
	WrapAfter() TWrapAfter          // property
	SetWrapAfter(AValue TWrapAfter) // property
	AllowAdd() bool                 // function
	AllowDelete() bool              // function
}

// TFlowPanelControl Parent: TCollectionItem
type TFlowPanelControl struct {
	TCollectionItem
}

func NewFlowPanelControl(ACollection ICollection) IFlowPanelControl {
	r1 := LCL().SysCallN(3047, GetObjectUintptr(ACollection))
	return AsFlowPanelControl(r1)
}

func (m *TFlowPanelControl) Control() IControl {
	r1 := LCL().SysCallN(3046, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TFlowPanelControl) SetControl(AValue IControl) {
	LCL().SysCallN(3046, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFlowPanelControl) WrapAfter() TWrapAfter {
	r1 := LCL().SysCallN(3048, 0, m.Instance(), 0)
	return TWrapAfter(r1)
}

func (m *TFlowPanelControl) SetWrapAfter(AValue TWrapAfter) {
	LCL().SysCallN(3048, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanelControl) AllowAdd() bool {
	r1 := LCL().SysCallN(3043, m.Instance())
	return GoBool(r1)
}

func (m *TFlowPanelControl) AllowDelete() bool {
	r1 := LCL().SysCallN(3044, m.Instance())
	return GoBool(r1)
}

func FlowPanelControlClass() TClass {
	ret := LCL().SysCallN(3045)
	return TClass(ret)
}
