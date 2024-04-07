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

// 辅助工具-开发者工具 windows

package cef

// no windows
func createDevtoolsWindow(owner *LCLBrowserWindow, currentChromium ICEFChromiumBrowser) *devToolsWindow {
	return nil
}
