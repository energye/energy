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

// ICefCustomRenderProcessHandler Parent: ICefRenderProcessHandler
type ICefCustomRenderProcessHandler interface {
	ICefRenderProcessHandler
}

// TCefCustomRenderProcessHandler Parent: TCefRenderProcessHandler
type TCefCustomRenderProcessHandler struct {
	TCefRenderProcessHandler
}

func NewCefCustomRenderProcessHandler(aCefApp ICefApplicationCore) ICefCustomRenderProcessHandler {
	r1 := CEF().SysCallN(783, GetObjectUintptr(aCefApp))
	return AsCefCustomRenderProcessHandler(r1)
}
