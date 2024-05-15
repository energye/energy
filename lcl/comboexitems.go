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

// IComboExItems Parent: IListControlItems
type IComboExItems interface {
	IListControlItems
	ComboItems(AIndex int32) IComboExItem                                                                                                                       // property
	AddForComboExItem() IComboExItem                                                                                                                            // function
	AddItem(ACaption string, AImageIndex SmallInt, AOverlayImageIndex SmallInt, ASelectedImageIndex SmallInt, AIndent SmallInt, AData TCustomData) IComboExItem // function
	InsertForComboExItem(AIndex int32) IComboExItem                                                                                                             // function
}

// TComboExItems Parent: TListControlItems
type TComboExItems struct {
	TListControlItems
}

func NewComboExItems(AOwner IPersistent, AItemClass TCollectionItemClass) IComboExItems {
	r1 := LCL().SysCallN(858, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsComboExItems(r1)
}

func (m *TComboExItems) ComboItems(AIndex int32) IComboExItem {
	r1 := LCL().SysCallN(857, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddForComboExItem() IComboExItem {
	r1 := LCL().SysCallN(854, m.Instance())
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddItem(ACaption string, AImageIndex SmallInt, AOverlayImageIndex SmallInt, ASelectedImageIndex SmallInt, AIndent SmallInt, AData TCustomData) IComboExItem {
	r1 := LCL().SysCallN(855, m.Instance(), PascalStr(ACaption), uintptr(AImageIndex), uintptr(AOverlayImageIndex), uintptr(ASelectedImageIndex), uintptr(AIndent), uintptr(AData))
	return AsComboExItem(r1)
}

func (m *TComboExItems) InsertForComboExItem(AIndex int32) IComboExItem {
	r1 := LCL().SysCallN(859, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func ComboExItemsClass() TClass {
	ret := LCL().SysCallN(856)
	return TClass(ret)
}
