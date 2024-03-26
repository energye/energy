//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICustomBrowserViewDelegate Parent: ICefBrowserViewDelegate
//
//	This class handles all the ICefBrowserViewDelegate methods which call the ICefBrowserViewDelegateEvents methods.
//	ICefBrowserViewDelegateEvents will be implemented by the control receiving the ICefBrowserViewDelegate events.
type ICustomBrowserViewDelegate interface {
	ICefBrowserViewDelegate
}

// TCustomBrowserViewDelegate Parent: TCefBrowserViewDelegate
//
//	This class handles all the ICefBrowserViewDelegate methods which call the ICefBrowserViewDelegateEvents methods.
//	ICefBrowserViewDelegateEvents will be implemented by the control receiving the ICefBrowserViewDelegate events.
type TCustomBrowserViewDelegate struct {
	TCefBrowserViewDelegate
}

func NewCustomBrowserViewDelegate(events ICefBrowserViewDelegateEvents) ICustomBrowserViewDelegate {
	r1 := CEF().SysCallN(2128, GetObjectUintptr(events))
	return AsCustomBrowserViewDelegate(r1)
}
