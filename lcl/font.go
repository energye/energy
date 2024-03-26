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

// IFont Parent: IFPCustomFont
type IFont interface {
	IFPCustomFont
	FontData() (resultFontData TFontData) // property
	SetFontData(AValue *TFontData)        // property
	Handle() HFONT                        // property
	SetHandle(AValue HFONT)               // property
	IsMonoSpace() bool                    // property
	PixelsPerInch() int32                 // property
	SetPixelsPerInch(AValue int32)        // property
	CharSet() TFontCharSet                // property
	SetCharSet(AValue TFontCharSet)       // property
	Color() TColor                        // property
	SetColor(AValue TColor)               // property
	Height() int32                        // property
	SetHeight(AValue int32)               // property
	Pitch() TFontPitch                    // property
	SetPitch(AValue TFontPitch)           // property
	Quality() TFontQuality                // property
	SetQuality(AValue TFontQuality)       // property
	Style() TFontStyles                   // property
	SetStyle(AValue TFontStyles)          // property
	HandleAllocated() bool                // function
	IsDefault() bool                      // function
	IsEqual(AFont IFont) bool             // function
	BeginUpdate()                         // procedure
	EndUpdate()                           // procedure
	SetDefault()                          // procedure
}

// TFont Parent: TFPCustomFont
type TFont struct {
	TFPCustomFont
}

func NewFont() IFont {
	r1 := LCL().SysCallN(2841)
	return AsFont(r1)
}

func (m *TFont) FontData() (resultFontData TFontData) {
	r1 := LCL().SysCallN(2843, 0, m.Instance(), 0)
	return *(*TFontData)(getPointer(r1))
}

func (m *TFont) SetFontData(AValue *TFontData) {
	LCL().SysCallN(2843, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFont) Handle() HFONT {
	r1 := LCL().SysCallN(2844, 0, m.Instance(), 0)
	return HFONT(r1)
}

func (m *TFont) SetHandle(AValue HFONT) {
	LCL().SysCallN(2844, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) IsMonoSpace() bool {
	r1 := LCL().SysCallN(2849, m.Instance())
	return GoBool(r1)
}

func (m *TFont) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(2851, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetPixelsPerInch(AValue int32) {
	LCL().SysCallN(2851, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) CharSet() TFontCharSet {
	r1 := LCL().SysCallN(2838, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TFont) SetCharSet(AValue TFontCharSet) {
	LCL().SysCallN(2838, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Color() TColor {
	r1 := LCL().SysCallN(2840, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TFont) SetColor(AValue TColor) {
	LCL().SysCallN(2840, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Height() int32 {
	r1 := LCL().SysCallN(2846, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetHeight(AValue int32) {
	LCL().SysCallN(2846, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Pitch() TFontPitch {
	r1 := LCL().SysCallN(2850, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TFont) SetPitch(AValue TFontPitch) {
	LCL().SysCallN(2850, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Quality() TFontQuality {
	r1 := LCL().SysCallN(2852, 0, m.Instance(), 0)
	return TFontQuality(r1)
}

func (m *TFont) SetQuality(AValue TFontQuality) {
	LCL().SysCallN(2852, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Style() TFontStyles {
	r1 := LCL().SysCallN(2854, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TFont) SetStyle(AValue TFontStyles) {
	LCL().SysCallN(2854, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) HandleAllocated() bool {
	r1 := LCL().SysCallN(2845, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsDefault() bool {
	r1 := LCL().SysCallN(2847, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsEqual(AFont IFont) bool {
	r1 := LCL().SysCallN(2848, m.Instance(), GetObjectUintptr(AFont))
	return GoBool(r1)
}

func FontClass() TClass {
	ret := LCL().SysCallN(2839)
	return TClass(ret)
}

func (m *TFont) BeginUpdate() {
	LCL().SysCallN(2837, m.Instance())
}

func (m *TFont) EndUpdate() {
	LCL().SysCallN(2842, m.Instance())
}

func (m *TFont) SetDefault() {
	LCL().SysCallN(2853, m.Instance())
}
