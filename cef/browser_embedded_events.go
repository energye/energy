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
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types/messages"
)

func (m *TEmbeddedBrowser) initEmbeddedBrowserDefaultEvent() {
	logger.Debug("Browser.initEmbeddedBrowserDefaultEvent")

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)
	m.chromium.SetOnBeforeBrowse(m.chromiumOnBeforeBrowse)

	//m.chromium.SetOnDraggableRegionsChanged(m.chromiumOnDraggableRegionsChanged)

	// close browser
	m.chromium.SetOnBeforeClose(m.chromiumOnBeforeClose)
	m.chromium.SetOnClose(m.chromiumOnClose)

	m.ICEFWinControl.SetOnEnter(m.winControlOnEnter)
	m.ICEFWinControl.SetOnExit(m.winControlOnExit)
}

func (m *TEmbeddedBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.window != nil && m.window.BrowserId() == 0 {
		options := m.window.Options()
		m.browserId = uint32(browser.GetIdentifier())
		m.window.SetBrowserId(m.browserId)
		// ipc
		ipc.RegisterProcessMessage(m)
		// local load
		//m.schemeHandlerFactory = createSchemeHandlerFactory(browser)
		// pre-creates a window
		if options.AutoPopupWindow {
			if gPrePopupWindow == nil {
				lcl.RunOnMainThreadAsync(func(id uint32) {
					gPrePopupWindow = NewPopupWindow()
				})
			}
		}
	}
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(sender)
	}
}

func (m *TEmbeddedBrowser) chromiumOnBeforeBrowse(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, userGesture bool,
	isRedirect bool, outResult *bool) {
	logger.Debug("Chromium.OnBeforeBrowse")
	m.UpdateSize()
}

//func (m *TEmbeddedBrowser) chromiumOnStartDragging(sender lcl.IObject, browser cef.ICefBrowser, dragData cef.ICefDragData, allowedOps cefTypes.TCefDragOperations,
//	X int32, Y int32, outResult *bool) {
//	logger.Debug("Chromium.OnStartDragging", browser.GetIdentifier())
//}

//func (m *TEmbeddedBrowser) chromiumOnDraggableRegionsChanged(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, regionsCount cefTypes.NativeUInt,
//	regions cef.ICefDraggableRegionArray) {
//	logger.Debug("Chromium.OnDraggableRegionsChanged", browser.GetIdentifier())
//}

func (m *TEmbeddedBrowser) chromiumOnClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {
	logger.Debug("Chromium.OnClose", browser.GetIdentifier())
	if tool.IsDarwin() {
		ok := m.DestroyChildWindow()
		logger.Debug("Chromium.OnClose => winControl.DestroyChildWindow() Success:", ok)
		*action = cefTypes.CbaClose
	} else if tool.IsLinux() {
		*action = cefTypes.CbaClose
	} else if tool.IsWindows() {
		*action = cefTypes.CbaDelay
	}
	if tool.IsWindows() || tool.IsLinux() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.ICEFWinControl.Free()
		})
	}
}

func (m *TEmbeddedBrowser) chromiumOnBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnBeforeClose", "Current-BrowserID:", m.browserId, "Target-BrowserID:", browser.GetIdentifier())
	if m.browserId != uint32(browser.GetIdentifier()) {
		logger.Debug("Chromium.OnBeforeClose Non-current user browser")
		return
	}
	closeWindow := func() {
		if m.window != nil {
			m.canClose = true
			if tool.IsWindows() {
				rtl.PostMessage(m.window.Handle(), messages.WM_CLOSE, 0, 0)
			} else if tool.IsDarwin() || tool.IsLinux() {
				m.window.Close()
			}
		} else {
			logger.Error("Browser associated window is nil, failed to close browser")
		}
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		closeWindow()
	})
}

func (m *TEmbeddedBrowser) winControlOnEnter(sender lcl.IObject) {
	m.chromium.Initialized()
	m.chromium.FrameIsFocused()
	m.chromium.SetFocus(true)
}

func (m *TEmbeddedBrowser) winControlOnExit(sender lcl.IObject) {
	m.chromium.SendCaptureLostEvent()
}
