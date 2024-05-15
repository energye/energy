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

// IGridColumn Parent: ICollectionItem
type IGridColumn interface {
	ICollectionItem
	Grid() ICustomGrid                               // property
	DefaultWidth() int32                             // property
	StoredWidth() int32                              // property
	WidthChanged() bool                              // property
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	ButtonStyle() TColumnButtonStyle                 // property
	SetButtonStyle(AValue TColumnButtonStyle)        // property
	Color() TColor                                   // property
	SetColor(AValue TColor)                          // property
	DropDownRows() int32                             // property
	SetDropDownRows(AValue int32)                    // property
	Expanded() bool                                  // property
	SetExpanded(AValue bool)                         // property
	Font() IFont                                     // property
	SetFont(AValue IFont)                            // property
	Layout() TTextLayout                             // property
	SetLayout(AValue TTextLayout)                    // property
	MinSize() int32                                  // property
	SetMinSize(AValue int32)                         // property
	MaxSize() int32                                  // property
	SetMaxSize(AValue int32)                         // property
	PickList() IStrings                              // property
	SetPickList(AValue IStrings)                     // property
	ReadOnly() bool                                  // property
	SetReadOnly(AValue bool)                         // property
	SizePriority() int32                             // property
	SetSizePriority(AValue int32)                    // property
	Tag() uint32                                     // property
	SetTag(AValue uint32)                            // property
	Title() IGridColumnTitle                         // property
	SetTitle(AValue IGridColumnTitle)                // property
	Width() int32                                    // property
	SetWidth(AValue int32)                           // property
	Visible() bool                                   // property
	SetVisible(AValue bool)                          // property
	ValueChecked() string                            // property
	SetValueChecked(AValue string)                   // property
	ValueUnchecked() string                          // property
	SetValueUnchecked(AValue string)                 // property
	IsDefault() bool                                 // function
	FillDefaultFont()                                // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)          // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64) // procedure
}

// TGridColumn Parent: TCollectionItem
type TGridColumn struct {
	TCollectionItem
}

func NewGridColumn(ACollection ICollection) IGridColumn {
	r1 := LCL().SysCallN(3226, GetObjectUintptr(ACollection))
	return AsGridColumn(r1)
}

func (m *TGridColumn) Grid() ICustomGrid {
	r1 := LCL().SysCallN(3233, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumn) DefaultWidth() int32 {
	r1 := LCL().SysCallN(3227, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) StoredWidth() int32 {
	r1 := LCL().SysCallN(3242, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) WidthChanged() bool {
	r1 := LCL().SysCallN(3249, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumn) Alignment() TAlignment {
	r1 := LCL().SysCallN(3222, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumn) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3222, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) ButtonStyle() TColumnButtonStyle {
	r1 := LCL().SysCallN(3223, 0, m.Instance(), 0)
	return TColumnButtonStyle(r1)
}

func (m *TGridColumn) SetButtonStyle(AValue TColumnButtonStyle) {
	LCL().SysCallN(3223, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Color() TColor {
	r1 := LCL().SysCallN(3225, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumn) SetColor(AValue TColor) {
	LCL().SysCallN(3225, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) DropDownRows() int32 {
	r1 := LCL().SysCallN(3228, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetDropDownRows(AValue int32) {
	LCL().SysCallN(3228, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Expanded() bool {
	r1 := LCL().SysCallN(3229, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetExpanded(AValue bool) {
	LCL().SysCallN(3229, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) Font() IFont {
	r1 := LCL().SysCallN(3232, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumn) SetFont(AValue IFont) {
	LCL().SysCallN(3232, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Layout() TTextLayout {
	r1 := LCL().SysCallN(3235, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumn) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(3235, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MinSize() int32 {
	r1 := LCL().SysCallN(3237, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMinSize(AValue int32) {
	LCL().SysCallN(3237, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MaxSize() int32 {
	r1 := LCL().SysCallN(3236, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMaxSize(AValue int32) {
	LCL().SysCallN(3236, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) PickList() IStrings {
	r1 := LCL().SysCallN(3238, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TGridColumn) SetPickList(AValue IStrings) {
	LCL().SysCallN(3238, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) ReadOnly() bool {
	r1 := LCL().SysCallN(3239, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetReadOnly(AValue bool) {
	LCL().SysCallN(3239, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) SizePriority() int32 {
	r1 := LCL().SysCallN(3241, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetSizePriority(AValue int32) {
	LCL().SysCallN(3241, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Tag() uint32 {
	r1 := LCL().SysCallN(3243, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TGridColumn) SetTag(AValue uint32) {
	LCL().SysCallN(3243, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Title() IGridColumnTitle {
	r1 := LCL().SysCallN(3244, 0, m.Instance(), 0)
	return AsGridColumnTitle(r1)
}

func (m *TGridColumn) SetTitle(AValue IGridColumnTitle) {
	LCL().SysCallN(3244, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Width() int32 {
	r1 := LCL().SysCallN(3248, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetWidth(AValue int32) {
	LCL().SysCallN(3248, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Visible() bool {
	r1 := LCL().SysCallN(3247, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetVisible(AValue bool) {
	LCL().SysCallN(3247, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) ValueChecked() string {
	r1 := LCL().SysCallN(3245, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueChecked(AValue string) {
	LCL().SysCallN(3245, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) ValueUnchecked() string {
	r1 := LCL().SysCallN(3246, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueUnchecked(AValue string) {
	LCL().SysCallN(3246, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) IsDefault() bool {
	r1 := LCL().SysCallN(3234, m.Instance())
	return GoBool(r1)
}

func GridColumnClass() TClass {
	ret := LCL().SysCallN(3224)
	return TClass(ret)
}

func (m *TGridColumn) FillDefaultFont() {
	LCL().SysCallN(3230, m.Instance())
}

func (m *TGridColumn) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(3231, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumn) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(3240, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}
