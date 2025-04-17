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
	"github.com/cyber-xxm/energy/v2/consts/messages"
	"github.com/energye/golcl/lcl/win"
)

func (m *drag) drag() {
	window := m.window.AsLCLBrowserWindow().BrowserWindow()
	switch m.T {
	case dragUp:
	case dragDown:
	case dragMove:
		// 全屏时不能拖拽窗口
		if window.IsFullScreen() {
			return
		}
		// 此时是 down 事件, 拖拽窗口
		if win.ReleaseCapture() {
			win.PostMessage(m.window.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case dragResize:
		if window.IsFullScreen() || !window.WindowProperty().EnableResize {
			return
		}
		var borderHT uintptr
		switch m.HT {
		case "n-resize":
			borderHT = messages.HTTOP
		case "ne-resize":
			borderHT = messages.HTTOPRIGHT
		case "e-resize":
			borderHT = messages.HTRIGHT
		case "se-resize":
			borderHT = messages.HTBOTTOMRIGHT
		case "s-resize":
			borderHT = messages.HTBOTTOM
		case "sw-resize":
			borderHT = messages.HTBOTTOMLEFT
		case "w-resize":
			borderHT = messages.HTLEFT
		case "nw-resize":
			borderHT = messages.HTTOPLEFT
		}
		if borderHT != 0 {
			if win.ReleaseCapture() {
				win.PostMessage(window.Handle(), messages.WM_NCLBUTTONDOWN, borderHT, 0)
			}
		}
	case dragDblClick:
		if window.WindowProperty().EnableWebkitAppRegionDClk {
			window.Maximize()
		}
	}

}
