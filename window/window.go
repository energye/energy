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
	"github.com/energye/energy/v3/core"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
)

type TWindowCloseState int32

const (
	WcsOpen TWindowCloseState = iota
	WcsClosing
	WcsClosed
)

type IWindow interface {
	lcl.IEngForm
	// SetOptions Configure options for current window
	SetOptions(options application.Options)
	Options() *application.Options
	// UpdateWindowOption Configure window options for a *TWindow instance
	UpdateWindowOption()
	SetBrowserId(windowId uint32)
	BrowserId() uint32
	IsMain() bool // Whether current window is main window
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
	Close()
	AddOnWindowStateChange(fn lcl.TNotifyEvent)
	AddOnWindowResize(fn lcl.TNotifyEvent)
	AddOnWindowCreate(fn lcl.TNotifyEvent)
	AddOnWindowShow(fn lcl.TNotifyEvent)
	AddOnWindowClose(fn lcl.TCloseEvent)
	AddOnWindowCloseQuery(fn lcl.TCloseQueryEvent)
	SetOnThemeChange(fn core.TOnThemeChange)
}

type TEnergyWindow struct {
	lcl.TEngForm
	windowId                uint32 // Window ID maps to first browser ID in current window
	isMain                  bool
	closeState              TWindowCloseState
	flagFirstShow           bool
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
	onThemeChange           core.TOnThemeChange
}

func (m *TEnergyWindow) SetClose(v bool) {
	if v {
		m.closeState = WcsClosing
	}
}

func (m *TEnergyWindow) IsClose() bool {
	return m.closeState == WcsClosing
}

func (m *TEnergyWindow) Close() {
	if m.closeState == WcsClosed {
		return
	}
	m.closeState = WcsClosed
	m.TEngForm.Close()
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

func (m *TEnergyWindow) SetOnThemeChange(fn core.TOnThemeChange) {
	m.onThemeChange = fn
}

func (m *TEnergyWindow) doOnThemeChange(isDark bool) {
	if m.onThemeChange != nil {
		m.onThemeChange(isDark)
	}
}

func (m *TEnergyWindow) SetOnResize(fn lcl.TNotifyEvent) {
	m.onResize = fn
}

func (m *TEnergyWindow) SetOnWindowStateChange(fn lcl.TNotifyEvent) {
	m.onWindowStateChange = fn
}

func (m *TEnergyWindow) FormAfterCreate(sender lcl.IObject) {
	mainForm := lcl.Application.MainForm()
	m.isMain = mainForm == nil || mainForm.Instance() == 0
}

func (m *TWindow) SetBrowserId(windowId uint32) {
	m.windowId = windowId
}

func (m *TWindow) BrowserId() uint32 {
	return m.windowId
}

func (m *TEnergyWindow) IsMain() bool {
	return m.isMain
}

func (m *TWindow) Minimize() {
	m.minimize()
}

func (m *TWindow) Maximize() {
	if m.IsFullScreen() || (m.options != nil && m.options.DisableMaximize) {
		return
	}
	m.maximize()
}

func (m *TWindow) Restore() {
	// In the case of a title bar
	// If the current state is full screen and the extracted state is Ws Maximized,
	// So let's first perform IsFullScreen() judgment here
	m.restore()
}

func (m *TWindow) IsFullScreen() bool {
	return m.isFullScreen()
}

func (m *TWindow) IsMinimize() bool {
	return m.isMinimize()
}

func (m *TWindow) IsMaximize() bool {
	return m.isMaximize()
}

func (m *TWindow) SetOptions(options application.Options) {
	m.options = &options
	if m.options.BackgroundColor == nil {
		m.options.BackgroundColor = &colors.TARGB{R: 255, G: 255, B: 255, A: 255}
	}
}

func (m *TWindow) Options() *application.Options {
	return m.options
}

func (m *TWindow) FormCreate(sender lcl.IObject) {
	if m.options == nil {
		if application.GApplication != nil {
			m.SetOptions(application.GApplication.Options)
		}
	}
	m.TEngForm.SetOnResize(m.doOnResize)
	m.TEngForm.SetOnWindowStateChange(m.doOnWindowStateChange)
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
	if *canClose {
		m.closeState = WcsClosed
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
	if tool.IsDarwin() {
		// 窗口状态改变时获取窗口状态
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

func (m *TWindow) UpdateConfigProperty() {
	if m.options != nil {
		if !m.options.Frameless {
			if m.options.DisableResize {
				m.SetBorderStyleToFormBorderStyle(types.BsSingle)
				m.EnabledMaximize(false)
			}
			if m.options.DisableMinimize {
				m.EnabledMinimize(false)
			}
			if m.options.DisableMaximize {
				m.EnabledMaximize(false)
			}
			if m.options.DisableSystemMenu {
				m.EnabledSystemMenu(false)
			}
		}
		if m.Caption() == "" {
			m.SetCaption(m.options.Caption)
		}
		constr := m.Constraints()
		if m.options.MaxWidth > 0 || m.options.MaxHeight > 0 {
			constr.SetMaxWidth(m.options.MaxWidth)
			constr.SetMaxHeight(m.options.MaxHeight)
		}
		if m.options.MinWidth > 0 || m.options.MinHeight > 0 {
			constr.SetMinWidth(m.options.MinWidth)
			constr.SetMinHeight(m.options.MinHeight)
		}
		windowBr := m.BoundsRect()
		if m.options.Width <= 0 {
			m.options.Width = windowBr.Width()
		}
		if m.options.Height <= 0 {
			m.options.Height = windowBr.Height()
		}
		if m.options.X != 0 {
			windowBr.Left = m.options.X
		}
		if m.options.Y != 0 {
			windowBr.Top = m.options.Y
		}
		if m.options.Width > 0 {
			windowBr.SetWidth(m.options.Width)
		}
		if m.options.Height > 0 {
			windowBr.SetHeight(m.options.Height)
		}
		m.SetBoundsRect(windowBr)
	}
}
