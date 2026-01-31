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

// windows => LCL tray

package application

import (
	"github.com/energye/lcl/api/misc"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"syscall"
)

const (
	trayIconMsg = messages.WM_USER + 25
)

func (m *TTrayIcon) _HookWndProc() {
	if m.oldWndPrc == 0 {
		hWnd := misc.TrayIconHandle(m.trayIcon.Instance())
		newCallback := syscall.NewCallback(m.trayIconWndProc)
		m.oldWndPrc = win.SetWindowLongPtr(hWnd, win.GWL_WNDPROC, newCallback)
	}
}

func (m *TTrayIcon) trayIconWndProc(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	if message == trayIconMsg {

	}
	return win.CallWindowProc(m.oldWndPrc, hwnd, message, wParam, lParam)
}
