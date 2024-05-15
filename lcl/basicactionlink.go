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

// IBasicActionLink Parent: IObject
type IBasicActionLink interface {
	IObject
	Action() IBasicAction               // property
	SetAction(AValue IBasicAction)      // property
	Execute(AComponent IComponent) bool // function
	Update() bool                       // function
	SetOnChange(fn TNotifyEvent)        // property event
}

// TBasicActionLink Parent: TObject
type TBasicActionLink struct {
	TObject
	changePtr uintptr
}

func NewBasicActionLink(AClient IObject) IBasicActionLink {
	r1 := LCL().SysCallN(393, GetObjectUintptr(AClient))
	return AsBasicActionLink(r1)
}

func (m *TBasicActionLink) Action() IBasicAction {
	r1 := LCL().SysCallN(391, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TBasicActionLink) SetAction(AValue IBasicAction) {
	LCL().SysCallN(391, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBasicActionLink) Execute(AComponent IComponent) bool {
	r1 := LCL().SysCallN(394, m.Instance(), GetObjectUintptr(AComponent))
	return GoBool(r1)
}

func (m *TBasicActionLink) Update() bool {
	r1 := LCL().SysCallN(396, m.Instance())
	return GoBool(r1)
}

func BasicActionLinkClass() TClass {
	ret := LCL().SysCallN(392)
	return TClass(ret)
}

func (m *TBasicActionLink) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(395, m.Instance(), m.changePtr)
}
