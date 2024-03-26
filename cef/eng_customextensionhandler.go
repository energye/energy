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

// ICustomExtensionHandler Parent: ICefExtensionHandler
type ICustomExtensionHandler interface {
	ICefExtensionHandler
}

// TCustomExtensionHandler Parent: TCefExtensionHandler
type TCustomExtensionHandler struct {
	TCefExtensionHandler
}

func NewCustomExtensionHandler(events IChromiumEvents) ICustomExtensionHandler {
	r1 := CEF().SysCallN(2142, GetObjectUintptr(events))
	return AsCustomExtensionHandler(r1)
}
