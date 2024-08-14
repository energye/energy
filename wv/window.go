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

type OnCreate func(window IBrowserWindow)

var (
	globalBrowserWindows map[types.HWND]IBrowserWindow
)

func addBrowserWindow(window IBrowserWindow) {
	globalBrowserWindows[window.Handle()] = window
}

func deleteBrowserWindow(window IBrowserWindow) {
	delete(globalBrowserWindows, window.Handle())
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

// MainWindow app main window
type MainWindow struct {
	BrowserWindow
	onWindowCreate      OnCreate
	onWindowAfterCreate OnCreate
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

func (m *MainWindow) FormAfterCreate(sender lcl.IObject) {
	m.BrowserWindow.FormAfterCreate(sender)
	if m.onWindowAfterCreate != nil {
		m.onWindowAfterCreate(m)
	}
}

// NewBrowserWindow create browser window
func NewBrowserWindow(options Options) IBrowserWindow {
	var window = &BrowserWindow{options: options}
	lcl.Application.CreateForm(window)
	window.afterCreate()
	return window
}
