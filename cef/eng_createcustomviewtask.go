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

// ICefCreateCustomViewTask Parent: ICefTask
type ICefCreateCustomViewTask interface {
	ICefTask
}

// TCefCreateCustomViewTask Parent: TCefTask
type TCefCreateCustomViewTask struct {
	TCefTask
}

func NewCefCreateCustomViewTask(aEvents ICefViewDelegateEvents) ICefCreateCustomViewTask {
	r1 := CEF().SysCallN(775, GetObjectUintptr(aEvents))
	return AsCefCreateCustomViewTask(r1)
}

func CefCreateCustomViewTaskClass() TClass {
	ret := CEF().SysCallN(774)
	return TClass(ret)
}
