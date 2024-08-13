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
		if m.IsFullScreen() {
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
