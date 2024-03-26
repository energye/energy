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

// ICustomJsDialogHandler Parent: ICefJsDialogHandler
type ICustomJsDialogHandler interface {
	ICefJsDialogHandler
}

// TCustomJsDialogHandler Parent: TCefJsDialogHandler
type TCustomJsDialogHandler struct {
	TCefJsDialogHandler
}

func NewCustomJsDialogHandler(events IChromiumEvents) ICustomJsDialogHandler {
	r1 := CEF().SysCallN(2146, GetObjectUintptr(events))
	return AsCustomJsDialogHandler(r1)
}
