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
	r1 := LCL().SysCallN(3040, GetObjectUintptr(AOwner))
	return AsFlowPanelControlList(r1)
}

func (m *TFlowPanelControlList) ItemsForFlowPanelControl(Index int32) IFlowPanelControl {
	r1 := LCL().SysCallN(3042, 0, m.Instance(), uintptr(Index))
	return AsFlowPanelControl(r1)
}

func (m *TFlowPanelControlList) SetItemsForFlowPanelControl(Index int32, AValue IFlowPanelControl) {
	LCL().SysCallN(3042, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TFlowPanelControlList) IndexOf(AControl IControl) int32 {
	r1 := LCL().SysCallN(3041, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TFlowPanelControlList) AllowAdd() bool {
	r1 := LCL().SysCallN(3037, m.Instance())
	return GoBool(r1)
}

func (m *TFlowPanelControlList) AllowDelete() bool {
	r1 := LCL().SysCallN(3038, m.Instance())
	return GoBool(r1)
}

func FlowPanelControlListClass() TClass {
	ret := LCL().SysCallN(3039)
	return TClass(ret)
}
