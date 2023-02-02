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

//只适用于windows的无菜单托盘, 需使用web端技术实现
//
//不支持Linux和MacOSX
func (m *LCLBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	if m == nil {
		return nil
	}
	if m.tray == nil {
		m.tray = newLCLTrayWindow(m, width, height, url)
	}
	return m.tray
}

//只适用于windows的无菜单托盘, 需使用web端技术实现
//
//不支持Linux和MacOSX
func (m *ViewsFrameworkBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	if m == nil {
		return nil
	}
	if m.tray == nil {
		m.tray = newViewsFrameTray(m.component, width, height, url)
	}
	return m.tray
}
