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

// ICefCustomSetCookieCallback Parent: ICefSetCookieCallback
type ICefCustomSetCookieCallback interface {
	ICefSetCookieCallback
}

// TCefCustomSetCookieCallback Parent: TCefSetCookieCallback
type TCefCustomSetCookieCallback struct {
	TCefSetCookieCallback
}

func NewCefCustomSetCookieCallback(aEvents IChromiumEvents, aID int32) ICefCustomSetCookieCallback {
	r1 := CEF().SysCallN(785, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomSetCookieCallback(r1)
}
