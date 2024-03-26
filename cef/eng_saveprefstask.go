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

// ICefSavePrefsTask Parent: ICefTask
type ICefSavePrefsTask interface {
	ICefTask
}

// TCefSavePrefsTask Parent: TCefTask
type TCefSavePrefsTask struct {
	TCefTask
}

func NewCefSavePrefsTask(aEvents IChromiumEvents) ICefSavePrefsTask {
	r1 := CEF().SysCallN(1343, GetObjectUintptr(aEvents))
	return AsCefSavePrefsTask(r1)
}

func CefSavePrefsTaskClass() TClass {
	ret := CEF().SysCallN(1342)
	return TClass(ret)
}
