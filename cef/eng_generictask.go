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

// ICefGenericTask Parent: ICefTask
type ICefGenericTask interface {
	ICefTask
}

// TCefGenericTask Parent: TCefTask
type TCefGenericTask struct {
	TCefTask
}

func NewCefGenericTask(aEvents IChromiumEvents, aTaskID uint32) ICefGenericTask {
	r1 := CEF().SysCallN(982, GetObjectUintptr(aEvents), uintptr(aTaskID))
	return AsCefGenericTask(r1)
}

func CefGenericTaskClass() TClass {
	ret := CEF().SysCallN(981)
	return TClass(ret)
}
