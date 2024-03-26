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

// ICustomCefNavigationEntryVisitor Parent: ICefNavigationEntryVisitor
type ICustomCefNavigationEntryVisitor interface {
	ICefNavigationEntryVisitor
}

// TCustomCefNavigationEntryVisitor Parent: TCefNavigationEntryVisitor
type TCustomCefNavigationEntryVisitor struct {
	TCefNavigationEntryVisitor
}

func NewCustomCefNavigationEntryVisitor(aEvents IChromiumEvents) ICustomCefNavigationEntryVisitor {
	r1 := CEF().SysCallN(2130, GetObjectUintptr(aEvents))
	return AsCustomCefNavigationEntryVisitor(r1)
}
