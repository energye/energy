// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package wv

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/lcl/api/misc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"syscall"
)

type TWebSystray struct {
	oldWndPrc uintptr
	trayIcon  lcl.ITrayIcon
	window    lcl.IEngForm
}

func NewWebSystray() *TWebSystray {
	m := &TWebSystray{}
	return m
}

func (m *TWebSystray) BindLCLTray(appTray application.ILCLSystray) {
	if appTray == nil {
		return
	}
	m.trayIcon = appTray.LCLTray()
	m._TrayWndProc()
}

func (m *TWebSystray) Show() {

}

func (m *TWebSystray) Hide() {

}

const (
	trayIconMsg = messages.WM_USER + 25
)

func (m *TWebSystray) _TrayWndProc() {
	if m.oldWndPrc == 0 {
		hWnd := misc.TrayIconHandle(m.trayIcon.Instance())
		newCallback := syscall.NewCallback(m.trayWndProc)
		m.oldWndPrc = win.SetWindowLongPtr(hWnd, win.GWL_WNDPROC, newCallback)
	}
}

func (m *TWebSystray) trayWndProc(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	if message == trayIconMsg {

	}
	return win.CallWindowProc(m.oldWndPrc, hwnd, message, wParam, lParam)
}
