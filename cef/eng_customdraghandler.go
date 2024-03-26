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

// ICustomDragHandler Parent: ICefDragHandler
type ICustomDragHandler interface {
	ICefDragHandler
}

// TCustomDragHandler Parent: TCefDragHandler
type TCustomDragHandler struct {
	TCefDragHandler
}

func NewCustomDragHandler(events IChromiumEvents) ICustomDragHandler {
	r1 := CEF().SysCallN(2141, GetObjectUintptr(events))
	return AsCustomDragHandler(r1)
}
