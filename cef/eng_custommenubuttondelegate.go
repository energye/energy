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

// ICustomMenuButtonDelegate Parent: ICefMenuButtonDelegate
//
//	This class handles all the ICefMenuButtonDelegate methods which call the ICefMenuButtonDelegateEvents methods.
//	ICefMenuButtonDelegateEvents will be implemented by the control receiving the ICefMenuButtonDelegate events.
type ICustomMenuButtonDelegate interface {
	ICefMenuButtonDelegate
}

// TCustomMenuButtonDelegate Parent: TCefMenuButtonDelegate
//
//	This class handles all the ICefMenuButtonDelegate methods which call the ICefMenuButtonDelegateEvents methods.
//	ICefMenuButtonDelegateEvents will be implemented by the control receiving the ICefMenuButtonDelegate events.
type TCustomMenuButtonDelegate struct {
	TCefMenuButtonDelegate
}

func NewCustomMenuButtonDelegate(events ICefMenuButtonDelegateEvents) ICustomMenuButtonDelegate {
	r1 := CEF().SysCallN(2151, GetObjectUintptr(events))
	return AsCustomMenuButtonDelegate(r1)
}
