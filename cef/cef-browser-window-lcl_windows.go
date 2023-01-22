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
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

//定义四角和边框范围
var (
	angleRange  int32 = 10 //四角
	borderRange int32 = 5  //四边框
)

type customWindowCaption struct {
	canCaption           bool                  //当前鼠标是否在标题栏区域
	canBorder            bool                  //当前鼠标是否在边框
	borderHT, borderWMSZ int                   //borderHT: 鼠标所在边框位置, borderWMSZ: 窗口改变大小边框方向 borderMD:
	borderMD             bool                  //borderMD: 鼠标调整窗口大小，已按下后，再次接收到132消息应该忽略该消息
	regions              *TCefDraggableRegions //窗口内html拖拽区域
	rgn                  *HRGN                 //
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

//NC 非客户区鼠标移动
func (m *customWindowCaption) onNCMouseMove(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canCaption { // 当前在标题栏
	} else if m.canBorder { // 当前在边框
		*lResult = types.LRESULT(m.borderHT)
		*aHandled = true
	}
}

//设置鼠标图标
func (m *customWindowCaption) onSetCursor(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canBorder { //当前在边框
		switch LOWORD(message.LParam) {
		case HTBOTTOMRIGHT, HTTOPLEFT: //右下 左上
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			WinSetCursor(WinLoadCursor(0, IDC_SIZENWSE))
		case HTRIGHT, HTLEFT: //右 左
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			WinSetCursor(WinLoadCursor(0, IDC_SIZEWE))
		case HTTOPRIGHT, HTBOTTOMLEFT: //右上 左下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			WinSetCursor(WinLoadCursor(0, IDC_SIZENESW))
		case HTTOP, HTBOTTOM: //上 下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			WinSetCursor(WinLoadCursor(0, IDC_SIZENS))
		}
	}
}

//鼠标是否在边框
func (m *customWindowCaption) onCanBorder(x, y int32, rect *types.TRect) (int, bool) {
	if m.canBorder = x <= rect.Width() && x >= rect.Width()-angleRange && y <= angleRange; m.canBorder { // 右上
		m.borderWMSZ = WMSZ_TOPRIGHT
		m.borderHT = HTTOPRIGHT
		return m.borderHT, true
	} else if m.canBorder = x <= rect.Width() && x >= rect.Width()-angleRange && y <= rect.Height() && y >= rect.Height()-angleRange; m.canBorder { // 右下
		m.borderWMSZ = WMSZ_BOTTOMRIGHT
		m.borderHT = HTBOTTOMRIGHT
		return m.borderHT, true
	} else if m.canBorder = x <= angleRange && y <= angleRange; m.canBorder { //左上
		m.borderWMSZ = WMSZ_TOPLEFT
		m.borderHT = HTTOPLEFT
		return m.borderHT, true
	} else if m.canBorder = x <= angleRange && y >= rect.Height()-angleRange; m.canBorder { //左下
		m.borderWMSZ = WMSZ_BOTTOMLEFT
		m.borderHT = HTBOTTOMLEFT
		return m.borderHT, true
	} else if m.canBorder = x > angleRange && x < rect.Width()-angleRange && y <= borderRange; m.canBorder { //上
		m.borderWMSZ = WMSZ_TOP
		m.borderHT = HTTOP
		return m.borderHT, true
	} else if m.canBorder = x > angleRange && x < rect.Width()-angleRange && y >= rect.Height()-borderRange; m.canBorder { //下
		m.borderWMSZ = WMSZ_BOTTOM
		m.borderHT = HTBOTTOM
		return m.borderHT, true
	} else if m.canBorder = x <= borderRange && y > angleRange && y < rect.Height()-angleRange; m.canBorder { //左
		m.borderWMSZ = WMSZ_LEFT
		m.borderHT = HTLEFT
		return m.borderHT, true
	} else if m.canBorder = x <= rect.Width() && x >= rect.Width()-borderRange && y > angleRange && y < rect.Height()-angleRange; m.canBorder { // 右
		m.borderWMSZ = WMSZ_RIGHT
		m.borderHT = HTRIGHT
		return m.borderHT, true
	}
	return 0, false
}

//NC 鼠标左键按下
func (m *customWindowCaption) onNCLButtonDown(hWND types.HWND, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canCaption { // 标题栏
		*lResult = HTCAPTION
		m.borderMD = true
		*aHandled = true
		win.ReleaseCapture()
		rtl.PostMessage(hWND, WM_NCLBUTTONDOWN, HTCAPTION, rtl.MakeLParam(m.toPoint(message)))
	} else if m.canBorder { // 边框
		*lResult = types.LRESULT(m.borderHT)
		m.borderMD = true
		*aHandled = true
		win.ReleaseCapture()
		rtl.PostMessage(hWND, WM_SYSCOMMAND, uintptr(SC_SIZE|m.borderWMSZ), rtl.MakeLParam(m.toPoint(message)))
	}
}

//转换XY坐标
func (m *customWindowCaption) toPoint(message *types.TMessage) (x, y uint16) {
	return LOWORD(message.LParam), HIWORD(message.LParam)
}

//鼠标在标题栏区域
func (m *customWindowCaption) isCaption(hWND types.HWND, message *types.TMessage) (int32, int32, bool) {
	dx, dy := m.toPoint(message)
	p := &types.TPoint{
		X: int32(dx),
		Y: int32(dy),
	}
	WinScreenToClient(hWND, p)
	m.canCaption = WinPtInRegion(m.rgn, p.X, p.Y)
	return p.X, p.Y, m.canCaption
}

func (m *LCLBrowserWindow) doOnRenderCompMsg(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.cwcap.regions != nil && m.cwcap.regions.RegionsCount() > 0 {
		//fmt.Println("msg:", message.Msg)
		switch message.Msg {
		//case WM_SIZE, WM_SIZING:
		//	if m.cwcap.canBorder {
		//		*lResult = types.LRESULT(m.cwcap.borderHT)
		//		*aHandled = true
		//	}
		//case WM_NCRBUTTONDOWN: // nc r down
		//	if m.cwcap.rgn != nil && m.cwcap.canCaption {
		//	}
		//case WM_NCRBUTTONUP: // nc r up
		//	if m.cwcap.rgn != nil && m.cwcap.canCaption {
		//	}
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
			m.cwcap.onNCLButtonDown(m.Handle(), message, lResult, aHandled)
		case WM_NCLBUTTONUP: // 162 nc l up
			fmt.Println("nc l up")
			if m.cwcap.rgn != nil && m.cwcap.canCaption {
				*lResult = HTCAPTION
				*aHandled = true
			}
		case WM_NCMOUSEMOVE: // 160 nc mouse move
			m.cwcap.onNCMouseMove(message, lResult, aHandled)
		case WM_SETCURSOR: // 32 设置鼠标图标样式
			m.cwcap.onSetCursor(message, lResult, aHandled)
		case WM_NCHITTEST: // 132 NCHITTEST
			if m.cwcap.rgn != nil {
				if m.cwcap.borderMD { //TODO 测试windows7, 161消息之后再次处理132消息导致消息错误
					m.cwcap.borderMD = false
					return
				}
				x, y, caption := m.cwcap.isCaption(m.Handle(), message)
				//设置鼠标坐标是否在标题区域
				if caption { //窗口标题栏
					*lResult = HTCAPTION
					*aHandled = true
				} else if m.WindowProperty()._CanHideCaption && m.WindowProperty().CanResize && m.WindowState() == types.WsNormal { //1.窗口隐藏标题栏 2.启用了调整窗口大小 3.非最大化、最小化、全屏状态
					rect := m.BoundsRect()
					if result, handled := m.cwcap.onCanBorder(x, y, &rect); handled {
						*lResult = types.LRESULT(result)
						*aHandled = true
					}
				}
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
