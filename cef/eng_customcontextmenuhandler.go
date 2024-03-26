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

// ICustomContextMenuHandler Parent: ICefContextMenuHandler
type ICustomContextMenuHandler interface {
	ICefContextMenuHandler
}

// TCustomContextMenuHandler Parent: TCefContextMenuHandler
type TCustomContextMenuHandler struct {
	TCefContextMenuHandler
}

func NewCustomContextMenuHandler(events IChromiumEvents) ICustomContextMenuHandler {
	r1 := CEF().SysCallN(2136, GetObjectUintptr(events))
	return AsCustomContextMenuHandler(r1)
}
