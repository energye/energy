// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package window

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
)

type IWindow interface {
	lcl.IEngForm
	// 设置当前窗口配置选项
	SetOptions(options application.Options)
	Options() *application.Options
	// 更新当前窗口配置选项
	UpdateWindowOption()
	SetBrowserId(windowId uint32)
	BrowserId() uint32
	Restore()
	Minimize()
	Maximize()
	IsMinimize() bool
	IsMaximize() bool
	FullScreen()
	ExitFullScreen()
	IsFullScreen() bool
	SetClose(v bool)
	IsClose() bool
	AddOnWindowStateChange(fn lcl.TNotifyEvent)
	AddOnWindowResize(fn lcl.TNotifyEvent)
	AddOnWindowCreate(fn lcl.TNotifyEvent)
	AddOnWindowShow(fn lcl.TNotifyEvent)
	AddOnWindowClose(fn lcl.TCloseEvent)
	AddOnWindowCloseQuery(fn lcl.TCloseQueryEvent)
}

type TEnergyWindow struct {
	lcl.TEngForm
	windowId                uint32 // 窗口 ID 对应第一个浏览器 ID
	isClose                 bool
	options                 *application.Options
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
	onResize                lcl.TNotifyEvent
	onWindowStateChange     lcl.TNotifyEvent
	onWindowStateChangeList []lcl.TNotifyEvent
	onWindowResizeList      []lcl.TNotifyEvent
	onWindowCreateList      []lcl.TNotifyEvent
	onWindowShowList        []lcl.TNotifyEvent
	onWindowCloseList       []lcl.TCloseEvent
	onWindowCloseQueryList  []lcl.TCloseQueryEvent
}

func (m *TEnergyWindow) SetClose(v bool) {
	m.isClose = v
}

func (m *TEnergyWindow) IsClose() bool {
	return m.isClose
}

func (m *TEnergyWindow) AddOnWindowStateChange(fn lcl.TNotifyEvent) {
	m.onWindowStateChangeList = append(m.onWindowStateChangeList, fn)
}

func (m *TEnergyWindow) AddOnWindowResize(fn lcl.TNotifyEvent) {
	m.onWindowResizeList = append(m.onWindowResizeList, fn)
}

func (m *TEnergyWindow) AddOnWindowCreate(fn lcl.TNotifyEvent) {
	m.onWindowCreateList = append(m.onWindowCreateList, fn)
}

func (m *TEnergyWindow) AddOnWindowShow(fn lcl.TNotifyEvent) {
	m.onWindowShowList = append(m.onWindowShowList, fn)
}

func (m *TEnergyWindow) AddOnWindowClose(fn lcl.TCloseEvent) {
	m.onWindowCloseList = append(m.onWindowCloseList, fn)
}

func (m *TEnergyWindow) AddOnWindowCloseQuery(fn lcl.TCloseQueryEvent) {
	m.onWindowCloseQueryList = append(m.onWindowCloseQueryList, fn)
}

func (m *TEnergyWindow) SetOnResize(fn lcl.TNotifyEvent) {
	m.onResize = fn
}

func (m *TEnergyWindow) SetOnWindowStateChange(fn lcl.TNotifyEvent) {
	m.onWindowStateChange = fn
}

func (m *TWindow) SetBrowserId(windowId uint32) {
	m.windowId = windowId
}

func (m *TWindow) BrowserId() uint32 {
	return m.windowId
}

func (m *TWindow) Minimize() {
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.SetWindowState(types.WsMinimized)
	})
}

func (m *TWindow) Maximize() {
	if m.IsFullScreen() || (m.options != nil && m.options.DisableMaximize) {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	})
}

func (m *TWindow) Restore() {
	// In the case of a title bar
	// If the current state is full screen and the extracted state is Ws Maximized,
	// So let's first perform IsFullScreen() judgment here
	if m.IsFullScreen() {
		m.ExitFullScreen()
	} else if m.IsMinimize() || m.IsMaximize() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
		})
	}
}

func (m *TWindow) IsFullScreen() bool {
	return m.windowsState == types.WsFullScreen
}

func (m *TWindow) IsMinimize() bool {
	return m.WindowState() == types.WsMinimized
}

func (m *TWindow) IsMaximize() bool {
	return m.WindowState() == types.WsMaximized
}

func (m *TWindow) SetOptions(options application.Options) {
	m.options = &options
}

func (m *TWindow) Options() *application.Options {
	return m.options
}

func (m *TWindow) FormCreate(sender lcl.IObject) {
	if m.options == nil {
		// 获取全局配置
		if application.GApplication != nil {
			m.SetOptions(application.GApplication.Options)
		}
	}
	m.TEngForm.SetOnResize(m.doOnResize)
	m.TEngForm.SetOnWindowStateChange(m.doOnWindowStateChange)
	m._BeforeFormCreate()
	for _, fn := range m.onWindowCreateList {
		fn(sender)
	}
}

func (m *TWindow) OnShow(sender lcl.IObject) {
	m._BeforeFormShow()
	for _, fn := range m.onWindowShowList {
		fn(sender)
	}
}

func (m *TWindow) OnCloseQuery(sender lcl.IObject, canClose *bool) {
	for _, fn := range m.onWindowCloseQueryList {
		fn(sender, canClose)
	}
}

func (m *TWindow) OnClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	for _, fn := range m.onWindowCloseList {
		fn(sender, closeAction)
	}
}

func (m *TWindow) doOnResize(sender lcl.IObject) {
	for _, fn := range m.onWindowResizeList {
		fn(sender)
	}
	if m.TEnergyWindow.onResize != nil {
		m.TEnergyWindow.onResize(sender)
	}
}

func (m *TWindow) doOnWindowStateChange(sender lcl.IObject) {
	if !tool.IsWindows() {
		// 窗口状态改变时获取窗口状态
		// windows 不在此设置该状态, m.windowsState
		m.windowsState = m.TEngForm.WindowState()
	}
	for _, fn := range m.onWindowStateChangeList {
		fn(sender)
	}
	if m.TEnergyWindow.onWindowStateChange != nil {
		m.TEnergyWindow.onWindowStateChange(sender)
	}
}

func PtInRegion(x, y int32, rectX, rectY, rectWidth, rectHeight int32) bool {
	// 检查点(x, y)是否在矩形(rectX, rectY, rectWidth, rectHeight)内
	return x >= rectX && x <= rectX+rectWidth && y >= rectY && y <= rectY+rectHeight
}
