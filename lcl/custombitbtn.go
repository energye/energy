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

// ICustomBitBtn Parent: ICustomButton
type ICustomBitBtn interface {
	ICustomButton
	DefaultCaption() bool                                     // property
	SetDefaultCaption(AValue bool)                            // property
	DisabledImageIndex() TImageIndex                          // property
	SetDisabledImageIndex(AValue TImageIndex)                 // property
	Glyph() IBitmap                                           // property
	SetGlyph(AValue IBitmap)                                  // property
	NumGlyphs() int32                                         // property
	SetNumGlyphs(AValue int32)                                // property
	HotImageIndex() TImageIndex                               // property
	SetHotImageIndex(AValue TImageIndex)                      // property
	Images() ICustomImageList                                 // property
	SetImages(AValue ICustomImageList)                        // property
	ImageIndex() TImageIndex                                  // property
	SetImageIndex(AValue TImageIndex)                         // property
	ImageWidth() int32                                        // property
	SetImageWidth(AValue int32)                               // property
	Kind() TBitBtnKind                                        // property
	SetKind(AValue TBitBtnKind)                               // property
	Layout() TButtonLayout                                    // property
	SetLayout(AValue TButtonLayout)                           // property
	Margin() int32                                            // property
	SetMargin(AValue int32)                                   // property
	PressedImageIndex() TImageIndex                           // property
	SetPressedImageIndex(AValue TImageIndex)                  // property
	Spacing() int32                                           // property
	SetSpacing(AValue int32)                                  // property
	GlyphShowMode() TGlyphShowMode                            // property
	SetGlyphShowMode(AValue TGlyphShowMode)                   // property
	CanShowGlyph(AWithShowMode bool) bool                     // function
	LoadGlyphFromResourceName(Instance THandle, AName string) // procedure
	LoadGlyphFromLazarusResource(AName string)                // procedure
	LoadGlyphFromStock(idButton int32)                        // procedure
	LoadGlyphFromResource(idButton TButtonImage)              // procedure
}

// TCustomBitBtn Parent: TCustomButton
type TCustomBitBtn struct {
	TCustomButton
}

func NewCustomBitBtn(TheOwner IComponent) ICustomBitBtn {
	r1 := LCL().SysCallN(1117, GetObjectUintptr(TheOwner))
	return AsCustomBitBtn(r1)
}

func (m *TCustomBitBtn) DefaultCaption() bool {
	r1 := LCL().SysCallN(1118, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomBitBtn) SetDefaultCaption(AValue bool) {
	LCL().SysCallN(1118, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomBitBtn) DisabledImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1119, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetDisabledImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1119, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Glyph() IBitmap {
	r1 := LCL().SysCallN(1120, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomBitBtn) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(1120, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) NumGlyphs() int32 {
	r1 := LCL().SysCallN(1133, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(1133, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) HotImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1122, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetHotImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1122, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Images() ICustomImageList {
	r1 := LCL().SysCallN(1125, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomBitBtn) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1125, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1123, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1123, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) ImageWidth() int32 {
	r1 := LCL().SysCallN(1124, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetImageWidth(AValue int32) {
	LCL().SysCallN(1124, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Kind() TBitBtnKind {
	r1 := LCL().SysCallN(1126, 0, m.Instance(), 0)
	return TBitBtnKind(r1)
}

func (m *TCustomBitBtn) SetKind(AValue TBitBtnKind) {
	LCL().SysCallN(1126, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Layout() TButtonLayout {
	r1 := LCL().SysCallN(1127, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TCustomBitBtn) SetLayout(AValue TButtonLayout) {
	LCL().SysCallN(1127, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Margin() int32 {
	r1 := LCL().SysCallN(1132, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetMargin(AValue int32) {
	LCL().SysCallN(1132, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) PressedImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1134, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetPressedImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1134, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Spacing() int32 {
	r1 := LCL().SysCallN(1135, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetSpacing(AValue int32) {
	LCL().SysCallN(1135, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) GlyphShowMode() TGlyphShowMode {
	r1 := LCL().SysCallN(1121, 0, m.Instance(), 0)
	return TGlyphShowMode(r1)
}

func (m *TCustomBitBtn) SetGlyphShowMode(AValue TGlyphShowMode) {
	LCL().SysCallN(1121, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) CanShowGlyph(AWithShowMode bool) bool {
	r1 := LCL().SysCallN(1115, m.Instance(), PascalBool(AWithShowMode))
	return GoBool(r1)
}

func CustomBitBtnClass() TClass {
	ret := LCL().SysCallN(1116)
	return TClass(ret)
}

func (m *TCustomBitBtn) LoadGlyphFromResourceName(Instance THandle, AName string) {
	LCL().SysCallN(1130, m.Instance(), uintptr(Instance), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromLazarusResource(AName string) {
	LCL().SysCallN(1128, m.Instance(), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromStock(idButton int32) {
	LCL().SysCallN(1131, m.Instance(), uintptr(idButton))
}

func (m *TCustomBitBtn) LoadGlyphFromResource(idButton TButtonImage) {
	LCL().SysCallN(1129, m.Instance(), uintptr(idButton))
}
