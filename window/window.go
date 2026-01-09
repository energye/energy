// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package window

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type IWindow interface {
	lcl.IEngForm
	SetOptions()
	SetBrowserId(windowId uint32)
	BrowserId() uint32
	IsFullScreen() bool
	Restore()
	Maximize()
	IsMinimize() bool
	IsMaximize() bool
}

type TEnergyWindow struct {
	lcl.TEngForm
	windowId                uint32 // 窗口 ID 对应第一个浏览器 ID
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
}

func (m *TWindow) SetBrowserId(windowId uint32) {
	m.windowId = windowId
}

func (m *TWindow) BrowserId() uint32 {
	return m.windowId
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
