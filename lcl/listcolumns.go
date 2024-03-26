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

// IListColumns Parent: ICollection
type IListColumns interface {
	ICollection
	OwnerForCustomListView() ICustomListView                // property
	ItemsForListColumn(AIndex int32) IListColumn            // property
	SetItemsForListColumn(AIndex int32, AValue IListColumn) // property
	AddForListColumn() IListColumn                          // function
	Update(Item ICollectionItem)                            // procedure
}

// TListColumns Parent: TCollection
type TListColumns struct {
	TCollection
}

func NewListColumns(AOwner ICustomListView) IListColumns {
	r1 := LCL().SysCallN(3358, GetObjectUintptr(AOwner))
	return AsListColumns(r1)
}

func (m *TListColumns) OwnerForCustomListView() ICustomListView {
	r1 := LCL().SysCallN(3360, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListColumns) ItemsForListColumn(AIndex int32) IListColumn {
	r1 := LCL().SysCallN(3359, 0, m.Instance(), uintptr(AIndex))
	return AsListColumn(r1)
}

func (m *TListColumns) SetItemsForListColumn(AIndex int32, AValue IListColumn) {
	LCL().SysCallN(3359, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TListColumns) AddForListColumn() IListColumn {
	r1 := LCL().SysCallN(3356, m.Instance())
	return AsListColumn(r1)
}

func ListColumnsClass() TClass {
	ret := LCL().SysCallN(3357)
	return TClass(ret)
}

func (m *TListColumns) Update(Item ICollectionItem) {
	LCL().SysCallN(3361, m.Instance(), GetObjectUintptr(Item))
}
