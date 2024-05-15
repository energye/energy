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

// IClipboardFormats Parent: IStringList
type IClipboardFormats interface {
	IStringList
	Owner() IBaseVirtualTree // property
}

// TClipboardFormats Parent: TStringList
type TClipboardFormats struct {
	TStringList
}

func NewClipboardFormats(AOwner IBaseVirtualTree) IClipboardFormats {
	r1 := LCL().SysCallN(655, GetObjectUintptr(AOwner))
	return AsClipboardFormats(r1)
}

func (m *TClipboardFormats) Owner() IBaseVirtualTree {
	r1 := LCL().SysCallN(656, m.Instance())
	return AsBaseVirtualTree(r1)
}

func ClipboardFormatsClass() TClass {
	ret := LCL().SysCallN(654)
	return TClass(ret)
}
