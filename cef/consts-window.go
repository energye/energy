//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

type TCefMenuAccelerator = int32

const (
	MA_Shift                          = "SHIFT"
	MA_Shift_Code TCefMenuAccelerator = 0x10 //  16
	MA_Ctrl                           = "CTRL"
	MA_Ctrl_Code  TCefMenuAccelerator = 0x11 //  17
	MA_Alt                            = "ALT"
	MA_Alt_Code   TCefMenuAccelerator = 0x12 //  18
)

const (
	CEF_PREFERENCES_SAVED = WM_APP + 0x000A00
	CEF_DOONCLOSE         = WM_APP + 0x000A01
	CEF_STARTDRAGGING     = WM_APP + 0x000A02
	CEF_AFTERCREATED      = WM_APP + 0x000A03
	CEF_PENDINGRESIZE     = WM_APP + 0x000A04
	CEF_PUMPHAVEWORK      = WM_APP + 0x000A05
	CEF_DESTROY           = WM_APP + 0x000A06
	CEF_DOONBEFORECLOSE   = WM_APP + 0x000A07
	CEF_PENDINGINVALIDATE = WM_APP + 0x000A08
	CEF_IMERANGECHANGED   = WM_APP + 0x000A09
	CEF_SENTINEL_START    = WM_APP + 0x000A0A
	CEF_SENTINEL_DOCLOSE  = WM_APP + 0x000A0B
	CEF_BEFORECLOSE       = WM_APP + 0x000A0C
)
