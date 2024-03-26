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

// ICefSetZoomStepTask Parent: ICefTask
type ICefSetZoomStepTask interface {
	ICefTask
}

// TCefSetZoomStepTask Parent: TCefTask
type TCefSetZoomStepTask struct {
	TCefTask
}

func NewCefSetZoomStepTask(aEvents IChromiumEvents, aValue byte) ICefSetZoomStepTask {
	r1 := CEF().SysCallN(1363, GetObjectUintptr(aEvents), uintptr(aValue))
	return AsCefSetZoomStepTask(r1)
}

func CefSetZoomStepTaskClass() TClass {
	ret := CEF().SysCallN(1362)
	return TClass(ret)
}
