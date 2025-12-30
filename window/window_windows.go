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
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/win32"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
)

func (m *TWindow) borderFrameless() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

func (m *TWindow) platformCreate() {
	//if m.options.Windows.ICON == nil {
	//	lcl.Application.Icon().LoadFromBytes(assets.ICON.ICO())
	//} else {
	//	lcl.Application.Icon().LoadFromBytes(m.options.Windows.ICON)
	//}
	switch application.GApplication.Options.Windows.Theme {
	case application.SystemDefault:
		win32.ChangeTheme(m.Handle(), win32.IsCurrentlyDarkMode())
	case application.Light:
		win32.ChangeTheme(m.Handle(), false)
	case application.Dark:
		win32.ChangeTheme(m.Handle(), true)
	}
}

// SetOptions 设置webview窗口的选项配置
// 该方法用于配置*TWindow实例的各种选项参数
func (m *TWindow) SetOptions(windowId uint32) {
	m.windowId = windowId
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
	m.SetBounds(options.X, options.Y, options.Width, options.Height)
}

func (m *TWindow) WindowId() uint32 {
	return m.windowId
}

func (m *TWindow) Resize(ht string) {
	if m.IsFullScreen() || application.GApplication.Options.DisableResize {
		return
	}
	if win.ReleaseCapture() {
		var borderHT uintptr
		switch ht {
		case "n-resize":
			borderHT = messages.HTTOP
		case "ne-resize":
			borderHT = messages.HTTOPRIGHT
		case "e-resize":
			borderHT = messages.HTRIGHT
		case "se-resize":
			borderHT = messages.HTBOTTOMRIGHT
		case "s-resize":
			borderHT = messages.HTBOTTOM
		case "sw-resize":
			borderHT = messages.HTBOTTOMLEFT
		case "w-resize":
			borderHT = messages.HTLEFT
		case "nw-resize":
			borderHT = messages.HTTOPLEFT
		}
		win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, borderHT, 0)
	}
}

func (m *TWindow) Drag(message ipc.ProcessMessage) {
	if m.IsFullScreen() {
		return
	}
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
		if m.IsFullScreen() {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		m.Maximize()
	}
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

func (m *TWindow) Maximize() {
	if m.IsFullScreen() || application.GApplication.Options.DisableMaximize {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	})
}

func (m *TWindow) Restore() {
	// In the case of a title bar
	// If the current state is full screen and the extracted state is Ws Maximized,
	// So let's first perform IsFullScreen() judgment here
	if m.IsFullScreen() {
		m.ExitFullScreen()
	} else if m.IsMinimize() || m.IsMaximize() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
		})
	}
}

func (m *TWindow) IsFullScreen() bool {
	return m.windowsState == types.WsFullScreen
}

func (m *TWindow) IsMinimize() bool {
	return m.WindowState() == types.WsMinimized
}

func (m *TWindow) IsMaximize() bool {
	return m.WindowState() == types.WsMaximized
}
