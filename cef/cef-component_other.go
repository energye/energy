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

package cef

// 非windows系统不支持CefTray
func (m *LCLBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	return nil
}

// 非windows系统不支持CefTray
func (m *ViewsFrameworkBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	return nil
}
