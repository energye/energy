//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

const (
	view_source_name = "ViewSource"
)

func updateBrowserViewSource(browser *ICefBrowser, title string) {
	if browserWinInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); browserWinInfo != nil && browserWinInfo.Window != nil && browserWinInfo.Window.windowType == WT_VIEW_SOURCE {
		if browserWinInfo.Window != nil {
			QueueAsyncCall(func(id int) {
				browserWinInfo.Window.SetCaption(fmt.Sprintf("%s - %s", view_source_name, browser.MainFrame().Url))
			})
		}
	}
}

func viewSourceAfterCreate(browser *ICefBrowser) bool {
	if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
		if winInfo.Window.windowType == WT_VIEW_SOURCE && winInfo.auxTools.viewSourceWindow != nil {
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
		BrowserWindow.popupWindow.SetWindowType(WT_VIEW_SOURCE)
		BrowserWindow.popupWindow.ChromiumCreate(nil, viewSourceUrl)
		BrowserWindow.popupWindow.putChromiumWindowInfo()
		BrowserWindow.popupWindow.defaultChromiumEvent()
		if winInfo := BrowserWindow.GetWindowInfo(BrowserWindow.popupWindow.windowId); winInfo != nil {
			winInfo.auxTools.viewSourceUrl = viewSourceUrl
			winInfo.auxTools.viewSourceWindow = BrowserWindow.popupWindow
		}
		m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool { return false })
		m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction) {})
		m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {})
		m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {})
		m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {})
		BrowserWindow.popupWindow.Show()
	})
}
