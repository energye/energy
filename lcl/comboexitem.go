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

// IComboExItem Parent: IListControlItem
type IComboExItem interface {
	IListControlItem
	Indent() int32                            // property
	SetIndent(AValue int32)                   // property
	OverlayImageIndex() TImageIndex           // property
	SetOverlayImageIndex(AValue TImageIndex)  // property
	SelectedImageIndex() TImageIndex          // property
	SetSelectedImageIndex(AValue TImageIndex) // property
}

// TComboExItem Parent: TListControlItem
type TComboExItem struct {
	TListControlItem
}

func NewComboExItem(ACollection ICollection) IComboExItem {
	r1 := LCL().SysCallN(660, GetObjectUintptr(ACollection))
	return AsComboExItem(r1)
}

func (m *TComboExItem) Indent() int32 {
	r1 := LCL().SysCallN(661, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboExItem) SetIndent(AValue int32) {
	LCL().SysCallN(661, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboExItem) OverlayImageIndex() TImageIndex {
	r1 := LCL().SysCallN(662, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TComboExItem) SetOverlayImageIndex(AValue TImageIndex) {
	LCL().SysCallN(662, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboExItem) SelectedImageIndex() TImageIndex {
	r1 := LCL().SysCallN(663, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TComboExItem) SetSelectedImageIndex(AValue TImageIndex) {
	LCL().SysCallN(663, 1, m.Instance(), uintptr(AValue))
}

func ComboExItemClass() TClass {
	ret := LCL().SysCallN(659)
	return TClass(ret)
}
