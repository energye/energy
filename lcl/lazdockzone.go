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

// ILazDockZone Parent: IDockZone
type ILazDockZone interface {
	IDockZone
	Splitter() ILazDockSplitter          // property
	SetSplitter(AValue ILazDockSplitter) // property
	Pages() ILazDockPages                // property
	SetPages(AValue ILazDockPages)       // property
	Page() ILazDockPage                  // property
	SetPage(AValue ILazDockPage)         // property
	GetCaption() string                  // function
	GetParentControl() IWinControl       // function
	FreeSubComponents()                  // procedure
}

// TLazDockZone Parent: TDockZone
type TLazDockZone struct {
	TDockZone
}

func NewLazDockZone(TheTree IDockTree, TheChildControl IControl) ILazDockZone {
	r1 := LCL().SysCallN(3543, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsLazDockZone(r1)
}

func (m *TLazDockZone) Splitter() ILazDockSplitter {
	r1 := LCL().SysCallN(3549, 0, m.Instance(), 0)
	return AsLazDockSplitter(r1)
}

func (m *TLazDockZone) SetSplitter(AValue ILazDockSplitter) {
	LCL().SysCallN(3549, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Pages() ILazDockPages {
	r1 := LCL().SysCallN(3548, 0, m.Instance(), 0)
	return AsLazDockPages(r1)
}

func (m *TLazDockZone) SetPages(AValue ILazDockPages) {
	LCL().SysCallN(3548, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Page() ILazDockPage {
	r1 := LCL().SysCallN(3547, 0, m.Instance(), 0)
	return AsLazDockPage(r1)
}

func (m *TLazDockZone) SetPage(AValue ILazDockPage) {
	LCL().SysCallN(3547, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) GetCaption() string {
	r1 := LCL().SysCallN(3545, m.Instance())
	return GoStr(r1)
}

func (m *TLazDockZone) GetParentControl() IWinControl {
	r1 := LCL().SysCallN(3546, m.Instance())
	return AsWinControl(r1)
}

func LazDockZoneClass() TClass {
	ret := LCL().SysCallN(3542)
	return TClass(ret)
}

func (m *TLazDockZone) FreeSubComponents() {
	LCL().SysCallN(3544, m.Instance())
}
