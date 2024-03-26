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

// ICustomLifeSpanHandler Parent: ICefLifeSpanHandler
type ICustomLifeSpanHandler interface {
	ICefLifeSpanHandler
}

// TCustomLifeSpanHandler Parent: TCefLifeSpanHandler
type TCustomLifeSpanHandler struct {
	TCefLifeSpanHandler
}

func NewCustomLifeSpanHandler(events IChromiumEvents) ICustomLifeSpanHandler {
	r1 := CEF().SysCallN(2148, GetObjectUintptr(events))
	return AsCustomLifeSpanHandler(r1)
}
