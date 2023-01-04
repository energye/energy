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
	dev_tools_name = "DevTools"
)

func updateBrowserDevTools(browser *ICefBrowser, title string) {
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil {
		if browserWinInfo.auxTools != nil && browserWinInfo.auxTools.devToolsWindow != nil {
			QueueAsyncCall(func(id int) {
				browserWinInfo.auxTools.devToolsWindow.SetCaption(fmt.Sprintf("%s - %s", dev_tools_name, browser.MainFrame().Url))
			})
		}
	}
}
func createBrowserDevTools(browser *ICefBrowser, browserWinInfo *TCefWindowInfo) {
	QueueAsyncCall(func(id int) {
		winAuxTools := browserWinInfo.auxTools
		if winAuxTools.devToolsWindow != nil {
			winAuxTools.devToolsWindow.Show()
			return
		}
		winAuxTools.devToolsWindow = &LCLBrowserWindow{}
		winAuxTools.devToolsWindow.SetWindowType(WT_DEV_TOOLS)
		winAuxTools.devToolsWindow.TForm = lcl.NewForm(browserWinInfo.Window)
		winAuxTools.devToolsWindow.SetCaption(fmt.Sprintf("%s - %s", dev_tools_name, browser.MainFrame().Url))
		winAuxTools.devToolsWindow.FormCreate()
		winAuxTools.devToolsWindow.defaultWindowEvent()
		winAuxTools.devToolsWindow.defaultWindowCloseEvent()
		winAuxTools.devToolsWindow.SetWidth(1024)
		winAuxTools.devToolsWindow.SetHeight(768)
		winAuxTools.devToolsWindow.SetShowInTaskBar()

		winAuxTools.devToolsWindow.SetOnResize(func(sender lcl.IObject) {
			winAuxTools.devToolsX = winAuxTools.devToolsWindow.Left()
			winAuxTools.devToolsY = winAuxTools.devToolsWindow.Top()
			winAuxTools.devToolsWidth = winAuxTools.devToolsWindow.Width()
			winAuxTools.devToolsHeight = winAuxTools.devToolsWindow.Height()

			if winAuxTools.devToolsWindow.isClosing {
				return
			}
			if winAuxTools.devToolsWindow.chromium != nil {
				winAuxTools.devToolsWindow.chromium.NotifyMoveOrResizeStarted()
			}
			if winAuxTools.devToolsWindow.windowParent != nil {
				winAuxTools.devToolsWindow.windowParent.UpdateSize()
			}
		})
		winAuxTools.devToolsWindow.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
			if winAuxTools.devToolsWindow.isClosing {
				return
			}
			*action = types.CaFree
		})
		winAuxTools.devToolsWindow.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
			if winAuxTools.devToolsWindow.isClosing {
				return
			}
			winAuxTools.devToolsWindow.isClosing = true
			BrowserWindow.removeWindowInfo(winAuxTools.devToolsWindow.windowId)
		})

		winAuxTools.devToolsWindow.ChromiumCreate(nil, "")
		winAuxTools.devToolsWindow.putChromiumWindowInfo()
		winAuxTools.devToolsWindow.defaultChromiumEvent()
		winAuxTools.devToolsWindow.windowInfo = browserWinInfo
		winAuxTools.devToolsWindow.Show()
		//明确的生成下一个窗体序号
		BrowserWindow.setOrIncNextWindowNum()
		_CEFBrowser_ShowDevTools(winAuxTools.devToolsWindow.chromium.Instance(), uintptr(browser.Identifier()), winAuxTools.devToolsWindow.windowParent.Instance(), api.PascalStr(dev_tools_name))
	})
}
