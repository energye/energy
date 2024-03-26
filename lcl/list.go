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

// IList Parent: IObject
type IList interface {
	IObject
	Capacity() int32                                                                   // property
	SetCapacity(AValue int32)                                                          // property
	Count() int32                                                                      // property
	SetCount(AValue int32)                                                             // property
	Items(Index int32) uintptr                                                         // property
	SetItems(Index int32, AValue uintptr)                                              // property
	List() uintptr                                                                     // property
	Add(Item uintptr) int32                                                            // function
	Expand() IList                                                                     // function
	Extract(item uintptr) uintptr                                                      // function
	First() uintptr                                                                    // function
	GetEnumerator() IListEnumerator                                                    // function
	IndexOf(Item uintptr) int32                                                        // function
	Last() uintptr                                                                     // function
	Remove(Item uintptr) int32                                                         // function
	FPOAttachObserver(AObserver IObject)                                               // procedure
	FPODetachObserver(AObserver IObject)                                               // procedure
	FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) // procedure
	AddList(AList IList)                                                               // procedure
	Clear()                                                                            // procedure
	Delete(Index int32)                                                                // procedure
	Exchange(Index1, Index2 int32)                                                     // procedure
	Insert(Index int32, Item uintptr)                                                  // procedure
	Move(CurIndex, NewIndex int32)                                                     // procedure
	Pack()                                                                             // procedure
	Sort(fn TListSortCompare)                                                          // procedure
}

// TList Parent: TObject
type TList struct {
	TObject
	sortPtr uintptr
}

func NewList() IList {
	r1 := LCL().SysCallN(3506)
	return AsList(r1)
}

func (m *TList) Capacity() int32 {
	r1 := LCL().SysCallN(3502, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TList) SetCapacity(AValue int32) {
	LCL().SysCallN(3502, 1, m.Instance(), uintptr(AValue))
}

func (m *TList) Count() int32 {
	r1 := LCL().SysCallN(3505, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TList) SetCount(AValue int32) {
	LCL().SysCallN(3505, 1, m.Instance(), uintptr(AValue))
}

func (m *TList) Items(Index int32) uintptr {
	r1 := LCL().SysCallN(3518, 0, m.Instance(), uintptr(Index))
	return uintptr(r1)
}

func (m *TList) SetItems(Index int32, AValue uintptr) {
	LCL().SysCallN(3518, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TList) List() uintptr {
	r1 := LCL().SysCallN(3520, m.Instance())
	return uintptr(r1)
}

func (m *TList) Add(Item uintptr) int32 {
	r1 := LCL().SysCallN(3500, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TList) Expand() IList {
	r1 := LCL().SysCallN(3509, m.Instance())
	return AsList(r1)
}

func (m *TList) Extract(item uintptr) uintptr {
	r1 := LCL().SysCallN(3510, m.Instance(), uintptr(item))
	return uintptr(r1)
}

func (m *TList) First() uintptr {
	r1 := LCL().SysCallN(3514, m.Instance())
	return uintptr(r1)
}

func (m *TList) GetEnumerator() IListEnumerator {
	r1 := LCL().SysCallN(3515, m.Instance())
	return AsListEnumerator(r1)
}

func (m *TList) IndexOf(Item uintptr) int32 {
	r1 := LCL().SysCallN(3516, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TList) Last() uintptr {
	r1 := LCL().SysCallN(3519, m.Instance())
	return uintptr(r1)
}

func (m *TList) Remove(Item uintptr) int32 {
	r1 := LCL().SysCallN(3523, m.Instance(), uintptr(Item))
	return int32(r1)
}

func ListClass() TClass {
	ret := LCL().SysCallN(3503)
	return TClass(ret)
}

func (m *TList) FPOAttachObserver(AObserver IObject) {
	LCL().SysCallN(3511, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TList) FPODetachObserver(AObserver IObject) {
	LCL().SysCallN(3512, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TList) FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) {
	LCL().SysCallN(3513, m.Instance(), GetObjectUintptr(ASender), uintptr(AOperation), uintptr(Data))
}

func (m *TList) AddList(AList IList) {
	LCL().SysCallN(3501, m.Instance(), GetObjectUintptr(AList))
}

func (m *TList) Clear() {
	LCL().SysCallN(3504, m.Instance())
}

func (m *TList) Delete(Index int32) {
	LCL().SysCallN(3507, m.Instance(), uintptr(Index))
}

func (m *TList) Exchange(Index1, Index2 int32) {
	LCL().SysCallN(3508, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TList) Insert(Index int32, Item uintptr) {
	LCL().SysCallN(3517, m.Instance(), uintptr(Index), uintptr(Item))
}

func (m *TList) Move(CurIndex, NewIndex int32) {
	LCL().SysCallN(3521, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TList) Pack() {
	LCL().SysCallN(3522, m.Instance())
}

func (m *TList) Sort(fn TListSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3524, m.Instance(), m.sortPtr)
}
