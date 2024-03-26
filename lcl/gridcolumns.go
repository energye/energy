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

// IGridColumns Parent: ICollection
type IGridColumns interface {
	ICollection
	Grid() ICustomGrid                                     // property
	ItemsForGridColumn(Index int32) IGridColumn            // property
	SetItemsForGridColumn(Index int32, AValue IGridColumn) // property
	VisibleCount() int32                                   // property
	Enabled() bool                                         // property
	AddForGridColumn() IGridColumn                         // function
	ColumnByTitle(aTitle string) IGridColumn               // function
	RealIndex(Index int32) int32                           // function
	IndexOf(Column IGridColumn) int32                      // function
	IsDefault() bool                                       // function
	HasIndex(Index int32) bool                             // function
	VisibleIndex(Index int32) int32                        // function
}

// TGridColumns Parent: TCollection
type TGridColumns struct {
	TCollection
}

func NewGridColumns(AGrid ICustomGrid, aItemClass TCollectionItemClass) IGridColumns {
	r1 := LCL().SysCallN(3010, GetObjectUintptr(AGrid), uintptr(aItemClass))
	return AsGridColumns(r1)
}

func (m *TGridColumns) Grid() ICustomGrid {
	r1 := LCL().SysCallN(3012, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumns) ItemsForGridColumn(Index int32) IGridColumn {
	r1 := LCL().SysCallN(3016, 0, m.Instance(), uintptr(Index))
	return AsGridColumn(r1)
}

func (m *TGridColumns) SetItemsForGridColumn(Index int32, AValue IGridColumn) {
	LCL().SysCallN(3016, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TGridColumns) VisibleCount() int32 {
	r1 := LCL().SysCallN(3018, m.Instance())
	return int32(r1)
}

func (m *TGridColumns) Enabled() bool {
	r1 := LCL().SysCallN(3011, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) AddForGridColumn() IGridColumn {
	r1 := LCL().SysCallN(3007, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumns) ColumnByTitle(aTitle string) IGridColumn {
	r1 := LCL().SysCallN(3009, m.Instance(), PascalStr(aTitle))
	return AsGridColumn(r1)
}

func (m *TGridColumns) RealIndex(Index int32) int32 {
	r1 := LCL().SysCallN(3017, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TGridColumns) IndexOf(Column IGridColumn) int32 {
	r1 := LCL().SysCallN(3014, m.Instance(), GetObjectUintptr(Column))
	return int32(r1)
}

func (m *TGridColumns) IsDefault() bool {
	r1 := LCL().SysCallN(3015, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) HasIndex(Index int32) bool {
	r1 := LCL().SysCallN(3013, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TGridColumns) VisibleIndex(Index int32) int32 {
	r1 := LCL().SysCallN(3019, m.Instance(), uintptr(Index))
	return int32(r1)
}

func GridColumnsClass() TClass {
	ret := LCL().SysCallN(3008)
	return TClass(ret)
}
