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
	r1 := LCL().SysCallN(4058, GetObjectUintptr(AOwner))
	return AsListItems(r1)
}

func (m *TListItems) Flags() TListItemsFlags {
	r1 := LCL().SysCallN(4065, m.Instance())
	return TListItemsFlags(r1)
}

func (m *TListItems) Count() int32 {
	r1 := LCL().SysCallN(4057, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItems) SetCount(AValue int32) {
	LCL().SysCallN(4057, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItems) Item(AIndex int32) IListItem {
	r1 := LCL().SysCallN(4070, 0, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func (m *TListItems) SetItem(AIndex int32, AValue IListItem) {
	LCL().SysCallN(4070, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TListItems) Owner() ICustomListView {
	r1 := LCL().SysCallN(4072, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListItems) Add() IListItem {
	r1 := LCL().SysCallN(4052, m.Instance())
	return AsListItem(r1)
}

func (m *TListItems) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := LCL().SysCallN(4062, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TListItems) FindData(AData uintptr) IListItem {
	r1 := LCL().SysCallN(4063, m.Instance(), uintptr(AData))
	return AsListItem(r1)
}

func (m *TListItems) FindData1(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := LCL().SysCallN(4064, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TListItems) GetEnumerator() IListItemsEnumerator {
	r1 := LCL().SysCallN(4066, m.Instance())
	return AsListItemsEnumerator(r1)
}

func (m *TListItems) IndexOf(AItem IListItem) int32 {
	r1 := LCL().SysCallN(4067, m.Instance(), GetObjectUintptr(AItem))
	return int32(r1)
}

func (m *TListItems) Insert(AIndex int32) IListItem {
	r1 := LCL().SysCallN(4068, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func ListItemsClass() TClass {
	ret := LCL().SysCallN(4055)
	return TClass(ret)
}

func (m *TListItems) AddItem(AItem IListItem) {
	LCL().SysCallN(4053, m.Instance(), GetObjectUintptr(AItem))
}

func (m *TListItems) BeginUpdate() {
	LCL().SysCallN(4054, m.Instance())
}

func (m *TListItems) Clear() {
	LCL().SysCallN(4056, m.Instance())
}

func (m *TListItems) Delete(AIndex int32) {
	LCL().SysCallN(4059, m.Instance(), uintptr(AIndex))
}

func (m *TListItems) EndUpdate() {
	LCL().SysCallN(4060, m.Instance())
}

func (m *TListItems) Exchange(AIndex1, AIndex2 int32) {
	LCL().SysCallN(4061, m.Instance(), uintptr(AIndex1), uintptr(AIndex2))
}

func (m *TListItems) Move(AFromIndex, AToIndex int32) {
	LCL().SysCallN(4071, m.Instance(), uintptr(AFromIndex), uintptr(AToIndex))
}

func (m *TListItems) InsertItem(AItem IListItem, AIndex int32) {
	LCL().SysCallN(4069, m.Instance(), GetObjectUintptr(AItem), uintptr(AIndex))
}
