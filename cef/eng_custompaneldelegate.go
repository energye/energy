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

// ICustomPanelDelegate Parent: ICefPanelDelegate
//
//	This class handles all the ICefPanelDelegate methods which call the ICefPanelDelegateEvents methods.
//	ICefPanelDelegateEvents will be implemented by the control receiving the ICefPanelDelegate events.
type ICustomPanelDelegate interface {
	ICefPanelDelegate
}

// TCustomPanelDelegate Parent: TCefPanelDelegate
//
//	This class handles all the ICefPanelDelegate methods which call the ICefPanelDelegateEvents methods.
//	ICefPanelDelegateEvents will be implemented by the control receiving the ICefPanelDelegate events.
type TCustomPanelDelegate struct {
	TCefPanelDelegate
}

func NewCustomPanelDelegate(events ICefPanelDelegateEvents) ICustomPanelDelegate {
	r1 := CEF().SysCallN(2152, GetObjectUintptr(events))
	return AsCustomPanelDelegate(r1)
}
