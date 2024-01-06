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

// CEF 扩展组件定义-windows

package cef

// NewCefTray
// LCL窗口组件 只适用于windows的无菜单托盘, 需使用web端技术实现
//
// 不支持Linux和MacOSX
func (m *LCLBrowserWindow) NewCefTray(width, height int32, url string) (tray ITray) {
	if m == nil {
		return nil
	}
	tray = newLCLTrayWindow(m, width, height, url)
	m.tray = append(m.tray, tray)
	return tray
}

// NewCefTray
// VF窗口组件 只适用于windows的无菜单托盘, 需使用web端技术实现
//
// 不支持Linux和MacOSX
func (m *ViewsFrameworkBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	if m == nil {
		return nil
	}
	if m.tray == nil {
		m.tray = newViewsFrameTray(m.component, width, height, url)
	}
	return m.tray
}
