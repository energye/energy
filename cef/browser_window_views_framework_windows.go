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

// VF窗口组件定义和实现-windows平台

package cef

import (
	"github.com/energye/golcl/lcl/win"
)

// ShowTitle 显示标题栏-无法动态控制, 在 CreateTopLevelWindow 之前调用
func (m *ViewsFrameworkBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	//handle := m.WindowComponent().WindowHandle()
	//win.SetWindowLong(handle.ToPtr(), win.GWL_STYLE, uintptr(win.GetWindowLong(handle.ToPtr(), win.GWL_STYLE)|win.WS_CAPTION))
	//win.SetWindowPos(handle.ToPtr(), 0, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.WindowComponent().SetOnIsFrameless(func(window *ICefWindow, aResult *bool) {
		*aResult = false
	})
}

// HideTitle 隐藏标题栏-无法动态控制, 在 CreateTopLevelWindow 之前调用
func (m *ViewsFrameworkBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	//handle := m.WindowComponent().WindowHandle()
	//win.SetWindowLong(handle.ToPtr(), win.GWL_STYLE, uintptr(win.GetWindowLong(handle.ToPtr(), win.GWL_STYLE)&^win.WS_CAPTION))
	//win.SetWindowPos(handle.ToPtr(), 0, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.WindowComponent().SetOnIsFrameless(func(window *ICefWindow, aResult *bool) {
		*aResult = true
	})
}

// SetDefaultInTaskBar 默认窗口在任务栏上显示按钮, 在 CreateTopLevelWindow 之后调用
func (m *ViewsFrameworkBrowserWindow) SetDefaultInTaskBar() {
	m.SetShowInTaskBar()
}

// SetShowInTaskBar 强制窗口在任务栏上显示按钮, 在 CreateTopLevelWindow 之后调用
func (m *ViewsFrameworkBrowserWindow) SetShowInTaskBar() {
	handle := m.WindowComponent().WindowHandle()
	win.ShowWindow(handle.ToPtr(), win.SW_SHOW)
	win.SetWindowLong(handle.ToPtr(), win.GWL_EXSTYLE, win.WS_EX_APPWINDOW)
}

// SetNotInTaskBar 强制窗口不在任务栏上显示按钮, 在 CreateTopLevelWindow 之后调用
func (m *ViewsFrameworkBrowserWindow) SetNotInTaskBar() {
	handle := m.WindowComponent().WindowHandle()
	win.ShowWindow(handle.ToPtr(), win.SW_HIDE)
	win.SetWindowLong(handle.ToPtr(), win.GWL_EXSTYLE, win.WS_EX_TOOLWINDOW)
}
