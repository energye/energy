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
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//Cef托盘
type tViewsFrameTrayWindow struct {
	trayWindow *ViewsFrameworkBrowserWindow
	trayIcon   *lcl.TTrayIcon
	x, y, w, h int32
	mouseUp    TMouseEvent
	isClosing  bool
}

func newViewsFrameTray(owner lcl.IComponent, width, height int32, url string) *tViewsFrameTrayWindow {
	var tray = &tViewsFrameTrayWindow{}
	cc := NewChromiumConfig()
	cc.SetEnableMenu(false)
	wp := NewWindowProperty()
	wp.Title = ""
	wp.Url = url
	wp.Width = width
	wp.Height = height
	wp.X = -width
	wp.Y = -height
	wp.AlwaysOnTop = true
	wp.CanMaximize = false
	wp.CanMinimize = false
	wp.CanResize = false
	wp.CenterWindow = false
	tray.trayWindow = NewViewsFrameworkBrowserWindow(cc, wp)
	tray.trayWindow.resetWindowPropertyEvent()
	tray.trayWindow.windowId = BrowserWindow.GetNextWindowNum()
	tray.trayWindow.putChromiumWindowInfo()
	tray.trayIcon = lcl.NewTrayIcon(owner)
	tray.trayIcon.SetVisible(true)
	tray.x = wp.X
	tray.y = wp.Y
	tray.w = wp.Width
	tray.h = wp.Height
	tray.registerMouseEvent()
	tray.registerChromiumEvent()
	return tray
}

func (m *tViewsFrameTrayWindow) Tray() *Tray {
	return nil
}

func (m *tViewsFrameTrayWindow) Show() {
	m.trayWindow.Show()
}

func (m *tViewsFrameTrayWindow) Hide() {
	m.trayWindow.Hide()
}

func (m *tViewsFrameTrayWindow) close() {
	if m.isClosing {
		return
	}
	m.trayIcon.SetVisible(false)
	m.Hide()
}

func (m *tViewsFrameTrayWindow) SetOnDblClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnDblClick(fn)
}

func (m *tViewsFrameTrayWindow) SetOnClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnClick(fn)
}

func (m *tViewsFrameTrayWindow) SetOnMouseUp(fn TMouseEvent) {
	m.mouseUp = fn
}

func (m *tViewsFrameTrayWindow) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}

func (m *tViewsFrameTrayWindow) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

func (m *tViewsFrameTrayWindow) Visible() bool {
	return false
}

func (m *tViewsFrameTrayWindow) SetVisible(v bool) {
	m.trayIcon.SetVisible(v)
}

func (m *tViewsFrameTrayWindow) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

func (m *tViewsFrameTrayWindow) SetTitle(title string) {
}

func (m *tViewsFrameTrayWindow) registerMouseEvent() {
	m.trayWindow.WindowComponent().SetOnWindowActivationChanged(func(sender lcl.IObject, window *ICefWindow, active bool) {
		if active {
			m.trayWindow.Show()
		} else {
			m.trayWindow.Hide()
		}
	})
	var IsCreateTopLevelWindow = true
	m.trayIcon.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if IsCreateTopLevelWindow {
			IsCreateTopLevelWindow = false
			m.trayWindow.CreateTopLevelWindow()
			m.trayWindow.HideTitle()
			m.trayWindow.SetNotInTaskBar()
			m.trayWindow.WindowComponent().SetAlwaysOnTop(true)
		}
		display := m.trayWindow.WindowComponent().Display()
		bounds := display.Bounds()
		var monitorWidth = bounds.Width
		width, height := m.w, m.h
		var mx = x + width
		var my = y + height
		if mx < monitorWidth {
			mx = x
		} else {
			mx = x - width
		}
		if my > m.h {
			my = y
		}
		if my > height {
			my = y - height
		}
		var ret bool
		if m.mouseUp != nil {
			ret = m.mouseUp(sender, button, shift, x, y)
		}
		if !ret {
			if button == types.MbRight {
				m.trayWindow.WindowComponent().SetBounds(NewCefRect(mx, my, width, height))
				m.trayWindow.Show()
				m.trayWindow.BrowserViewComponent().RequestFocus()
			}
		}
	})
}

func (m *tViewsFrameTrayWindow) registerChromiumEvent() {
	m.trayWindow.Chromium().SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		model.Clear()
	})
	m.trayWindow.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool {
		BrowserWindow.setOrIncNextWindowNum(browser.Identifier() + 1)
		return false
	})
	m.trayWindow.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
		}
	})
	m.trayWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("tray.chromium.onBeforeClose")
		m.close()
	})
	m.trayWindow.Chromium().SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool {
		return false
	})
}

//设置托盘气泡
//title 气泡标题
//content 气泡内容
//timeout 显示时间(毫秒)
func (m *tViewsFrameTrayWindow) SetBalloon(title, content string, timeout int32) ITray {
	m.trayIcon.SetBalloonTitle(title)
	m.trayIcon.SetBalloonHint(content)
	m.trayIcon.SetBalloonTimeout(timeout)
	return m
}

//显示托盘气泡
func (m *tViewsFrameTrayWindow) ShowBalloon() {
	m.trayIcon.ShowBalloonHint()
}

//设置托盘图标
func (m *tViewsFrameTrayWindow) SetIconFS(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
}

//设置托盘图标
func (m *tViewsFrameTrayWindow) SetIcon(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFile(iconResourcePath)
}
