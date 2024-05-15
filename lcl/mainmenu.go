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

// IMainMenu Parent: IMenu
type IMainMenu interface {
	IMenu
	Height() int32                   // property
	WindowHandle() HWND              // property
	SetWindowHandle(AValue HWND)     // property
	Merge(Menu IMainMenu)            // procedure
	Unmerge(Menu IMainMenu)          // procedure
	SetOnChange(fn TMenuChangeEvent) // property event
}

// TMainMenu Parent: TMenu
type TMainMenu struct {
	TMenu
	changePtr uintptr
}

func NewMainMenu(AOwner IComponent) IMainMenu {
	r1 := LCL().SysCallN(4168, GetObjectUintptr(AOwner))
	return AsMainMenu(r1)
}

func (m *TMainMenu) Height() int32 {
	r1 := LCL().SysCallN(4169, m.Instance())
	return int32(r1)
}

func (m *TMainMenu) WindowHandle() HWND {
	r1 := LCL().SysCallN(4173, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMainMenu) SetWindowHandle(AValue HWND) {
	LCL().SysCallN(4173, 1, m.Instance(), uintptr(AValue))
}

func MainMenuClass() TClass {
	ret := LCL().SysCallN(4167)
	return TClass(ret)
}

func (m *TMainMenu) Merge(Menu IMainMenu) {
	LCL().SysCallN(4170, m.Instance(), GetObjectUintptr(Menu))
}

func (m *TMainMenu) Unmerge(Menu IMainMenu) {
	LCL().SysCallN(4172, m.Instance(), GetObjectUintptr(Menu))
}

func (m *TMainMenu) SetOnChange(fn TMenuChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4171, m.Instance(), m.changePtr)
}
