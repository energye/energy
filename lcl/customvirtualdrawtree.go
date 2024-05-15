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

// ICustomVirtualDrawTree Parent: IBaseVirtualTree
// Tree descendant to let an application draw its stuff itself.
type ICustomVirtualDrawTree interface {
	IBaseVirtualTree
}

// TCustomVirtualDrawTree Parent: TBaseVirtualTree
// Tree descendant to let an application draw its stuff itself.
type TCustomVirtualDrawTree struct {
	TBaseVirtualTree
}

func NewCustomVirtualDrawTree(AOwner IComponent) ICustomVirtualDrawTree {
	r1 := LCL().SysCallN(2479, GetObjectUintptr(AOwner))
	return AsCustomVirtualDrawTree(r1)
}

func CustomVirtualDrawTreeClass() TClass {
	ret := LCL().SysCallN(2478)
	return TClass(ret)
}
