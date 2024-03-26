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

// IBitmap Parent: IFPImageBitmap
type IBitmap interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
}

// TBitmap Parent: TFPImageBitmap
type TBitmap struct {
	TFPImageBitmap
}

func NewBitmap() IBitmap {
	r1 := LCL().SysCallN(257)
	return AsBitmap(r1)
}

func BitmapClass() TClass {
	ret := LCL().SysCallN(256)
	return TClass(ret)
}
