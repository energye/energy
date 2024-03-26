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

// ICustomSpeedButton Parent: IGraphicControl
type ICustomSpeedButton interface {
	IGraphicControl
	Alignment() TAlignment                                    // property
	SetAlignment(AValue TAlignment)                           // property
	AllowAllUp() bool                                         // property
	SetAllowAllUp(AValue bool)                                // property
	DisabledImageIndex() TImageIndex                          // property
	SetDisabledImageIndex(AValue TImageIndex)                 // property
	Down() bool                                               // property
	SetDown(AValue bool)                                      // property
	Flat() bool                                               // property
	SetFlat(AValue bool)                                      // property
	Glyph() IBitmap                                           // property
	SetGlyph(AValue IBitmap)                                  // property
	GroupIndex() int32                                        // property
	SetGroupIndex(AValue int32)                               // property
	HotImageIndex() TImageIndex                               // property
	SetHotImageIndex(AValue TImageIndex)                      // property
	Images() ICustomImageList                                 // property
	SetImages(AValue ICustomImageList)                        // property
	ImageIndex() TImageIndex                                  // property
	SetImageIndex(AValue TImageIndex)                         // property
	ImageWidth() int32                                        // property
	SetImageWidth(AValue int32)                               // property
	Layout() TButtonLayout                                    // property
	SetLayout(AValue TButtonLayout)                           // property
	Margin() int32                                            // property
	SetMargin(AValue int32)                                   // property
	NumGlyphs() int32                                         // property
	SetNumGlyphs(AValue int32)                                // property
	PressedImageIndex() TImageIndex                           // property
	SetPressedImageIndex(AValue TImageIndex)                  // property
	SelectedImageIndex() TImageIndex                          // property
	SetSelectedImageIndex(AValue TImageIndex)                 // property
	ShowAccelChar() bool                                      // property
	SetShowAccelChar(AValue bool)                             // property
	ShowCaption() bool                                        // property
	SetShowCaption(AValue bool)                               // property
	Spacing() int32                                           // property
	SetSpacing(AValue int32)                                  // property
	Transparent() bool                                        // property
	SetTransparent(AValue bool)                               // property
	FindDownButton() ICustomSpeedButton                       // function
	Click()                                                   // procedure
	LoadGlyphFromResourceName(Instance THandle, AName string) // procedure
	LoadGlyphFromLazarusResource(AName string)                // procedure
}

// TCustomSpeedButton Parent: TGraphicControl
type TCustomSpeedButton struct {
	TGraphicControl
}

func NewCustomSpeedButton(AOwner IComponent) ICustomSpeedButton {
	r1 := LCL().SysCallN(2030, GetObjectUintptr(AOwner))
	return AsCustomSpeedButton(r1)
}

func (m *TCustomSpeedButton) Alignment() TAlignment {
	r1 := LCL().SysCallN(2026, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomSpeedButton) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2026, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) AllowAllUp() bool {
	r1 := LCL().SysCallN(2027, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetAllowAllUp(AValue bool) {
	LCL().SysCallN(2027, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) DisabledImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2031, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetDisabledImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2031, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Down() bool {
	r1 := LCL().SysCallN(2032, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetDown(AValue bool) {
	LCL().SysCallN(2032, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Flat() bool {
	r1 := LCL().SysCallN(2034, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetFlat(AValue bool) {
	LCL().SysCallN(2034, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Glyph() IBitmap {
	r1 := LCL().SysCallN(2035, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomSpeedButton) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(2035, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSpeedButton) GroupIndex() int32 {
	r1 := LCL().SysCallN(2036, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetGroupIndex(AValue int32) {
	LCL().SysCallN(2036, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) HotImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2037, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetHotImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2037, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Images() ICustomImageList {
	r1 := LCL().SysCallN(2040, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomSpeedButton) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2040, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSpeedButton) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2038, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2038, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) ImageWidth() int32 {
	r1 := LCL().SysCallN(2039, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetImageWidth(AValue int32) {
	LCL().SysCallN(2039, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Layout() TButtonLayout {
	r1 := LCL().SysCallN(2041, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TCustomSpeedButton) SetLayout(AValue TButtonLayout) {
	LCL().SysCallN(2041, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Margin() int32 {
	r1 := LCL().SysCallN(2044, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetMargin(AValue int32) {
	LCL().SysCallN(2044, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) NumGlyphs() int32 {
	r1 := LCL().SysCallN(2045, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(2045, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) PressedImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2046, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetPressedImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2046, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) SelectedImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2047, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetSelectedImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2047, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) ShowAccelChar() bool {
	r1 := LCL().SysCallN(2048, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(2048, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) ShowCaption() bool {
	r1 := LCL().SysCallN(2049, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(2049, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Spacing() int32 {
	r1 := LCL().SysCallN(2050, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetSpacing(AValue int32) {
	LCL().SysCallN(2050, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Transparent() bool {
	r1 := LCL().SysCallN(2051, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetTransparent(AValue bool) {
	LCL().SysCallN(2051, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) FindDownButton() ICustomSpeedButton {
	r1 := LCL().SysCallN(2033, m.Instance())
	return AsCustomSpeedButton(r1)
}

func CustomSpeedButtonClass() TClass {
	ret := LCL().SysCallN(2028)
	return TClass(ret)
}

func (m *TCustomSpeedButton) Click() {
	LCL().SysCallN(2029, m.Instance())
}

func (m *TCustomSpeedButton) LoadGlyphFromResourceName(Instance THandle, AName string) {
	LCL().SysCallN(2043, m.Instance(), uintptr(Instance), PascalStr(AName))
}

func (m *TCustomSpeedButton) LoadGlyphFromLazarusResource(AName string) {
	LCL().SysCallN(2042, m.Instance(), PascalStr(AName))
}
