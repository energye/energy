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

// IContainedAction Parent: IBasicAction
type IContainedAction interface {
	IBasicAction
	ActionList() ICustomActionList          // property
	SetActionList(AValue ICustomActionList) // property
	Index() int32                           // property
	SetIndex(AValue int32)                  // property
	Category() string                       // property
	SetCategory(AValue string)              // property
}

// TContainedAction Parent: TBasicAction
type TContainedAction struct {
	TBasicAction
}

func NewContainedAction(AOwner IComponent) IContainedAction {
	r1 := LCL().SysCallN(718, GetObjectUintptr(AOwner))
	return AsContainedAction(r1)
}

func (m *TContainedAction) ActionList() ICustomActionList {
	r1 := LCL().SysCallN(715, 0, m.Instance(), 0)
	return AsCustomActionList(r1)
}

func (m *TContainedAction) SetActionList(AValue ICustomActionList) {
	LCL().SysCallN(715, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TContainedAction) Index() int32 {
	r1 := LCL().SysCallN(719, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TContainedAction) SetIndex(AValue int32) {
	LCL().SysCallN(719, 1, m.Instance(), uintptr(AValue))
}

func (m *TContainedAction) Category() string {
	r1 := LCL().SysCallN(716, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TContainedAction) SetCategory(AValue string) {
	LCL().SysCallN(716, 1, m.Instance(), PascalStr(AValue))
}

func ContainedActionClass() TClass {
	ret := LCL().SysCallN(717)
	return TClass(ret)
}
