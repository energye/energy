//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

//适用于 windows linux macos 系统托盘
func (m *TCefWindowInfo) NewTray() ITray {
	return newTray(m.Window)
}
