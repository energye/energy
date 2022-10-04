//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

// 非windows系统不支持CefTray
func (m *TCefWindowInfo) NewCefTray(width, height int32, url string) ITray {
	return nil
}
