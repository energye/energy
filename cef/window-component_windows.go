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

//只适用于windows的无菜单托盘
func (m *LCLBrowserWindow) NewCefTray(width, height int32, url string) ITray {
	if BrowserWindow.mainBrowserWindow.tray == nil {
		BrowserWindow.mainBrowserWindow.tray = newCefTray(m, width, height, url)
	}
	return BrowserWindow.mainBrowserWindow.tray
}
