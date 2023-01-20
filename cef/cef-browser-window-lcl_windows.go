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
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

var ov = version.OSVersion

type customWindowCaption struct {
	canCaption bool                  //当前鼠标是否在标题栏区域
	regions    *TCefDraggableRegions //窗口内html拖拽区域
	rgn        *HRGN                 //
}

//显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	m.WindowProperty()._CanHideCaption = false
	//win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)|win.WS_CAPTION))
	//win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	if m.WindowProperty().CanMaximize {
		m.EnabledMaximize(true)
	}
	m.SetBorderStyle(types.BsSizeable)

}

//隐藏标题栏
func (m *LCLBrowserWindow) HideTitle() {
	m.WindowProperty()._CanHideCaption = true
	//win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)&^win.WS_CAPTION))
	//win.SetWindowPos(m.Handle(), 0, 0, 0, m.Width(), m.Height()+500, win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED|win.SWP_DRAWFRAME)
	//无标题栏情况会导致任务栏不能切换窗口，不知道为什么要这样设置一下
	m.EnabledMaximize(false)
	m.SetBorderStyle(types.BsNone)

}

func (m *customWindowCaption) freeRgn() {
	if m.rgn != nil {
		WinSetRectRgn(m.rgn, 0, 0, 0, 0)
		WinDeleteObject(m.rgn)
		m.rgn.Free()
	}
}
func (m *customWindowCaption) freeRegions() {
	if m.regions != nil {
		m.regions.regions = nil
		m.regions = nil
	}
}
func (m *customWindowCaption) free() {
	if m != nil {
		m.freeRgn()
		m.freeRegions()
	}
}

func (m *customWindowCaption) toPoint(message *types.TMessage) (x, y int32) {
	return int32(message.LParam & 0xFFFF), int32(message.LParam & 0xFFFF0000 >> 16)
}

//鼠标在标题栏区域
func (m *customWindowCaption) isCaption(hWND types.HWND, rgn *HRGN, message *types.TMessage) (x, y int32, caption bool) {
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
	if m.cwcap.regions != nil && m.cwcap.regions.RegionsCount() > 0 {
		switch message.Msg {
		case WM_NCLBUTTONDBLCLK: /*-- NC l d click --*/
			if !m.WindowProperty().CanCaptionDClkMaximize {
				return
			}
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
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
				} else {
					rtl.PostMessage(m.Handle(), WM_SYSCOMMAND, SC_RESTORE, 0)
				}
				rtl.SendMessage(m.Handle(), WM_NCLBUTTONUP, HTCAPTION, 0)
			}
		case WM_NCLBUTTONDOWN: //nc l down
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
				win.ReleaseCapture()
				rtl.PostMessage(m.Handle(), WM_NCLBUTTONDOWN, HTCAPTION, 0)
			}
		case WM_NCLBUTTONUP: //nc l up
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
			}
		case WM_NCRBUTTONDOWN: //nc r down
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
			}
		case WM_NCRBUTTONUP: //nc r up
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
			}
		case WM_NCHITTEST: /*-- NCHITTEST --*/
			if m.cwcap.rgn != nil {
				_, _, caption := m.cwcap.isCaption(m.Handle(), m.cwcap.rgn, message)
				//设置鼠标坐标是否在标题区域
				m.cwcap.canCaption = caption
				if caption {
					//如果光标在一个可拖动区域内，返回HTCAPTION允许拖动。
					*lResult = HTCAPTION
					*aHandled = true
				}
			}
		}
	}
}

func (m *LCLBrowserWindow) setDraggableRegions() {
	QueueAsyncCall(func(id int) {
		if m.cwcap.rgn == nil {
			m.cwcap.rgn = WinCreateRectRgn(0, 0, 0, 0)
		} else {
			WinSetRectRgn(m.cwcap.rgn, 0, 0, 0, 0)
		}
		for i := 0; i < m.cwcap.regions.RegionsCount(); i++ {
			region := m.cwcap.regions.Region(i)
			creRGN := WinCreateRectRgn(region.Bounds.X, region.Bounds.Y, region.Bounds.X+region.Bounds.Width, region.Bounds.Y+region.Bounds.Height)
			if region.Draggable {
				WinCombineRgn(m.cwcap.rgn, m.cwcap.rgn, creRGN, consts.RGN_OR)
			} else {
				WinCombineRgn(m.cwcap.rgn, m.cwcap.rgn, creRGN, consts.RGN_DIFF)
			}
			WinDeleteObject(creRGN)
		}
	})
}

// default event register: windows CompMsgEvent
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
		m.windowResize = func(sender lcl.IObject) bool {
			if m.WindowState() == types.WsMaximized && (m.WindowProperty()._CanHideCaption || m.BorderStyle() == types.BsNone || m.BorderStyle() == types.BsSingle) {
				var monitor = m.Monitor().WorkareaRect()
				m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left, monitor.Bottom-monitor.Top)
				m.SetWindowState(types.WsMaximized)
			}
			return false
		}
	} else {
		if bwEvent.onRenderCompMsg != nil {
			m.chromium.SetOnRenderCompMsg(bwEvent.onRenderCompMsg)
		}
	}
}

//for windows maximize and restore
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		win.ReleaseCapture()
		m.windowsState = m.WindowState()
		if m.windowsState == types.WsNormal {
			rtl.PostMessage(m.Handle(), WM_SYSCOMMAND, SC_MAXIMIZE, 0)
		} else {
			rtl.SendMessage(m.Handle(), WM_SYSCOMMAND, SC_RESTORE, 0)
		}
	})
}
