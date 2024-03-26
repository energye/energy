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

// ICefSetZoomPctTask Parent: ICefTask
type ICefSetZoomPctTask interface {
	ICefTask
}

// TCefSetZoomPctTask Parent: TCefTask
type TCefSetZoomPctTask struct {
	TCefTask
}

func NewCefSetZoomPctTask(aEvents IChromiumEvents, aValue float64) ICefSetZoomPctTask {
	r1 := CEF().SysCallN(1361, GetObjectUintptr(aEvents), uintptr(unsafePointer(&aValue)))
	return AsCefSetZoomPctTask(r1)
}

func CefSetZoomPctTaskClass() TClass {
	ret := CEF().SysCallN(1360)
	return TClass(ret)
}
