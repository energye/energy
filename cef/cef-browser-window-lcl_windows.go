//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

var ov = version.OSVersion

//显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	m.WindowProperty()._CanHideCaption = false
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)|win.WS_CAPTION))
	win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
}

//隐藏标题栏
func (m *LCLBrowserWindow) HideTitle() {
	m.WindowProperty()._CanHideCaption = true
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)&^win.WS_CAPTION))
	win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
}

//windows 窗口标题栏管理
var wdrs = &windowCaption{}

type windowCaption struct {
	canCaption bool
}

func (m *windowCaption) toPoint(message *types.TMessage) (x, y int32) {
	return int32(message.LParam & 0xFFFF), int32(message.LParam & 0xFFFF0000 >> 16)
}

//鼠标在标题栏区域
func (m *windowCaption) isCaption(hWND types.HWND, rgn *HRGN, message *types.TMessage) (x, y int32, caption bool) {
	dx, dy := m.toPoint(message)
	p := &types.TPoint{
		X: dx,
		Y: dy,
	}
	WinScreenToClient(hWND, p)
	m.canCaption = WinPtInRegion(rgn, p.X, p.Y)
	return p.X, p.Y, m.canCaption
}

func (m *LCLBrowserWindow) doOnRenderCompMsg(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.regions != nil && m.regions.RegionsCount() > 0 {
		switch message.Msg {
		case WM_LBUTTONDBLCLK:
		case WM_LBUTTONDOWN:
			fmt.Println("l up", wdrs.canCaption)
		case WM_NCLBUTTONDBLCLK: /*-- NC --*/
			if !m.WindowProperty().CanCaptionDClkMaximize {
				return
			}
			fmt.Println("nc ld click", m.windowsState, m.WindowState())
			if m.rgn != nil && wdrs.canCaption {
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_MAXIMIZE, 0); // 最大化
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_MINIMIZE, 0); // 最小化
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_CLOSE, 0); // 关闭
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_RESTORE, 0); // 最大化状态还原
				*lResult = HTCAPTION
				*aHandled = true
				win.ReleaseCapture()
				m.windowsState = m.WindowState()
				if m.windowsState == types.WsNormal {
					rtl.PostMessage(m.Handle(), WM_SYSCOMMAND, SC_MAXIMIZE, 0)
					//var monitor = m.Monitor().WorkareaRect()
					//m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left-1, monitor.Bottom-monitor.Top-1)
				} else {
					rtl.SendMessage(m.Handle(), WM_SYSCOMMAND, SC_RESTORE, 0)
				}
				rtl.SendMessage(m.Handle(), WM_NCLBUTTONUP, 0, 0)
			}
		case WM_NCLBUTTONDOWN:
			fmt.Println("nc l down", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
				win.ReleaseCapture()
				rtl.PostMessage(m.Handle(), WM_NCLBUTTONDOWN, HTCAPTION, 0)
				rtl.SendMessage(m.Handle(), WM_NCLBUTTONUP, 0, 0)
			}
		case WM_NCLBUTTONUP:
			fmt.Println("nc l up", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
			}
		case WM_NCRBUTTONDOWN:
			fmt.Println("nc r down", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
			}
		case WM_NCRBUTTONUP:
			fmt.Println("nc r up", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
			}
		case WM_NCHITTEST: /*-- NCHITTEST --*/
			if m.rgn != nil {
				_, _, caption := wdrs.isCaption(m.Handle(), m.rgn, message)
				//设置鼠标坐标是否在标题区域
				wdrs.canCaption = caption
				if caption {
					//如果光标在一个可拖动区域内，返回HTCAPTION允许拖动。
					*lResult = HTCAPTION
					*aHandled = true
				}
			}
		}
	}
}

// 默认事件注册 windows CompMsgEvent
func (m *LCLBrowserWindow) registerWindowsCompMsgEvent() {
	var bwEvent = BrowserWindow.browserEvent
	if m.WindowProperty().CanWebkitAppRegion {
		m.chromium.SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
			if bwEvent.onRenderCompMsg != nil {
				bwEvent.onRenderCompMsg(sender, message, lResult, aHandled)
			}
			if !*aHandled {
				m.doOnRenderCompMsg(message, lResult, aHandled)
			}
		})
		m.TForm.SetOnPaint(func(sender lcl.IObject) {
			//m.chromium.NotifyMoveOrResizeStarted()
			//m.windowParent.UpdateSize()
			var ret bool
			if m.onPaint != nil {
				ret = m.onPaint(sender)
			}
			if !ret {
				if m.WindowState() == types.WsMaximized && m.WindowProperty()._CanHideCaption {
					var monitor = m.Monitor().WorkareaRect()
					m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left, monitor.Bottom-monitor.Top)
					m.SetWindowState(types.WsMaximized)
					//rtl.PostMessage(m.Handle(), WM_NCLBUTTONDOWN, HTCAPTION, 0) //WM_PAINT
				}
			}
		})
	} else {
		if bwEvent.onRenderCompMsg != nil {
			m.chromium.SetOnRenderCompMsg(bwEvent.onRenderCompMsg)
		}
		if m.onPaint != nil {
			m.TForm.SetOnPaint(func(sender lcl.IObject) {
				m.onPaint(sender)
			})
		}
	}
}
