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
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func WinFindWindow(lpClassName string, lpWindowName string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(api.PascalStr(lpClassName), api.PascalStr(lpWindowName))
	return types.HWND(r1)
}

func WinFindWindowEx(_para1 types.HWND, _para2 types.HWND, _para3 string, _para4 string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(_para1), uintptr(_para2), api.PascalStr(_para3), api.PascalStr(_para4))
	return types.HWND(r1)
}

func WinSetWindowText(hWnd types.HWND, lpString string) types.LongBool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd), api.PascalStr(lpString))
	return types.LongBool(api.GoBool(r1))
}

func WinGetWindowText(hWnd types.HWND, lpString string, nMaxCount types.LongInt) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd), api.PascalStr(lpString), uintptr(nMaxCount))
	return types.LongInt(r1)
}

func WinGetWindowTextLength(hWnd types.HWND) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd))
	return types.LongInt(r1)
}
