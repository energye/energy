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
	return newTray(m)
}
