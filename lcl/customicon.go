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

// ICustomIcon Parent: IRasterImage
type ICustomIcon interface {
	IRasterImage
	Current() int32                                                                  // property
	SetCurrent(AValue int32)                                                         // property
	Count() int32                                                                    // property
	GetBestIndexForSize(ASize *TSize) int32                                          // function
	Add(AFormat TPixelFormat, AHeight, AWidth Word)                                  // procedure
	AssignImage(ASource IRasterImage)                                                // procedure
	Delete(Aindex int32)                                                             // procedure
	Remove(AFormat TPixelFormat, AHeight, AWidth Word)                               // procedure
	GetDescription(Aindex int32, OutFormat *TPixelFormat, OutHeight, OutWidth *Word) // procedure
	SetSize(AWidth, AHeight int32)                                                   // procedure
	LoadFromResourceHandle(Instance THandle, ResHandle TFPResourceHandle)            // procedure
	Sort()                                                                           // procedure
}

// TCustomIcon Parent: TRasterImage
type TCustomIcon struct {
	TRasterImage
}

func NewCustomIcon() ICustomIcon {
	r1 := LCL().SysCallN(1805)
	return AsCustomIcon(r1)
}

func (m *TCustomIcon) Current() int32 {
	r1 := LCL().SysCallN(1806, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomIcon) SetCurrent(AValue int32) {
	LCL().SysCallN(1806, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIcon) Count() int32 {
	r1 := LCL().SysCallN(1804, m.Instance())
	return int32(r1)
}

func (m *TCustomIcon) GetBestIndexForSize(ASize *TSize) int32 {
	r1 := LCL().SysCallN(1808, m.Instance(), uintptr(unsafePointer(ASize)))
	return int32(r1)
}

func CustomIconClass() TClass {
	ret := LCL().SysCallN(1803)
	return TClass(ret)
}

func (m *TCustomIcon) Add(AFormat TPixelFormat, AHeight, AWidth Word) {
	LCL().SysCallN(1801, m.Instance(), uintptr(AFormat), uintptr(AHeight), uintptr(AWidth))
}

func (m *TCustomIcon) AssignImage(ASource IRasterImage) {
	LCL().SysCallN(1802, m.Instance(), GetObjectUintptr(ASource))
}

func (m *TCustomIcon) Delete(Aindex int32) {
	LCL().SysCallN(1807, m.Instance(), uintptr(Aindex))
}

func (m *TCustomIcon) Remove(AFormat TPixelFormat, AHeight, AWidth Word) {
	LCL().SysCallN(1811, m.Instance(), uintptr(AFormat), uintptr(AHeight), uintptr(AWidth))
}

func (m *TCustomIcon) GetDescription(Aindex int32, OutFormat *TPixelFormat, OutHeight, OutWidth *Word) {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	LCL().SysCallN(1809, m.Instance(), uintptr(Aindex), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*OutFormat = TPixelFormat(result1)
	*OutHeight = Word(result2)
	*OutWidth = Word(result3)
}

func (m *TCustomIcon) SetSize(AWidth, AHeight int32) {
	LCL().SysCallN(1812, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}

func (m *TCustomIcon) LoadFromResourceHandle(Instance THandle, ResHandle TFPResourceHandle) {
	LCL().SysCallN(1810, m.Instance(), uintptr(Instance), uintptr(ResHandle))
}

func (m *TCustomIcon) Sort() {
	LCL().SysCallN(1813, m.Instance())
}
