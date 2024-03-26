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

// IComponentEnumerator Parent: IObject
type IComponentEnumerator interface {
	IObject
	Current() IComponent    // property
	GetCurrent() IComponent // function
	MoveNext() bool         // function
}

// TComponentEnumerator Parent: TObject
type TComponentEnumerator struct {
	TObject
}

func NewComponentEnumerator(AComponent IComponent) IComponentEnumerator {
	r1 := LCL().SysCallN(685, GetObjectUintptr(AComponent))
	return AsComponentEnumerator(r1)
}

func (m *TComponentEnumerator) Current() IComponent {
	r1 := LCL().SysCallN(686, m.Instance())
	return AsComponent(r1)
}

func (m *TComponentEnumerator) GetCurrent() IComponent {
	r1 := LCL().SysCallN(687, m.Instance())
	return AsComponent(r1)
}

func (m *TComponentEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(688, m.Instance())
	return GoBool(r1)
}

func ComponentEnumeratorClass() TClass {
	ret := LCL().SysCallN(684)
	return TClass(ret)
}
