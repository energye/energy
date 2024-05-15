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
	r1 := LCL().SysCallN(3211, GetObjectUintptr(TheColumn))
	return AsGridColumnTitle(r1)
}

func (m *TGridColumnTitle) Column() IGridColumn {
	r1 := LCL().SysCallN(3210, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumnTitle) Alignment() TAlignment {
	r1 := LCL().SysCallN(3206, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumnTitle) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3206, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Caption() string {
	r1 := LCL().SysCallN(3207, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumnTitle) SetCaption(AValue string) {
	LCL().SysCallN(3207, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumnTitle) Color() TColor {
	r1 := LCL().SysCallN(3209, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumnTitle) SetColor(AValue TColor) {
	LCL().SysCallN(3209, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Font() IFont {
	r1 := LCL().SysCallN(3214, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumnTitle) SetFont(AValue IFont) {
	LCL().SysCallN(3214, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumnTitle) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3215, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TGridColumnTitle) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3215, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) ImageLayout() TButtonLayout {
	r1 := LCL().SysCallN(3216, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TGridColumnTitle) SetImageLayout(AValue TButtonLayout) {
	LCL().SysCallN(3216, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Layout() TTextLayout {
	r1 := LCL().SysCallN(3218, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumnTitle) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(3218, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) MultiLine() bool {
	r1 := LCL().SysCallN(3219, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumnTitle) SetMultiLine(AValue bool) {
	LCL().SysCallN(3219, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumnTitle) PrefixOption() TPrefixOption {
	r1 := LCL().SysCallN(3220, 0, m.Instance(), 0)
	return TPrefixOption(r1)
}

func (m *TGridColumnTitle) SetPrefixOption(AValue TPrefixOption) {
	LCL().SysCallN(3220, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) IsDefault() bool {
	r1 := LCL().SysCallN(3217, m.Instance())
	return GoBool(r1)
}

func GridColumnTitleClass() TClass {
	ret := LCL().SysCallN(3208)
	return TClass(ret)
}

func (m *TGridColumnTitle) FillTitleDefaultFont() {
	LCL().SysCallN(3212, m.Instance())
}

func (m *TGridColumnTitle) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(3213, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumnTitle) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(3221, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}
