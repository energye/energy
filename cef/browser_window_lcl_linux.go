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
// +build linux

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	et "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func (m *LCLBrowserWindow) initDragEventListeners() {
	// no impl
}

func (m *LCLBrowserWindow) frameless() {
	// no impl
}

func (m *LCLBrowserWindow) SetRoundRectRgn(rgn int) {
	if m.rgn == 0 && rgn > 0 {
		m.rgn = rgn
		m.SetOnPaint(func(sender lcl.IObject) {
			hnd := winapi.CreateRoundRectRgn(0, 0, et.LongInt(m.Width()), et.LongInt(m.Height()), et.LongInt(m.rgn), et.LongInt(m.rgn))
			winapi.SetWindowRgn(et.HWND(m.Handle()), hnd, true)
		})
	}
}

// FullScreen 窗口全屏
func (m *LCLBrowserWindow) FullScreen() {
	if m.WindowProperty().EnableHideCaption {
		if m.IsFullScreen() {
			return
		}
		RunOnMainThread(func() {
			m.WindowProperty().current.windowState = types.WsFullScreen
			m.WindowProperty().current.previousWindowPlacement = m.BoundsRect()
			m.SetWindowState(types.WsFullScreen)
			m.SetBoundsRect(m.Monitor().BoundsRect())
		})
	}
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	// 恢复窗口大小
	if m.IsFullScreen() {
		wp := m.WindowProperty()
		RunOnMainThread(func() {
			wp.current.windowState = types.WsNormal
			m.SetBoundsRect(m.WindowProperty().current.previousWindowPlacement)
			m.SetWindowState(types.WsNormal)
		})
	}
}
