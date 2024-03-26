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

// ICefEnableFocusTask Parent: ICefTask
type ICefEnableFocusTask interface {
	ICefTask
}

// TCefEnableFocusTask Parent: TCefTask
type TCefEnableFocusTask struct {
	TCefTask
}

func NewCefEnableFocusTask(aEvents IChromiumEvents) ICefEnableFocusTask {
	r1 := CEF().SysCallN(933, GetObjectUintptr(aEvents))
	return AsCefEnableFocusTask(r1)
}

func CefEnableFocusTaskClass() TClass {
	ret := CEF().SysCallN(932)
	return TClass(ret)
}
