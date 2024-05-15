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

// IBasicAction Parent: IComponent
type IBasicAction interface {
	IComponent
	ActionComponent() IComponent              // property
	SetActionComponent(AValue IComponent)     // property
	HandlesTarget(Target IObject) bool        // function
	Execute() bool                            // function
	Update() bool                             // function
	UpdateTarget(Target IObject)              // procedure
	ExecuteTarget(Target IObject)             // procedure
	RegisterChanges(Value IBasicActionLink)   // procedure
	UnRegisterChanges(Value IBasicActionLink) // procedure
	SetOnExecute(fn TNotifyEvent)             // property event
	SetOnUpdate(fn TNotifyEvent)              // property event
}

// TBasicAction Parent: TComponent
type TBasicAction struct {
	TComponent
	executePtr uintptr
	updatePtr  uintptr
}

func NewBasicAction(AOwner IComponent) IBasicAction {
	r1 := LCL().SysCallN(399, GetObjectUintptr(AOwner))
	return AsBasicAction(r1)
}

func (m *TBasicAction) ActionComponent() IComponent {
	r1 := LCL().SysCallN(397, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TBasicAction) SetActionComponent(AValue IComponent) {
	LCL().SysCallN(397, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBasicAction) HandlesTarget(Target IObject) bool {
	r1 := LCL().SysCallN(402, m.Instance(), GetObjectUintptr(Target))
	return GoBool(r1)
}

func (m *TBasicAction) Execute() bool {
	r1 := LCL().SysCallN(400, m.Instance())
	return GoBool(r1)
}

func (m *TBasicAction) Update() bool {
	r1 := LCL().SysCallN(407, m.Instance())
	return GoBool(r1)
}

func BasicActionClass() TClass {
	ret := LCL().SysCallN(398)
	return TClass(ret)
}

func (m *TBasicAction) UpdateTarget(Target IObject) {
	LCL().SysCallN(408, m.Instance(), GetObjectUintptr(Target))
}

func (m *TBasicAction) ExecuteTarget(Target IObject) {
	LCL().SysCallN(401, m.Instance(), GetObjectUintptr(Target))
}

func (m *TBasicAction) RegisterChanges(Value IBasicActionLink) {
	LCL().SysCallN(403, m.Instance(), GetObjectUintptr(Value))
}

func (m *TBasicAction) UnRegisterChanges(Value IBasicActionLink) {
	LCL().SysCallN(406, m.Instance(), GetObjectUintptr(Value))
}

func (m *TBasicAction) SetOnExecute(fn TNotifyEvent) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(404, m.Instance(), m.executePtr)
}

func (m *TBasicAction) SetOnUpdate(fn TNotifyEvent) {
	if m.updatePtr != 0 {
		RemoveEventElement(m.updatePtr)
	}
	m.updatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(405, m.Instance(), m.updatePtr)
}
