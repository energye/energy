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

// ICustomDownloadHandler Parent: ICefDownloadHandler
type ICustomDownloadHandler interface {
	ICefDownloadHandler
}

// TCustomDownloadHandler Parent: TCefDownloadHandler
type TCustomDownloadHandler struct {
	TCefDownloadHandler
}

func NewCustomDownloadHandler(events IChromiumEvents) ICustomDownloadHandler {
	r1 := CEF().SysCallN(2140, GetObjectUintptr(events))
	return AsCustomDownloadHandler(r1)
}
