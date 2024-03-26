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

// ICustomUpDown Parent: ICustomControl
type ICustomUpDown interface {
	ICustomControl
}

// TCustomUpDown Parent: TCustomControl
type TCustomUpDown struct {
	TCustomControl
}

func NewCustomUpDown(AOwner IComponent) ICustomUpDown {
	r1 := LCL().SysCallN(2285, GetObjectUintptr(AOwner))
	return AsCustomUpDown(r1)
}

func CustomUpDownClass() TClass {
	ret := LCL().SysCallN(2284)
	return TClass(ret)
}
