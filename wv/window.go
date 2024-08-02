//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import "github.com/energye/lcl/types"

var (
	globalBrowserWindows map[types.HWND]IBrowserWindow
)

func addBrowserWindow(window IBrowserWindow) {
	globalBrowserWindows[window.Handle()] = window
}

func deleteBrowserWindow(hwnd types.HWND) {
	delete(globalBrowserWindows, hwnd)
}

func getBrowserWindow(hwnd types.HWND) IBrowserWindow {
	if window, ok := globalBrowserWindows[hwnd]; ok {
		return window
	}
	return nil
}

func init() {
	globalBrowserWindows = make(map[types.HWND]IBrowserWindow)
}
