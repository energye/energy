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
	"github.com/energye/lcl/types"
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
	m.window.SetOnWindowActivationChanged(m.windowOnWindowActivationChanged)
	m.window.SetOnWindowBoundsChanged(m.windowOnWindowBoundsChanged)
	m.window.SetOnCanMaximize(m.windowOnCanMaximize)
	m.window.SetOnCanMinimize(m.windowOnCanMinimize)
	m.window.SetOnCanResize(m.windowOnCanResize)
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

func (m *TViewsBrowser) windowOnAdapterThemeChanged(sender lcl.IObject, isDark bool) {
	logger.Debug("Window.OnAdapterThemeChanged isDark:", isDark)
	if m.onThemeChange != nil {
		m.onThemeChange(isDark)
	}
}

func (m *TViewsBrowser) windowOnWindowClosing(sender lcl.IObject, window cef.ICefWindow) {
	logger.Debug("Window.OnWindowClosing")
	delete(GApplication.windowList, m.browserId)
}

func (m *TViewsBrowser) windowOnCanClose(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	logger.Debug("Window.OnCanClose", m.canClose, *result)
	if !m.canClose {
		canClose := true
		closeQueryHandle := lcl.CallFormCloseQuery(m.self, sender, &canClose)
		if closeQueryHandle || !canClose {
			return
		}
		closeAction := types.CaFree
		closeHandle := lcl.CallFormClose(m.self, sender, &closeAction)
		if closeHandle {
			return
		}
		switch closeAction {
		case types.CaNone:
			*result = false
			return
		case types.CaHide:
			*result = false
			m.Hide()
			return
		case types.CaMinimize:
			*result = false
			m.Minimize()
			return
		}
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
		m.CenterWindow()
	}
	m.browserView.RequestFocus()
}

func (m *TViewsBrowser) windowOnGetInitialBounds(sender lcl.IObject, window cef.ICefWindow, result *cef.TCefRect) {
	if m.options == nil || result == nil {
		return
	}
	m.bounds.X = m.options.X
	m.bounds.Y = m.options.Y
	m.bounds.Width = m.options.Width
	m.bounds.Height = m.options.Height

	if m.bounds.X > 0 {
		result.X = m.bounds.X
	}
	if m.bounds.Y > 0 {
		result.Y = m.bounds.Y
	}
	if m.bounds.Width > 0 {
		result.Width = m.bounds.Width
	}
	if m.bounds.Height > 0 {
		result.Height = m.bounds.Height
	}
}

func (m *TViewsBrowser) windowOnGetInitialShowState(sender lcl.IObject, window cef.ICefWindow, result *cefTypes.TCefShowState) {
	logger.Debug("Window.OnGetInitialShowState")
	if m.options != nil {
		switch m.options.DefaultWindowStatus {
		case types.WsMinimized:
			*result = cefTypes.CEF_SHOW_STATE_MINIMIZED
		case types.WsMaximized:
			*result = cefTypes.CEF_SHOW_STATE_MAXIMIZED
		case types.WsFullScreen:
			*result = cefTypes.CEF_SHOW_STATE_FULLSCREEN
		default:
			*result = cefTypes.CEF_SHOW_STATE_NORMAL
		}
	}
}

func (m *TViewsBrowser) windowOnIsFrameless(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	logger.Debug("Window.OnIsFrameless")
	if m.options != nil {
		*result = m.options.Frameless
	}
}

func (m *TViewsBrowser) windowOnWindowActivationChanged(sender lcl.IObject, window cef.ICefWindow, active bool) {
	logger.Debug("Window.OnWindowActivationChanged")
	if !m.isFirstShow {
		m.isFirstShow = true
		lcl.CallFormShow(m.self, m)
	}
	if m.onActivate != nil {
		m.onActivate(sender)
	}
}

func (m *TViewsBrowser) windowOnWindowBoundsChanged(sender lcl.IObject, window cef.ICefWindow, newBounds cef.TCefRect) {
	m.bounds.X = newBounds.X
	m.bounds.Y = newBounds.Y
	m.bounds.Width = newBounds.Width
	m.bounds.Height = newBounds.Height
	if m.onResize != nil {
		m.onResize(sender)
	}
}

func (m *TViewsBrowser) windowOnCanMaximize(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	*result = !m.options.DisableMaximize
}

func (m *TViewsBrowser) windowOnCanMinimize(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	*result = !m.options.DisableMinimize
}

func (m *TViewsBrowser) windowOnCanResize(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	*result = !m.options.DisableResize
}
