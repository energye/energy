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
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
)

const (
	dev_tools_name = "dev-tools"
)

func updateBrowserDevTools(browser *ICefBrowser, title string) {
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil {
		if browserWinInfo.IsLCL() {
			window := browserWinInfo.AsLCLBrowserWindow().BrowserWindow()
			if window.getAuxTools() != nil && window.getAuxTools().devToolsWindow != nil {
				window.getAuxTools().devToolsWindow.SetTitle(fmt.Sprintf("%s - %s", dev_tools_name, browser.MainFrame().Url()))
			}
		}
	}
}

func (m *ICefBrowser) createBrowserDevTools(browserWindow IBrowserWindow) {
	if browserWindow.IsLCL() {
		window := browserWindow.AsLCLBrowserWindow().BrowserWindow()
		window.createAuxTools()
		if window.getAuxTools().devToolsWindow != nil {
			dtw := window.getAuxTools().devToolsWindow.AsLCLBrowserWindow().BrowserWindow()
			QueueAsyncCall(func(id int) {
				if dtw.WindowState() == types.WsMinimized {
					dtw.SetWindowState(types.WsNormal)
				}
				dtw.Show()
				dtw.Active()
				dtw.Focused()
			})
			return
		}
		wp := NewWindowProperty()
		wp.Url = m.MainFrame().Url()
		wp.Title = fmt.Sprintf("%s - %s", dev_tools_name, m.MainFrame().Url())
		wp.WindowType = WT_DEV_TOOLS
		devToolsWindow := NewLCLBrowserWindow(nil, wp)
		devToolsWindow.SetWidth(1024)
		devToolsWindow.SetHeight(768)
		window.auxTools.devToolsWindow = devToolsWindow
		devToolsWindow.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
			if devToolsWindow.isClosing {
				return false
			}
			*action = types.CaFree
			return true
		})
		devToolsWindow.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
			if devToolsWindow.isClosing {
				return true
			}
			devToolsWindow.isClosing = true
			window.auxTools.devToolsWindow = nil
			return true
		})
		devToolsWindow.chromiumBrowser.CreateBrowser()
		devToolsWindow.Show()
		imports.Proc(internale_CEFBrowser_ShowDevTools).Call(m.Instance(), devToolsWindow.Chromium().Instance(), devToolsWindow.WindowParent().Instance(), api.PascalStr(dev_tools_name))
	} else if browserWindow.IsViewsFramework() {
		if application.RemoteDebuggingPort() > 1024 && application.RemoteDebuggingPort() < 65535 {
			wp := NewWindowProperty()
			wp.Url = fmt.Sprintf("http://127.0.0.1:%d", application.RemoteDebuggingPort())
			wp.Title = dev_tools_name
			wp.IconFS = BrowserWindow.Config.IconFS
			wp.Icon = BrowserWindow.Config.Icon
			wp.WindowType = WT_DEV_TOOLS
			devToolsWindow := NewViewsFrameworkBrowserWindow(nil, wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
			devToolsWindow.ResetWindowPropertyForEvent()
			devToolsWindow.CreateTopLevelWindow()
		}
	}
}
