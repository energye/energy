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
)

// OnCreate TForm create callback
type OnCreate func(window IBrowserWindow)

var (
	// global browser list
	globalBrowserWindows map[uint32]IBrowserWindow
)

func init() {
	globalBrowserWindows = make(map[uint32]IBrowserWindow)
}

func addBrowserWindow(window IBrowserWindow) {
	globalBrowserWindows[window.BrowserId()] = window
}

func deleteBrowserWindow(window IBrowserWindow) {
	delete(globalBrowserWindows, window.BrowserId())
}

func GetBrowserWindow(windowId uint32) IBrowserWindow {
	if window, ok := globalBrowserWindows[windowId]; ok {
		return window
	}
	return nil
}

// MainWindow app main window
type MainWindow struct {
	BrowserWindow
	onWindowCreate      OnCreate
	onWindowAfterCreate OnCreate
}

func (m *MainWindow) FormCreate(sender lcl.IObject) {
	m.BrowserWindow.FormCreate(sender)
	ipc.SetMainDefaultBrowserId(m.BrowserId())
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
