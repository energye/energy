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
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types/messages"
	"github.com/energye/golcl/lcl/win"
	"syscall"
	"unsafe"
)

func (m *LCLBrowserWindow) wndProc(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case messages.WM_DPICHANGED:
		if !lcl.Application.Scaled() {
			newWindowSize := (*types.Rect)(unsafe.Pointer(lParam))
			win.SetWindowPos(m.Handle(), uintptr(0),
				newWindowSize.Left, newWindowSize.Top, newWindowSize.Right-newWindowSize.Left, newWindowSize.Bottom-newWindowSize.Top,
				win.SWP_NOZORDER|win.SWP_NOACTIVATE)
		}
	}
	if m.WindowProperty().EnableHideCaption {
		switch message {
		case messages.WM_ACTIVATE:
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			winapi.ExtendFrameIntoClientArea(m.Handle(), winapi.Margins{CxLeftWidth: 1, CxRightWidth: 1, CyTopHeight: 1, CyBottomHeight: 1})
		case messages.WM_NCCALCSIZE:
			// Trigger condition: Change the window size
			// Disable the standard frame by allowing the client area to take the full window size.
			// See: https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize#remarks
			// This hides the titlebar and also disables the resizing from user interaction because the standard frame is not
			// shown. We still need the WS_THICKFRAME style to enable resizing from the frontend.
			if wParam != 0 {
				// Content overflow screen issue when maximizing borderless windows
				// See: https://github.com/MicrosoftEdge/WebView2Feedback/issues/2549
				//isMinimize := uint32(win.GetWindowLong(m.Handle(), win.GWL_STYLE))&win.WS_MINIMIZE != 0
				isMaximize := uint32(win.GetWindowLong(m.Handle(), win.GWL_STYLE))&win.WS_MAXIMIZE != 0
				if isMaximize {
					rect := (*types.Rect)(unsafe.Pointer(lParam))
					// m.Monitor().WorkareaRect(): When minimizing windows and restoring windows on multiple monitors, the main monitor is obtained.
					// Need to obtain correct monitor information to prevent error freezing message loops from occurring
					monitor := winapi.MonitorFromRect(*rect, winapi.MONITOR_DEFAULTTONULL)
					if monitor != 0 {
						var monitorInfo types.TagMonitorInfo
						monitorInfo.CbSize = types.DWORD(unsafe.Sizeof(monitorInfo))
						if winapi.GetMonitorInfo(monitor, &monitorInfo) {
							*rect = monitorInfo.RcWork
						}
					}
				}
				return 0
			}
		}
	}

	return win.CallWindowProc(m.oldWndPrc, uintptr(hwnd), message, wParam, lParam)
}

// 该函数调用可能会影响窗口的一些默认行为，需要知道在合适的时机调用它
func (m *LCLBrowserWindow) _HookWndProcMessage() {
	switch m.WindowType() {
	case consts.WT_DEV_TOOLS, consts.WT_VIEW_SOURCE:
		return
	}
	wndProcCallback := syscall.NewCallback(m.wndProc)
	m.oldWndPrc = win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, wndProcCallback)
}

func (m *LCLBrowserWindow) _RestoreWndProc() {
	if m.oldWndPrc != 0 {
		win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, m.oldWndPrc)
		m.oldWndPrc = 0
	}
}
