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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
)

const (
	dev_tools_name = "dev-tools"
)

type devToolsWindow struct {
	*lcl.TForm
	parent ICEFWindowParent
}

func updateBrowserDevTools(browser *ICefBrowser, title string) {
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil {
		if browserWinInfo.IsLCL() {
			window := browserWinInfo.AsLCLBrowserWindow().BrowserWindow()
			if window.getAuxTools() != nil && window.getAuxTools().devToolsWindow != nil {
				window.getAuxTools().devToolsWindow.SetCaption(fmt.Sprintf("%s - %s", dev_tools_name, browser.MainFrame().Url()))
			}
		}
	}
}

func (m *ICefBrowser) createBrowserDevTools(browserWindow IBrowserWindow) {
	if browserWindow.IsLCL() {
		window := browserWindow.AsLCLBrowserWindow().BrowserWindow()
		QueueAsyncCall(func(id int) { // show window, run is main ui thread
			window.getAuxTools().devToolsWindow.Show()
			window.getAuxTools().devToolsWindow.parent.UpdateSize() // 更新CEFWindow大小到当前窗口大小
		})
		imports.Proc(def.CEFBrowser_ShowDevTools).Call(m.Instance(), browserWindow.Chromium().Instance(), window.getAuxTools().devToolsWindow.parent.Instance(), api.PascalStr(dev_tools_name))
	} else if browserWindow.IsViewsFramework() {
		if application.RemoteDebuggingPort() > 1024 && application.RemoteDebuggingPort() < 65535 {
			wp := NewWindowProperty()
			wp.Url = fmt.Sprintf("http://127.0.0.1:%d", application.RemoteDebuggingPort())
			wp.Title = dev_tools_name
			wp.IconFS = BrowserWindow.Config.IconFS
			wp.Icon = BrowserWindow.Config.Icon
			wp.WindowType = WT_DEV_TOOLS
			window := NewViewsFrameworkBrowserWindow(nil, wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
			window.ResetWindowPropertyForEvent()
			window.CreateTopLevelWindow()
		}
	}
}

// 创建开发者工具窗口，在 UI 线程中创建
// 在显示开发者工具时，需要在显示的窗口中初始化 CEFWindow
// 步骤
//  1. 创建窗口, 设置窗口宽高为0, 展示位置在屏幕之外
//	2. 创建 CEFWindow
//  3. show & hide, 先显示窗口让CEF初始化CEFWindow, 紧跟着隐藏掉
//  4. 设置默认的窗口宽高、居中显示在桌面并显示在任务栏
func createDevtoolsWindow(owner *LCLBrowserWindow) *devToolsWindow {
	window := &devToolsWindow{}
	window.TForm = lcl.NewForm(owner)
	window.SetCaption(dev_tools_name)
	window.SetIcon(owner.Icon())
	window.SetWidth(0)
	window.SetWidth(0)
	window.SetLeft(-500)
	window.SetTop(-500)
	parent := NewCEFWindow(window)
	parent.SetParent(window)
	parent.SetAlign(types.AlClient)
	window.parent = parent
	window.Show()
	window.Hide()
	window.SetWidth(1024)
	window.SetHeight(768)
	parent.SetAlign(types.AlClient)
	window.ScreenCenter()
	window.SetShowInTaskBar(types.StAlways)
	// 关闭流程
	//  1. 当前浏览器窗口关闭后触发 closeQuery 事件, 调用父窗口的chromium.CloseDevTools 关闭开发者工具
	//  2. 默认 aAanClose = true, 然后触发  OnClose 事件, 如果关闭的是开发者工具窗口，我们什么都不做,默认隐藏
	window.TForm.SetOnCloseQuery(func(sender lcl.IObject, aAanClose *bool) {
		owner.Chromium().CloseDevTools(window.parent) // close devtools
	})
	//	3. 如果关闭的是浏览器窗口 CaFree
	window.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		if owner.isClosing {
			*action = types.CaFree
		}
	})
	return window
}
