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

// ICustomVirtualTreeOptions Parent: IPersistent
type ICustomVirtualTreeOptions interface {
	IPersistent
	Owner() IBaseVirtualTree   // property
	AssignTo(Dest IPersistent) // procedure
}

// TCustomVirtualTreeOptions Parent: TPersistent
type TCustomVirtualTreeOptions struct {
	TPersistent
}

func NewCustomVirtualTreeOptions(AOwner IBaseVirtualTree) ICustomVirtualTreeOptions {
	r1 := LCL().SysCallN(2501, GetObjectUintptr(AOwner))
	return AsCustomVirtualTreeOptions(r1)
}

func (m *TCustomVirtualTreeOptions) Owner() IBaseVirtualTree {
	r1 := LCL().SysCallN(2502, m.Instance())
	return AsBaseVirtualTree(r1)
}

func CustomVirtualTreeOptionsClass() TClass {
	ret := LCL().SysCallN(2500)
	return TClass(ret)
}

func (m *TCustomVirtualTreeOptions) AssignTo(Dest IPersistent) {
	LCL().SysCallN(2499, m.Instance(), GetObjectUintptr(Dest))
}
