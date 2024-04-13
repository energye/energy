//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

// VF窗口组件定义和实现-非windows平台

package cef

// ShowTitle 显示标题栏-无法动态控制, 在 CreateTopLevelWindow 之前调用
func (m *ViewsFrameworkBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	m.WindowComponent().SetOnIsFrameless(func(window *ICefWindow, aResult *bool) {
		*aResult = false
	})
}

// HideTitle 隐藏标题栏-无法动态控制, 在 CreateTopLevelWindow 之前调用
func (m *ViewsFrameworkBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	m.WindowComponent().SetOnIsFrameless(func(window *ICefWindow, aResult *bool) {
		*aResult = true
	})
}

// SetDefaultInTaskBar 空函数
func (m *ViewsFrameworkBrowserWindow) SetDefaultInTaskBar() {

}

// SetShowInTaskBar 空函数
func (m *ViewsFrameworkBrowserWindow) SetShowInTaskBar() {

}

// SetNotInTaskBar 空函数
func (m *ViewsFrameworkBrowserWindow) SetNotInTaskBar() {

}
