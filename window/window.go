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
	SetOptions()
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
	SetOnWindowCreate(fn lcl.TNotifyEvent)
	SetOnWindowShow(fn lcl.TNotifyEvent)
	SetOnWindowClose(fn lcl.TCloseEvent)
	SetOnWindowCloseQuery(fn lcl.TCloseQueryEvent)
}

type TEnergyWindow struct {
	lcl.TEngForm
	windowId                uint32 // 窗口 ID 对应第一个浏览器 ID
	isClose                 bool
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
	onWindowCreate          []lcl.TNotifyEvent
	onWindowShow            []lcl.TNotifyEvent
	onWindowClose           []lcl.TCloseEvent
	onWindowCloseQuery      []lcl.TCloseQueryEvent
}

func (m *TEnergyWindow) SetClose(v bool) {
	m.isClose = v
}

func (m *TEnergyWindow) IsClose() bool {
	return m.isClose
}

func (m *TEnergyWindow) SetOnWindowCreate(fn lcl.TNotifyEvent) {
	m.onWindowCreate = append(m.onWindowCreate, fn)
}

func (m *TEnergyWindow) SetOnWindowShow(fn lcl.TNotifyEvent) {
	m.onWindowShow = append(m.onWindowShow, fn)
}

func (m *TEnergyWindow) SetOnWindowClose(fn lcl.TCloseEvent) {
	m.onWindowClose = append(m.onWindowClose, fn)
}

func (m *TEnergyWindow) SetOnWindowCloseQuery(fn lcl.TCloseQueryEvent) {
	m.onWindowCloseQuery = append(m.onWindowCloseQuery, fn)
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
	if m.IsFullScreen() || application.GApplication.Options.DisableMaximize {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
			if tool.IsDarwin() { //要这样重复设置2次不然不启作用
				m.SetWindowState(types.WsMaximized)
				m.SetWindowState(types.WsNormal)
			}
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

func (m *TWindow) FormCreate(sender lcl.IObject) {
	m._BeforeFormCreate()
	for _, fn := range m.onWindowCreate {
		fn(sender)
	}
}

func (m *TWindow) OnShow(sender lcl.IObject) {
	for _, fn := range m.onWindowShow {
		fn(sender)
	}
}

func (m *TWindow) OnCloseQuery(sender lcl.IObject, canClose *bool) {
	for _, fn := range m.onWindowCloseQuery {
		fn(sender, canClose)
	}
}

func (m *TWindow) OnClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	for _, fn := range m.onWindowClose {
		fn(sender, closeAction)
	}
}

func PtInRegion(x, y int32, rectX, rectY, rectWidth, rectHeight int32) bool {
	// 检查点(x, y)是否在矩形(rectX, rectY, rectWidth, rectHeight)内
	return x >= rectX && x <= rectX+rectWidth && y >= rectY && y <= rectY+rectHeight
}
