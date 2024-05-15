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

// IDockZone Parent: IObject
type IDockZone interface {
	IObject
	ChildControl() IControl                        // property
	ChildCount() int32                             // property
	FirstChild() IDockZone                         // property
	Height() int32                                 // property
	SetHeight(AValue int32)                        // property
	Left() int32                                   // property
	SetLeft(AValue int32)                          // property
	LimitBegin() int32                             // property
	SetLimitBegin(AValue int32)                    // property
	LimitSize() int32                              // property
	SetLimitSize(AValue int32)                     // property
	Orientation() TDockOrientation                 // property
	SetOrientation(AValue TDockOrientation)        // property
	Parent() IDockZone                             // property
	Top() int32                                    // property
	SetTop(AValue int32)                           // property
	Tree() IDockTree                               // property
	Visible() bool                                 // property
	VisibleChildCount() int32                      // property
	Width() int32                                  // property
	SetWidth(AValue int32)                         // property
	NextSibling() IDockZone                        // property
	PrevSibling() IDockZone                        // property
	FindZone(AControl IControl) IDockZone          // function
	FirstVisibleChild() IDockZone                  // function
	GetNextVisibleZone() IDockZone                 // function
	NextVisible() IDockZone                        // function
	PrevVisible() IDockZone                        // function
	GetLastChild() IDockZone                       // function
	GetIndex() int32                               // function
	AddSibling(NewZone IDockZone, InsertAt TAlign) // procedure
	AddAsFirstChild(NewChildZone IDockZone)        // procedure
	AddAsLastChild(NewChildZone IDockZone)         // procedure
	ReplaceChild(OldChild, NewChild IDockZone)     // procedure
	Remove(ChildZone IDockZone)                    // procedure
}

// TDockZone Parent: TObject
type TDockZone struct {
	TObject
}

func NewDockZone(TheTree IDockTree, TheChildControl IControl) IDockZone {
	r1 := LCL().SysCallN(2656, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsDockZone(r1)
}

func (m *TDockZone) ChildControl() IControl {
	r1 := LCL().SysCallN(2653, m.Instance())
	return AsControl(r1)
}

func (m *TDockZone) ChildCount() int32 {
	r1 := LCL().SysCallN(2654, m.Instance())
	return int32(r1)
}

func (m *TDockZone) FirstChild() IDockZone {
	r1 := LCL().SysCallN(2658, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Height() int32 {
	r1 := LCL().SysCallN(2663, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetHeight(AValue int32) {
	LCL().SysCallN(2663, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Left() int32 {
	r1 := LCL().SysCallN(2664, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLeft(AValue int32) {
	LCL().SysCallN(2664, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitBegin() int32 {
	r1 := LCL().SysCallN(2665, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitBegin(AValue int32) {
	LCL().SysCallN(2665, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitSize() int32 {
	r1 := LCL().SysCallN(2666, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitSize(AValue int32) {
	LCL().SysCallN(2666, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Orientation() TDockOrientation {
	r1 := LCL().SysCallN(2669, 0, m.Instance(), 0)
	return TDockOrientation(r1)
}

func (m *TDockZone) SetOrientation(AValue TDockOrientation) {
	LCL().SysCallN(2669, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Parent() IDockZone {
	r1 := LCL().SysCallN(2670, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Top() int32 {
	r1 := LCL().SysCallN(2675, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetTop(AValue int32) {
	LCL().SysCallN(2675, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Tree() IDockTree {
	r1 := LCL().SysCallN(2676, m.Instance())
	return AsDockTree(r1)
}

func (m *TDockZone) Visible() bool {
	r1 := LCL().SysCallN(2677, m.Instance())
	return GoBool(r1)
}

func (m *TDockZone) VisibleChildCount() int32 {
	r1 := LCL().SysCallN(2678, m.Instance())
	return int32(r1)
}

func (m *TDockZone) Width() int32 {
	r1 := LCL().SysCallN(2679, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetWidth(AValue int32) {
	LCL().SysCallN(2679, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) NextSibling() IDockZone {
	r1 := LCL().SysCallN(2667, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevSibling() IDockZone {
	r1 := LCL().SysCallN(2671, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) FindZone(AControl IControl) IDockZone {
	r1 := LCL().SysCallN(2657, m.Instance(), GetObjectUintptr(AControl))
	return AsDockZone(r1)
}

func (m *TDockZone) FirstVisibleChild() IDockZone {
	r1 := LCL().SysCallN(2659, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetNextVisibleZone() IDockZone {
	r1 := LCL().SysCallN(2662, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) NextVisible() IDockZone {
	r1 := LCL().SysCallN(2668, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevVisible() IDockZone {
	r1 := LCL().SysCallN(2672, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetLastChild() IDockZone {
	r1 := LCL().SysCallN(2661, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetIndex() int32 {
	r1 := LCL().SysCallN(2660, m.Instance())
	return int32(r1)
}

func DockZoneClass() TClass {
	ret := LCL().SysCallN(2655)
	return TClass(ret)
}

func (m *TDockZone) AddSibling(NewZone IDockZone, InsertAt TAlign) {
	LCL().SysCallN(2652, m.Instance(), GetObjectUintptr(NewZone), uintptr(InsertAt))
}

func (m *TDockZone) AddAsFirstChild(NewChildZone IDockZone) {
	LCL().SysCallN(2650, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) AddAsLastChild(NewChildZone IDockZone) {
	LCL().SysCallN(2651, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) ReplaceChild(OldChild, NewChild IDockZone) {
	LCL().SysCallN(2674, m.Instance(), GetObjectUintptr(OldChild), GetObjectUintptr(NewChild))
}

func (m *TDockZone) Remove(ChildZone IDockZone) {
	LCL().SysCallN(2673, m.Instance(), GetObjectUintptr(ChildZone))
}
