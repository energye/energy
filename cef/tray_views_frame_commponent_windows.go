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

// 基于VF组件系统托盘 - windows 平台
// Html + CSS + JavaScript实现

package cef

import (
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// ViewsFrameTray VF(views framework)组件+html 托盘
type ViewsFrameTray struct {
	trayWindow *ViewsFrameworkBrowserWindow
	trayIcon   *lcl.TTrayIcon
	visible    bool
	x, y, w, h int32
	mouseUp    TMouseEvent
	isClosing  bool
}

// 创建系统托盘
func newViewsFrameTray(owner lcl.IComponent, width, height int32, url string) *ViewsFrameTray {
	var tray = &ViewsFrameTray{}
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
	wp.EnableMaximize = false
	wp.EnableMinimize = false
	wp.EnableResize = false
	wp.EnableCenterWindow = false
	tray.trayWindow = NewViewsFrameworkBrowserWindow(cc, wp, nil)
	tray.trayWindow.ResetWindowPropertyForEvent()
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

func (m *ViewsFrameTray) registerMouseEvent() {
	m.trayWindow.WindowComponent().SetOnWindowActivationChanged(func(window *ICefWindow, active bool) {
		if active {
		} else {
			m.trayWindow.Hide()
		}
	})
	var isCreateTopLevelWindow = true
	m.trayIcon.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if isCreateTopLevelWindow {
			isCreateTopLevelWindow = false
			m.trayWindow.HideTitle()
			m.trayWindow.CreateTopLevelWindow()
			m.trayWindow.WindowComponent().SetAlwaysOnTop(true)
			m.trayWindow.SetNotInTaskBar()
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

func (m *ViewsFrameTray) registerChromiumEvent() {
	m.trayWindow.Chromium().SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		model.Clear()
	})
	m.trayWindow.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			if application.Is49() {
				headerMap := request.GetHeaderMap()
				headerMap.Append(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue)
				request.SetHeaderMap(headerMap)
				headerMap.Free()
			} else {
				request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
			}
		}
	})
	m.trayWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("tray.chromium.onBeforeClose")
		m.close()
	})
	m.trayWindow.Chromium().SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
		return false
	})
}

// AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
func (m *ViewsFrameTray) AsSysTray() *SysTray {
	return nil
}

// AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *ViewsFrameTray) AsViewsFrameTray() *ViewsFrameTray {
	return m
}

// AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *ViewsFrameTray) AsCEFTray() *CEFTray {
	return nil
}

// AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *ViewsFrameTray) AsLCLTray() *LCLTray {
	return nil
}

// Show 显示/启动 托盘
func (m *ViewsFrameTray) Show() {
	m.trayWindow.Show()
}

// Hide 隐藏 托盘
func (m *ViewsFrameTray) Hide() {
	m.trayWindow.Hide()
}

func (m *ViewsFrameTray) close() {
	if m.isClosing {
		return
	}
	m.isClosing = true
	m.trayIcon.SetVisible(false)
	m.Hide()
	m.trayIcon.Free()
}

// SetOnDblClick 设置双击事件
func (m *ViewsFrameTray) SetOnDblClick(fn TrayICONClick) {
	m.trayIcon.SetOnDblClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnClick 设置单击事件
func (m *ViewsFrameTray) SetOnClick(fn TrayICONClick) {
	m.trayIcon.SetOnClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnMouseUp 鼠标抬起事件
func (m *ViewsFrameTray) SetOnMouseUp(fn TMouseEvent) {
	m.mouseUp = fn
}

// SetOnMouseDown 鼠标按下事件
func (m *ViewsFrameTray) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}

// SetOnMouseMove 鼠标移动事件
func (m *ViewsFrameTray) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

// Visible 显示状态
func (m *ViewsFrameTray) Visible() bool {
	return m.visible
}

// SetVisible 设置显示状态
func (m *ViewsFrameTray) SetVisible(v bool) {
	m.visible = v
	m.trayIcon.SetVisible(v)
}

// SetHint 设置提示
func (m *ViewsFrameTray) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

// SetTitle 设置标题 - 空函数
func (m *ViewsFrameTray) SetTitle(title string) {
}

// 显示系统通知
//
// title 标题
//
// content 内容
//
// timeout 显示时间(毫秒)
func (m *ViewsFrameTray) Notice(title, content string, timeout int32) {
	if common.IsWindows() {
		notification(m.trayIcon, title, content, timeout)
	} else {
		notification(nil, title, content, timeout)
	}
}

// 设置托盘图标
func (m *ViewsFrameTray) SetIconFS(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
}

// 设置托盘图标
func (m *ViewsFrameTray) SetIcon(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFile(iconResourcePath)
}
