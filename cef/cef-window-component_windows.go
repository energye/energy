//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

//只适用于windows的无菜单托盘
func (m *TCefWindowInfo) NewCefTray(width, height int32, url string) ITray {
	if BrowserWindow.browserWindow.tray == nil {
		BrowserWindow.browserWindow.tray = newCefTray(m.Window, width, height, url)
	}
	return BrowserWindow.browserWindow.tray
}
