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

// ICustomLoadHandler Parent: ICefLoadHandler
type ICustomLoadHandler interface {
	ICefLoadHandler
}

// TCustomLoadHandler Parent: TCefLoadHandler
type TCustomLoadHandler struct {
	TCefLoadHandler
}

func NewCustomLoadHandler(events IChromiumEvents) ICustomLoadHandler {
	r1 := CEF().SysCallN(2149, GetObjectUintptr(events))
	return AsCustomLoadHandler(r1)
}
