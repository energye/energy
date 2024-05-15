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

// IFPImageBitmap Parent: ICustomBitmap
type IFPImageBitmap interface {
	ICustomBitmap
}

// TFPImageBitmap Parent: TCustomBitmap
type TFPImageBitmap struct {
	TCustomBitmap
}

func NewFPImageBitmap() IFPImageBitmap {
	r1 := LCL().SysCallN(2952)
	return AsFPImageBitmap(r1)
}

func FPImageBitmapClass() TClass {
	ret := LCL().SysCallN(2951)
	return TClass(ret)
}
