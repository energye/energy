//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

//适用于 windows linux macos 系统托盘
func (m *LCLBrowserWindow) NewTray() ITray {
	if m.tray == nil {
		m.tray = newTray(m.TForm)
	}
	return m.tray
}

func (m *LCLBrowserWindow) NewSysTray() ITray {
	if m.tray == nil {
		m.tray = newSysTray()
	}
	return m.tray
}

//只适用于windows的无菜单托盘
func (m *ViewsFrameworkBrowserWindow) NewSysTray() ITray {
	if m == nil {
		return nil
	}
	if m.tray == nil {
		m.tray = newSysTray()
	}
	return m.tray
}
