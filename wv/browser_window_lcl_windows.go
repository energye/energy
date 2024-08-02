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
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"syscall"
)

func (m *BrowserWindow) Resize(border uintptr) {
	if win.ReleaseCapture() {
		win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, border, 0)
	}
}

func (m *BrowserWindow) Drag(message ipc.ProcessMessage) {
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
		//fmt.Println("MT_DRAG_MOVE", m.WindowState())
		if m.WindowState() == types.WsFullScreen {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	}
}

var (
	wndProcCallback = syscall.NewCallback(_WndProcCallback)
)

func _WndProcCallback(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	//fmt.Println("HookMsg msg:", message)
	if window := getBrowserWindow(hwnd); window != nil {
		return window._WndProc(message, wParam, lParam)
	}
	//return win.CallWindowProc(oldWndPrc, hwnd, message, wParam, lParam)
	return win.DefWindowProc(hwnd, message, wParam, lParam)
}

func (m *BrowserWindow) _WndProc(message uint32, wParam, lParam uintptr) uintptr {
	if m.options.Frameless {
		switch message {
		case messages.WM_ACTIVATE:
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			win.ExtendFrameIntoClientArea(m.Handle(), win.Margins{CxLeftWidth: 1, CxRightWidth: 1, CyTopHeight: 1, CyBottomHeight: 1})
		case messages.WM_NCCALCSIZE:
			// Disable the standard frame by allowing the client area to take the full
			// window size.
			// See: https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize#remarks
			// This hides the titlebar and also disables the resizing from user interaction because the standard frame is not
			// shown. We still need the WS_THICKFRAME style to enable resizing from the frontend.
			if wParam != 0 {
				return 0
			}
		}
		//fmt.Println("message:", message)
	}
	return win.CallWindowProc(m.oldWndPrc, m.Handle(), message, wParam, lParam)
}

func (m *BrowserWindow) _HookWndProcMessage() {
	m.oldWndPrc = win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, wndProcCallback)
}

func (m *BrowserWindow) _RestoreWndProc() {
	if m.oldWndPrc != 0 {
		win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, m.oldWndPrc)
	}
}
