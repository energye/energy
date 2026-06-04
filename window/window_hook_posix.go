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

package window

<<<<<<<< HEAD:cef/browser_window_hook_other.go
func (m *LCLBrowserWindow) _HookWndProcMessage() {
	// no impl
}

func (m *LCLBrowserWindow) _RestoreWndProc() {
	// no impl
========
func (m *TWindow) _HookWndProcMessage() {

}

func (m *TWindow) _RestoreWndProc() {

>>>>>>>> v3-alpha:window/window_hook_posix.go
}
