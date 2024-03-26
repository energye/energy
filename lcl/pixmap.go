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

// IPixmap Parent: IFPImageBitmap
type IPixmap interface {
	IFPImageBitmap
}

// TPixmap Parent: TFPImageBitmap
type TPixmap struct {
	TFPImageBitmap
}

func NewPixmap() IPixmap {
	r1 := LCL().SysCallN(3912)
	return AsPixmap(r1)
}

func PixmapClass() TClass {
	ret := LCL().SysCallN(3911)
	return TClass(ret)
}
