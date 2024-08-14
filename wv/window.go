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

import (
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

var (
	globalBrowserWindows map[types.HWND]IBrowserWindow
)

func addBrowserWindow(window IBrowserWindow) {
	globalBrowserWindows[window.Handle()] = window
}

func deleteBrowserWindow(windowId types.HWND) {
	delete(globalBrowserWindows, windowId)
}

func getBrowserWindow(windowId types.HWND) IBrowserWindow {
	if window, ok := globalBrowserWindows[windowId]; ok {
		return window
	}
	return nil
}

func init() {
	globalBrowserWindows = make(map[types.HWND]IBrowserWindow)
}

// MainWindow app main window
type MainWindow struct {
	BrowserWindow
	onWindowCreate OnWindowCreate
}

func (m *MainWindow) FormCreate(sender lcl.IObject) {
	m.BrowserWindow.FormCreate(sender)
	ipc.SetMainWindowId(m.WindowId())
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
	m.afterCreate()
}

// NewBrowserWindow create browser window
func NewBrowserWindow(options Options) IBrowserWindow {
	var window = &BrowserWindow{options: options}
	lcl.Application.CreateForm(window)
	window.afterCreate()
	return window
}
