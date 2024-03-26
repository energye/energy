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
	r1 := LCL().SysCallN(2983, GetObjectUintptr(ACollection))
	return AsGridColumn(r1)
}

func (m *TGridColumn) Grid() ICustomGrid {
	r1 := LCL().SysCallN(2990, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumn) DefaultWidth() int32 {
	r1 := LCL().SysCallN(2984, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) StoredWidth() int32 {
	r1 := LCL().SysCallN(2999, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) WidthChanged() bool {
	r1 := LCL().SysCallN(3006, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumn) Alignment() TAlignment {
	r1 := LCL().SysCallN(2979, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumn) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2979, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) ButtonStyle() TColumnButtonStyle {
	r1 := LCL().SysCallN(2980, 0, m.Instance(), 0)
	return TColumnButtonStyle(r1)
}

func (m *TGridColumn) SetButtonStyle(AValue TColumnButtonStyle) {
	LCL().SysCallN(2980, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Color() TColor {
	r1 := LCL().SysCallN(2982, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumn) SetColor(AValue TColor) {
	LCL().SysCallN(2982, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) DropDownRows() int32 {
	r1 := LCL().SysCallN(2985, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetDropDownRows(AValue int32) {
	LCL().SysCallN(2985, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Expanded() bool {
	r1 := LCL().SysCallN(2986, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetExpanded(AValue bool) {
	LCL().SysCallN(2986, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) Font() IFont {
	r1 := LCL().SysCallN(2989, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumn) SetFont(AValue IFont) {
	LCL().SysCallN(2989, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Layout() TTextLayout {
	r1 := LCL().SysCallN(2992, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumn) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(2992, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MinSize() int32 {
	r1 := LCL().SysCallN(2994, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMinSize(AValue int32) {
	LCL().SysCallN(2994, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MaxSize() int32 {
	r1 := LCL().SysCallN(2993, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMaxSize(AValue int32) {
	LCL().SysCallN(2993, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) PickList() IStrings {
	r1 := LCL().SysCallN(2995, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TGridColumn) SetPickList(AValue IStrings) {
	LCL().SysCallN(2995, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) ReadOnly() bool {
	r1 := LCL().SysCallN(2996, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetReadOnly(AValue bool) {
	LCL().SysCallN(2996, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) SizePriority() int32 {
	r1 := LCL().SysCallN(2998, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetSizePriority(AValue int32) {
	LCL().SysCallN(2998, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Tag() uint32 {
	r1 := LCL().SysCallN(3000, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TGridColumn) SetTag(AValue uint32) {
	LCL().SysCallN(3000, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Title() IGridColumnTitle {
	r1 := LCL().SysCallN(3001, 0, m.Instance(), 0)
	return AsGridColumnTitle(r1)
}

func (m *TGridColumn) SetTitle(AValue IGridColumnTitle) {
	LCL().SysCallN(3001, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Width() int32 {
	r1 := LCL().SysCallN(3005, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetWidth(AValue int32) {
	LCL().SysCallN(3005, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Visible() bool {
	r1 := LCL().SysCallN(3004, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetVisible(AValue bool) {
	LCL().SysCallN(3004, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) ValueChecked() string {
	r1 := LCL().SysCallN(3002, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueChecked(AValue string) {
	LCL().SysCallN(3002, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) ValueUnchecked() string {
	r1 := LCL().SysCallN(3003, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueUnchecked(AValue string) {
	LCL().SysCallN(3003, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) IsDefault() bool {
	r1 := LCL().SysCallN(2991, m.Instance())
	return GoBool(r1)
}

func GridColumnClass() TClass {
	ret := LCL().SysCallN(2981)
	return TClass(ret)
}

func (m *TGridColumn) FillDefaultFont() {
	LCL().SysCallN(2987, m.Instance())
}

func (m *TGridColumn) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(2988, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumn) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(2997, m.Instance(), uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}
