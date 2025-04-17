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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"syscall"
	"unsafe"
)

var (
	user32dll                     = syscall.NewLazyDLL("user32.dll")
	dwmAPI                        = syscall.NewLazyDLL("dwmapi.dll")
	_BeginDeferWindowPos          = user32dll.NewProc("BeginDeferWindowPos")
	_DeferWindowPos               = user32dll.NewProc("DeferWindowPos")
	_EndDeferWindowPos            = user32dll.NewProc("EndDeferWindowPos")
	_GetDpiForWindow              = user32dll.NewProc("GetDpiForWindow")
	_DwmExtendFrameIntoClientArea = dwmAPI.NewProc("DwmExtendFrameIntoClientArea")
)
var (
	gdi32dll  = syscall.NewLazyDLL("gdi32.dll")
	_FrameRgn = gdi32dll.NewProc("FrameRgn")
)

type Margins struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

func ExtendFrameIntoClientArea(hwnd uintptr, margins Margins) {
	_, _, _ = _DwmExtendFrameIntoClientArea.Call(hwnd, uintptr(unsafe.Pointer(&margins)))
}

func GetDpiForWindow(hwnd types.HWND) (types.UINT, error) {
	if err := _GetDpiForWindow.Find(); err == nil {
		dpi, _, _ := _GetDpiForWindow.Call(uintptr(hwnd))
		return types.UINT(dpi), nil
	} else {
		return 0, err
	}
}

func FrameRgn(hdc types.HDC, hrgn *types.HRGN, hbr types.HBRUSH, w, h int) bool {
	r1, _, _ := _FrameRgn.Call(hdc.ToPtr(), hrgn.Instance(), hbr.ToPtr(), uintptr(w), uintptr(h))
	return r1 > 0
}

func BeginDeferWindowPos(nNumWindows int) types.HDWP {
	r1, _, _ := _BeginDeferWindowPos.Call(uintptr(nNumWindows))
	return types.HDWP(r1)
}

func DeferWindowPos(hWinPosInfo types.HDWP, hWnd types.HWND, hWndInsertAfter types.HWND, x, y, cx, cy int, uFlags uint) types.HDWP {
	r1, _, _ := _DeferWindowPos.Call(uintptr(hWinPosInfo), hWnd.ToPtr(), hWndInsertAfter.ToPtr(), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return types.HDWP(r1)
}

func EndDeferWindowPos(hWinPosInfo types.HDWP) bool {
	r1, _, _ := _EndDeferWindowPos.Call(uintptr(hWinPosInfo))
	return r1 > 0
}

func GetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func SetWindowLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func GetClassLongPtr(hWnd types.HWND, nIndex types.LongInt) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return types.LongPtr(r1)
}

func SetClassLongPtr(hWnd types.HWND, nIndex types.LongInt, dwNewLong types.LongPtr) types.LongPtr {
	r1, _, _ := imports.Proc(def.CEF_Win_SetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return types.LongPtr(r1)
}

func FindWindow(lpClassName string, lpWindowName string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_FindWindow).Call(api.PascalStr(lpClassName), api.PascalStr(lpWindowName))
	return types.HWND(r1)
}

func FindWindowEx(_para1 types.HWND, _para2 types.HWND, _para3 string, _para4 string) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_FindWindowEx).Call(uintptr(_para1), uintptr(_para2), api.PascalStr(_para3), api.PascalStr(_para4))
	return types.HWND(r1)
}

func SetWindowText(hWnd types.HWND, lpString string) types.LongBool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString))
	return types.LongBool(api.GoBool(r1))
}

func GetWindowText(hWnd types.HWND, lpString string, nMaxCount types.LongInt) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString), uintptr(nMaxCount))
	return types.LongInt(r1)
}

func GetWindowTextLength(hWnd types.HWND) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowTextLength).Call(uintptr(hWnd))
	return types.LongInt(r1)
}
