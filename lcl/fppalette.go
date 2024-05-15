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

// IFPPalette Parent: IObject
type IFPPalette interface {
	IObject
	Color(Index int32) (resultFPColor TFPColor) // property
	SetColor(Index int32, AValue *TFPColor)     // property
	Count() int32                               // property
	SetCount(AValue int32)                      // property
	Capacity() int32                            // property
	SetCapacity(AValue int32)                   // property
	IndexOf(AColor *TFPColor) int32             // function
	Add(Value *TFPColor) int32                  // function
	Build(Img IFPCustomImage)                   // procedure
	Copy(APalette IFPPalette)                   // procedure
	Merge(pal IFPPalette)                       // procedure
	Clear()                                     // procedure
}

// TFPPalette Parent: TObject
type TFPPalette struct {
	TObject
}

func NewFPPalette(ACount int32) IFPPalette {
	r1 := LCL().SysCallN(2990, uintptr(ACount))
	return AsFPPalette(r1)
}

func (m *TFPPalette) Color(Index int32) (resultFPColor TFPColor) {
	r1 := LCL().SysCallN(2987, 0, m.Instance(), uintptr(Index))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPPalette) SetColor(Index int32, AValue *TFPColor) {
	LCL().SysCallN(2987, 1, m.Instance(), uintptr(Index), uintptr(unsafePointer(AValue)))
}

func (m *TFPPalette) Count() int32 {
	r1 := LCL().SysCallN(2989, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPPalette) SetCount(AValue int32) {
	LCL().SysCallN(2989, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPPalette) Capacity() int32 {
	r1 := LCL().SysCallN(2984, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPPalette) SetCapacity(AValue int32) {
	LCL().SysCallN(2984, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPPalette) IndexOf(AColor *TFPColor) int32 {
	r1 := LCL().SysCallN(2991, m.Instance(), uintptr(unsafePointer(AColor)))
	return int32(r1)
}

func (m *TFPPalette) Add(Value *TFPColor) int32 {
	r1 := LCL().SysCallN(2982, m.Instance(), uintptr(unsafePointer(Value)))
	return int32(r1)
}

func FPPaletteClass() TClass {
	ret := LCL().SysCallN(2985)
	return TClass(ret)
}

func (m *TFPPalette) Build(Img IFPCustomImage) {
	LCL().SysCallN(2983, m.Instance(), GetObjectUintptr(Img))
}

func (m *TFPPalette) Copy(APalette IFPPalette) {
	LCL().SysCallN(2988, m.Instance(), GetObjectUintptr(APalette))
}

func (m *TFPPalette) Merge(pal IFPPalette) {
	LCL().SysCallN(2992, m.Instance(), GetObjectUintptr(pal))
}

func (m *TFPPalette) Clear() {
	LCL().SysCallN(2986, m.Instance())
}
