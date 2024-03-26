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

// ICustomKeyboardHandler Parent: ICefKeyboardHandler
type ICustomKeyboardHandler interface {
	ICefKeyboardHandler
}

// TCustomKeyboardHandler Parent: TCefKeyboardHandler
type TCustomKeyboardHandler struct {
	TCefKeyboardHandler
}

func NewCustomKeyboardHandler(events IChromiumEvents) ICustomKeyboardHandler {
	r1 := CEF().SysCallN(2147, GetObjectUintptr(events))
	return AsCustomKeyboardHandler(r1)
}
