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
// export in Lazarus ascinc

//go:build windows
// +build windows

package winapi

import (
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/types"
	"syscall"
)

var (
	user32dll            = syscall.NewLazyDLL("user32.dll")
	_BeginDeferWindowPos = user32dll.NewProc("BeginDeferWindowPos")
	_DeferWindowPos      = user32dll.NewProc("DeferWindowPos")
	_EndDeferWindowPos   = user32dll.NewProc("EndDeferWindowPos")
	_GetDpiForWindow     = user32dll.NewProc("GetDpiForWindow")
)
var (
	gdi32dll  = syscall.NewLazyDLL("gdi32.dll")
	_FrameRgn = gdi32dll.NewProc("FrameRgn")
)

func GetDpiForWindow(hwnd types.HWND) (types.UINT, error) {
	if err := _GetDpiForWindow.Find(); err == nil {
		dpi, _, _ := _GetDpiForWindow.Call(uintptr(hwnd))
		return types.UINT(dpi), nil
	} else {
		return 0, err
	}
}

func FrameRgn(hdc types.HDC, hrgn types.HRGN, hbr types.HBRUSH, w, h int) bool {
	r1, _, _ := _FrameRgn.Call(hdc, hrgn, hbr, uintptr(w), uintptr(h))
	return r1 > 0
}

func BeginDeferWindowPos(nNumWindows int) types.HDWP {
	r1, _, _ := _BeginDeferWindowPos.Call(uintptr(nNumWindows))
	return types.HDWP(r1)
}

func DeferWindowPos(hWinPosInfo types.HDWP, hWnd types.HWND, hWndInsertAfter types.HWND, x, y, cx, cy int, uFlags uint) types.HDWP {
	r1, _, _ := _DeferWindowPos.Call(uintptr(hWinPosInfo), hWnd, hWndInsertAfter, uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return types.HDWP(r1)
}

func EndDeferWindowPos(hWinPosInfo types.HDWP) bool {
	r1, _, _ := _EndDeferWindowPos.Call(uintptr(hWinPosInfo))
	return r1 > 0
}

func GetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := api.WinAPI().Proc(cef_win_GetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func SetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := api.WinAPI().Proc(cef_win_SetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func GetClassLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := api.WinAPI().Proc(cef_win_GetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func SetClassLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := api.WinAPI().Proc(cef_win_SetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func FindWindow(lpClassName string, lpWindowName string) types.HWND {
	r1, _, _ := api.WinAPI().Proc(cef_win_FindWindow).Call(api.PascalStr(lpClassName), api.PascalStr(lpWindowName))
	return types.HWND(r1)
}

func FindWindowEx(_para1 types.HWND, _para2 types.HWND, _para3 string, _para4 string) types.HWND {
	r1, _, _ := api.WinAPI().Proc(cef_win_FindWindowEx).Call(uintptr(_para1), uintptr(_para2), api.PascalStr(_para3), api.PascalStr(_para4))
	return types.HWND(r1)
}

func SetWindowText(hWnd types.HWND, lpString string) types.LongBool {
	r1, _, _ := api.WinAPI().Proc(cef_win_SetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString))
	return types.LongBool(api.GoBool(r1))
}

func GetWindowText(hWnd types.HWND, lpString string, nMaxCount types.LongInt) types.LongInt {
	r1, _, _ := api.WinAPI().Proc(cef_win_GetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString), uintptr(nMaxCount))
	return types.LongInt(r1)
}

func GetWindowTextLength(hWnd types.HWND) types.LongInt {
	r1, _, _ := api.WinAPI().Proc(cef_win_GetWindowTextLength).Call(uintptr(hWnd))
	return types.LongInt(r1)
}
