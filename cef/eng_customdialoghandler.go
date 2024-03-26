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

// ICustomDialogHandler Parent: ICefDialogHandler
type ICustomDialogHandler interface {
	ICefDialogHandler
}

// TCustomDialogHandler Parent: TCefDialogHandler
type TCustomDialogHandler struct {
	TCefDialogHandler
}

func NewCustomDialogHandler(events IChromiumEvents) ICustomDialogHandler {
	r1 := CEF().SysCallN(2138, GetObjectUintptr(events))
	return AsCustomDialogHandler(r1)
}
