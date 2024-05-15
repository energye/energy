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
	r1 := LCL().SysCallN(1307, GetObjectUintptr(TheOwner))
	return AsCustomBitBtn(r1)
}

func (m *TCustomBitBtn) DefaultCaption() bool {
	r1 := LCL().SysCallN(1308, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomBitBtn) SetDefaultCaption(AValue bool) {
	LCL().SysCallN(1308, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomBitBtn) DisabledImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1309, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetDisabledImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1309, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Glyph() IBitmap {
	r1 := LCL().SysCallN(1310, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomBitBtn) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(1310, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) NumGlyphs() int32 {
	r1 := LCL().SysCallN(1323, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(1323, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) HotImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1312, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetHotImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1312, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Images() ICustomImageList {
	r1 := LCL().SysCallN(1315, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomBitBtn) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1315, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1313, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1313, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) ImageWidth() int32 {
	r1 := LCL().SysCallN(1314, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetImageWidth(AValue int32) {
	LCL().SysCallN(1314, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Kind() TBitBtnKind {
	r1 := LCL().SysCallN(1316, 0, m.Instance(), 0)
	return TBitBtnKind(r1)
}

func (m *TCustomBitBtn) SetKind(AValue TBitBtnKind) {
	LCL().SysCallN(1316, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Layout() TButtonLayout {
	r1 := LCL().SysCallN(1317, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TCustomBitBtn) SetLayout(AValue TButtonLayout) {
	LCL().SysCallN(1317, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Margin() int32 {
	r1 := LCL().SysCallN(1322, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetMargin(AValue int32) {
	LCL().SysCallN(1322, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) PressedImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1324, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetPressedImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1324, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Spacing() int32 {
	r1 := LCL().SysCallN(1325, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetSpacing(AValue int32) {
	LCL().SysCallN(1325, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) GlyphShowMode() TGlyphShowMode {
	r1 := LCL().SysCallN(1311, 0, m.Instance(), 0)
	return TGlyphShowMode(r1)
}

func (m *TCustomBitBtn) SetGlyphShowMode(AValue TGlyphShowMode) {
	LCL().SysCallN(1311, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) CanShowGlyph(AWithShowMode bool) bool {
	r1 := LCL().SysCallN(1305, m.Instance(), PascalBool(AWithShowMode))
	return GoBool(r1)
}

func CustomBitBtnClass() TClass {
	ret := LCL().SysCallN(1306)
	return TClass(ret)
}

func (m *TCustomBitBtn) LoadGlyphFromResourceName(Instance THandle, AName string) {
	LCL().SysCallN(1320, m.Instance(), uintptr(Instance), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromLazarusResource(AName string) {
	LCL().SysCallN(1318, m.Instance(), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromStock(idButton int32) {
	LCL().SysCallN(1321, m.Instance(), uintptr(idButton))
}

func (m *TCustomBitBtn) LoadGlyphFromResource(idButton TButtonImage) {
	LCL().SysCallN(1319, m.Instance(), uintptr(idButton))
}
