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

// ITextAttributes Parent: IPersistent
type ITextAttributes interface {
	IPersistent
	Name() string                   // property
	SetName(AValue string)          // property
	Pitch() TFontPitch              // property
	SetPitch(AValue TFontPitch)     // property
	Charset() TFontCharSet          // property
	SetCharset(AValue TFontCharSet) // property
	Color() TColor                  // property
	SetColor(AValue TColor)         // property
	Size() int32                    // property
	SetSize(AValue int32)           // property
	Style() TFontStyles             // property
	SetStyle(AValue TFontStyles)    // property
	Height() int32                  // property
	SetHeight(AValue int32)         // property
}

// TTextAttributes Parent: TPersistent
type TTextAttributes struct {
	TPersistent
}

func NewTextAttributes(AOwner IRichMemo, AType TAttributeType) ITextAttributes {
	r1 := LCL().SysCallN(5370, GetObjectUintptr(AOwner), uintptr(AType))
	return AsTextAttributes(r1)
}

func (m *TTextAttributes) Name() string {
	r1 := LCL().SysCallN(5372, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTextAttributes) SetName(AValue string) {
	LCL().SysCallN(5372, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTextAttributes) Pitch() TFontPitch {
	r1 := LCL().SysCallN(5373, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TTextAttributes) SetPitch(AValue TFontPitch) {
	LCL().SysCallN(5373, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Charset() TFontCharSet {
	r1 := LCL().SysCallN(5367, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TTextAttributes) SetCharset(AValue TFontCharSet) {
	LCL().SysCallN(5367, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Color() TColor {
	r1 := LCL().SysCallN(5369, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTextAttributes) SetColor(AValue TColor) {
	LCL().SysCallN(5369, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Size() int32 {
	r1 := LCL().SysCallN(5374, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetSize(AValue int32) {
	LCL().SysCallN(5374, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Style() TFontStyles {
	r1 := LCL().SysCallN(5375, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TTextAttributes) SetStyle(AValue TFontStyles) {
	LCL().SysCallN(5375, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Height() int32 {
	r1 := LCL().SysCallN(5371, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetHeight(AValue int32) {
	LCL().SysCallN(5371, 1, m.Instance(), uintptr(AValue))
}

func TextAttributesClass() TClass {
	ret := LCL().SysCallN(5368)
	return TClass(ret)
}
