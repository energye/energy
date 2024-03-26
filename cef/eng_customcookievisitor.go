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

// ICefCustomCookieVisitor Parent: ICefCookieVisitor
type ICefCustomCookieVisitor interface {
	ICefCookieVisitor
}

// TCefCustomCookieVisitor Parent: TCefCookieVisitor
type TCefCustomCookieVisitor struct {
	TCefCookieVisitor
}

func NewCefCustomCookieVisitor(aEvents IChromiumEvents, aID int32) ICefCustomCookieVisitor {
	r1 := CEF().SysCallN(777, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomCookieVisitor(r1)
}
