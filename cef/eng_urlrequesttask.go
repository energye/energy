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

// ICefURLRequestTask Parent: ICefTask
type ICefURLRequestTask interface {
	ICefTask
}

// TCefURLRequestTask Parent: TCefTask
type TCefURLRequestTask struct {
	TCefTask
}

func NewCefURLRequestTask(aEvents ICEFUrlRequestClientEvents) ICefURLRequestTask {
	r1 := CEF().SysCallN(1458, GetObjectUintptr(aEvents))
	return AsCefURLRequestTask(r1)
}

func CefURLRequestTaskClass() TClass {
	ret := CEF().SysCallN(1457)
	return TClass(ret)
}
