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

// ICustomCommandHandler Parent: ICefCommandHandler
type ICustomCommandHandler interface {
	ICefCommandHandler
}

// TCustomCommandHandler Parent: TCefCommandHandler
type TCustomCommandHandler struct {
	TCefCommandHandler
}

func NewCustomCommandHandler(events IChromiumEvents) ICustomCommandHandler {
	r1 := CEF().SysCallN(2135, GetObjectUintptr(events))
	return AsCustomCommandHandler(r1)
}
