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

// ICefCustomCompletionCallback Parent: ICefCompletionCallback
type ICefCustomCompletionCallback interface {
	ICefCompletionCallback
}

// TCefCustomCompletionCallback Parent: TCefCompletionCallback
type TCefCustomCompletionCallback struct {
	TCefCompletionCallback
}

func NewCefCustomCompletionCallback(aEvents IChromiumEvents) ICefCustomCompletionCallback {
	r1 := CEF().SysCallN(776, GetObjectUintptr(aEvents))
	return AsCefCustomCompletionCallback(r1)
}
