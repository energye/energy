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

// IGridColumnTitle Parent: IPersistent
type IGridColumnTitle interface {
	IPersistent
	Column() IGridColumn                             // property
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	Caption() string                                 // property
	SetCaption(AValue string)                        // property
	Color() TColor                                   // property
	SetColor(AValue TColor)                          // property
	Font() IFont                                     // property
	SetFont(AValue IFont)                            // property
	ImageIndex() TImageIndex                         // property
	SetImageIndex(AValue TImageIndex)                // property
	ImageLayout() TButtonLayout                      // property
	SetImageLayout(AValue TButtonLayout)             // property
	Layout() TTextLayout                             // property
	SetLayout(AValue TTextLayout)                    // property
	MultiLine() bool                                 // property
	SetMultiLine(AValue bool)                        // property
	PrefixOption() TPrefixOption                     // property
	SetPrefixOption(AValue TPrefixOption)            // property
	IsDefault() bool                                 // function
	FillTitleDefaultFont()                           // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)          // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64) // procedure
}

// TGridColumnTitle Parent: TPersistent
type TGridColumnTitle struct {
	TPersistent
}

func NewGridColumnTitle(TheColumn IGridColumn) IGridColumnTitle {
	r1 := LCL().SysCallN(2968, GetObjectUintptr(TheColumn))
	return AsGridColumnTitle(r1)
}

func (m *TGridColumnTitle) Column() IGridColumn {
	r1 := LCL().SysCallN(2967, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumnTitle) Alignment() TAlignment {
	r1 := LCL().SysCallN(2963, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumnTitle) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2963, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Caption() string {
	r1 := LCL().SysCallN(2964, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumnTitle) SetCaption(AValue string) {
	LCL().SysCallN(2964, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumnTitle) Color() TColor {
	r1 := LCL().SysCallN(2966, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumnTitle) SetColor(AValue TColor) {
	LCL().SysCallN(2966, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Font() IFont {
	r1 := LCL().SysCallN(2971, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumnTitle) SetFont(AValue IFont) {
	LCL().SysCallN(2971, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumnTitle) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2972, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TGridColumnTitle) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2972, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) ImageLayout() TButtonLayout {
	r1 := LCL().SysCallN(2973, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TGridColumnTitle) SetImageLayout(AValue TButtonLayout) {
	LCL().SysCallN(2973, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Layout() TTextLayout {
	r1 := LCL().SysCallN(2975, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumnTitle) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(2975, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) MultiLine() bool {
	r1 := LCL().SysCallN(2976, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumnTitle) SetMultiLine(AValue bool) {
	LCL().SysCallN(2976, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumnTitle) PrefixOption() TPrefixOption {
	r1 := LCL().SysCallN(2977, 0, m.Instance(), 0)
	return TPrefixOption(r1)
}

func (m *TGridColumnTitle) SetPrefixOption(AValue TPrefixOption) {
	LCL().SysCallN(2977, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) IsDefault() bool {
	r1 := LCL().SysCallN(2974, m.Instance())
	return GoBool(r1)
}

func GridColumnTitleClass() TClass {
	ret := LCL().SysCallN(2965)
	return TClass(ret)
}

func (m *TGridColumnTitle) FillTitleDefaultFont() {
	LCL().SysCallN(2969, m.Instance())
}

func (m *TGridColumnTitle) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(2970, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumnTitle) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(2978, m.Instance(), uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}
