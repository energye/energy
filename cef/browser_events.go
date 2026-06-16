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
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types/messages"
)

func (m *TBrowser) initDefaultEvent() {
	logger.Debug("Browser.initDefaultEvent")

	m.chromium.SetOnProcessMessageReceived(m.chromiumOnProcessMessageReceived)

	m.chromium.SetOnBeforeResourceLoad(m.chromiumOnBeforeResourceLoad)
	m.chromium.SetOnGetResourceHandler(m.chromiumOnGetResourceHandler)

	m.chromium.SetOnBeforeContextMenu(m.chromiumOnBeforeContextMenu)
	m.chromium.SetOnContextMenuCommand(m.chromiumOnContextMenuCommand)

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)
	m.chromium.SetOnBeforeBrowse(m.chromiumOnBeforeBrowse)

	m.chromium.SetOnBeforeDownload(m.chromiumOnBeforeDownload)
	m.chromium.SetOnLoadStart(m.chromiumOnLoadStart)
	m.chromium.SetOnKeyEvent(m.chromiumOnKeyEvent)

	m.chromium.SetOnTitleChange(m.chromiumOnTitleChange)
	m.chromium.SetOnDragEnter(m.chromiumOnDragEnter)
	m.chromium.SetOnDraggableRegionsChanged(m.chromiumOnDraggableRegionsChanged)

	// new tab or popup browser
	m.chromium.SetOnOpenUrlFromTab(m.chromiumOnOpenUrlFromTab)
	m.chromium.SetOnBeforePopup(m.chromiumOnBeforePopup)

	// close browser
	m.chromium.SetOnBeforeClose(m.chromiumOnBeforeClose)
	m.chromium.SetOnClose(m.chromiumOnClose)

	m.ICEFWinControl.SetOnEnter(m.winControlOnEnter)
	m.ICEFWinControl.SetOnExit(m.winControlOnExit)
}

func (m *TBrowser) winControlOnEnter(sender lcl.IObject) {
	m.chromium.Initialized()
	m.chromium.FrameIsFocused()
	m.chromium.SetFocus(true)
}

func (m *TBrowser) winControlOnExit(sender lcl.IObject) {
	m.chromium.SendCaptureLostEvent()
}

func (m *TBrowser) chromiumOnProcessMessageReceived(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, outResult *bool) {
}

func (m *TBrowser) chromiumOnBeforeResourceLoad(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	callback cef.ICefCallback, outResult *cefTypes.TCefReturnValue) {
}

func (m *TBrowser) chromiumOnGetResourceHandler(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	resourceHandler *cef.IEngResourceHandler) {

}

func (m *TBrowser) chromiumOnBeforeContextMenu(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, model cef.ICefMenuModel) {

}

func (m *TBrowser) chromiumOnContextMenuCommand(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams,
	commandId int32, eventFlags cefTypes.TCefEventFlags, outResult *bool) {

}

func (m *TBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.window != nil && m.window.BrowserId() == 0 {
		m.browserId = uint32(browser.GetIdentifier())
		m.window.SetBrowserId(m.browserId)
	}
}

func (m *TBrowser) chromiumOnBeforeBrowse(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, userGesture bool,
	isRedirect bool, outResult *bool) {
	logger.Debug("Chromium.OnBeforeBrowse")
	m.UpdateSize()
}

func (m *TBrowser) chromiumOnAdapterBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {

}
func (m *TBrowser) chromiumOnLoadStart(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, transitionType cefTypes.TCefTransitionType) {

}

func (m *TBrowser) chromiumOnKeyEvent(sender lcl.IObject, browser cef.ICefBrowser, event cef.TCefKeyEvent, osEvent cefTypes.TCefEventHandle, outResult *bool) {

}

func (m *TBrowser) chromiumOnTitleChange(sender lcl.IObject, browser cef.ICefBrowser, title string) {

}

func (m *TBrowser) chromiumOnDragEnter(sender lcl.IObject, browser cef.ICefBrowser, dragData cef.ICefDragData, mask cefTypes.TCefDragOperations, outResult *bool) {

}

func (m *TBrowser) chromiumOnDraggableRegionsChanged(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, regionsCount cefTypes.NativeUInt,
	regions cef.ICefDraggableRegionArray) {

}

func (m *TBrowser) chromiumOnClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {
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

func (m *TBrowser) chromiumOnBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnBeforeClose", browser.GetIdentifier())
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

func (m *TBrowser) chromiumOnOpenUrlFromTab(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, targetUrl string,
	targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, outResult *bool) {

}

func (m *TBrowser) chromiumOnAdapterBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, popupId int32, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {

}
