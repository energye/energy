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
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)|win.WS_CAPTION))
	win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
}

//隐藏标题栏
func (m *LCLBrowserWindow) HideTitle() {
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)&^win.WS_CAPTION))
	win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
}

//windows 窗口拖拽区域管理
var wdrs = &windowDragRegions{}

type windowDragRegions struct {
	canCaption bool
}

func (m *windowDragRegions) toPoint(message *types.TMessage) (x, y int32) {
	return int32(message.LParam & 0xFFFF), int32(message.LParam & 0xFFFF0000 >> 16)
}

//鼠标在标题栏区域
func (m *windowDragRegions) isCaption(hWND types.HWND, rgn *HRGN, message *types.TMessage) (x, y int32, caption bool) {
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
		case WM_GETMINMAXINFO:
			fmt.Println("WM_GETMINMAXINFO")
		case WM_MOUSEMOVE:
			//fmt.Println("move")
		case WM_LBUTTONDBLCLK:
		case WM_LBUTTONDOWN:
		case WM_LBUTTONUP:
			fmt.Println("l up", wdrs.canCaption)
		case WM_MOVE, WM_MOVING:
			fmt.Println("move")

		case WM_NCLBUTTONDBLCLK: /*-- NC --*/
			if !m.WindowProperty().CanCaptionDClkMaximize {
				break
			}
			fmt.Println("nc ld click", m.windowsState)
			if m.rgn != nil && wdrs.canCaption {
				if ov.Major == 6 {
					//win.ReleaseCapture()
				}
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_MAXIMIZE, 0); // 最大化
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_MINIMIZE, 0); // 最小化
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_CLOSE, 0); // 关闭
				//SendMessage(hwnd, WM_SYSCOMMAND, SC_RESTORE, 0); // 最大化状态还原

				if m.windowsState == 0 {
					m.windowsState = types.WsMaximized
					m.SetWindowState(types.WsMaximized)
					//var monitor = m.Monitor().WorkareaRect()
					//m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left-1, monitor.Bottom-monitor.Top-1)
				} else {
					//需要先设置一次-不然不生效
					m.SetWindowState(m.windowsState)
					m.windowsState = types.WsNormal
					m.SetWindowState(types.WsNormal)
				}
				return
			}
		case WM_NCLBUTTONDOWN:
			fmt.Println("nc l down", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
				//1 这个在windows7不正确
				//WinDefWindowProc(m.hWnd, message.Msg, message.WParam, message.LParam)

				//2 这个在windows7下正常
				win.ReleaseCapture()
				rtl.SendMessage(m.Handle(), WM_NCLBUTTONDOWN, HTCAPTION, 0)
				rtl.SendMessage(m.Handle(), WM_NCLBUTTONUP, 0, 0)
				//rtl.SendMessage(m.Handle(), WM_NCLBUTTONUP, 0, 0)
				return
			}
		case WM_NCLBUTTONUP:
			fmt.Println("nc l up", wdrs.canCaption)
		case WM_NCMOUSEMOVE:
			fmt.Println("nc mouse move", wdrs.canCaption)
			if m.rgn != nil && wdrs.canCaption {
				WinDefWindowProc(m.Handle(), message.Msg, message.WParam, message.LParam)
				*lResult = HTCAPTION
				*aHandled = true
				break
			}
		case WM_NCRBUTTONDOWN:
			fmt.Println("nc r down", wdrs.canCaption)
		case WM_NCRBUTTONUP:
			fmt.Println("nc r up", wdrs.canCaption)
		case WM_NCHITTEST: /*-- NCHITTEST --*/
			if m.rgn != nil {
				//var hit = WinDefWindowProc(m.Handle(), message.Msg, message.WParam, message.LParam)
				_, _, caption := wdrs.isCaption(m.Handle(), m.rgn, message)
				//设置鼠标坐标是否在标题区域
				wdrs.canCaption = caption
				if caption {
					//如果光标在一个可拖动区域内，返回HTCAPTION允许拖动。
					*lResult = HTCAPTION
					*aHandled = true
					return
				}
				//*lResult = hit
				return
			}
		}
		//other message -> WinDefWindowProc
		*lResult = WinDefWindowProc(m.Handle(), message.Msg, message.WParam, message.LParam)
		//m.InheritedWndProc(message)
	}
}

// 默认事件注册 windows 消息事件
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
		m.TForm.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
			m.windowsState = m.WindowState()
			if m.onConstrainedResizeEvent != nil {
				m.onConstrainedResizeEvent(sender, minWidth, minHeight, maxWidth, maxHeight)
			}
		})
		m.TForm.SetOnPaint(func(sender lcl.IObject) {
			fmt.Println("OnPaint", m.WindowState()) //TODO test
			//m.chromium.NotifyMoveOrResizeStarted()
			//m.windowParent.UpdateSize()
		})
	} else {
		if bwEvent.onRenderCompMsg != nil {
			m.chromium.SetOnRenderCompMsg(bwEvent.onRenderCompMsg)
		}
		if m.onConstrainedResizeEvent != nil {
			m.TForm.SetOnConstrainedResize(m.onConstrainedResizeEvent)
		}
	}
}
