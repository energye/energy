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

// ICefSetAudioMutedTask Parent: ICefTask
type ICefSetAudioMutedTask interface {
	ICefTask
}

// TCefSetAudioMutedTask Parent: TCefTask
type TCefSetAudioMutedTask struct {
	TCefTask
}

func NewCefSetAudioMutedTask(aEvents IChromiumEvents, aValue bool) ICefSetAudioMutedTask {
	r1 := CEF().SysCallN(1357, GetObjectUintptr(aEvents), PascalBool(aValue))
	return AsCefSetAudioMutedTask(r1)
}

func CefSetAudioMutedTaskClass() TClass {
	ret := CEF().SysCallN(1356)
	return TClass(ret)
}
