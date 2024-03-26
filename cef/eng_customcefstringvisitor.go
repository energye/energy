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

// ICustomCefStringVisitor Parent: ICefStringVisitor
type ICustomCefStringVisitor interface {
	ICefStringVisitor
}

// TCustomCefStringVisitor Parent: TCefStringVisitor
type TCustomCefStringVisitor struct {
	TCefStringVisitor
}

func NewCustomCefStringVisitor(aEvents IChromiumEvents) ICustomCefStringVisitor {
	r1 := CEF().SysCallN(2131, GetObjectUintptr(aEvents))
	return AsCustomCefStringVisitor(r1)
}
