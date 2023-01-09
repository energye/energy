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
	"github.com/energye/energy/logger"
)

const (
	view_source_name = "ViewSource"
)

func updateBrowserViewSource(browser *ICefBrowser, title string) {
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil && browserWinInfo.WindowType() == WT_VIEW_SOURCE {
		QueueAsyncCall(func(id int) {
			if mainFrame := browser.MainFrame(); mainFrame != nil {
				browserWinInfo.SetTitle(fmt.Sprintf("%s - %s", view_source_name, mainFrame.Url))
			} else {
				logger.Error("failed to get main frame")
			}
		})
	}
}

func viewSourceAfterCreate(browser *ICefBrowser) bool {
	if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
		if winInfo.WindowType() == WT_VIEW_SOURCE && winInfo.getAuxTools().viewSourceWindow != nil {
			winInfo.getAuxTools().viewSourceWindow.Chromium().LoadUrl(winInfo.getAuxTools().viewSourceUrl)
			return true
		}
	}
	return false
}

func (m *ICefBrowser) createBrowserViewSource(frame *ICefFrame) {
	var viewSourceUrl = fmt.Sprintf("view-source:%s", frame.Url)
	if currentWindowInfo := BrowserWindow.GetWindowInfo(m.Identifier()); currentWindowInfo != nil {
		if currentWindowInfo.IsLCL() {
			QueueAsyncCall(func(id int) {
				var bw = BrowserWindow.popupWindow.AsLCLBrowserWindow().BrowserWindow()
				if bw != nil {
					bw.SetShowInTaskBar()
					bw.SetWindowType(WT_VIEW_SOURCE)
					bw.ChromiumCreate(nil, viewSourceUrl)
					bw.putChromiumWindowInfo()
					bw.defaultChromiumEvent()
					bw.SetSize(1024, 768)
					if winInfo := BrowserWindow.GetWindowInfo(bw.windowId); winInfo != nil {
						winInfo.createAuxTools()
						winInfo.getAuxTools().viewSourceUrl = viewSourceUrl
						winInfo.getAuxTools().viewSourceWindow = bw
					}
					bw.Show()
				} else {
					logger.Fatal("Window not initialized successfully")
				}
			})
		} else if currentWindowInfo.IsViewsFramework() {
			frame.ViewSource()
		}
	}
}
