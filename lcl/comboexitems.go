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
	r1 := LCL().SysCallN(668, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsComboExItems(r1)
}

func (m *TComboExItems) ComboItems(AIndex int32) IComboExItem {
	r1 := LCL().SysCallN(667, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddForComboExItem() IComboExItem {
	r1 := LCL().SysCallN(664, m.Instance())
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddItem(ACaption string, AImageIndex SmallInt, AOverlayImageIndex SmallInt, ASelectedImageIndex SmallInt, AIndent SmallInt, AData TCustomData) IComboExItem {
	r1 := LCL().SysCallN(665, m.Instance(), PascalStr(ACaption), uintptr(AImageIndex), uintptr(AOverlayImageIndex), uintptr(ASelectedImageIndex), uintptr(AIndent), uintptr(AData))
	return AsComboExItem(r1)
}

func (m *TComboExItems) InsertForComboExItem(AIndex int32) IComboExItem {
	r1 := LCL().SysCallN(669, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func ComboExItemsClass() TClass {
	ret := LCL().SysCallN(666)
	return TClass(ret)
}
