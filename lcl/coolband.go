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
	r1 := LCL().SysCallN(1152, GetObjectUintptr(aCollection))
	return AsCoolBand(r1)
}

func (m *TCoolBand) Height() int32 {
	r1 := LCL().SysCallN(1155, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Left() int32 {
	r1 := LCL().SysCallN(1159, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Right() int32 {
	r1 := LCL().SysCallN(1164, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Top() int32 {
	r1 := LCL().SysCallN(1166, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Bitmap() IBitmap {
	r1 := LCL().SysCallN(1146, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCoolBand) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(1146, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1147, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCoolBand) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1147, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Break() bool {
	r1 := LCL().SysCallN(1148, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetBreak(AValue bool) {
	LCL().SysCallN(1148, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Color() TColor {
	r1 := LCL().SysCallN(1150, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCoolBand) SetColor(AValue TColor) {
	LCL().SysCallN(1150, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Control() IControl {
	r1 := LCL().SysCallN(1151, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCoolBand) SetControl(AValue IControl) {
	LCL().SysCallN(1151, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) FixedBackground() bool {
	r1 := LCL().SysCallN(1153, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedBackground(AValue bool) {
	LCL().SysCallN(1153, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) FixedSize() bool {
	r1 := LCL().SysCallN(1154, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedSize(AValue bool) {
	LCL().SysCallN(1154, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) HorizontalOnly() bool {
	r1 := LCL().SysCallN(1156, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetHorizontalOnly(AValue bool) {
	LCL().SysCallN(1156, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1157, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCoolBand) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1157, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinHeight() int32 {
	r1 := LCL().SysCallN(1160, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinHeight(AValue int32) {
	LCL().SysCallN(1160, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinWidth() int32 {
	r1 := LCL().SysCallN(1161, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinWidth(AValue int32) {
	LCL().SysCallN(1161, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) ParentColor() bool {
	r1 := LCL().SysCallN(1163, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentColor(AValue bool) {
	LCL().SysCallN(1163, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ParentBitmap() bool {
	r1 := LCL().SysCallN(1162, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentBitmap(AValue bool) {
	LCL().SysCallN(1162, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Text() string {
	r1 := LCL().SysCallN(1165, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoolBand) SetText(AValue string) {
	LCL().SysCallN(1165, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoolBand) Visible() bool {
	r1 := LCL().SysCallN(1167, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetVisible(AValue bool) {
	LCL().SysCallN(1167, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Width() int32 {
	r1 := LCL().SysCallN(1168, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetWidth(AValue int32) {
	LCL().SysCallN(1168, 1, m.Instance(), uintptr(AValue))
}

func CoolBandClass() TClass {
	ret := LCL().SysCallN(1149)
	return TClass(ret)
}

func (m *TCoolBand) AutosizeWidth() {
	LCL().SysCallN(1145, m.Instance())
}

func (m *TCoolBand) InvalidateCoolBar(Sender IObject) {
	LCL().SysCallN(1158, m.Instance(), GetObjectUintptr(Sender))
}
