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

// IPortableAnyMapGraphic Parent: IFPImageBitmap
type IPortableAnyMapGraphic interface {
	IFPImageBitmap
}

// TPortableAnyMapGraphic Parent: TFPImageBitmap
type TPortableAnyMapGraphic struct {
	TFPImageBitmap
}

func NewPortableAnyMapGraphic() IPortableAnyMapGraphic {
	r1 := LCL().SysCallN(4569)
	return AsPortableAnyMapGraphic(r1)
}

func PortableAnyMapGraphicClass() TClass {
	ret := LCL().SysCallN(4568)
	return TClass(ret)
}
