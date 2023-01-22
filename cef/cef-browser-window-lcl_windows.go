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
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

var ov = version.OSVersion

type customWindowCaption struct {
	canCaption      bool                  //当前鼠标是否在标题栏区域
	canBorder       bool                  //当前鼠标是否在边框
	borderDirection int                   //当前鼠标所在边框的方向
	regions         *TCefDraggableRegions //窗口内html拖拽区域
	rgn             *HRGN                 //
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

//NC 鼠标移动
func (m *customWindowCaption) onNCMouseMove(hWND types.HWND, message *types.TMessage) {

}

//NC 鼠标左键按下
func (m *customWindowCaption) onNCLButtonDown(hWND types.HWND, message *types.TMessage) {

}

//设置鼠标图标
func (m *customWindowCaption) onSetCursor(hWND types.HWND, message *types.TMessage) {

}

//鼠标是否在边框
func (m *customWindowCaption) onCanBorder(hWND types.HWND, message *types.TMessage) {

}

//转换XY坐标
func (m *customWindowCaption) toPoint(message *types.TMessage) (x, y uint16) {
	return LOWORD(message.LParam), HIWORD(message.LParam)
}

//鼠标在标题栏区域
func (m *customWindowCaption) isCaption(hWND types.HWND, rgn *HRGN, message *types.TMessage) (int32, int32, bool) {
	dx, dy := m.toPoint(message)
	p := &types.TPoint{
		X: int32(dx),
		Y: int32(dy),
	}
	WinScreenToClient(hWND, p)
	m.canCaption = WinPtInRegion(rgn, p.X, p.Y)
	return p.X, p.Y, m.canCaption
}

//定义四角和边框范围
var (
	angleRange  int32 = 8
	broderRange int32 = 4
)

func (m *LCLBrowserWindow) doOnRenderCompMsg(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.cwcap.regions != nil && m.cwcap.regions.RegionsCount() > 0 {
		switch message.Msg {
		case WM_NCLBUTTONDBLCLK: // 163 NC left dclick
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
		case WM_NCLBUTTONDOWN: // 161 nc left down
			if m.cwcap.rgn != nil && m.cwcap.canCaption { //拖拽区域已设置，并且在标题栏
				*lResult = HTCAPTION
				*aHandled = true
				win.ReleaseCapture()
				rtl.PostMessage(m.Handle(), WM_NCLBUTTONDOWN, HTCAPTION, 0)
			} else {
				if m.cwcap.canBorder { //当前在边框
					*lResult = types.LRESULT(m.cwcap.borderDirection)
					*aHandled = true
					win.ReleaseCapture()
					rtl.PostMessage(m.Handle(), WM_SYSCOMMAND, SC_SIZE|WMSZ_BOTTOMRIGHT, rtl.MakeLParam(m.cwcap.toPoint(message)))
				}
			}
		case WM_NCLBUTTONUP: // 162 nc l up
			fmt.Println("nc l up")
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
			}
		case WM_NCRBUTTONDOWN: // nc r down
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
			}
		case WM_NCRBUTTONUP: // nc r up
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
			}
		case WM_NCMOUSEMOVE: // 160 nc mouse move
			//非客户区边框鼠标移动
			if m.cwcap.canBorder { //当前在边框
				//mx, my := m.cwcap.toPoint(message)
				//fmt.Println("MOVE 当前在边框:", m.cwcap.borderDirection, "dx-dy:", dx, dy, "mx-my:", mx, my, "dm:", dx-mx, dy-my)
				*lResult = types.LRESULT(m.cwcap.borderDirection)
				*aHandled = true
			}
		case WM_SETCURSOR: // 32 设置鼠标图标样式
			//设置鼠标图标样式
			if m.cwcap.canBorder { //当前在边框
				switch LOWORD(message.LParam) {
				case HTBOTTOMRIGHT: //右下
					fmt.Println("WM_SETCURSOR", LOWORD(message.LParam))
					WinSetCursor(WinLoadCursor(0, IDC_SIZENWSE))
					*lResult = HTBOTTOMRIGHT
					*aHandled = true
				}
			}
		case WM_NCHITTEST: // 132 NCHITTEST
			if m.cwcap.rgn != nil {
				x, y, caption := m.cwcap.isCaption(m.Handle(), m.cwcap.rgn, message)
				//设置鼠标坐标是否在标题区域
				m.cwcap.canCaption = caption
				if caption { //窗口标题栏
					*lResult = HTCAPTION
					*aHandled = true
				} else if m.WindowProperty()._CanHideCaption && m.WindowProperty().CanResize && m.WindowState() == types.WsNormal { //1.窗口隐藏标题栏 2.启用了调整窗口大小 3.非最大化、最小化、全屏状态
					if m.cwcap.canBorder = x <= m.Width() && x >= m.Width()-angleRange && y <= m.Height() && y >= m.Height()-angleRange; m.cwcap.canBorder { // 1.右下
						m.cwcap.borderDirection = HTBOTTOMRIGHT
						*lResult = HTBOTTOMRIGHT
						*aHandled = true
					} else if m.cwcap.canBorder = x <= m.Width() && x >= m.Width()-5 && y <= m.Height() && y >= m.Height()-5; m.cwcap.canBorder { // 2.右

					}
					return
				}
				m.cwcap.canBorder = false
			}
		}
	}
}

//每一次拖拽区域改变都需要重新设置
func (m *LCLBrowserWindow) setDraggableRegions() {
	//在主线程中运行
	QueueAsyncCall(func(id int) {
		if m.cwcap.rgn == nil {
			//第一次时创建RGN
			m.cwcap.rgn = WinCreateRectRgn(0, 0, 0, 0)
		} else {
			//每次重置RGN
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
