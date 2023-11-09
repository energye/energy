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

// LCL窗口组件定义和实现-非windows平台

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/common"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// TODO no
type customWindowCaption struct {
	bw      *LCLBrowserWindow
	regions *TCefDraggableRegions
}

func (m *customWindowCaption) free() {
	//TODO no
}

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

// Restore 非Windows平台，窗口还原
func (m *LCLBrowserWindow) Restore() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		m.SetWindowState(types.WsNormal)
	})
}

// Minimize 非Windows平台，窗口最小化
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		m.SetWindowState(types.WsMinimized)
	})
}

// Maximize 非Windows平台，窗口最大化/还原
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		//var bs = m.BorderStyle()
		//var monitor = m.Monitor().WorkareaRect()
		//这个if判断应该不会执行了，windows有自己的，linux是VF的，mac走else
		//if (bs == types.BsNone || bs == types.BsSingle) && !common.IsDarwin() { //不能调整窗口大，所以手动控制窗口状态
		//	var ws = m.WindowState()
		//	var redWindowState types.TWindowState
		//	//默认状态0
		//	if m.windowProperty.windowState == types.WsNormal && m.windowProperty.windowState == ws {
		//		redWindowState = types.WsMaximized
		//	} else {
		//		if m.windowProperty.windowState == types.WsNormal {
		//			redWindowState = types.WsMaximized
		//		} else if m.windowProperty.windowState == types.WsMaximized {
		//			redWindowState = types.WsNormal
		//		}
		//	}
		//	m.windowProperty.windowState = redWindowState
		//	if redWindowState == types.WsMaximized {
		//		m.windowProperty.X = m.Left()
		//		m.windowProperty.Y = m.Top()
		//		m.windowProperty.Width = m.Width()
		//		m.windowProperty.Height = m.Height()
		//		m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left-1, monitor.Bottom-monitor.Top-1)
		//	} else if redWindowState == types.WsNormal {
		//		m.SetBounds(m.windowProperty.X, m.windowProperty.Y, m.windowProperty.Width, m.windowProperty.Height)
		//	}
		//	m.SetWindowState(redWindowState)
		//} else {
		if m.WindowState() == types.WsMaximized {
			m.SetWindowState(types.WsNormal)
			if common.IsDarwin() { //要这样重复设置2次不然不启作用
				m.SetWindowState(types.WsMaximized)
				m.SetWindowState(types.WsNormal)
			}
		} else if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		}
		//m.windowProperty.windowState = m.WindowState()
		//}
	})
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
