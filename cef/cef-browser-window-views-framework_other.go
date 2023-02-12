//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build !windows
// +build !windows

// 非windows的未实现
package cef

// 显示标题栏
func (m *ViewsFrameworkBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	m.WindowComponent().SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = false
	})
}

// 隐藏标题栏
func (m *ViewsFrameworkBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	m.WindowComponent().SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = true
	})
}

func (m *ViewsFrameworkBrowserWindow) SetDefaultInTaskBar() {

}

func (m *ViewsFrameworkBrowserWindow) SetShowInTaskBar() {

}

func (m *ViewsFrameworkBrowserWindow) SetNotInTaskBar() {

}
