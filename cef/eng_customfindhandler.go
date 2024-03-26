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

// ICustomFindHandler Parent: ICefFindHandler
type ICustomFindHandler interface {
	ICefFindHandler
}

// TCustomFindHandler Parent: TCefFindHandler
type TCustomFindHandler struct {
	TCefFindHandler
}

func NewCustomFindHandler(events IChromiumEvents) ICustomFindHandler {
	r1 := CEF().SysCallN(2143, GetObjectUintptr(events))
	return AsCustomFindHandler(r1)
}
