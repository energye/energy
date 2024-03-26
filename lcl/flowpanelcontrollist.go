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

// IFlowPanelControlList Parent: IOwnedCollection
type IFlowPanelControlList interface {
	IOwnedCollection
	ItemsForFlowPanelControl(Index int32) IFlowPanelControl            // property
	SetItemsForFlowPanelControl(Index int32, AValue IFlowPanelControl) // property
	IndexOf(AControl IControl) int32                                   // function
	AllowAdd() bool                                                    // function
	AllowDelete() bool                                                 // function
}

// TFlowPanelControlList Parent: TOwnedCollection
type TFlowPanelControlList struct {
	TOwnedCollection
}

func NewFlowPanelControlList(AOwner IPersistent) IFlowPanelControlList {
	r1 := LCL().SysCallN(2797, GetObjectUintptr(AOwner))
	return AsFlowPanelControlList(r1)
}

func (m *TFlowPanelControlList) ItemsForFlowPanelControl(Index int32) IFlowPanelControl {
	r1 := LCL().SysCallN(2799, 0, m.Instance(), uintptr(Index))
	return AsFlowPanelControl(r1)
}

func (m *TFlowPanelControlList) SetItemsForFlowPanelControl(Index int32, AValue IFlowPanelControl) {
	LCL().SysCallN(2799, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TFlowPanelControlList) IndexOf(AControl IControl) int32 {
	r1 := LCL().SysCallN(2798, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TFlowPanelControlList) AllowAdd() bool {
	r1 := LCL().SysCallN(2794, m.Instance())
	return GoBool(r1)
}

func (m *TFlowPanelControlList) AllowDelete() bool {
	r1 := LCL().SysCallN(2795, m.Instance())
	return GoBool(r1)
}

func FlowPanelControlListClass() TClass {
	ret := LCL().SysCallN(2796)
	return TClass(ret)
}
