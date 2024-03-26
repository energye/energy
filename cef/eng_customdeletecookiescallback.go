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

// ICefCustomDeleteCookiesCallback Parent: ICefDeleteCookiesCallback
type ICefCustomDeleteCookiesCallback interface {
	ICefDeleteCookiesCallback
}

// TCefCustomDeleteCookiesCallback Parent: TCefDeleteCookiesCallback
type TCefCustomDeleteCookiesCallback struct {
	TCefDeleteCookiesCallback
}

func NewCefCustomDeleteCookiesCallback(aEvents IChromiumEvents) ICefCustomDeleteCookiesCallback {
	r1 := CEF().SysCallN(778, GetObjectUintptr(aEvents))
	return AsCefCustomDeleteCookiesCallback(r1)
}
