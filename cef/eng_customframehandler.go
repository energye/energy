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

// ICustomFrameHandler Parent: ICefFrameHandler
type ICustomFrameHandler interface {
	ICefFrameHandler
}

// TCustomFrameHandler Parent: TCefFrameHandler
type TCustomFrameHandler struct {
	TCefFrameHandler
}

func NewCustomFrameHandler(events IChromiumEvents) ICustomFrameHandler {
	r1 := CEF().SysCallN(2145, GetObjectUintptr(events))
	return AsCustomFrameHandler(r1)
}
