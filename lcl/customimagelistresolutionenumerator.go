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

// ICustomImageListResolutionEnumerator Parent: IObject
type ICustomImageListResolutionEnumerator interface {
	IObject
	Current() ICustomImageListResolution                 // property
	GetEnumerator() ICustomImageListResolutionEnumerator // function
	MoveNext() bool                                      // function
}

// TCustomImageListResolutionEnumerator Parent: TObject
type TCustomImageListResolutionEnumerator struct {
	TObject
}

func NewCustomImageListResolutionEnumerator(AImgList ICustomImageList, ADesc bool) ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1815, GetObjectUintptr(AImgList), PascalBool(ADesc))
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageListResolutionEnumerator) Current() ICustomImageListResolution {
	r1 := LCL().SysCallN(1816, m.Instance())
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageListResolutionEnumerator) GetEnumerator() ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1817, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageListResolutionEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(1818, m.Instance())
	return GoBool(r1)
}

func CustomImageListResolutionEnumeratorClass() TClass {
	ret := LCL().SysCallN(1814)
	return TClass(ret)
}
