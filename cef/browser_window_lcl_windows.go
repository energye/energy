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

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/consts/messages"
	et "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

func (m *LCLBrowserWindow) initDragEventListeners() {
	// no impl
}

func (m *LCLBrowserWindow) frameless() {
	// no impl
}

// SetRoundRectRgn 窗口无边框时圆角设置
//
//	如果 rgn 值设置的过大同时开启GPU加速窗口会卡顿
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
			if m.WindowState() == types.WsMinimized || m.WindowState() == types.WsMaximized {
				if win.ReleaseCapture() {
					win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
				}
			}
			m.WindowProperty().current.windowState = types.WsFullScreen
			m.WindowProperty().current.previousWindowPlacement = m.BoundsRect()
			//style := uint32(win.GetWindowLongPtr(m.Handle(), win.GWL_STYLE))
			monitorRect := m.Monitor().BoundsRect()
			win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
		})
	}
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		wp := m.WindowProperty()
		RunOnMainThread(func() {
			wp.current.windowState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.WindowProperty().current.previousWindowPlacement)
		})
	}
}

// 窗口透明
//func (m *LCLBrowserWindow) SetTransparentColor() {
//	m.SetColor(colors.ClNavy)
//	Exstyle := win.GetWindowLong(m.Handle(), win.GWL_EXSTYLE)
//	Exstyle = Exstyle | win.WS_EX_LAYERED&^win.WS_EX_TRANSPARENT
//	win.SetWindowLong(m.Handle(), win.GWL_EXSTYLE, uintptr(Exstyle))
//	win.SetLayeredWindowAttributes(m.Handle(),
//		colors.ClNavy, //crKey 指定需要透明的背景颜色值
//		255,           //bAlpha 设置透明度,0完全透明，255不透明
//		//LWA_ALPHA: crKey无效,bAlpha有效
//		//LWA_COLORKEY: 窗体中的所有颜色为crKey的地方全透明，bAlpha无效
//		//LWA_ALPHA | LWA_COLORKEY: crKey的地方全透明，其它地方根据bAlpha确定透明度
//		win.LWA_ALPHA|win.LWA_COLORKEY)
//}
