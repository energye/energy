//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

func (m *LCLBrowserWindow) _HookWndProcMessage() {
	// no impl
}

func (m *LCLBrowserWindow) _RestoreWndProc() {
	// no impl
}
