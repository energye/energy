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

// ILazDockSplitter Parent: ICustomSplitter
type ILazDockSplitter interface {
	ICustomSplitter
}

// TLazDockSplitter Parent: TCustomSplitter
type TLazDockSplitter struct {
	TCustomSplitter
}

func NewLazDockSplitter(AOwner IComponent) ILazDockSplitter {
	r1 := LCL().SysCallN(3291, GetObjectUintptr(AOwner))
	return AsLazDockSplitter(r1)
}

func LazDockSplitterClass() TClass {
	ret := LCL().SysCallN(3290)
	return TClass(ret)
}
