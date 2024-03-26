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

// IAction Parent: ICustomAction
type IAction interface {
	ICustomAction
}

// TAction Parent: TCustomAction
type TAction struct {
	TCustomAction
}

func NewAction(AOwner IComponent) IAction {
	r1 := LCL().SysCallN(86, GetObjectUintptr(AOwner))
	return AsAction(r1)
}

func ActionClass() TClass {
	ret := LCL().SysCallN(85)
	return TClass(ret)
}
