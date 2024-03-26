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

// ICefSetZoomLevelTask Parent: ICefTask
type ICefSetZoomLevelTask interface {
	ICefTask
}

// TCefSetZoomLevelTask Parent: TCefTask
type TCefSetZoomLevelTask struct {
	TCefTask
}

func NewCefSetZoomLevelTask(aEvents IChromiumEvents, aValue float64) ICefSetZoomLevelTask {
	r1 := CEF().SysCallN(1359, GetObjectUintptr(aEvents), uintptr(unsafePointer(&aValue)))
	return AsCefSetZoomLevelTask(r1)
}

func CefSetZoomLevelTaskClass() TClass {
	ret := CEF().SysCallN(1358)
	return TClass(ret)
}
