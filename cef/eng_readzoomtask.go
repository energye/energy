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

// ICefReadZoomTask Parent: ICefTask
type ICefReadZoomTask interface {
	ICefTask
}

// TCefReadZoomTask Parent: TCefTask
type TCefReadZoomTask struct {
	TCefTask
}

func NewCefReadZoomTask(aEvents IChromiumEvents) ICefReadZoomTask {
	r1 := CEF().SysCallN(1250, GetObjectUintptr(aEvents))
	return AsCefReadZoomTask(r1)
}

func CefReadZoomTaskClass() TClass {
	ret := CEF().SysCallN(1249)
	return TClass(ret)
}
