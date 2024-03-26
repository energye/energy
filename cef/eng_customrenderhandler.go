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

// ICustomRenderHandler Parent: ICefRenderHandler
type ICustomRenderHandler interface {
	ICefRenderHandler
}

// TCustomRenderHandler Parent: TCefRenderHandler
type TCustomRenderHandler struct {
	TCefRenderHandler
}

func NewCustomRenderHandler(events IChromiumEvents) ICustomRenderHandler {
	r1 := CEF().SysCallN(2155, GetObjectUintptr(events))
	return AsCustomRenderHandler(r1)
}
