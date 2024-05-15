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

// ILazDockPages Parent: ICustomTabControl
type ILazDockPages interface {
	ICustomTabControl
	PageForLazDockPage(Index int32) ILazDockPage // property
	ActivePageComponent() ILazDockPage           // property
	SetActivePageComponent(AValue ILazDockPage)  // property
}

// TLazDockPages Parent: TCustomTabControl
type TLazDockPages struct {
	TCustomTabControl
}

func NewLazDockPages(TheOwner IComponent) ILazDockPages {
	r1 := LCL().SysCallN(3531, GetObjectUintptr(TheOwner))
	return AsLazDockPages(r1)
}

func (m *TLazDockPages) PageForLazDockPage(Index int32) ILazDockPage {
	r1 := LCL().SysCallN(3532, m.Instance(), uintptr(Index))
	return AsLazDockPage(r1)
}

func (m *TLazDockPages) ActivePageComponent() ILazDockPage {
	r1 := LCL().SysCallN(3529, 0, m.Instance(), 0)
	return AsLazDockPage(r1)
}

func (m *TLazDockPages) SetActivePageComponent(AValue ILazDockPage) {
	LCL().SysCallN(3529, 1, m.Instance(), GetObjectUintptr(AValue))
}

func LazDockPagesClass() TClass {
	ret := LCL().SysCallN(3530)
	return TClass(ret)
}
