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
	"unsafe"
)

// ICustomLabel Parent: IGraphicControl
type ICustomLabel interface {
	IGraphicControl
	CalcFittingFontHeight(TheText string, MaxWidth, MaxHeight int32, OutFontHeight, OutNeededWidth, OutNeededHeight *int32) bool // function
	AdjustFontForOptimalFill() bool                                                                                              // function
	Paint()                                                                                                                      // procedure
}

// TCustomLabel Parent: TGraphicControl
type TCustomLabel struct {
	TGraphicControl
}

func NewCustomLabel(TheOwner IComponent) ICustomLabel {
	r1 := LCL().SysCallN(1792, GetObjectUintptr(TheOwner))
	return AsCustomLabel(r1)
}

func (m *TCustomLabel) CalcFittingFontHeight(TheText string, MaxWidth, MaxHeight int32, OutFontHeight, OutNeededWidth, OutNeededHeight *int32) bool {
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := LCL().SysCallN(1790, m.Instance(), PascalStr(TheText), uintptr(MaxWidth), uintptr(MaxHeight), uintptr(unsafe.Pointer(&result2)), uintptr(unsafe.Pointer(&result3)), uintptr(unsafe.Pointer(&result4)))
	*OutFontHeight = int32(result2)
	*OutNeededWidth = int32(result3)
	*OutNeededHeight = int32(result4)
	return GoBool(r1)
}

func (m *TCustomLabel) AdjustFontForOptimalFill() bool {
	r1 := LCL().SysCallN(1789, m.Instance())
	return GoBool(r1)
}

func CustomLabelClass() TClass {
	ret := LCL().SysCallN(1791)
	return TClass(ret)
}

func (m *TCustomLabel) Paint() {
	LCL().SysCallN(1793, m.Instance())
}
