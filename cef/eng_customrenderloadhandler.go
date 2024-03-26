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

// ICustomRenderLoadHandler Parent: ICefLoadHandler
type ICustomRenderLoadHandler interface {
	ICefLoadHandler
}

// TCustomRenderLoadHandler Parent: TCefLoadHandler
type TCustomRenderLoadHandler struct {
	TCefLoadHandler
}

func NewCustomRenderLoadHandler(aCefApp ICefApplicationCore) ICustomRenderLoadHandler {
	r1 := CEF().SysCallN(2156, GetObjectUintptr(aCefApp))
	return AsCustomRenderLoadHandler(r1)
}
