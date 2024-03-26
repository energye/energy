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

// ICoolBand Parent: ICollectionItem
type ICoolBand interface {
	ICollectionItem
	Height() int32                      // property
	Left() int32                        // property
	Right() int32                       // property
	Top() int32                         // property
	Bitmap() IBitmap                    // property
	SetBitmap(AValue IBitmap)           // property
	BorderStyle() TBorderStyle          // property
	SetBorderStyle(AValue TBorderStyle) // property
	Break() bool                        // property
	SetBreak(AValue bool)               // property
	Color() TColor                      // property
	SetColor(AValue TColor)             // property
	Control() IControl                  // property
	SetControl(AValue IControl)         // property
	FixedBackground() bool              // property
	SetFixedBackground(AValue bool)     // property
	FixedSize() bool                    // property
	SetFixedSize(AValue bool)           // property
	HorizontalOnly() bool               // property
	SetHorizontalOnly(AValue bool)      // property
	ImageIndex() TImageIndex            // property
	SetImageIndex(AValue TImageIndex)   // property
	MinHeight() int32                   // property
	SetMinHeight(AValue int32)          // property
	MinWidth() int32                    // property
	SetMinWidth(AValue int32)           // property
	ParentColor() bool                  // property
	SetParentColor(AValue bool)         // property
	ParentBitmap() bool                 // property
	SetParentBitmap(AValue bool)        // property
	Text() string                       // property
	SetText(AValue string)              // property
	Visible() bool                      // property
	SetVisible(AValue bool)             // property
	Width() int32                       // property
	SetWidth(AValue int32)              // property
	AutosizeWidth()                     // procedure
	InvalidateCoolBar(Sender IObject)   // procedure
}

// TCoolBand Parent: TCollectionItem
type TCoolBand struct {
	TCollectionItem
}

func NewCoolBand(aCollection ICollection) ICoolBand {
	r1 := LCL().SysCallN(962, GetObjectUintptr(aCollection))
	return AsCoolBand(r1)
}

func (m *TCoolBand) Height() int32 {
	r1 := LCL().SysCallN(965, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Left() int32 {
	r1 := LCL().SysCallN(969, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Right() int32 {
	r1 := LCL().SysCallN(974, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Top() int32 {
	r1 := LCL().SysCallN(976, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Bitmap() IBitmap {
	r1 := LCL().SysCallN(956, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCoolBand) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(956, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(957, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCoolBand) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(957, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Break() bool {
	r1 := LCL().SysCallN(958, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetBreak(AValue bool) {
	LCL().SysCallN(958, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Color() TColor {
	r1 := LCL().SysCallN(960, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCoolBand) SetColor(AValue TColor) {
	LCL().SysCallN(960, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Control() IControl {
	r1 := LCL().SysCallN(961, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCoolBand) SetControl(AValue IControl) {
	LCL().SysCallN(961, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) FixedBackground() bool {
	r1 := LCL().SysCallN(963, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedBackground(AValue bool) {
	LCL().SysCallN(963, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) FixedSize() bool {
	r1 := LCL().SysCallN(964, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedSize(AValue bool) {
	LCL().SysCallN(964, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) HorizontalOnly() bool {
	r1 := LCL().SysCallN(966, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetHorizontalOnly(AValue bool) {
	LCL().SysCallN(966, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(967, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCoolBand) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(967, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinHeight() int32 {
	r1 := LCL().SysCallN(970, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinHeight(AValue int32) {
	LCL().SysCallN(970, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinWidth() int32 {
	r1 := LCL().SysCallN(971, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinWidth(AValue int32) {
	LCL().SysCallN(971, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) ParentColor() bool {
	r1 := LCL().SysCallN(973, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentColor(AValue bool) {
	LCL().SysCallN(973, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ParentBitmap() bool {
	r1 := LCL().SysCallN(972, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentBitmap(AValue bool) {
	LCL().SysCallN(972, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Text() string {
	r1 := LCL().SysCallN(975, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoolBand) SetText(AValue string) {
	LCL().SysCallN(975, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoolBand) Visible() bool {
	r1 := LCL().SysCallN(977, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetVisible(AValue bool) {
	LCL().SysCallN(977, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Width() int32 {
	r1 := LCL().SysCallN(978, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetWidth(AValue int32) {
	LCL().SysCallN(978, 1, m.Instance(), uintptr(AValue))
}

func CoolBandClass() TClass {
	ret := LCL().SysCallN(959)
	return TClass(ret)
}

func (m *TCoolBand) AutosizeWidth() {
	LCL().SysCallN(955, m.Instance())
}

func (m *TCoolBand) InvalidateCoolBar(Sender IObject) {
	LCL().SysCallN(968, m.Instance(), GetObjectUintptr(Sender))
}
