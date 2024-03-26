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
	r1 := LCL().SysCallN(2711, GetObjectUintptr(AList))
	return AsFPListEnumerator(r1)
}

func (m *TFPListEnumerator) Current() uintptr {
	r1 := LCL().SysCallN(2712, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) GetCurrent() uintptr {
	r1 := LCL().SysCallN(2713, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(2714, m.Instance())
	return GoBool(r1)
}

func FPListEnumeratorClass() TClass {
	ret := LCL().SysCallN(2710)
	return TClass(ret)
}
