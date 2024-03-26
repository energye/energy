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

// IListColumn Parent: ICollectionItem
type IListColumn interface {
	ICollectionItem
	WidthType() int32                       // property
	Alignment() TAlignment                  // property
	SetAlignment(AValue TAlignment)         // property
	AutoSize() bool                         // property
	SetAutoSize(AValue bool)                // property
	Caption() string                        // property
	SetCaption(AValue string)               // property
	ImageIndex() TImageIndex                // property
	SetImageIndex(AValue TImageIndex)       // property
	MaxWidth() int32                        // property
	SetMaxWidth(AValue int32)               // property
	MinWidth() int32                        // property
	SetMinWidth(AValue int32)               // property
	Tag() uint32                            // property
	SetTag(AValue uint32)                   // property
	Visible() bool                          // property
	SetVisible(AValue bool)                 // property
	Width() int32                           // property
	SetWidth(AValue int32)                  // property
	SortIndicator() TSortIndicator          // property
	SetSortIndicator(AValue TSortIndicator) // property
}

// TListColumn Parent: TCollectionItem
type TListColumn struct {
	TCollectionItem
}

func NewListColumn(ACollection ICollection) IListColumn {
	r1 := LCL().SysCallN(3347, GetObjectUintptr(ACollection))
	return AsListColumn(r1)
}

func (m *TListColumn) WidthType() int32 {
	r1 := LCL().SysCallN(3355, m.Instance())
	return int32(r1)
}

func (m *TListColumn) Alignment() TAlignment {
	r1 := LCL().SysCallN(3343, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TListColumn) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3343, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) AutoSize() bool {
	r1 := LCL().SysCallN(3344, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListColumn) SetAutoSize(AValue bool) {
	LCL().SysCallN(3344, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListColumn) Caption() string {
	r1 := LCL().SysCallN(3345, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListColumn) SetCaption(AValue string) {
	LCL().SysCallN(3345, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListColumn) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3348, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListColumn) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3348, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) MaxWidth() int32 {
	r1 := LCL().SysCallN(3349, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetMaxWidth(AValue int32) {
	LCL().SysCallN(3349, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) MinWidth() int32 {
	r1 := LCL().SysCallN(3350, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetMinWidth(AValue int32) {
	LCL().SysCallN(3350, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) Tag() uint32 {
	r1 := LCL().SysCallN(3352, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TListColumn) SetTag(AValue uint32) {
	LCL().SysCallN(3352, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) Visible() bool {
	r1 := LCL().SysCallN(3353, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListColumn) SetVisible(AValue bool) {
	LCL().SysCallN(3353, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListColumn) Width() int32 {
	r1 := LCL().SysCallN(3354, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetWidth(AValue int32) {
	LCL().SysCallN(3354, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) SortIndicator() TSortIndicator {
	r1 := LCL().SysCallN(3351, 0, m.Instance(), 0)
	return TSortIndicator(r1)
}

func (m *TListColumn) SetSortIndicator(AValue TSortIndicator) {
	LCL().SysCallN(3351, 1, m.Instance(), uintptr(AValue))
}

func ListColumnClass() TClass {
	ret := LCL().SysCallN(3346)
	return TClass(ret)
}
