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

// IFPCustomImageHandler Parent: IObject
type IFPCustomImageHandler interface {
	IObject
	SetOnProgress(fn TFPImgProgressEvent) // property event
}

// TFPCustomImageHandler Parent: TObject
type TFPCustomImageHandler struct {
	TObject
	progressPtr uintptr
}

func NewFPCustomImageHandler() IFPCustomImageHandler {
	r1 := LCL().SysCallN(2910)
	return AsFPCustomImageHandler(r1)
}

func FPCustomImageHandlerClass() TClass {
	ret := LCL().SysCallN(2909)
	return TClass(ret)
}

func (m *TFPCustomImageHandler) SetOnProgress(fn TFPImgProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2911, m.Instance(), m.progressPtr)
}
