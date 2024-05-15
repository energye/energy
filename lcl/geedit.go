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

// IGEEdit Parent: ICustomMaskEdit
type IGEEdit interface {
	ICustomMaskEdit
}

// TGEEdit Parent: TCustomMaskEdit
type TGEEdit struct {
	TCustomMaskEdit
}

func NewGEEdit(TheOwner IComponent) IGEEdit {
	r1 := LCL().SysCallN(3161, GetObjectUintptr(TheOwner))
	return AsGEEdit(r1)
}

func GEEditClass() TClass {
	ret := LCL().SysCallN(3160)
	return TClass(ret)
}
