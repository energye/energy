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

// IFPCustomImageWriter Is Abstract Class Parent: IFPCustomImageHandler
type IFPCustomImageWriter interface {
	IFPCustomImageHandler
	ImageWrite(Str IStream, Img IFPCustomImage) // procedure
}

// TFPCustomImageWriter Is Abstract Class Parent: TFPCustomImageHandler
type TFPCustomImageWriter struct {
	TFPCustomImageHandler
}

func FPCustomImageWriterClass() TClass {
	ret := LCL().SysCallN(2673)
	return TClass(ret)
}

func (m *TFPCustomImageWriter) ImageWrite(Str IStream, Img IFPCustomImage) {
	LCL().SysCallN(2674, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Img))
}
