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
	"github.com/energye/energy/v3/pkgs/gtk3"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type TWindow struct {
	TEnergyWindow
	gtkWindow *gtk3.Window
}

func (m *TWindow) CreateParams(params *types.TCreateParams) {

}

func (m *TWindow) _BeforeFormCreate() {
	m.UpdateWindowOption()
}

func (m *TWindow) _BeforeFormShow() {
	if m.flagFirstShow {
		return
	}
	m.flagFirstShow = true
}

// UpdateWindowOption 设置webview窗口的选项配置
// 该方法用于配置*TWindow实例的各种选项参数
func (m *TWindow) UpdateWindowOption() {
	gtkHandle := lcl.PlatformHandle(m.Handle())
	m.gtkWindow = gtk3.ToGtkWindow(uintptr(gtkHandle.Gtk3Window()))
	if m.options != nil {
		if m.options.Width <= 0 {
			m.options.Width = m.Width()
		}
		if m.options.Height <= 0 {
			m.options.Height = m.Height()
		}
		m.SetCaption(m.options.Caption)
		m.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
		if m.options.Frameless {
			m.gtkWindow.SetDecorated(false)
		}
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
		monitorRect := m.Monitor().WorkareaRect()
		if !m.options.Frameless {
			// save current window style, use ExitFullScreen
			m.SetWindowState(types.WsFullScreen)
		}
		m.SetBounds(monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height())
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !m.options.Frameless {
				m.SetBoundsRect(m.previousWindowPlacement)
			}
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.previousWindowPlacement)
		})
	}
}
