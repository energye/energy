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
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil && browserWinInfo.Window != nil && browserWinInfo.Window.WindowType() == WT_VIEW_SOURCE {
		if browserWinInfo.Window != nil {
			QueueAsyncCall(func(id int) {
				if mainFrame := browser.MainFrame(); mainFrame != nil {
					browserWinInfo.Window.SetCaption(fmt.Sprintf("%s - %s", view_source_name, mainFrame.Url))
				} else {
					logger.Error("failed to get main frame")
				}
			})
		}
	}
}

func viewSourceAfterCreate(browser *ICefBrowser) bool {
	if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
		if winInfo.Window.WindowType() == WT_VIEW_SOURCE && winInfo.auxTools.viewSourceWindow != nil {
			winInfo.auxTools.viewSourceWindow.chromium.LoadUrl(winInfo.auxTools.viewSourceUrl)
			return true
		}
	}
	return false
}

func createBrowserViewSource(browser *ICefBrowser, frame *ICefFrame) {
	BrowserWindow.uiLock.Lock()
	defer BrowserWindow.uiLock.Unlock()
	var viewSourceUrl = fmt.Sprintf("view-source:%s", frame.Url)
	QueueAsyncCall(func(id int) {
		var m = BrowserWindow.popupWindow
		if m != nil {
			m.SetWindowType(WT_VIEW_SOURCE)
			m.ChromiumCreate(nil, viewSourceUrl)
			m.chromium.EnableIndependentEvent()
			m.putChromiumWindowInfo()
			m.defaultChromiumEvent()
			m.SetWidth(1024)
			m.SetHeight(768)
			if winInfo := BrowserWindow.GetWindowInfo(m.windowId); winInfo != nil {
				winInfo.auxTools.viewSourceUrl = viewSourceUrl
				winInfo.auxTools.viewSourceWindow = m
			}
			m.Show()
		} else {
			logger.Fatal("Window not initialized successfully")
		}
	})
}
