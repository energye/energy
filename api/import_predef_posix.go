//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package api

import (
	"github.com/energye/energy/v2/api/internal/lcl"
	"github.com/energye/energy/v2/types"
	"unsafe"
)

func DSendMessage(hWd types.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return defSyscallN(lcl.DSENDMESSAGE, hWd, uintptr(msg), wParam, lParam)
}

func DPostMessage(hWd types.HWND, msg uint32, wParam, lParam uintptr) bool {
	return GoBool(defSyscallN(lcl.DPOSTMESSAGE, hWd, uintptr(msg), wParam, lParam))
}

func DIsIconic(hWnd types.HWND) bool {
	return GoBool(defSyscallN(lcl.DISICONIC, hWnd))
}

func DIsWindow(hWnd types.HWND) bool {
	return GoBool(defSyscallN(lcl.DISWINDOW, hWnd))
}

func DIsZoomed(hWnd types.HWND) bool {
	return GoBool(defSyscallN(lcl.DISZOOMED, hWnd))
}

func DIsWindowVisible(hWnd types.HWND) bool {
	return GoBool(defSyscallN(lcl.DISWINDOWVISIBLE, hWnd))
}

func DGetDC(hWnd types.HWND) types.HDC {
	return defSyscallN(lcl.DGETDC, hWnd)
}

func DReleaseDC(hWnd types.HWND, dc types.HDC) int {
	return int(defSyscallN(lcl.DRELEASEDC, hWnd, dc))
}

func DSetForegroundWindow(hWnd types.HWND) bool {
	return GoBool(defSyscallN(lcl.DSETFOREGROUNDWINDOW, hWnd))
}

func DWindowFromPoint(point types.TPoint) types.HWND {
	return defSyscallN(lcl.DWINDOWFROMPOINT, uintptr(unsafe.Pointer(&point)))
}
