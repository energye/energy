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

// ICustomStringTreeOptions Parent: ICustomVirtualTreeOptions
type ICustomStringTreeOptions interface {
	ICustomVirtualTreeOptions
}

// TCustomStringTreeOptions Parent: TCustomVirtualTreeOptions
type TCustomStringTreeOptions struct {
	TCustomVirtualTreeOptions
}

func NewCustomStringTreeOptions(AOwner IBaseVirtualTree) ICustomStringTreeOptions {
	r1 := LCL().SysCallN(2292, GetObjectUintptr(AOwner))
	return AsCustomStringTreeOptions(r1)
}

func CustomStringTreeOptionsClass() TClass {
	ret := LCL().SysCallN(2291)
	return TClass(ret)
}
