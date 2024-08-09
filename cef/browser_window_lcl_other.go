//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

// LCL窗口组件定义和实现-非windows平台

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/common"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableHideCaption = false
	m.SetBorderStyle(types.BsSizeable)
}

// 隐藏标题栏
func (m *LCLBrowserWindow) HideTitle() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableHideCaption = true
	m.SetBorderStyle(types.BsNone)
}

// 默认事件注册 windows 消息事件
func (m *LCLBrowserWindow) registerWindowsCompMsgEvent() {
	//TODO no impl
}

func (m *LCLBrowserWindow) setDraggableRegions() {
	//TODO no impl
}

// FramelessForLine 窗口四边框是一条细线
func (m *LCLBrowserWindow) FramelessForLine() {
	//TODO no impl
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

func (m *LCLBrowserWindow) Frameless() {

}

func (m *LCLBrowserWindow) taskMenu() {

}

// Restore 非Windows平台，窗口还原
func (m *LCLBrowserWindow) Restore() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		m.SetWindowState(types.WsNormal)
	})
}

// Minimize 非Windows平台，窗口最小化
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		m.SetWindowState(types.WsMinimized)
	})
}

// Maximize 非Windows平台，窗口最大化/还原
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		if m.WindowState() == types.WsMaximized {
			// 当前窗口是最大化状态 > 恢复窗口
			// 此时记录窗口状态
			m.WindowProperty().current.windowState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			if common.IsDarwin() { //要这样重复设置2次不然不启作用
				m.SetWindowState(types.WsMaximized)
				m.SetWindowState(types.WsNormal)
			}
			// 当前窗口如果是无标题栏窗口需要恢复到之前记录的窗口属性
			wp := m.WindowProperty()
			if wp.EnableHideCaption {
				m.SetBoundsRect(wp.current.previousWindowPlacement)
			}
		} else if m.WindowState() == types.WsNormal {
			// 当前状态是正常的 > 将窗口最大化
			// 在无标题栏窗口时，最大化和全屏正常是无法改变窗口状态
			// 因此需要自己处理窗口大小, 在此之前需要记录窗口状态
			m.WindowProperty().current.windowState = types.WsMaximized
			if m.WindowProperty().EnableHideCaption {
				// 无标题窗口时调整窗口大小，设置为工作窗口大小
				m.SetBoundsRect(m.Monitor().WorkareaRect())
			}
			// 触发窗口最大化
			m.SetWindowState(types.WsMaximized)
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
			//m.setCurrentProperty()
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
	wp := m.WindowProperty()
	if wp.EnableHideCaption {
		if m.IsFullScreen() {
			RunOnMainThread(func() {
				wp.current.windowState = types.WsNormal
				//m.SetBounds(wp.current.x, wp.current.y, wp.current.w, wp.current.h)
				m.SetBoundsRect(m.WindowProperty().current.previousWindowPlacement)
				m.SetWindowState(types.WsNormal)
			})
		}
	}
}

// IsFullScreen 是否全屏
func (m *LCLBrowserWindow) IsFullScreen() bool {
	return m.WindowProperty().current.windowState == types.WsFullScreen
}

// SetFocus 设置窗口焦点
func (m *LCLBrowserWindow) SetFocus() {
	if m.TForm != nil {
		m.TForm.SetFocus()
	}
}

func (m *LCLBrowserWindow) doDrag() {
	// MacOS/Linux Drag Window
	if m.drag != nil {
		m.drag.drag()
	}
}
