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

// IFPList Parent: IObject
type IFPList interface {
	IObject
	Capacity() int32                      // property
	SetCapacity(AValue int32)             // property
	Count() int32                         // property
	SetCount(AValue int32)                // property
	Items(Index int32) uintptr            // property
	SetItems(Index int32, AValue uintptr) // property
	List() uintptr                        // property
	Add(Item uintptr) int32               // function
	Expand() IFPList                      // function
	Extract(Item uintptr) uintptr         // function
	First() uintptr                       // function
	GetEnumerator() IFPListEnumerator     // function
	IndexOf(Item uintptr) int32           // function
	Last() uintptr                        // function
	Remove(Item uintptr) int32            // function
	AddList(AList IFPList)                // procedure
	Clear()                               // procedure
	Delete(Index int32)                   // procedure
	Exchange(Index1, Index2 int32)        // procedure
	Insert(Index int32, Item uintptr)     // procedure
	Move(CurIndex, NewIndex int32)        // procedure
	Pack()                                // procedure
	Sort(fn TListSortCompare)             // procedure
}

// TFPList Parent: TObject
type TFPList struct {
	TObject
	sortPtr uintptr
}

func NewFPList() IFPList {
	r1 := LCL().SysCallN(2721)
	return AsFPList(r1)
}

func (m *TFPList) Capacity() int32 {
	r1 := LCL().SysCallN(2717, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCapacity(AValue int32) {
	LCL().SysCallN(2717, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Count() int32 {
	r1 := LCL().SysCallN(2720, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCount(AValue int32) {
	LCL().SysCallN(2720, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Items(Index int32) uintptr {
	r1 := LCL().SysCallN(2730, 0, m.Instance(), uintptr(Index))
	return uintptr(r1)
}

func (m *TFPList) SetItems(Index int32, AValue uintptr) {
	LCL().SysCallN(2730, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TFPList) List() uintptr {
	r1 := LCL().SysCallN(2732, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Add(Item uintptr) int32 {
	r1 := LCL().SysCallN(2715, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Expand() IFPList {
	r1 := LCL().SysCallN(2724, m.Instance())
	return AsFPList(r1)
}

func (m *TFPList) Extract(Item uintptr) uintptr {
	r1 := LCL().SysCallN(2725, m.Instance(), uintptr(Item))
	return uintptr(r1)
}

func (m *TFPList) First() uintptr {
	r1 := LCL().SysCallN(2726, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) GetEnumerator() IFPListEnumerator {
	r1 := LCL().SysCallN(2727, m.Instance())
	return AsFPListEnumerator(r1)
}

func (m *TFPList) IndexOf(Item uintptr) int32 {
	r1 := LCL().SysCallN(2728, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Last() uintptr {
	r1 := LCL().SysCallN(2731, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Remove(Item uintptr) int32 {
	r1 := LCL().SysCallN(2735, m.Instance(), uintptr(Item))
	return int32(r1)
}

func FPListClass() TClass {
	ret := LCL().SysCallN(2718)
	return TClass(ret)
}

func (m *TFPList) AddList(AList IFPList) {
	LCL().SysCallN(2716, m.Instance(), GetObjectUintptr(AList))
}

func (m *TFPList) Clear() {
	LCL().SysCallN(2719, m.Instance())
}

func (m *TFPList) Delete(Index int32) {
	LCL().SysCallN(2722, m.Instance(), uintptr(Index))
}

func (m *TFPList) Exchange(Index1, Index2 int32) {
	LCL().SysCallN(2723, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TFPList) Insert(Index int32, Item uintptr) {
	LCL().SysCallN(2729, m.Instance(), uintptr(Index), uintptr(Item))
}

func (m *TFPList) Move(CurIndex, NewIndex int32) {
	LCL().SysCallN(2733, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TFPList) Pack() {
	LCL().SysCallN(2734, m.Instance())
}

func (m *TFPList) Sort(fn TListSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2736, m.Instance(), m.sortPtr)
}
