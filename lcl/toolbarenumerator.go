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

// IToolBarEnumerator Parent: IObject
type IToolBarEnumerator interface {
	IObject
	Current() IToolButton // property
	MoveNext() bool       // function
}

// TToolBarEnumerator Parent: TObject
type TToolBarEnumerator struct {
	TObject
}

func NewToolBarEnumerator(AToolBar IToolBar) IToolBarEnumerator {
	r1 := LCL().SysCallN(5455, GetObjectUintptr(AToolBar))
	return AsToolBarEnumerator(r1)
}

func (m *TToolBarEnumerator) Current() IToolButton {
	r1 := LCL().SysCallN(5456, m.Instance())
	return AsToolButton(r1)
}

func (m *TToolBarEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5457, m.Instance())
	return GoBool(r1)
}

func ToolBarEnumeratorClass() TClass {
	ret := LCL().SysCallN(5454)
	return TClass(ret)
}
