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

// ICustomPrintHandler Parent: ICefPrintHandler
type ICustomPrintHandler interface {
	ICefPrintHandler
}

// TCustomPrintHandler Parent: TCefPrintHandler
type TCustomPrintHandler struct {
	TCefPrintHandler
}

func NewCustomPrintHandler(events IChromiumEvents) ICustomPrintHandler {
	r1 := CEF().SysCallN(2154, GetObjectUintptr(events))
	return AsCustomPrintHandler(r1)
}
