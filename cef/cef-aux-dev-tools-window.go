//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
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
				QueueAsyncCall(func(id int) {
					window.getAuxTools().devToolsWindow.SetTitle(fmt.Sprintf("%s - %s", dev_tools_name, browser.MainFrame().Url))
				})
			}
		}
	}
}

func (m *ICefBrowser) createBrowserDevTools(browserWinInfo IBrowserWindow) {
	if browserWinInfo.IsLCL() {
		QueueAsyncCall(func(id int) {
			window := browserWinInfo.AsLCLBrowserWindow().BrowserWindow()
			window.createAuxTools()
			if window.getAuxTools().devToolsWindow != nil {
				dtw := window.getAuxTools().devToolsWindow.AsLCLBrowserWindow().BrowserWindow()
				if dtw.WindowState() == types.WsMinimized {
					dtw.SetWindowState(types.WsNormal)
				}
				dtw.Show()
				dtw.Active()
				dtw.Focused()
				return
			}
			wp := NewWindowProperty()
			wp.Title = fmt.Sprintf("%s - %s", dev_tools_name, m.MainFrame().Url)
			wp.WindowType = WT_DEV_TOOLS
			wp.Url = m.MainFrame().Url
			devToolsWindow := NewLCLBrowserWindow(nil, wp)
			window.auxTools.devToolsWindow = devToolsWindow
			devToolsWindow.SetWidth(800)
			devToolsWindow.SetHeight(600)
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
				BrowserWindow.removeWindowInfo(devToolsWindow.Id())
				window.auxTools.devToolsWindow = nil
				return true
			})
			devToolsWindow.EnableDefaultCloseEvent()
			devToolsWindow.Show()
			BrowserWindow.setOrIncNextWindowNum() //明确的生成下一个窗体序号
			_CEFBrowser_ShowDevTools(devToolsWindow.chromium.Instance(), uintptr(m.Identifier()), devToolsWindow.windowParent.Instance(), api.PascalStr(dev_tools_name))
		})
	} else if browserWinInfo.IsViewsFramework() {
		if application.cfg.remoteDebuggingPort > 1024 && application.cfg.remoteDebuggingPort < 65535 {
			wp := NewWindowProperty()
			wp.Url = fmt.Sprintf("http://127.0.0.1:%d", application.cfg.remoteDebuggingPort)
			wp.Title = dev_tools_name
			wp.IconFS = BrowserWindow.Config.IconFS
			wp.Icon = BrowserWindow.Config.Icon
			devToolsWindow := NewViewsFrameworkBrowserWindow(nil, wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
			devToolsWindow.ResetWindowPropertyForEvent()
			devToolsWindow.SetWindowType(WT_DEV_TOOLS)
			BrowserWindow.setOrIncNextWindowNum() //明确的生成下一个窗口序号
			devToolsWindow.windowComponent.CreateTopLevelWindow()
		}
	}
}
