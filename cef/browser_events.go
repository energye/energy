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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/keys"
	"github.com/energye/lcl/types/messages"
	"net/url"
)

func (m *TBrowser) initDefaultEvent() {
	logger.Debug("Browser.initDefaultEvent")

	m.chromium.SetOnProcessMessageReceived(m.chromiumOnProcessMessageReceived)

	m.chromium.SetOnGetResourceHandler(m.chromiumOnGetResourceHandler)
	m.chromium.SetOnResourceLoadComplete(m.chromiumOnResourceLoadComplete)

	m.chromium.SetOnBeforeContextMenu(m.chromiumOnBeforeContextMenu)
	m.chromium.SetOnContextMenuCommand(m.chromiumOnContextMenuCommand)

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)
	m.chromium.SetOnBeforeBrowse(m.chromiumOnBeforeBrowse)

	m.chromium.SetOnBeforeDownload(m.chromiumOnBeforeDownload)

	m.chromium.SetOnLoadStart(m.chromiumOnLoadStart)
	m.chromium.SetOnLoadEnd(m.chromiumOnLoadEnd)

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
	if m.onProcessMessage != nil {
		m.onProcessMessage("")
	}
}

func (m *TBrowser) chromiumOnGetResourceHandler(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	resourceHandler *cef.IEngResourceHandler) {
	logger.Debug("Chromium.OnGetResourceHandler")

	var (
		uri          = request.GetUrl()
		reqUrl, err  = url.Parse(uri)
		path, method string
		resource     string
		handle       bool
		data         []byte
	)

	if err == nil {
		if application.GApplication == nil || application.GApplication.LocalLoad == nil || reqUrl.Scheme != application.GApplication.LocalLoad.Scheme {
			return
		}
		if m.resourceHandlerList == nil {
			m.resourceHandlerList = make(map[string]*source)
		}
		if m.onResourceRequest != nil {
			header := make(map[string]string)
			path = reqUrl.Path
			method = request.GetMethod()
			headerMap := cef.NewStringMultimapOwn()
			intfHeaderMap := cef.AsCefStringMultimapOwn(headerMap.AsIntfStringMultimap())
			request.GetHeaderMap(intfHeaderMap)
			for i := 0; i < int(intfHeaderMap.GetSize()); i++ {
				key := intfHeaderMap.GetKey(uint32(i))
				value := intfHeaderMap.GetValue(uint32(i))
				header[key] = value
			}
			intfHeaderMap.Release()
			headerMap.Free()
			resource, handle = m.onResourceRequest(uri, path, method, header)
		}
		if handle && resource != "" {
			data = []byte(resource)
		}
		src, err := makeSource(browser, frame, request)
		if err != nil {
			logger.Error("Chromium.OnGetResourceHandler makeSource:", err.Error())
			return
		}
		src.data = data
		*resourceHandler = cef.AsEngResourceHandler(src.resourceHandler.AsIntfResourceHandler())
		m.resourceHandlerList[uri] = src
	}
}

func (m *TBrowser) chromiumOnResourceLoadComplete(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	response cef.ICefResponse, status cefTypes.TCefUrlRequestStatus, receivedContentLength int64) {
	uri := request.GetUrl()
	logger.Debug("Chromium.OnResourceLoadComplete uri:", uri, "status:", status, "receivedContentLength:", receivedContentLength)
	if src, ok := m.resourceHandlerList[uri]; ok {
		src.resourceHandler.Free()
		delete(m.resourceHandlerList, uri)
	}
}

func (m *TBrowser) chromiumOnBeforeContextMenu(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, model cef.ICefMenuModel) {

}

func (m *TBrowser) chromiumOnContextMenuCommand(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams,
	commandId int32, eventFlags cefTypes.TCefEventFlags, outResult *bool) {

}

func (m *TBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
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

func (m *TBrowser) chromiumOnBeforeBrowse(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, userGesture bool,
	isRedirect bool, outResult *bool) {
	logger.Debug("Chromium.OnBeforeBrowse")
	m.UpdateSize()
}

func (m *TBrowser) chromiumOnAdapterBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {

}

func (m *TBrowser) chromiumOnLoadStart(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, transitionType cefTypes.TCefTransitionType) {
	m.loadingURL = frame.GetUrl()
	m.loadingState = core.LcStart
	logger.Debug("Chromium.OnLoadStart", m.loadingURL)
	if m.onLoadChange != nil {
		m.onLoadChange(m.loadingURL, m.loadingTitle, m.loadingState)
	}
}

func (m *TBrowser) chromiumOnLoadEnd(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, httpStatusCode int32) {
	logger.Debug("Chromium.OnLoadEnd")
	m.loadingState = core.LcFinish
	if m.onLoadChange != nil {
		m.onLoadChange(m.loadingURL, m.loadingTitle, m.loadingState)
	}
}

func (m *TBrowser) chromiumOnKeyEvent(sender lcl.IObject, browser cef.ICefBrowser, event cef.TCefKeyEvent, osEvent cefTypes.TCefEventHandle, outResult *bool) {
	if event.WindowsKeyCode == keys.VkF12 {
		m.chromium.ShowDevToolsWithPointWinControl(types.Point(0, 0), nil)
	}
}

func (m *TBrowser) chromiumOnTitleChange(sender lcl.IObject, browser cef.ICefBrowser, title string) {
	logger.Debug("Chromium.OnTitleChange title:", title)
	m.loadingTitle = title
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.window.Caption() == "" {
			m.window.SetCaption(title)
		}
	})
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

func (m *TBrowser) chromiumOnOpenUrlFromTab(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, targetUrl string,
	targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, outResult *bool) {
	logger.Debug("Chromium.OnOpenUrlFromTab", "targetUrl:", targetUrl)
	*outResult = true
}

func (m *TBrowser) chromiumOnAdapterBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, popupId int32, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {
	logger.Debug("Chromium.OnAdapterBeforePopup", "popupId:", popupId, "targetUrl:", targetUrl)
	*result = true
	var handle bool
	if m.onPopupWindow != nil {
		handle = m.onPopupWindow(targetUrl)
	}
	if !handle && m.window != nil {
		options := m.window.Options()
		if options.AutoPopupWindow && gPrePopupWindow != nil {
			lcl.RunOnMainThreadAsync(func(id uint32) {
				gPrePopupWindow.Browser().Chromium().SetDefaultUrl(targetUrl)
				gPrePopupWindow.Show()
				gPrePopupWindow = nil
			})
		}
	}
}
