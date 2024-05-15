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

// IFPMemoryImage Parent: IFPCustomImage
type IFPMemoryImage interface {
	IFPCustomImage
}

// TFPMemoryImage Parent: TFPCustomImage
type TFPMemoryImage struct {
	TFPCustomImage
}

func NewFPMemoryImage(AWidth, AHeight int32) IFPMemoryImage {
	r1 := LCL().SysCallN(2981, uintptr(AWidth), uintptr(AHeight))
	return AsFPMemoryImage(r1)
}

func FPMemoryImageClass() TClass {
	ret := LCL().SysCallN(2980)
	return TClass(ret)
}
