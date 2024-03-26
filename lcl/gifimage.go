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

// IGIFImage Parent: IFPImageBitmap
type IGIFImage interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Interlaced() bool   // property
	BitsPerPixel() byte // property
}

// TGIFImage Parent: TFPImageBitmap
type TGIFImage struct {
	TFPImageBitmap
}

func NewGIFImage() IGIFImage {
	r1 := LCL().SysCallN(2921)
	return AsGIFImage(r1)
}

func (m *TGIFImage) Interlaced() bool {
	r1 := LCL().SysCallN(2922, m.Instance())
	return GoBool(r1)
}

func (m *TGIFImage) BitsPerPixel() byte {
	r1 := LCL().SysCallN(2919, m.Instance())
	return byte(r1)
}

func GIFImageClass() TClass {
	ret := LCL().SysCallN(2920)
	return TClass(ret)
}
