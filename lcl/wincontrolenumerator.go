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

// IWinControlEnumerator Parent: IObject
type IWinControlEnumerator interface {
	IObject
	Current() IControl                    // property
	GetEnumerator() IWinControlEnumerator // function
	MoveNext() bool                       // function
}

// TWinControlEnumerator Parent: TObject
type TWinControlEnumerator struct {
	TObject
}

func NewWinControlEnumerator(Parent IWinControl, aLowToHigh bool) IWinControlEnumerator {
	r1 := LCL().SysCallN(5197, GetObjectUintptr(Parent), PascalBool(aLowToHigh))
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) Current() IControl {
	r1 := LCL().SysCallN(5198, m.Instance())
	return AsControl(r1)
}

func (m *TWinControlEnumerator) GetEnumerator() IWinControlEnumerator {
	r1 := LCL().SysCallN(5199, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5200, m.Instance())
	return GoBool(r1)
}

func WinControlEnumeratorClass() TClass {
	ret := LCL().SysCallN(5196)
	return TClass(ret)
}
