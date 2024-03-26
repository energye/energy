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

// ICustomViewDelegate Parent: ICefViewDelegate
//
//	This class handles all the ICefViewDelegate methods which call the ICefViewDelegateEvents methods.
//	ICefViewDelegateEvents will be implemented by the control receiving the ICefViewDelegate events.
type ICustomViewDelegate interface {
	ICefViewDelegate
}

// TCustomViewDelegate Parent: TCefViewDelegate
//
//	This class handles all the ICefViewDelegate methods which call the ICefViewDelegateEvents methods.
//	ICefViewDelegateEvents will be implemented by the control receiving the ICefViewDelegate events.
type TCustomViewDelegate struct {
	TCefViewDelegate
}

func NewCustomViewDelegate(events ICefViewDelegateEvents) ICustomViewDelegate {
	r1 := CEF().SysCallN(2168, GetObjectUintptr(events))
	return AsCustomViewDelegate(r1)
}
