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

// ICustomFocusHandler Parent: ICefFocusHandler
type ICustomFocusHandler interface {
	ICefFocusHandler
}

// TCustomFocusHandler Parent: TCefFocusHandler
type TCustomFocusHandler struct {
	TCefFocusHandler
}

func NewCustomFocusHandler(events IChromiumEvents) ICustomFocusHandler {
	r1 := CEF().SysCallN(2144, GetObjectUintptr(events))
	return AsCustomFocusHandler(r1)
}
