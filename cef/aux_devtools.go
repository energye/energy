//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 辅助工具-开发者工具

package cef

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/energye/golcl/lcl"
)

const (
	devToolsName = "DevTools"
)

type devToolsWindow struct {
	*lcl.TForm
	parent ICEFWindowParent
}

func updateBrowserDevTools(window IBrowserWindow, browser *ICefBrowser, title string) {
	if window.IsLCL() {
		window := window.AsLCLBrowserWindow().BrowserWindow()
		if window.GetAuxTools() != nil && window.GetAuxTools().DevTools() != nil {
			window.GetAuxTools().DevTools().SetCaption(fmt.Sprintf("%s - %s", devToolsName, browser.MainFrame().Url()))
		}
	}
}

func (m *ICefBrowser) createBrowserDevTools(currentWindow IBrowserWindow, currentChromium ICEFChromiumBrowser) {
	if currentWindow.IsLCL() {
		if common.IsWindows() {
			// 如果开启开发者工具, 需要在IU线程中创建window
			currentWindow.AsLCLBrowserWindow().BrowserWindow().createAuxTools()
			devTools := currentWindow.AsLCLBrowserWindow().BrowserWindow().GetAuxTools()
			if devTools.DevTools() == nil {
				devTools.SetDevTools(createDevtoolsWindow(currentWindow.AsLCLBrowserWindow().BrowserWindow(), currentChromium))
				devTools.DevTools().SetCaption(fmt.Sprintf("%s - %s", devToolsName, m.MainFrame().Url()))
			}
			RunOnMainThread(func() { // show window, run is main ui thread
				devTools.DevTools().Show()
			})
		} else {
			currentChromium.Chromium().ShowDevTools(nil)
		}
	} else if currentWindow.IsViewsFramework() {
		currentChromium.Chromium().ShowDevTools(nil)
	}
}

func (m *devToolsWindow) WindowParent() ICEFWindowParent {
	return m.parent
}

func (m *devToolsWindow) SetWindowParent(parent ICEFWindowParent) {
	m.parent = parent
}
