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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/win"
)

// 显示标题栏
func (m *ViewsFrameworkBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	//handle := m.WindowComponent().WindowHandle()
	//win.SetWindowLong(handle.ToPtr(), win.GWL_STYLE, uintptr(win.GetWindowLong(handle.ToPtr(), win.GWL_STYLE)|win.WS_CAPTION))
	//win.SetWindowPos(handle.ToPtr(), 0, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.WindowComponent().SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = false
	})
}

// 隐藏标题栏
func (m *ViewsFrameworkBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	//handle := m.WindowComponent().WindowHandle()
	//win.SetWindowLong(handle.ToPtr(), win.GWL_STYLE, uintptr(win.GetWindowLong(handle.ToPtr(), win.GWL_STYLE)&^win.WS_CAPTION))
	//win.SetWindowPos(handle.ToPtr(), 0, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.WindowComponent().SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = true
	})
}

func (m *ViewsFrameworkBrowserWindow) SetDefaultInTaskBar() {
	m.SetShowInTaskBar()
}

func (m *ViewsFrameworkBrowserWindow) SetShowInTaskBar() {
	handle := m.WindowComponent().WindowHandle()
	win.ShowWindow(handle.ToPtr(), win.SW_SHOW)
	win.SetWindowLong(handle.ToPtr(), win.GWL_EXSTYLE, win.WS_EX_APPWINDOW)
}

func (m *ViewsFrameworkBrowserWindow) SetNotInTaskBar() {
	handle := m.WindowComponent().WindowHandle()
	win.ShowWindow(handle.ToPtr(), win.SW_HIDE)
	win.SetWindowLong(handle.ToPtr(), win.GWL_EXSTYLE, win.WS_EX_TOOLWINDOW)
}
