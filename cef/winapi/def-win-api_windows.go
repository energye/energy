//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package winapi energy - lcl window api
// export in Lazarus ascdef.inc

//go:build windows
// +build windows

package winapi

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
)

// Window class styles
const (
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_KEYCVTWINDOW    = 0x00000004
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOKEYCVT        = 0x00000100
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

const (
	GCL_CBCLSEXTRA    = -20
	GCL_CBWNDEXTRA    = -18
	GCL_HBRBACKGROUND = -10
	GCL_HCURSOR       = -12
	GCL_HICON         = -14
	GCL_HICONSM       = -34
	GCL_HMODULE       = -16
	GCL_MENUNAME      = -8
	GCL_STYLE         = -26
	GCL_WNDPROC       = -24
)

func WinGetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func WinSetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func WinGetClassLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func WinSetClassLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_SetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func WinFindWindow(lpClassName string, lpWindowName string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_FindWindow).Call(api.PascalStr(lpClassName), api.PascalStr(lpWindowName))
	return types.HWND(r1)
}

func WinFindWindowEx(_para1 types.HWND, _para2 types.HWND, _para3 string, _para4 string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_FindWindowEx).Call(uintptr(_para1), uintptr(_para2), api.PascalStr(_para3), api.PascalStr(_para4))
	return types.HWND(r1)
}

func WinSetWindowText(hWnd types.HWND, lpString string) types.LongBool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString))
	return types.LongBool(api.GoBool(r1))
}

func WinGetWindowText(hWnd types.HWND, lpString string, nMaxCount types.LongInt) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString), uintptr(nMaxCount))
	return types.LongInt(r1)
}

func WinGetWindowTextLength(hWnd types.HWND) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowTextLength).Call(uintptr(hWnd))
	return types.LongInt(r1)
}
