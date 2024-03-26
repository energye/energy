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

// IListEnumerator Parent: IObject
type IListEnumerator interface {
	IObject
	Current() uintptr    // property
	GetCurrent() uintptr // function
	MoveNext() bool      // function
}

// TListEnumerator Parent: TObject
type TListEnumerator struct {
	TObject
}

func NewListEnumerator(AList IList) IListEnumerator {
	r1 := LCL().SysCallN(3377, GetObjectUintptr(AList))
	return AsListEnumerator(r1)
}

func (m *TListEnumerator) Current() uintptr {
	r1 := LCL().SysCallN(3378, m.Instance())
	return uintptr(r1)
}

func (m *TListEnumerator) GetCurrent() uintptr {
	r1 := LCL().SysCallN(3379, m.Instance())
	return uintptr(r1)
}

func (m *TListEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(3380, m.Instance())
	return GoBool(r1)
}

func ListEnumeratorClass() TClass {
	ret := LCL().SysCallN(3376)
	return TClass(ret)
}
