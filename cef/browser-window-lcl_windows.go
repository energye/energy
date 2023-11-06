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

// LCL窗口组件定义和实现-windows平台

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/consts/messages"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

// 定义四角和边框范围
var (
	angleRange  int32 = 10 //四角
	borderRange int32 = 5  //四边框
)

// customWindowCaption 自定义窗口标题栏
//
// 隐藏窗口标题栏，通过html+css实现自定义窗口标题栏，实现窗口拖拽等
type customWindowCaption struct {
	bw                   *LCLBrowserWindow     //
	canCaption           bool                  //当前鼠标是否在标题栏区域
	canBorder            bool                  //当前鼠标是否在边框
	borderHT, borderWMSZ int                   //borderHT: 鼠标所在边框位置, borderWMSZ: 窗口改变大小边框方向 borderMD:
	borderMD             bool                  //borderMD: 鼠标调整窗口大小，已按下后，再次接收到132消息应该忽略该消息
	regions              *TCefDraggableRegions //窗口内html拖拽区域
	rgn                  *et.HRGN              //
}

// ShowTitle 显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	//win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)|win.WS_CAPTION))
	//win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
	m.SetBorderStyle(types.BsSizeable)
}

// HideTitle 隐藏标题栏 无边框样式
func (m *LCLBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	m.SetBorderStyle(types.BsNone)
	//m.RoundRectRgn()
}

// RoundRectRgn 窗口无边框时圆角
func (m *LCLBrowserWindow) RoundRectRgn() {
	hnd := winapi.WinCreateRoundRectRgn(0, 0, et.LongInt(m.Width()), et.LongInt(m.Height()), 10, 10)
	winapi.WinSetWindowRgn(et.HWND(m.Handle()), hnd, true)
}

// FramelessForDefault 窗口四边框系统默认样式
//  TODO 窗口顶部有条线
func (m *LCLBrowserWindow) FramelessForDefault() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_BORDER|win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))

	// SetClassLong(this.Handle, GCL_STYLE, GetClassLong(this.Handle, GCL_STYLE) | CS_DropSHADOW);
	//gclStyle := winapi.WinGetClassLongPtr(et.HWND(m.Handle()), messages.GCL_STYLE)
	//winapi.WinSetClassLongPtr(et.HWND(m.Handle()), messages.GCL_STYLE, gclStyle|messages.CS_DROPSHADOW)
}

// FramelessForLine 窗口四边框是一条细线
func (m *LCLBrowserWindow) FramelessForLine() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME|win.WS_BORDER))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

// SetFocus
//  在窗口 (Visible = true) 显示之后设置窗口焦点
//  https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-showwindow
//  https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setfocus
func (m *LCLBrowserWindow) SetFocus() {
	if m.TForm != nil {
		m.Visible()
		//窗口\激活在Z序中的下个顶层窗口
		m.Minimize()
		//激活窗口出现在前景
		m.Restore()
		//窗口设置焦点
		m.TForm.SetFocus()
	}
}

//func (m *LCLBrowserWindow) Frameless() {
//	var rect = &types.TRect{}
//	win.GetWindowRect(m.Handle(), rect)
//	win.SetWindowPos(m.Handle(), 0, rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, win.SWP_FRAMECHANGED)
//}

// freeRgn
func (m *customWindowCaption) freeRgn() {
	if m.rgn != nil {
		winapi.WinSetRectRgn(m.rgn, 0, 0, 0, 0)
		winapi.WinDeleteObject(m.rgn)
		m.rgn.Free()
	}
}

// freeRegions
func (m *customWindowCaption) freeRegions() {
	if m.regions != nil {
		m.regions.regions = nil
		m.regions = nil
	}
}

// free
func (m *customWindowCaption) free() {
	if m != nil {
		m.freeRgn()
		m.freeRegions()
	}
}

// onNCMouseMove NC 非客户区鼠标移动
func (m *customWindowCaption) onNCMouseMove(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canCaption { // 当前在标题栏
	} else if m.canBorder { // 当前在边框
		*lResult = types.LRESULT(m.borderHT)
		*aHandled = true
	}
}

// onSetCursor 设置鼠标图标
func (m *customWindowCaption) onSetCursor(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canBorder { //当前在边框
		switch winapi.LOWORD(uint32(message.LParam)) {
		case messages.HTBOTTOMRIGHT, messages.HTTOPLEFT: //右下 左上
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.WinSetCursor(winapi.WinLoadCursor(0, messages.IDC_SIZENWSE))
		case messages.HTRIGHT, messages.HTLEFT: //右 左
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.WinSetCursor(winapi.WinLoadCursor(0, messages.IDC_SIZEWE))
		case messages.HTTOPRIGHT, messages.HTBOTTOMLEFT: //右上 左下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.WinSetCursor(winapi.WinLoadCursor(0, messages.IDC_SIZENESW))
		case messages.HTTOP, messages.HTBOTTOM: //上 下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.WinSetCursor(winapi.WinLoadCursor(0, messages.IDC_SIZENS))
		}
	}
}

// onCanBorder 鼠标是否在边框
func (m *customWindowCaption) onCanBorder(x, y int32, rect *types.TRect) (int, bool) {
	if m.canBorder = x <= rect.Width() && x >= rect.Width()-angleRange && y <= angleRange; m.canBorder { // 右上
		m.borderWMSZ = messages.WMSZ_TOPRIGHT
		m.borderHT = messages.HTTOPRIGHT
		return m.borderHT, true
	} else if m.canBorder = x <= rect.Width() && x >= rect.Width()-angleRange && y <= rect.Height() && y >= rect.Height()-angleRange; m.canBorder { // 右下
		m.borderWMSZ = messages.WMSZ_BOTTOMRIGHT
		m.borderHT = messages.HTBOTTOMRIGHT
		return m.borderHT, true
	} else if m.canBorder = x <= angleRange && y <= angleRange; m.canBorder { //左上
		m.borderWMSZ = messages.WMSZ_TOPLEFT
		m.borderHT = messages.HTTOPLEFT
		return m.borderHT, true
	} else if m.canBorder = x <= angleRange && y >= rect.Height()-angleRange; m.canBorder { //左下
		m.borderWMSZ = messages.WMSZ_BOTTOMLEFT
		m.borderHT = messages.HTBOTTOMLEFT
		return m.borderHT, true
	} else if m.canBorder = x > angleRange && x < rect.Width()-angleRange && y <= borderRange; m.canBorder { //上
		m.borderWMSZ = messages.WMSZ_TOP
		m.borderHT = messages.HTTOP
		return m.borderHT, true
	} else if m.canBorder = x > angleRange && x < rect.Width()-angleRange && y >= rect.Height()-borderRange; m.canBorder { //下
		m.borderWMSZ = messages.WMSZ_BOTTOM
		m.borderHT = messages.HTBOTTOM
		return m.borderHT, true
	} else if m.canBorder = x <= borderRange && y > angleRange && y < rect.Height()-angleRange; m.canBorder { //左
		m.borderWMSZ = messages.WMSZ_LEFT
		m.borderHT = messages.HTLEFT
		return m.borderHT, true
	} else if m.canBorder = x <= rect.Width() && x >= rect.Width()-borderRange && y > angleRange && y < rect.Height()-angleRange; m.canBorder { // 右
		m.borderWMSZ = messages.WMSZ_RIGHT
		m.borderHT = messages.HTRIGHT
		return m.borderHT, true
	}
	return 0, false
}

// onNCLButtonDown NC 鼠标左键按下
func (m *customWindowCaption) onNCLButtonDown(hWND types.HWND, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canCaption { // 标题栏
		*lResult = messages.HTCAPTION
		*aHandled = true
		//全屏时不能移动窗口
		if m.bw.WindowProperty().current.ws == types.WsFullScreen {
			return
		}
		x, y := m.toPoint(message)
		m.borderMD = true
		if win.ReleaseCapture() {
			win.PostMessage(hWND, messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, rtl.MakeLParam(uint16(x), uint16(y)))
		}
	} else if m.canBorder { // 边框
		x, y := m.toPoint(message)
		*lResult = types.LRESULT(m.borderHT)
		*aHandled = true
		m.borderMD = true
		if win.ReleaseCapture() {
			win.PostMessage(hWND, messages.WM_SYSCOMMAND, uintptr(messages.SC_SIZE|m.borderWMSZ), rtl.MakeLParam(uint16(x), uint16(y)))
		}
	}
}

// toPoint 转换XY坐标
func (m *customWindowCaption) toPoint(message *types.TMessage) (x, y int32) {
	return winapi.GET_X_LPARAM(message.LParam), winapi.GET_Y_LPARAM(message.LParam)
}

// isCaption
// 鼠标是否在标题栏区域
//
// 如果启用了css拖拽则校验拖拽区域,否则只返回相对于浏览器窗口的x,y坐标
func (m *customWindowCaption) isCaption(hWND et.HWND, message *types.TMessage) (int32, int32, bool) {
	dx, dy := m.toPoint(message)
	p := &et.Point{
		X: dx,
		Y: dy,
	}
	winapi.WinScreenToClient(hWND, p)
	p.X -= m.bw.WindowParent().Left()
	p.Y -= m.bw.WindowParent().Top()
	if m.bw.WindowProperty().EnableWebkitAppRegion && m.rgn != nil {
		m.canCaption = winapi.WinPtInRegion(m.rgn, p.X, p.Y)
	} else {
		m.canCaption = false
	}
	return p.X, p.Y, m.canCaption
}

// doOnRenderCompMsg
func (m *LCLBrowserWindow) doOnRenderCompMsg(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	switch message.Msg {
	case messages.WM_NCLBUTTONDBLCLK: // 163 NC left dclick
		//标题栏拖拽区域 双击最大化和还原
		if m.cwcap.canCaption && m.WindowProperty().EnableWebkitAppRegionDClk {
			*lResult = messages.HTCAPTION
			*aHandled = true
			if win.ReleaseCapture() {
				if m.WindowState() == types.WsNormal {
					win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
				} else {
					win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
				}
				win.SendMessage(m.Handle(), messages.WM_NCLBUTTONUP, messages.HTCAPTION, 0)
			}
		}
	case messages.WM_NCLBUTTONDOWN: // 161 nc left down
		m.cwcap.onNCLButtonDown(m.Handle(), message, lResult, aHandled)
	case messages.WM_NCLBUTTONUP: // 162 nc l up
		if m.cwcap.canCaption {
			*lResult = messages.HTCAPTION
			*aHandled = true
		}
	case messages.WM_NCMOUSEMOVE: // 160 nc mouse move
		m.cwcap.onNCMouseMove(message, lResult, aHandled)
	case messages.WM_SETCURSOR: // 32 设置鼠标图标样式
		m.cwcap.onSetCursor(message, lResult, aHandled)
	case messages.WM_NCHITTEST: // 132 NCHITTEST
		if m.cwcap.borderMD { //TODO 测试windows7, 161消息之后再次处理132消息导致消息错误
			m.cwcap.borderMD = false
			return
		}
		//鼠标坐标是否在标题区域
		x, y, caption := m.cwcap.isCaption(et.HWND(m.Handle()), message)
		if caption { //窗口标题栏
			*lResult = messages.HTCAPTION
			*aHandled = true
		} else if m.WindowProperty().EnableHideCaption && m.WindowProperty().EnableResize && m.WindowState() == types.WsNormal { //1.窗口隐藏标题栏 2.启用了调整窗口大小 3.非最大化、最小化、全屏状态
			//全屏时不能调整窗口大小
			if m.WindowProperty().current.ws == types.WsFullScreen {
				return
			}
			rect := m.BoundsRect()
			if result, handled := m.cwcap.onCanBorder(x, y, &rect); handled {
				*lResult = types.LRESULT(result)
				*aHandled = true
			}
		}
	}
}

// setDraggableRegions
// 每一次拖拽区域改变都需要重新设置
func (m *LCLBrowserWindow) setDraggableRegions() {
	//在主线程中运行
	m.RunOnMainThread(func() {
		if m.cwcap.rgn == nil {
			//第一次时创建RGN
			m.cwcap.rgn = winapi.WinCreateRectRgn(0, 0, 0, 0)
		} else {
			//每次重置RGN
			winapi.WinSetRectRgn(m.cwcap.rgn, 0, 0, 0, 0)
		}
		for i := 0; i < m.cwcap.regions.RegionsCount(); i++ {
			region := m.cwcap.regions.Region(i)
			creRGN := winapi.WinCreateRectRgn(region.Bounds.X, region.Bounds.Y, region.Bounds.X+region.Bounds.Width, region.Bounds.Y+region.Bounds.Height)
			if region.Draggable {
				winapi.WinCombineRgn(m.cwcap.rgn, m.cwcap.rgn, creRGN, consts.RGN_OR)
			} else {
				winapi.WinCombineRgn(m.cwcap.rgn, m.cwcap.rgn, creRGN, consts.RGN_DIFF)
			}
			winapi.WinDeleteObject(creRGN)
		}
	})
}

// registerWindowsCompMsgEvent
// 注册windows下CompMsg事件
func (m *LCLBrowserWindow) registerWindowsCompMsgEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.Chromium().SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
		if bwEvent.onRenderCompMsg != nil {
			bwEvent.onRenderCompMsg(sender, message, lResult, aHandled)
		}
		if !*aHandled {
			m.doOnRenderCompMsg(message, lResult, aHandled)
		}
	})
	if m.WindowProperty().EnableWebkitAppRegion && m.WindowProperty().EnableWebkitAppRegionDClk {
		m.windowResize = func(sender lcl.IObject) bool {
			if m.WindowState() == types.WsMaximized && (m.WindowProperty().EnableHideCaption || m.BorderStyle() == types.BsNone || m.BorderStyle() == types.BsSingle) {
				var monitor = m.Monitor().WorkareaRect()
				m.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left, monitor.Bottom-monitor.Top)
				m.SetWindowState(types.WsMaximized)
			}
			return false
		}
	}

	//if m.WindowProperty().EnableWebkitAppRegion {
	//
	//} else {
	//	if bwEvent.onRenderCompMsg != nil {
	//		m.chromium.SetOnRenderCompMsg(bwEvent.onRenderCompMsg)
	//	}
	//}
}

// Restore Windows平台，窗口还原
func (m *LCLBrowserWindow) Restore() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		if win.ReleaseCapture() {
			win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
		}
	})
}

// Minimize Windows平台，窗口最小化
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MINIMIZE, 0)
		}
	})
}

// Maximize Windows平台，窗口最大化/还原
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		if win.ReleaseCapture() {
			if m.WindowState() == types.WsNormal {
				win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
			} else {
				win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
			}
		}
	})
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
	// Windows Drag Window
	// m.drag != nil 时，这里处理的是 up 事件, 给标题栏标记为false
	if m.drag != nil {
		m.drag.drag()
	} else {
		// 全屏时不能拖拽窗口
		if m.WindowProperty().current.ws == types.WsFullScreen {
			return
		}
		// 此时是 down 事件, 拖拽窗口
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
			m.cwcap.canCaption = true
		}
	}
}
