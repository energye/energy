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

// ICollection Parent: IPersistent
type ICollection interface {
	IPersistent
	Count() int32                                 // property
	ItemClass() TCollectionItemClass              // property
	Items(Index int32) ICollectionItem            // property
	SetItems(Index int32, AValue ICollectionItem) // property
	Owner() IPersistent                           // function
	Add() ICollectionItem                         // function
	GetEnumerator() ICollectionEnumerator         // function
	Insert(Index int32) ICollectionItem           // function
	FindItemID(ID int32) ICollectionItem          // function
	BeginUpdate()                                 // procedure
	Clear()                                       // procedure
	EndUpdate()                                   // procedure
	Delete(Index int32)                           // procedure
	Exchange(Index1, index2 int32)                // procedure
	Move(Index1, index2 int32)                    // procedure
	Sort(fn TCollectionSortCompare)               // procedure
}

// TCollection Parent: TPersistent
type TCollection struct {
	TPersistent
	sortPtr uintptr
}

func NewCollection(AItemClass TCollectionItemClass) ICollection {
	r1 := LCL().SysCallN(703, uintptr(AItemClass))
	return AsCollection(r1)
}

func (m *TCollection) Count() int32 {
	r1 := LCL().SysCallN(702, m.Instance())
	return int32(r1)
}

func (m *TCollection) ItemClass() TCollectionItemClass {
	r1 := LCL().SysCallN(710, m.Instance())
	return TCollectionItemClass(r1)
}

func (m *TCollection) Items(Index int32) ICollectionItem {
	r1 := LCL().SysCallN(711, 0, m.Instance(), uintptr(Index))
	return AsCollectionItem(r1)
}

func (m *TCollection) SetItems(Index int32, AValue ICollectionItem) {
	LCL().SysCallN(711, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCollection) Owner() IPersistent {
	r1 := LCL().SysCallN(713, m.Instance())
	return AsPersistent(r1)
}

func (m *TCollection) Add() ICollectionItem {
	r1 := LCL().SysCallN(698, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollection) GetEnumerator() ICollectionEnumerator {
	r1 := LCL().SysCallN(708, m.Instance())
	return AsCollectionEnumerator(r1)
}

func (m *TCollection) Insert(Index int32) ICollectionItem {
	r1 := LCL().SysCallN(709, m.Instance(), uintptr(Index))
	return AsCollectionItem(r1)
}

func (m *TCollection) FindItemID(ID int32) ICollectionItem {
	r1 := LCL().SysCallN(707, m.Instance(), uintptr(ID))
	return AsCollectionItem(r1)
}

func CollectionClass() TClass {
	ret := LCL().SysCallN(700)
	return TClass(ret)
}

func (m *TCollection) BeginUpdate() {
	LCL().SysCallN(699, m.Instance())
}

func (m *TCollection) Clear() {
	LCL().SysCallN(701, m.Instance())
}

func (m *TCollection) EndUpdate() {
	LCL().SysCallN(705, m.Instance())
}

func (m *TCollection) Delete(Index int32) {
	LCL().SysCallN(704, m.Instance(), uintptr(Index))
}

func (m *TCollection) Exchange(Index1, index2 int32) {
	LCL().SysCallN(706, m.Instance(), uintptr(Index1), uintptr(index2))
}

func (m *TCollection) Move(Index1, index2 int32) {
	LCL().SysCallN(712, m.Instance(), uintptr(Index1), uintptr(index2))
}

func (m *TCollection) Sort(fn TCollectionSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(714, m.Instance(), m.sortPtr)
}
