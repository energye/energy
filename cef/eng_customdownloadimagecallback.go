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

// ICefCustomDownloadImageCallback Parent: ICefDownloadImageCallback
type ICefCustomDownloadImageCallback interface {
	ICefDownloadImageCallback
}

// TCefCustomDownloadImageCallback Parent: TCefDownloadImageCallback
type TCefCustomDownloadImageCallback struct {
	TCefDownloadImageCallback
}

func NewCefCustomDownloadImageCallback(aEvents IChromiumEvents) ICefCustomDownloadImageCallback {
	r1 := CEF().SysCallN(779, GetObjectUintptr(aEvents))
	return AsCefCustomDownloadImageCallback(r1)
}
