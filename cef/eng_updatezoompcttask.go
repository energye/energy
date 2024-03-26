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

// ICefUpdateZoomPctTask Parent: ICefTask
type ICefUpdateZoomPctTask interface {
	ICefTask
}

// TCefUpdateZoomPctTask Parent: TCefTask
type TCefUpdateZoomPctTask struct {
	TCefTask
}

func NewCefUpdateZoomPctTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomPctTask {
	r1 := CEF().SysCallN(1462, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomPctTask(r1)
}

func CefUpdateZoomPctTaskClass() TClass {
	ret := CEF().SysCallN(1461)
	return TClass(ret)
}
