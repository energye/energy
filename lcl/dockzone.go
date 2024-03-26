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
	r1 := LCL().SysCallN(2413, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsDockZone(r1)
}

func (m *TDockZone) ChildControl() IControl {
	r1 := LCL().SysCallN(2410, m.Instance())
	return AsControl(r1)
}

func (m *TDockZone) ChildCount() int32 {
	r1 := LCL().SysCallN(2411, m.Instance())
	return int32(r1)
}

func (m *TDockZone) FirstChild() IDockZone {
	r1 := LCL().SysCallN(2415, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Height() int32 {
	r1 := LCL().SysCallN(2420, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetHeight(AValue int32) {
	LCL().SysCallN(2420, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Left() int32 {
	r1 := LCL().SysCallN(2421, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLeft(AValue int32) {
	LCL().SysCallN(2421, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitBegin() int32 {
	r1 := LCL().SysCallN(2422, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitBegin(AValue int32) {
	LCL().SysCallN(2422, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitSize() int32 {
	r1 := LCL().SysCallN(2423, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitSize(AValue int32) {
	LCL().SysCallN(2423, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Orientation() TDockOrientation {
	r1 := LCL().SysCallN(2426, 0, m.Instance(), 0)
	return TDockOrientation(r1)
}

func (m *TDockZone) SetOrientation(AValue TDockOrientation) {
	LCL().SysCallN(2426, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Parent() IDockZone {
	r1 := LCL().SysCallN(2427, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Top() int32 {
	r1 := LCL().SysCallN(2432, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetTop(AValue int32) {
	LCL().SysCallN(2432, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Tree() IDockTree {
	r1 := LCL().SysCallN(2433, m.Instance())
	return AsDockTree(r1)
}

func (m *TDockZone) Visible() bool {
	r1 := LCL().SysCallN(2434, m.Instance())
	return GoBool(r1)
}

func (m *TDockZone) VisibleChildCount() int32 {
	r1 := LCL().SysCallN(2435, m.Instance())
	return int32(r1)
}

func (m *TDockZone) Width() int32 {
	r1 := LCL().SysCallN(2436, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetWidth(AValue int32) {
	LCL().SysCallN(2436, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) NextSibling() IDockZone {
	r1 := LCL().SysCallN(2424, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevSibling() IDockZone {
	r1 := LCL().SysCallN(2428, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) FindZone(AControl IControl) IDockZone {
	r1 := LCL().SysCallN(2414, m.Instance(), GetObjectUintptr(AControl))
	return AsDockZone(r1)
}

func (m *TDockZone) FirstVisibleChild() IDockZone {
	r1 := LCL().SysCallN(2416, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetNextVisibleZone() IDockZone {
	r1 := LCL().SysCallN(2419, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) NextVisible() IDockZone {
	r1 := LCL().SysCallN(2425, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevVisible() IDockZone {
	r1 := LCL().SysCallN(2429, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetLastChild() IDockZone {
	r1 := LCL().SysCallN(2418, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetIndex() int32 {
	r1 := LCL().SysCallN(2417, m.Instance())
	return int32(r1)
}

func DockZoneClass() TClass {
	ret := LCL().SysCallN(2412)
	return TClass(ret)
}

func (m *TDockZone) AddSibling(NewZone IDockZone, InsertAt TAlign) {
	LCL().SysCallN(2409, m.Instance(), GetObjectUintptr(NewZone), uintptr(InsertAt))
}

func (m *TDockZone) AddAsFirstChild(NewChildZone IDockZone) {
	LCL().SysCallN(2407, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) AddAsLastChild(NewChildZone IDockZone) {
	LCL().SysCallN(2408, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) ReplaceChild(OldChild, NewChild IDockZone) {
	LCL().SysCallN(2431, m.Instance(), GetObjectUintptr(OldChild), GetObjectUintptr(NewChild))
}

func (m *TDockZone) Remove(ChildZone IDockZone) {
	LCL().SysCallN(2430, m.Instance(), GetObjectUintptr(ChildZone))
}
