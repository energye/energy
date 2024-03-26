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

// ICefCustomMediaRouteCreateCallback Parent: ICefMediaRouteCreateCallback
type ICefCustomMediaRouteCreateCallback interface {
	ICefMediaRouteCreateCallback
}

// TCefCustomMediaRouteCreateCallback Parent: TCefMediaRouteCreateCallback
type TCefCustomMediaRouteCreateCallback struct {
	TCefMediaRouteCreateCallback
}

func NewCefCustomMediaRouteCreateCallback(aEvents IChromiumEvents) ICefCustomMediaRouteCreateCallback {
	r1 := CEF().SysCallN(780, GetObjectUintptr(aEvents))
	return AsCefCustomMediaRouteCreateCallback(r1)
}
