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

// IListControlItem Parent: ICollectionItem
type IListControlItem interface {
	ICollectionItem
	Data() TCustomData                // property
	SetData(AValue TCustomData)       // property
	Caption() string                  // property
	SetCaption(AValue string)         // property
	ImageIndex() TImageIndex          // property
	SetImageIndex(AValue TImageIndex) // property
}

// TListControlItem Parent: TCollectionItem
type TListControlItem struct {
	TCollectionItem
}

func NewListControlItem(ACollection ICollection) IListControlItem {
	r1 := LCL().SysCallN(4006, GetObjectUintptr(ACollection))
	return AsListControlItem(r1)
}

func (m *TListControlItem) Data() TCustomData {
	r1 := LCL().SysCallN(4007, 0, m.Instance(), 0)
	return TCustomData(r1)
}

func (m *TListControlItem) SetData(AValue TCustomData) {
	LCL().SysCallN(4007, 1, m.Instance(), uintptr(AValue))
}

func (m *TListControlItem) Caption() string {
	r1 := LCL().SysCallN(4004, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListControlItem) SetCaption(AValue string) {
	LCL().SysCallN(4004, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListControlItem) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(4008, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListControlItem) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(4008, 1, m.Instance(), uintptr(AValue))
}

func ListControlItemClass() TClass {
	ret := LCL().SysCallN(4005)
	return TClass(ret)
}
