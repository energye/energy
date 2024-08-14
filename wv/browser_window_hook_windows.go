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

package wv

import (
	"github.com/energye/energy/v3/pkgs/win32"
	"github.com/energye/lcl/api/winapi"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"syscall"
	"unsafe"
)

func (m *BrowserWindow) wndProc(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case messages.WM_SETTINGCHANGE:
		settingChanged := win32.UTF16PtrToString((*uint16)(unsafe.Pointer(lParam)))
		if settingChanged == "ImmersiveColorSet" {
			//fmt.Println("WM_SETTINGCHANGE settingChanged:", settingChanged)
		}

	case messages.WM_POWERBROADCAST:
		//switch wParam {
		//case messages.PBT_APMPOWERSTATUSCHANGE:
		//case messages.PBT_APMSUSPEND:
		//case messages.PBT_APMRESUMEAUTOMATIC:
		//case messages.PBT_APMRESUMESUSPEND:
		//case messages.PBT_POWERSETTINGCHANGE:
		//}
	}
	if m.options.Frameless {
		switch message {
		case messages.WM_ACTIVATE:
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			win.ExtendFrameIntoClientArea(m.Handle(), win.Margins{CxLeftWidth: 1, CxRightWidth: 1, CyTopHeight: 1, CyBottomHeight: 1})
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
					rect := (*types.TRect)(unsafe.Pointer(lParam))
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

func (m *BrowserWindow) _HookWndProcMessage() {
	if m.oldWndPrc == 0 {
		wndProcCallback := syscall.NewCallback(m.wndProc)
		m.oldWndPrc = win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, wndProcCallback)
		// trigger WM_NCCALCSIZE
		// https://learn.microsoft.com/en-us/windows/win32/dwm/customframe#removing-the-standard-frame
		clientRect := m.BoundsRect()
		win.SetWindowPos(m.Handle(), 0, clientRect.Left, clientRect.Top, clientRect.Right-clientRect.Left, clientRect.Bottom-clientRect.Top, win.SWP_FRAMECHANGED)
	}
}

func (m *BrowserWindow) _RestoreWndProc() {
	if m.oldWndPrc != 0 {
		win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, m.oldWndPrc)
		m.oldWndPrc = 0
	}
}
