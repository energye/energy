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

// IIconOptions Parent: IPersistent
type IIconOptions interface {
	IPersistent
	Arrangement() TIconArrangement          // property
	SetArrangement(AValue TIconArrangement) // property
	AutoArrange() bool                      // property
	SetAutoArrange(AValue bool)             // property
	WrapText() bool                         // property
	SetWrapText(AValue bool)                // property
}

// TIconOptions Parent: TPersistent
type TIconOptions struct {
	TPersistent
}

func NewIconOptions(AOwner ICustomListView) IIconOptions {
	r1 := LCL().SysCallN(3346, GetObjectUintptr(AOwner))
	return AsIconOptions(r1)
}

func (m *TIconOptions) Arrangement() TIconArrangement {
	r1 := LCL().SysCallN(3343, 0, m.Instance(), 0)
	return TIconArrangement(r1)
}

func (m *TIconOptions) SetArrangement(AValue TIconArrangement) {
	LCL().SysCallN(3343, 1, m.Instance(), uintptr(AValue))
}

func (m *TIconOptions) AutoArrange() bool {
	r1 := LCL().SysCallN(3344, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIconOptions) SetAutoArrange(AValue bool) {
	LCL().SysCallN(3344, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIconOptions) WrapText() bool {
	r1 := LCL().SysCallN(3347, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIconOptions) SetWrapText(AValue bool) {
	LCL().SysCallN(3347, 1, m.Instance(), PascalBool(AValue))
}

func IconOptionsClass() TClass {
	ret := LCL().SysCallN(3345)
	return TClass(ret)
}
