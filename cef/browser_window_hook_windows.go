//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	"github.com/energye/golcl/lcl/win"
	"syscall"
	"unsafe"
)

var (
	dwmAPI                        = syscall.NewLazyDLL("dwmapi.dll")
	_DwmExtendFrameIntoClientArea = dwmAPI.NewProc("DwmExtendFrameIntoClientArea")
)

type margins struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

func extendFrameIntoClientArea(hwnd uintptr, margins margins) {
	_, _, _ = _DwmExtendFrameIntoClientArea.Call(hwnd, uintptr(unsafe.Pointer(&margins)))
}

func (m *LCLBrowserWindow) wndProc(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	if m.WindowProperty().EnableHideCaption {
		switch message {
		case messages.WM_ACTIVATE:
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			extendFrameIntoClientArea(m.Handle(), margins{CxLeftWidth: 1, CxRightWidth: 1, CyTopHeight: 1, CyBottomHeight: 1})
		case messages.WM_NCCALCSIZE:
			// Disable the standard frame by allowing the client area to take the full
			// window size.
			// See: https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize#remarks
			// This hides the titlebar and also disables the resizing from user interaction because the standard frame is not
			// shown. We still need the WS_THICKFRAME style to enable resizing from the frontend.
			if wParam != 0 {
				//cycaption := win.GetSystemMetrics(4)
				//rect := (*types.TRect)(unsafe.Pointer(lParam))
				//rect.Bottom += -1
				//rect.Right += -1
				return 0
			}
		}
	}
	return win.CallWindowProc(m.oldWndPrc, hwnd, message, wParam, lParam)
}

func (m *LCLBrowserWindow) _HookWndProcMessage() {
	wndProcCallback := syscall.NewCallback(m.wndProc)
	m.oldWndPrc = win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, wndProcCallback)
}

func (m *LCLBrowserWindow) _RestoreWndProc() {
	if m.oldWndPrc != 0 {
		win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, m.oldWndPrc)
		m.oldWndPrc = 0
	}
}
