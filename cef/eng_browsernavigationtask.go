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

// ICefBrowserNavigationTask Parent: ICefTask
type ICefBrowserNavigationTask interface {
	ICefTask
}

// TCefBrowserNavigationTask Parent: TCefTask
type TCefBrowserNavigationTask struct {
	TCefTask
}

func NewCefBrowserNavigationTask(aEvents IChromiumEvents, aTask TCefBrowserNavigation) ICefBrowserNavigationTask {
	r1 := CEF().SysCallN(674, GetObjectUintptr(aEvents), uintptr(aTask))
	return AsCefBrowserNavigationTask(r1)
}

func CefBrowserNavigationTaskClass() TClass {
	ret := CEF().SysCallN(673)
	return TClass(ret)
}
