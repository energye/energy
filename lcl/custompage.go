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

// ICustomPage Parent: IWinControl
type ICustomPage interface {
	IWinControl
	PageIndex() int32                 // property
	SetPageIndex(AValue int32)        // property
	TabVisible() bool                 // property
	SetTabVisible(AValue bool)        // property
	ImageIndex() TImageIndex          // property
	SetImageIndex(AValue TImageIndex) // property
	CanTab() bool                     // function
	VisibleIndex() int32              // function
	SetOnHide(fn TNotifyEvent)        // property event
	SetOnShow(fn TNotifyEvent)        // property event
}

// TCustomPage Parent: TWinControl
type TCustomPage struct {
	TWinControl
	hidePtr uintptr
	showPtr uintptr
}

func NewCustomPage(TheOwner IComponent) ICustomPage {
	r1 := LCL().SysCallN(1919, GetObjectUintptr(TheOwner))
	return AsCustomPage(r1)
}

func (m *TCustomPage) PageIndex() int32 {
	r1 := LCL().SysCallN(1921, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPage) SetPageIndex(AValue int32) {
	LCL().SysCallN(1921, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPage) TabVisible() bool {
	r1 := LCL().SysCallN(1924, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPage) SetTabVisible(AValue bool) {
	LCL().SysCallN(1924, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPage) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1920, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomPage) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1920, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPage) CanTab() bool {
	r1 := LCL().SysCallN(1917, m.Instance())
	return GoBool(r1)
}

func (m *TCustomPage) VisibleIndex() int32 {
	r1 := LCL().SysCallN(1925, m.Instance())
	return int32(r1)
}

func CustomPageClass() TClass {
	ret := LCL().SysCallN(1918)
	return TClass(ret)
}

func (m *TCustomPage) SetOnHide(fn TNotifyEvent) {
	if m.hidePtr != 0 {
		RemoveEventElement(m.hidePtr)
	}
	m.hidePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1922, m.Instance(), m.hidePtr)
}

func (m *TCustomPage) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1923, m.Instance(), m.showPtr)
}
