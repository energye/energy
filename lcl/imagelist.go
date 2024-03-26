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

// IImageList Parent: IDragImageList
type IImageList interface {
	IDragImageList
}

// TImageList Parent: TDragImageList
type TImageList struct {
	TDragImageList
}

func NewImageList(AOwner IComponent) IImageList {
	r1 := LCL().SysCallN(3136, GetObjectUintptr(AOwner))
	return AsImageList(r1)
}

func ImageListClass() TClass {
	ret := LCL().SysCallN(3135)
	return TClass(ret)
}
