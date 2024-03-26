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

// IFPCustomFont Parent: IFPCanvasHelper
type IFPCustomFont interface {
	IFPCanvasHelper
	Name() string                         // property
	SetName(AValue string)                // property
	Size() int32                          // property
	SetSize(AValue int32)                 // property
	Bold() bool                           // property
	SetBold(AValue bool)                  // property
	Italic() bool                         // property
	SetItalic(AValue bool)                // property
	Underline() bool                      // property
	SetUnderline(AValue bool)             // property
	StrikeThrough() bool                  // property
	SetStrikeThrough(AValue bool)         // property
	Orientation() int32                   // property
	SetOrientation(AValue int32)          // property
	CopyFont() IFPCustomFont              // function
	GetTextHeight(text string) int32      // function
	GetTextWidth(text string) int32       // function
	GetTextSize(text string, w, h *int32) // procedure
}

// TFPCustomFont Parent: TFPCanvasHelper
type TFPCustomFont struct {
	TFPCanvasHelper
}

func NewFPCustomFont() IFPCustomFont {
	r1 := LCL().SysCallN(2656)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) Name() string {
	r1 := LCL().SysCallN(2661, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFPCustomFont) SetName(AValue string) {
	LCL().SysCallN(2661, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFPCustomFont) Size() int32 {
	r1 := LCL().SysCallN(2663, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetSize(AValue int32) {
	LCL().SysCallN(2663, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) Bold() bool {
	r1 := LCL().SysCallN(2653, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetBold(AValue bool) {
	LCL().SysCallN(2653, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Italic() bool {
	r1 := LCL().SysCallN(2660, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetItalic(AValue bool) {
	LCL().SysCallN(2660, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Underline() bool {
	r1 := LCL().SysCallN(2665, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetUnderline(AValue bool) {
	LCL().SysCallN(2665, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) StrikeThrough() bool {
	r1 := LCL().SysCallN(2664, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetStrikeThrough(AValue bool) {
	LCL().SysCallN(2664, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Orientation() int32 {
	r1 := LCL().SysCallN(2662, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetOrientation(AValue int32) {
	LCL().SysCallN(2662, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) CopyFont() IFPCustomFont {
	r1 := LCL().SysCallN(2655, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) GetTextHeight(text string) int32 {
	r1 := LCL().SysCallN(2657, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomFont) GetTextWidth(text string) int32 {
	r1 := LCL().SysCallN(2659, m.Instance(), PascalStr(text))
	return int32(r1)
}

func FPCustomFontClass() TClass {
	ret := LCL().SysCallN(2654)
	return TClass(ret)
}

func (m *TFPCustomFont) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2658, m.Instance(), PascalStr(text), uintptr(unsafe.Pointer(&result1)), uintptr(unsafe.Pointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}
