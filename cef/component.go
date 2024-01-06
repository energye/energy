//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 扩展组件

package cef

// NewTray 适用于 windows linux macos 系统托盘
func (m *LCLBrowserWindow) NewTray() (tray ITray) {
	if BrowserWindow.Config.EnableMainWindow {
		tray = newTray(m.TForm)
		m.tray = append(m.tray, tray)
	} else {
		// 禁用主窗口, 这时需要使用 disabledMainWindow, 因为它才是实际的主窗口
		tray = newTray(disabledMainWindow.TForm)
		m.tray = append(m.tray, tray)
	}
	return tray
}

// NewSysTray LCL窗口组件,系统托盘
func (m *LCLBrowserWindow) NewSysTray() (tray ITray) {
	tray = newSysTray()
	m.tray = append(m.tray, tray)
	return tray
}

// NewSysTray VF窗口组件,只适用于windows的无菜单托盘
func (m *ViewsFrameworkBrowserWindow) NewSysTray() ITray {
	if m == nil {
		return nil
	}
	if m.tray == nil {
		m.tray = newSysTray()
	}
	return m.tray
}
