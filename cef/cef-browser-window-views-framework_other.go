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

// VF窗口组件定义和实现-非windows平台
package cef

import "github.com/energye/golcl/lcl"

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
