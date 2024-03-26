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

// IListItems Parent: IPersistent
type IListItems interface {
	IPersistent
	Flags() TListItemsFlags                                                                              // property
	Count() int32                                                                                        // property
	SetCount(AValue int32)                                                                               // property
	Item(AIndex int32) IListItem                                                                         // property
	SetItem(AIndex int32, AValue IListItem)                                                              // property
	Owner() ICustomListView                                                                              // property
	Add() IListItem                                                                                      // function
	FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem // function
	FindData(AData uintptr) IListItem                                                                    // function
	FindData1(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem                           // function
	GetEnumerator() IListItemsEnumerator                                                                 // function
	IndexOf(AItem IListItem) int32                                                                       // function
	Insert(AIndex int32) IListItem                                                                       // function
	AddItem(AItem IListItem)                                                                             // procedure
	BeginUpdate()                                                                                        // procedure
	Clear()                                                                                              // procedure
	Delete(AIndex int32)                                                                                 // procedure
	EndUpdate()                                                                                          // procedure
	Exchange(AIndex1, AIndex2 int32)                                                                     // procedure
	Move(AFromIndex, AToIndex int32)                                                                     // procedure
	InsertItem(AItem IListItem, AIndex int32)                                                            // procedure
}

// TListItems Parent: TPersistent
type TListItems struct {
	TPersistent
}

func NewListItems(AOwner ICustomListView) IListItems {
	r1 := LCL().SysCallN(3416, GetObjectUintptr(AOwner))
	return AsListItems(r1)
}

func (m *TListItems) Flags() TListItemsFlags {
	r1 := LCL().SysCallN(3423, m.Instance())
	return TListItemsFlags(r1)
}

func (m *TListItems) Count() int32 {
	r1 := LCL().SysCallN(3415, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItems) SetCount(AValue int32) {
	LCL().SysCallN(3415, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItems) Item(AIndex int32) IListItem {
	r1 := LCL().SysCallN(3428, 0, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func (m *TListItems) SetItem(AIndex int32, AValue IListItem) {
	LCL().SysCallN(3428, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TListItems) Owner() ICustomListView {
	r1 := LCL().SysCallN(3430, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListItems) Add() IListItem {
	r1 := LCL().SysCallN(3410, m.Instance())
	return AsListItem(r1)
}

func (m *TListItems) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := LCL().SysCallN(3420, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TListItems) FindData(AData uintptr) IListItem {
	r1 := LCL().SysCallN(3421, m.Instance(), uintptr(AData))
	return AsListItem(r1)
}

func (m *TListItems) FindData1(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := LCL().SysCallN(3422, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TListItems) GetEnumerator() IListItemsEnumerator {
	r1 := LCL().SysCallN(3424, m.Instance())
	return AsListItemsEnumerator(r1)
}

func (m *TListItems) IndexOf(AItem IListItem) int32 {
	r1 := LCL().SysCallN(3425, m.Instance(), GetObjectUintptr(AItem))
	return int32(r1)
}

func (m *TListItems) Insert(AIndex int32) IListItem {
	r1 := LCL().SysCallN(3426, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func ListItemsClass() TClass {
	ret := LCL().SysCallN(3413)
	return TClass(ret)
}

func (m *TListItems) AddItem(AItem IListItem) {
	LCL().SysCallN(3411, m.Instance(), GetObjectUintptr(AItem))
}

func (m *TListItems) BeginUpdate() {
	LCL().SysCallN(3412, m.Instance())
}

func (m *TListItems) Clear() {
	LCL().SysCallN(3414, m.Instance())
}

func (m *TListItems) Delete(AIndex int32) {
	LCL().SysCallN(3417, m.Instance(), uintptr(AIndex))
}

func (m *TListItems) EndUpdate() {
	LCL().SysCallN(3418, m.Instance())
}

func (m *TListItems) Exchange(AIndex1, AIndex2 int32) {
	LCL().SysCallN(3419, m.Instance(), uintptr(AIndex1), uintptr(AIndex2))
}

func (m *TListItems) Move(AFromIndex, AToIndex int32) {
	LCL().SysCallN(3429, m.Instance(), uintptr(AFromIndex), uintptr(AToIndex))
}

func (m *TListItems) InsertItem(AItem IListItem, AIndex int32) {
	LCL().SysCallN(3427, m.Instance(), GetObjectUintptr(AItem), uintptr(AIndex))
}
