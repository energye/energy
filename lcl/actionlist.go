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

// IActionList Parent: ICustomActionList
type IActionList interface {
	ICustomActionList
	SetOnChange(fn TNotifyEvent)  // property event
	SetOnExecute(fn TActionEvent) // property event
	SetOnUpdate(fn TActionEvent)  // property event
}

// TActionList Parent: TCustomActionList
type TActionList struct {
	TCustomActionList
	changePtr  uintptr
	executePtr uintptr
	updatePtr  uintptr
}

func NewActionList(AOwner IComponent) IActionList {
	r1 := LCL().SysCallN(81, GetObjectUintptr(AOwner))
	return AsActionList(r1)
}

func ActionListClass() TClass {
	ret := LCL().SysCallN(80)
	return TClass(ret)
}

func (m *TActionList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(82, m.Instance(), m.changePtr)
}

func (m *TActionList) SetOnExecute(fn TActionEvent) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(83, m.Instance(), m.executePtr)
}

func (m *TActionList) SetOnUpdate(fn TActionEvent) {
	if m.updatePtr != 0 {
		RemoveEventElement(m.updatePtr)
	}
	m.updatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(84, m.Instance(), m.updatePtr)
}
