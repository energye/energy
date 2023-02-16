//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

import (
	"github.com/energye/energy/common"
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

}

func (m *LCLBrowserWindow) setDraggableRegions() {
}

// for other platform maximize and restore
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		var bs = m.BorderStyle()
		var monitor = m.Monitor().WorkareaRect()
		if m.windowProperty == nil {
			m.windowProperty = &WindowProperty{}
		}
		//这个if判断应该不会执行了，windows有自己的，linux是VF的，mac走else
		if (bs == types.BsNone || bs == types.BsSingle) && !common.IsDarwin() { //不能调整窗口大，所以手动控制窗口状态
			var ws = m.WindowState()
			var redWindowState types.TWindowState
			//默认状态0
			if m.windowProperty.windowState == types.WsNormal && m.windowProperty.windowState == ws {
				redWindowState = types.WsMaximized
			} else {
				if m.windowProperty.windowState == types.WsNormal {
					redWindowState = types.WsMaximized
				} else if m.windowProperty.windowState == types.WsMaximized {
					redWindowState = types.WsNormal
				}
			}
			m.windowProperty.windowState = redWindowState
			if redWindowState == types.WsMaximized {
				m.windowProperty.X = m.Left()
				m.windowProperty.Y = m.Top()
				m.windowProperty.Width = m.Width()
				m.windowProperty.Height = m.Height()
				m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left-1, monitor.Bottom-monitor.Top-1)
			} else if redWindowState == types.WsNormal {
				m.SetBounds(m.windowProperty.X, m.windowProperty.Y, m.windowProperty.Width, m.windowProperty.Height)
			}
			m.SetWindowState(redWindowState)
		} else {
			if m.WindowState() == types.WsMaximized {
				m.SetWindowState(types.WsNormal)
				if common.IsDarwin() { //要这样重复设置2次不然不启作用
					m.SetWindowState(types.WsMaximized)
					m.SetWindowState(types.WsNormal)
				}
			} else if m.WindowState() == types.WsNormal {
				m.SetWindowState(types.WsMaximized)
			}
			m.windowProperty.windowState = m.WindowState()
		}
	})
}
