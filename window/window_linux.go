//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package window

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

// SetOptions 设置webview窗口的选项配置
// 该方法用于配置*TWindow实例的各种选项参数
func (m *TWindow) SetOptions() {
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

func (m *TWindow) Resize(ht string) {
	if m.IsFullScreen() || application.GApplication.Options.DisableResize {
		return
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
		//if win.ReleaseCapture() {
		//	win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		//}
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
		//monitorRect := m.Monitor().BoundsRect()
		if !application.GApplication.Options.Frameless {
			// save current window style, use ExitFullScreen
			//m.oldWindowStyle = uintptr(win.GetWindowLongPtr(m.Handle(), win.GWL_STYLE))
			//m.borderFrameless()
			m.SetWindowState(types.WsFullScreen)
		}
		//win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !application.GApplication.Options.Frameless {
				//win.SetWindowLong(m.Handle(), win.GWL_STYLE, m.oldWindowStyle)
			}
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.previousWindowPlacement)
		})
	}
}
