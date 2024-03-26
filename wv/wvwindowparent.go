//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IWVWindowParent Parent: IWVWinControl
//
//	Parent control used by VCL and LCL applications to show the web contents.
type IWVWindowParent interface {
	IWVWinControl
	// Browser
	//  Browser associated to this control to show web contents.
	Browser() IWVBrowserBase // property
	// SetBrowser Set Browser
	SetBrowser(AValue IWVBrowserBase) // property
}

// TWVWindowParent Parent: TWVWinControl
//
//	Parent control used by VCL and LCL applications to show the web contents.
type TWVWindowParent struct {
	TWVWinControl
}

func NewWVWindowParent(AOwner IComponent) IWVWindowParent {
	r1 := WV().SysCallN(1140, GetObjectUintptr(AOwner))
	return AsWVWindowParent(r1)
}

func (m *TWVWindowParent) Browser() IWVBrowserBase {
	var resultWVBrowserBase uintptr
	WV().SysCallN(1138, 0, m.Instance(), 0, uintptr(unsafePointer(&resultWVBrowserBase)))
	return AsWVBrowserBase(resultWVBrowserBase)
}

func (m *TWVWindowParent) SetBrowser(AValue IWVBrowserBase) {
	WV().SysCallN(1138, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func WVWindowParentClass() TClass {
	ret := WV().SysCallN(1139)
	return TClass(ret)
}
