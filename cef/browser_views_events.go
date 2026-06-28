//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
)

func (m *TViewsBrowser) initViewsBrowserDefaultEvent() {
	logger.Debug("Browser.initViewsBrowserDefaultEvent")

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)

	// close browser
	m.chromium.SetOnBeforeClose(m.chromiumOnBeforeClose)
	m.chromium.SetOnClose(m.chromiumOnClose)
}

func (m *TViewsBrowser) initViewsWindowDefaultEvent() {
	logger.Debug("Browser.initViewsWindowDefaultEvent")
	m.window.SetOnWindowCreated(m.windowOnWindowCreated)
	m.window.SetOnGetInitialShowState(m.windowOnGetInitialShowState)
	m.window.SetOnIsFrameless(m.windowOnIsFrameless)
	m.window.SetOnGetInitialBounds(m.windowOnGetInitialBounds)
	m.window.SetOnCanClose(m.windowOnCanClose)
	m.window.SetOnWindowClosing(m.windowOnWindowClosing)
}

func (m *TViewsBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.browserId == 0 {
		m.browserId = uint32(browser.GetIdentifier())
		if mainWindow, ok := GApplication.windowList[0]; ok {
			delete(GApplication.windowList, 0)
			GApplication.windowList[m.browserId] = mainWindow
		} else {
			GApplication.windowList[m.browserId] = m
		}
		// ipc
		ipc.RegisterProcessMessage(m)
		// pre-creates a window
		if m.options != nil && m.options.AutoPopupWindow {

		}
	}
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(sender)
	}
}

func (m *TViewsBrowser) chromiumOnBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnBeforeClose", "Current-BrowserID:", m.browserId, "Target-BrowserID:", browser.GetIdentifier())
	if m.browserId != uint32(browser.GetIdentifier()) {
		logger.Debug("Chromium.OnBeforeClose Non-current user browser")
		return
	}
	if len(GApplication.windowList) == 0 {
		GApplication.QuitMessageLoop()
	}
}

func (m *TViewsBrowser) chromiumOnClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {
	logger.Debug("Chromium.OnClose", browser.GetIdentifier())
	*action = cefTypes.CbaClose
	m.canClose = true
}

func (m *TViewsBrowser) windowOnWindowClosing(sender lcl.IObject, window cef.ICefWindow) {
	logger.Debug("Window.OnWindowClosing")
	delete(GApplication.windowList, m.browserId)
}

func (m *TViewsBrowser) windowOnCanClose(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	logger.Debug("Window.OnCanClose", m.canClose)
	if result == nil {
		return
	}
	*result = m.canClose
	if !m.canClose && m.chromium != nil {
		m.chromium.CloseBrowser(true)
	}
}

func (m *TViewsBrowser) windowOnWindowCreated(sender lcl.IObject, window cef.ICefWindow) {
	logger.Debug("Window.OnWindowCreated")
	if m.created {
		return
	}
	m.created = true
	lcl.CallFormCreate(m.self, m)

	m.UpdateBrowserOptions()
	url := m.defaultURL
	if url == "" {
		url = "about:blank"
	}

	created := m.chromium.CreateBrowserWithStrBVComponentRContextDValue(url, m.browserView, m.context, m.extraInfo)
	logger.Debug("Window.OnWindowCreated chromium.CreateBrowser:", created)
	if !created {
		return
	}

	m.window.SetToFillLayout()
	browserView := m.browserView.BrowserView()
	if browserView != nil && browserView.IsValid() {
		m.window.AddChildView(browserView)
	}

	if m.options != nil {
		if m.options.Caption != "" {
			m.window.SetTitle(m.options.Caption)
		}
		if m.options.Width > 0 && m.options.Height > 0 {
			m.window.CenterWindow(cef.TCefSize{
				Width:  m.options.Width,
				Height: m.options.Height,
			})
		}
	}

	m.browserView.RequestFocus()
}

func (m *TViewsBrowser) windowOnGetInitialBounds(sender lcl.IObject, window cef.ICefWindow, result *cef.TCefRect) {
	if m.options == nil || result == nil {
		return
	}
	if m.options.X > 0 {
		result.X = m.options.X
	}
	if m.options.Y > 0 {
		result.Y = m.options.Y
	}
	if m.options.Width > 0 {
		result.Width = m.options.Width
	}
	if m.options.Height > 0 {
		result.Height = m.options.Height
	}
}

func (m *TViewsBrowser) windowOnGetInitialShowState(sender lcl.IObject, window cef.ICefWindow, result *cefTypes.TCefShowState) {
	logger.Debug("Window.OnGetInitialShowState")
}

func (m *TViewsBrowser) windowOnIsFrameless(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	logger.Debug("Window.OnIsFrameless")
	if m.options != nil {
		*result = m.options.Frameless
	}
}
