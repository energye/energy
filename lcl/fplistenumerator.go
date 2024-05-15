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

// IFPListEnumerator Parent: IObject
type IFPListEnumerator interface {
	IObject
	Current() uintptr    // property
	GetCurrent() uintptr // function
	MoveNext() bool      // function
}

// TFPListEnumerator Parent: TObject
type TFPListEnumerator struct {
	TObject
}

func NewFPListEnumerator(AList IFPList) IFPListEnumerator {
	r1 := LCL().SysCallN(2954, GetObjectUintptr(AList))
	return AsFPListEnumerator(r1)
}

func (m *TFPListEnumerator) Current() uintptr {
	r1 := LCL().SysCallN(2955, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) GetCurrent() uintptr {
	r1 := LCL().SysCallN(2956, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(2957, m.Instance())
	return GoBool(r1)
}

func FPListEnumeratorClass() TClass {
	ret := LCL().SysCallN(2953)
	return TClass(ret)
}
