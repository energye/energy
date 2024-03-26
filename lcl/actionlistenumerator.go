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

// IActionListEnumerator Parent: IObject
type IActionListEnumerator interface {
	IObject
	Current() IContainedAction // property
	MoveNext() bool            // function
}

// TActionListEnumerator Parent: TObject
type TActionListEnumerator struct {
	TObject
}

func NewActionListEnumerator(AList ICustomActionList) IActionListEnumerator {
	r1 := LCL().SysCallN(77, GetObjectUintptr(AList))
	return AsActionListEnumerator(r1)
}

func (m *TActionListEnumerator) Current() IContainedAction {
	r1 := LCL().SysCallN(78, m.Instance())
	return AsContainedAction(r1)
}

func (m *TActionListEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(79, m.Instance())
	return GoBool(r1)
}

func ActionListEnumeratorClass() TClass {
	ret := LCL().SysCallN(76)
	return TClass(ret)
}
