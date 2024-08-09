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

// LCL窗口组件定义和实现-windows平台

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts/messages"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

// 定义四角和边框范围
var (
	angleRange  int32 = 10 //四角
	borderRange int32 = 5  //四边框
)

// 组件消息类型
type compMessageType int8

const (
	cmtCEF compMessageType = iota
	cmtLCL
)

// ShowTitle 显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
}

// HideTitle 隐藏标题栏 无边框样式
func (m *LCLBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
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

// SetFocus
//
//	在窗口 (Visible = true) 显示之后设置窗口焦点
//	https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-showwindow
//	https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setfocus
func (m *LCLBrowserWindow) SetFocus() {
	if m.TForm != nil {
		m.Visible()
		//窗口激活在Z序中的下个顶层窗口
		m.Minimize()
		//激活窗口出现在前景
		m.Restore()
		//窗口设置焦点
		m.TForm.SetFocus()
	}
}

// Restore Windows平台，窗口还原
func (m *LCLBrowserWindow) Restore() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		//if win.ReleaseCapture() {
		//	win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
		//}
		m.SetWindowState(types.WsNormal)
	})
}

// Minimize Windows平台，窗口最小化
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		//if win.ReleaseCapture() {
		//	win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MINIMIZE, 0)
		//}
		m.SetWindowState(types.WsMinimized)
	})
}

// Maximize Windows平台，窗口最大化/还原
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil || m.IsFullScreen() {
		return
	}
	RunOnMainThread(func() {
		//if win.ReleaseCapture() {
		//	if m.WindowState() == types.WsNormal {
		//		win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
		//	} else {
		//		win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
		//	}
		//}
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	})
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
			//m.setCurrentProperty()
			//style := uint32(win.GetWindowLongPtr(m.Handle(), win.GWL_STYLE))
			monitorRect := m.Monitor().BoundsRect()
			win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
		})
	}
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	wp := m.WindowProperty()
	if wp.EnableHideCaption {
		if m.IsFullScreen() {
			RunOnMainThread(func() {
				wp.current.windowState = types.WsNormal
				m.SetWindowState(types.WsNormal)
				//m.SetBounds(wp.current.x, wp.current.y, wp.current.w, wp.current.h)
				m.SetBoundsRect(m.WindowProperty().current.previousWindowPlacement)
			})
		}
	}
}

// IsFullScreen 是否全屏
func (m *LCLBrowserWindow) IsFullScreen() bool {
	return m.WindowProperty().current.windowState == types.WsFullScreen
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

func (m *LCLBrowserWindow) doDrag() {
	if m.drag != nil {
		m.drag.drag()
	}
}
