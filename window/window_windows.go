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

package window

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/pkgs/win32"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
)

type TWindow struct {
	TEnergyWindow
}

func (m *TWindow) borderFrameless() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

func (m *TWindow) platformCreate() {
	options := application.GApplication.Options
	switch options.Windows.Theme {
	case application.SystemDefault:
		win32.ChangeTheme(m.Handle(), win32.IsCurrentlyDarkMode())
	case application.Light:
		win32.ChangeTheme(m.Handle(), false)
	case application.Dark:
		win32.ChangeTheme(m.Handle(), true)
	}
}

func (m *TWindow) _BeforeFormCreate() {

}

// SetOptions 设置webview窗口的选项配置
// 该方法用于配置*TWindow实例的各种选项参数
func (m *TWindow) SetOptions() {
	m.platformCreate()
	m._HookWndProcMessage()
	m._AfterCreate()
	options := application.GApplication.Options
	if options.Width <= 0 {
		options.Width = m.Width()
	}
	if options.Height <= 0 {
		options.Height = m.Height()
	}
	m.SetCaption(options.Caption)
	m.SetBounds(options.X, options.Y, options.Width, options.Height)
}

func (m *TWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.IsMinimize() || m.IsMaximize() {
			m.Restore()
		}
		m.windowsState = types.WsFullScreen
		// save current window rect, use ExitFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		monitorRect := m.Monitor().BoundsRect()
		if !application.GApplication.Options.Frameless {
			// save current window style, use ExitFullScreen
			m.oldWindowStyle = uintptr(win.GetWindowLongPtr(m.Handle(), win.GWL_STYLE))
			m.borderFrameless()
			m.SetWindowState(types.WsFullScreen)
		}
		win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !application.GApplication.Options.Frameless {
				win.SetWindowLong(m.Handle(), win.GWL_STYLE, m.oldWindowStyle)
			}
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.previousWindowPlacement)
		})
	}
}
