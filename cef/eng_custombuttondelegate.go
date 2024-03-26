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

// ICustomButtonDelegate Parent: ICefButtonDelegate
//
//	This class handles all the ICefButtonDelegate methods which call the ICefButtonDelegateEvents methods.
//	ICefButtonDelegateEvents will be implemented by the control receiving the ICefButtonDelegate events.
type ICustomButtonDelegate interface {
	ICefButtonDelegate
}

// TCustomButtonDelegate Parent: TCefButtonDelegate
//
//	This class handles all the ICefButtonDelegate methods which call the ICefButtonDelegateEvents methods.
//	ICefButtonDelegateEvents will be implemented by the control receiving the ICefButtonDelegate events.
type TCustomButtonDelegate struct {
	TCefButtonDelegate
}

func NewCustomButtonDelegate(events ICefButtonDelegateEvents) ICustomButtonDelegate {
	r1 := CEF().SysCallN(2129, GetObjectUintptr(events))
	return AsCustomButtonDelegate(r1)
}
