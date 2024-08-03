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

package wv

import (
	"bytes"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"github.com/energye/wv/wv"
	"strconv"
	"syscall"
)

func (m *BrowserWindow) Resize(ht string) {
	if m.IsFullScreen() {
		return
	}
	if win.ReleaseCapture() {
		var borderHT uintptr
		switch ht {
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
		win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, borderHT, 0)
	}
}

func (m *BrowserWindow) Drag(message ipc.ProcessMessage) {
	if m.IsFullScreen() {
		return
	}
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
		//fmt.Println("MT_DRAG_MOVE", m.WindowState())
		if m.WindowState() == types.WsFullScreen {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	}
}

var (
	wndProcCallback = syscall.NewCallback(_WndProcCallback)
)

func _WndProcCallback(hwnd types.HWND, message uint32, wParam, lParam uintptr) uintptr {
	//fmt.Println("HookMsg msg:", message)
	if window := getBrowserWindow(hwnd); window != nil {
		return window._WndProc(message, wParam, lParam)
	}
	//return win.CallWindowProc(oldWndPrc, hwnd, message, wParam, lParam)
	return win.DefWindowProc(hwnd, message, wParam, lParam)
}

func (m *BrowserWindow) _WndProc(message uint32, wParam, lParam uintptr) uintptr {
	//fmt.Println("_WndProc message:", message)
	if m.options.Frameless {
		switch message {
		case messages.WM_ACTIVATE:
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			win.ExtendFrameIntoClientArea(m.Handle(), win.Margins{CxLeftWidth: 4, CxRightWidth: 4, CyTopHeight: 4, CyBottomHeight: 4})
		case messages.WM_NCCALCSIZE:
			// Disable the standard frame by allowing the client area to take the full
			// window size.
			// See: https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize#remarks
			// This hides the titlebar and also disables the resizing from user interaction because the standard frame is not
			// shown. We still need the WS_THICKFRAME style to enable resizing from the frontend.
			if wParam != 0 {
				//cycaption := win.GetSystemMetrics(4)
				//rect := (*types.TRect)(unsafe.Pointer(lParam))
				//rect.Bottom += -1
				//rect.Right += -1
				return 0
			}
		}
		//fmt.Println("message:", message)
	}
	return win.CallWindowProc(m.oldWndPrc, m.Handle(), message, wParam, lParam)
}

func (m *BrowserWindow) _HookWndProcMessage() {
	m.oldWndPrc = win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, wndProcCallback)
}

//func (m *BrowserWindow) _SetCursor(idc int) {
//	switch idc {
//	case messages.HTBOTTOMRIGHT, messages.HTTOPLEFT: //右下 左上
//		winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENWSE))
//	case messages.HTRIGHT, messages.HTLEFT: //右 左
//		winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZEWE))
//	case messages.HTTOPRIGHT, messages.HTBOTTOMLEFT: //右上 左下
//		winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENESW))
//	case messages.HTTOP, messages.HTBOTTOM: //上 下
//		winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_SIZENS))
//	default:
//		winapi.SetCursor(winapi.LoadCursor(0, messages.IDC_HAND))
//	}
//}

func (m *BrowserWindow) _RestoreWndProc() {
	if m.oldWndPrc != 0 {
		win.SetWindowLongPtr(m.Handle(), win.GWL_WNDPROC, m.oldWndPrc)
	}
}

var (
	frameWidth  = win.GetSystemMetrics(32)
	frameHeight = win.GetSystemMetrics(33)
	frameCorner = frameWidth + frameHeight
)

func (m *BrowserWindow) navigationStarting(webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {
	jsCode := bytes.Buffer{}
	var envJS = func(key, value string) {
		jsCode.WriteString("window.energy.setEnv('")
		jsCode.WriteString(key)
		jsCode.WriteString("',")
		jsCode.WriteString(value)
		jsCode.WriteString(");")
	}
	envJS("frameWidth", strconv.Itoa(int(frameWidth)))
	envJS("frameHeight", strconv.Itoa(int(frameHeight)))
	envJS("frameCorner", strconv.Itoa(int(frameCorner)))
	m.browser.ExecuteScript(jsCode.String(), 0)
}
