//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

// 非windows系统不支持CefTray
func (m *TCefWindowInfo) NewCefTray(width, height int32, url string) ITray {
	return nil
}
