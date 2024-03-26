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

// IWVBrowser Parent: IWVBrowserBase
//
//	VCL and LCL version of TWVBrowserBase that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IWVBrowser interface {
	IWVBrowserBase
}

// TWVBrowser Parent: TWVBrowserBase
//
//	VCL and LCL version of TWVBrowserBase that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TWVBrowser struct {
	TWVBrowserBase
}

func NewWVBrowser(AOwner IComponent) IWVBrowser {
	r1 := WV().SysCallN(1051, GetObjectUintptr(AOwner))
	return AsWVBrowser(r1)
}

func WVBrowserClass() TClass {
	ret := WV().SysCallN(1050)
	return TClass(ret)
}
