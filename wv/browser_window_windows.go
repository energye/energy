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
	"encoding/json"
	"github.com/energye/energy/v3/internal/assets"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/win32"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
	"runtime"
)

func (m *BrowserWindow) platformCreate() {
	if m.options.Windows.ICON == nil {
		lcl.Application.Icon().LoadFromBytes(assets.ICON.ICO())
	} else {
		lcl.Application.Icon().LoadFromBytes(m.options.Windows.ICON)
	}
	switch m.options.Windows.Theme {
	case SystemDefault:
		win32.ChangeTheme(m.Handle(), win32.IsCurrentlyDarkMode())
	case Light:
		win32.ChangeTheme(m.Handle(), false)
	case Dark:
		win32.ChangeTheme(m.Handle(), true)
	}
}

var (
	frameWidth  = win.GetSystemMetrics(32)
	frameHeight = win.GetSystemMetrics(33)
	frameCorner = frameWidth + frameHeight
)

func (m *BrowserWindow) navigationStarting() {
	jsCode := &bytes.Buffer{}
	var envJS = func(json string) {
		jsCode.WriteString(`window.energy.setOptionsEnv(`)
		jsCode.WriteString(json)
		jsCode.WriteString(`);`)
	}
	optionsJSON, err := json.Marshal(m.options)
	if err == nil {
		envJS(string(optionsJSON))
	}
	env := make(map[string]interface{})
	env["frameWidth"] = frameWidth
	env["frameHeight"] = frameHeight
	env["frameCorner"] = frameCorner
	env["os"] = runtime.GOOS
	envJSON, err := json.Marshal(env)
	if err == nil {
		envJS(string(envJSON))
	}
	m.browser.ExecuteScript(jsCode.String(), 0)
}

func (m *BrowserWindow) Resize(ht string) {
	if m.IsFullScreen() || m.options.DisableResize {
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
		if m.IsFullScreen() {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		m.Maximize()
	}
}

func (m *BrowserWindow) borderFramelessForLine() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME|win.WS_BORDER))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

func (m *BrowserWindow) borderFrameless() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

func (m *BrowserWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.IsMinimize() || m.IsMaximize() {
			m.Restore()
		}
		m.windowsState = types.WsFullScreen
		// save current window rect, use ExitFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		monitorRect := m.Monitor().BoundsRect()
		if !m.options.Frameless {
			// save current window style, use ExitFullScreen
			m.oldWindowStyle = uintptr(win.GetWindowLongPtr(m.Handle(), win.GWL_STYLE))
			m.borderFrameless()
			m.SetWindowState(types.WsFullScreen)
		}
		win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *BrowserWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !m.options.Frameless {
				win.SetWindowLong(m.Handle(), win.GWL_STYLE, m.oldWindowStyle)
			}
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(&m.previousWindowPlacement)
		})
	}
}
