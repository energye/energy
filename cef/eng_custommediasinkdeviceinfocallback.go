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

// ICefCustomMediaSinkDeviceInfoCallback Parent: ICefMediaSinkDeviceInfoCallback
type ICefCustomMediaSinkDeviceInfoCallback interface {
	ICefMediaSinkDeviceInfoCallback
}

// TCefCustomMediaSinkDeviceInfoCallback Parent: TCefMediaSinkDeviceInfoCallback
type TCefCustomMediaSinkDeviceInfoCallback struct {
	TCefMediaSinkDeviceInfoCallback
}

func NewCefCustomMediaSinkDeviceInfoCallback(aEvents IChromiumEvents) ICefCustomMediaSinkDeviceInfoCallback {
	r1 := CEF().SysCallN(781, GetObjectUintptr(aEvents))
	return AsCefCustomMediaSinkDeviceInfoCallback(r1)
}
