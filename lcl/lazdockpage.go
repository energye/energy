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

// ILazDockPage Parent: ICustomPage
type ILazDockPage interface {
	ICustomPage
	DockZone() IDockZone        // property
	PageControl() ILazDockPages // property
}

// TLazDockPage Parent: TCustomPage
type TLazDockPage struct {
	TCustomPage
}

func NewLazDockPage(TheOwner IComponent) ILazDockPage {
	r1 := LCL().SysCallN(3526, GetObjectUintptr(TheOwner))
	return AsLazDockPage(r1)
}

func (m *TLazDockPage) DockZone() IDockZone {
	r1 := LCL().SysCallN(3527, m.Instance())
	return AsDockZone(r1)
}

func (m *TLazDockPage) PageControl() ILazDockPages {
	r1 := LCL().SysCallN(3528, m.Instance())
	return AsLazDockPages(r1)
}

func LazDockPageClass() TClass {
	ret := LCL().SysCallN(3525)
	return TClass(ret)
}
