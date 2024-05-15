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

// IDockTree Parent: IDockManager
type IDockTree interface {
	IDockManager
	DockZoneClass() TDockZoneClass                  // property
	DockSite() IWinControl                          // property
	SetDockSite(AValue IWinControl)                 // property
	RootZone() IDockZone                            // property
	AdjustDockRect(AControl IControl, ARect *TRect) // procedure
	DumpLayout(FileName string)                     // procedure
}

// TDockTree Parent: TDockManager
type TDockTree struct {
	TDockManager
}

func NewDockTree(TheDockSite IWinControl) IDockTree {
	r1 := LCL().SysCallN(2645, GetObjectUintptr(TheDockSite))
	return AsDockTree(r1)
}

func (m *TDockTree) DockZoneClass() TDockZoneClass {
	r1 := LCL().SysCallN(2647, m.Instance())
	return TDockZoneClass(r1)
}

func (m *TDockTree) DockSite() IWinControl {
	r1 := LCL().SysCallN(2646, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TDockTree) SetDockSite(AValue IWinControl) {
	LCL().SysCallN(2646, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDockTree) RootZone() IDockZone {
	r1 := LCL().SysCallN(2649, m.Instance())
	return AsDockZone(r1)
}

func DockTreeClass() TClass {
	ret := LCL().SysCallN(2644)
	return TClass(ret)
}

func (m *TDockTree) AdjustDockRect(AControl IControl, ARect *TRect) {
	var result1 uintptr
	LCL().SysCallN(2643, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)))
	*ARect = *(*TRect)(getPointer(result1))
}

func (m *TDockTree) DumpLayout(FileName string) {
	LCL().SysCallN(2648, m.Instance(), PascalStr(FileName))
}
