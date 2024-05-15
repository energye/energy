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

// IStringsEnumerator Parent: IObject
type IStringsEnumerator interface {
	IObject
	Current() string    // property
	GetCurrent() string // function
	MoveNext() bool     // function
}

// TStringsEnumerator Parent: TObject
type TStringsEnumerator struct {
	TObject
}

func NewStringsEnumerator(AStrings IStrings) IStringsEnumerator {
	r1 := LCL().SysCallN(5246, GetObjectUintptr(AStrings))
	return AsStringsEnumerator(r1)
}

func (m *TStringsEnumerator) Current() string {
	r1 := LCL().SysCallN(5247, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) GetCurrent() string {
	r1 := LCL().SysCallN(5248, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5249, m.Instance())
	return GoBool(r1)
}

func StringsEnumeratorClass() TClass {
	ret := LCL().SysCallN(5245)
	return TClass(ret)
}
