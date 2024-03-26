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

// ICustomWindowDelegate Parent: ICefWindowDelegate
//
//	This class handles all the TCustomWindowDelegate methods which call the ICefWindowDelegateEvents methods.
//	ICefWindowDelegateEvents will be implemented by the control receiving the TCustomWindowDelegate events.
type ICustomWindowDelegate interface {
	ICefWindowDelegate
}

// TCustomWindowDelegate Parent: TCefWindowDelegate
//
//	This class handles all the TCustomWindowDelegate methods which call the ICefWindowDelegateEvents methods.
//	ICefWindowDelegateEvents will be implemented by the control receiving the TCustomWindowDelegate events.
type TCustomWindowDelegate struct {
	TCefWindowDelegate
}

func NewCustomWindowDelegate(events ICefWindowDelegateEvents) ICustomWindowDelegate {
	r1 := CEF().SysCallN(2169, GetObjectUintptr(events))
	return AsCustomWindowDelegate(r1)
}
