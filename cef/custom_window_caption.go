//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts/messages"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

// customWindowCaption 自定义窗口标题栏
//
// 当隐藏窗口标题栏时, 自定义窗口标题栏，通过html+css自定义窗口标题栏，实现窗口拖拽等
type customWindowCaption struct {
	bw                   IBrowserWindow //
	canCaption           bool           //当前鼠标是否在标题栏区域
	canBorder            bool           //当前鼠标是否在边框
	borderHT, borderWMSZ int            //borderHT: 鼠标所在边框位置, borderWMSZ: 窗口改变大小边框方向 borderMD:
	borderMD             bool           //borderMD: 鼠标调整窗口大小，已按下后，再次接收到132消息应该忽略该消息
}

// free
func (m *customWindowCaption) free() {
	if m != nil {
		cb := m.bw.(*LCLBrowserWindow).chromiumBrowser
		if cb != nil {
			cb.FreeRgn()
			cb.FreeRegions()
		}
	}
}

// onNCMouseMove NC 非客户区鼠标移动
func (m *customWindowCaption) onNCMouseMove(hWND types.HWND, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canCaption { // 当前在标题栏
	} else if m.canBorder { // 当前在边框
		*lResult = types.LRESULT(m.borderHT)
		*aHandled = true
		//全屏时不允许移动窗口
		// TODO 暂时不使用，配合 WndProc
		//if m.bw.WindowProperty().current.ws == types.WsFullScreen {
		//	return
		//}
		//s := winapi.GetKeyState(winapi.VK_LBUTTON) & 0x800
		//if winapi.GetKeyState(winapi.VK_LBUTTON) < 0 {
		//	x, y := m.toPoint(message)
		//	m.borderMD = true
		//	if win.ReleaseCapture() {
		//		win.PostMessage(hWND, messages.WM_SYSCOMMAND, uintptr(messages.SC_SIZE|m.borderWMSZ), rtl.MakeLParam(uint16(x), uint16(y)))
		//	}
		//}
	}
}

// onSetCursor 设置鼠标图标
func (m *customWindowCaption) onSetCursor(message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	if m.canBorder { //当前在边框
		switch winapi.LOWORD(uint32(message.LParam)) {
		case messages.HTBOTTOMRIGHT, messages.HTTOPLEFT: //右下 左上
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENWSE))
		case messages.HTRIGHT, messages.HTLEFT: //右 左
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZEWE))
		case messages.HTTOPRIGHT, messages.HTBOTTOMLEFT: //右上 左下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENESW))
		case messages.HTTOP, messages.HTBOTTOM: //上 下
			*lResult = types.LRESULT(m.borderHT)
			*aHandled = true
			winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENS))
		}
	}
}

// 鼠标是否在边框，并返回当前鼠标样式
func (m *customWindowCaption) onCanBorder(chromiumBrowser ICEFChromiumBrowser, x, y int32, windowRect *types.TRect) bool {
	width := windowRect.Width()
	height := windowRect.Height()
	bda := chromiumBrowser.BroderDirectionAdjustments()
	if m.canBorder = x <= width && x >= width-angleRange && y <= angleRange; m.canBorder && bda.In(et.BdaTopRight) { // 右上
		m.borderWMSZ = messages.WMSZ_TOPRIGHT
		m.borderHT = messages.HTTOPRIGHT
		return true
	} else if m.canBorder = x <= width && x >= width-angleRange && y <= height && y >= height-angleRange; m.canBorder && bda.In(et.BdaBottomRight) { // 右下
		m.borderWMSZ = messages.WMSZ_BOTTOMRIGHT
		m.borderHT = messages.HTBOTTOMRIGHT
		return true
	} else if m.canBorder = x <= angleRange && y <= angleRange; m.canBorder && bda.In(et.BdaTopLeft) { //左上
		m.borderWMSZ = messages.WMSZ_TOPLEFT
		m.borderHT = messages.HTTOPLEFT
		return true
	} else if m.canBorder = x <= angleRange && y >= height-angleRange; m.canBorder && bda.In(et.BdaBottomLeft) { //左下
		m.borderWMSZ = messages.WMSZ_BOTTOMLEFT
		m.borderHT = messages.HTBOTTOMLEFT
		return true
	} else if m.canBorder = x > angleRange && x < width-angleRange && y <= borderRange; m.canBorder && bda.In(et.BdaTop) { //上
		m.borderWMSZ = messages.WMSZ_TOP
		m.borderHT = messages.HTTOP
		return true
	} else if m.canBorder = x > angleRange && x < width-angleRange && y >= height-borderRange; m.canBorder && bda.In(et.BdaBottom) { //下
		m.borderWMSZ = messages.WMSZ_BOTTOM
		m.borderHT = messages.HTBOTTOM
		return true
	} else if m.canBorder = x <= borderRange && y > angleRange && y < height-angleRange; m.canBorder && bda.In(et.BdaLeft) { //左
		m.borderWMSZ = messages.WMSZ_LEFT
		m.borderHT = messages.HTLEFT
		return true
	} else if m.canBorder = x <= width && x >= width-borderRange && y > angleRange && y < height-angleRange; m.canBorder && bda.In(et.BdaRight) { // 右
		m.borderWMSZ = messages.WMSZ_RIGHT
		m.borderHT = messages.HTRIGHT
		return true
	}
	return false
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
		*lResult = types.LRESULT(m.borderHT)
		*aHandled = true
		x, y := m.toPoint(message)
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

// 鼠标是否在标题栏区域
//
// 如果启用了css拖拽则校验拖拽区域,否则只返回相对于浏览器窗口的x,y坐标
func (m *customWindowCaption) isCaption(chromiumBrowser ICEFChromiumBrowser, hWND et.HWND, message *types.TMessage) (int32, int32, bool) {
	dx, dy := m.toPoint(message)
	p := &et.Point{
		X: dx,
		Y: dy,
	}
	winapi.ScreenToClient(hWND, p)
	p.X -= chromiumBrowser.WindowParent().Left()
	p.Y -= chromiumBrowser.WindowParent().Top()
	if m.bw.WindowProperty().EnableWebkitAppRegion && chromiumBrowser.Rgn() != nil {
		m.canCaption = winapi.PtInRegion(chromiumBrowser.Rgn(), p.X, p.Y)
	} else {
		m.canCaption = false
	}
	return p.X, p.Y, m.canCaption
}
