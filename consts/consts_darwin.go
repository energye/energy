//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package consts

import "github.com/energye/energy/v2/types"

type TCefMenuAccelerator = types.Int32

const (
	MA_Shift                          = "SHIFT"
	MA_Shift_Code TCefMenuAccelerator = 0x10 //  16
	MA_Ctrl                           = "CTRL"
	MA_Ctrl_Code  TCefMenuAccelerator = 0x5B //  91
	MA_Alt                            = "ALT"
	MA_Alt_Code   TCefMenuAccelerator = 0x12 //  18
)

// 功能和消息常量
const (
	CEF_PREFERENCES_SAVED = 0x000A00
	CEF_DOONCLOSE         = 0x000A01
	CEF_STARTDRAGGING     = 0x000A02
	CEF_AFTERCREATED      = 0x000A03
	CEF_PENDINGRESIZE     = 0x000A04
	CEF_PUMPHAVEWORK      = 0x000A05
	CEF_DESTROY           = 0x000A06
	CEF_DOONBEFORECLOSE   = 0x000A07
	CEF_PENDINGINVALIDATE = 0x000A08
	CEF_IMERANGECHANGED   = 0x000A09
	CEF_SENTINEL_START    = 0x000A0A
	CEF_SENTINEL_DOCLOSE  = 0x000A0B
	CEF_BEFORECLOSE       = 0x000A0C
)

// TCefEventHandle
//
//	/include/internal/cef_types_mac.h (cef_event_handle_t)
type TCefEventHandle = uintptr

func EventHandle(ptr uintptr) TCefEventHandle {
	return TCefEventHandle(ptr)
}
