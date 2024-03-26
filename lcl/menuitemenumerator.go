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

// IMenuItemEnumerator Parent: IObject
type IMenuItemEnumerator interface {
	IObject
	Current() IMenuItem // property
	MoveNext() bool     // function
}

// TMenuItemEnumerator Parent: TObject
type TMenuItemEnumerator struct {
	TObject
}

func NewMenuItemEnumerator(AMenuItem IMenuItem) IMenuItemEnumerator {
	r1 := LCL().SysCallN(3595, GetObjectUintptr(AMenuItem))
	return AsMenuItemEnumerator(r1)
}

func (m *TMenuItemEnumerator) Current() IMenuItem {
	r1 := LCL().SysCallN(3596, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItemEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(3597, m.Instance())
	return GoBool(r1)
}

func MenuItemEnumeratorClass() TClass {
	ret := LCL().SysCallN(3594)
	return TClass(ret)
}
