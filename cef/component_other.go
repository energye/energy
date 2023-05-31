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

// CEF 扩展组件定义-非windows

package cef

// NewCefTray LCL窗口组件 非windows系统不支持CefTray
func (m *LCLBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	return nil
}

// NewCefTray VF窗口组件 非windows系统不支持CefTray
func (m *ViewsFrameworkBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	return nil
}
