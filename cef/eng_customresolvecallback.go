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

// ICefCustomResolveCallback Parent: ICefResolveCallback
type ICefCustomResolveCallback interface {
	ICefResolveCallback
}

// TCefCustomResolveCallback Parent: TCefResolveCallback
type TCefCustomResolveCallback struct {
	TCefResolveCallback
}

func NewCefCustomResolveCallback(aEvents IChromiumEvents) ICefCustomResolveCallback {
	r1 := CEF().SysCallN(784, GetObjectUintptr(aEvents))
	return AsCefCustomResolveCallback(r1)
}
