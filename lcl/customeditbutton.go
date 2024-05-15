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

// ICustomEditButton Parent: ICustomAbstractGroupedEdit
type ICustomEditButton interface {
	ICustomAbstractGroupedEdit
}

// TCustomEditButton Parent: TCustomAbstractGroupedEdit
type TCustomEditButton struct {
	TCustomAbstractGroupedEdit
}

func NewCustomEditButton(AOwner IComponent) ICustomEditButton {
	r1 := LCL().SysCallN(1610, GetObjectUintptr(AOwner))
	return AsCustomEditButton(r1)
}

func CustomEditButtonClass() TClass {
	ret := LCL().SysCallN(1609)
	return TClass(ret)
}
