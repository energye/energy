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

// ICefUpdatePrefsTask Parent: ICefTask
type ICefUpdatePrefsTask interface {
	ICefTask
}

// TCefUpdatePrefsTask Parent: TCefTask
type TCefUpdatePrefsTask struct {
	TCefTask
}

func NewCefUpdatePrefsTask(aEvents IChromiumEvents) ICefUpdatePrefsTask {
	r1 := CEF().SysCallN(1460, GetObjectUintptr(aEvents))
	return AsCefUpdatePrefsTask(r1)
}

func CefUpdatePrefsTaskClass() TClass {
	ret := CEF().SysCallN(1459)
	return TClass(ret)
}
