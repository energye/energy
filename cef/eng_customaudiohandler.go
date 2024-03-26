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

// ICustomAudioHandler Parent: ICefAudioHandler
type ICustomAudioHandler interface {
	ICefAudioHandler
}

// TCustomAudioHandler Parent: TCefAudioHandler
type TCustomAudioHandler struct {
	TCefAudioHandler
}

func NewCustomAudioHandler(events IChromiumEvents) ICustomAudioHandler {
	r1 := CEF().SysCallN(2127, GetObjectUintptr(events))
	return AsCustomAudioHandler(r1)
}
