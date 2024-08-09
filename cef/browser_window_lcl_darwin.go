//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package cef

import (
	"github.com/energye/golcl/lcl/types"
)

const (
	NSWindowStyleMaskBorderless     = 0 // 窗口没有标题栏和按钮
	NSWindowStyleMaskTitled         = 1 // 窗口具有标题栏
	NSWindowStyleMaskClosable       = 2 // 窗口具有关闭按钮
	NSWindowStyleMaskMiniaturizable = 4 // 窗口可以被最小化
	NSWindowStyleMaskResizable      = 8 // 窗口可以调整大小
)

func (m *LCLBrowserWindow) frameless() {
	nsWindow := m.PlatformWindow()
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	nsWindow.SetStyleMask(NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskResizable)
}

func (m *LCLBrowserWindow) SetRoundRectRgn(rgn int) {
	if m.rgn == 0 && rgn > 0 {
		m.rgn = rgn

	}
}

// FullScreen 窗口全屏
func (m *LCLBrowserWindow) FullScreen() {
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
