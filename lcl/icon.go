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

// IIcon Parent: ICustomIcon
type IIcon interface {
	ICustomIcon
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Handle() HICON          // property
	SetHandle(AValue HICON) // property
	ReleaseHandle() HICON   // function
}

// TIcon Parent: TCustomIcon
type TIcon struct {
	TCustomIcon
}

func NewIcon() IIcon {
	r1 := LCL().SysCallN(3106)
	return AsIcon(r1)
}

func (m *TIcon) Handle() HICON {
	r1 := LCL().SysCallN(3107, 0, m.Instance(), 0)
	return HICON(r1)
}

func (m *TIcon) SetHandle(AValue HICON) {
	LCL().SysCallN(3107, 1, m.Instance(), uintptr(AValue))
}

func (m *TIcon) ReleaseHandle() HICON {
	r1 := LCL().SysCallN(3108, m.Instance())
	return HICON(r1)
}

func IconClass() TClass {
	ret := LCL().SysCallN(3105)
	return TClass(ret)
}
