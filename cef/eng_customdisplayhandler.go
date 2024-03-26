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

// ICustomDisplayHandler Parent: ICefDisplayHandler
type ICustomDisplayHandler interface {
	ICefDisplayHandler
}

// TCustomDisplayHandler Parent: TCefDisplayHandler
type TCustomDisplayHandler struct {
	TCefDisplayHandler
}

func NewCustomDisplayHandler(events IChromiumEvents) ICustomDisplayHandler {
	r1 := CEF().SysCallN(2139, GetObjectUintptr(events))
	return AsCustomDisplayHandler(r1)
}
