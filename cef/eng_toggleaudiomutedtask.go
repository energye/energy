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

// ICefToggleAudioMutedTask Parent: ICefTask
type ICefToggleAudioMutedTask interface {
	ICefTask
}

// TCefToggleAudioMutedTask Parent: TCefTask
type TCefToggleAudioMutedTask struct {
	TCefTask
}

func NewCefToggleAudioMutedTask(aEvents IChromiumEvents) ICefToggleAudioMutedTask {
	r1 := CEF().SysCallN(1456, GetObjectUintptr(aEvents))
	return AsCefToggleAudioMutedTask(r1)
}

func CefToggleAudioMutedTaskClass() TClass {
	ret := CEF().SysCallN(1455)
	return TClass(ret)
}
