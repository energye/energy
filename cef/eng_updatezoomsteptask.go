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

// ICefUpdateZoomStepTask Parent: ICefTask
type ICefUpdateZoomStepTask interface {
	ICefTask
}

// TCefUpdateZoomStepTask Parent: TCefTask
type TCefUpdateZoomStepTask struct {
	TCefTask
}

func NewCefUpdateZoomStepTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomStepTask {
	r1 := CEF().SysCallN(1464, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomStepTask(r1)
}

func CefUpdateZoomStepTaskClass() TClass {
	ret := CEF().SysCallN(1463)
	return TClass(ret)
}
