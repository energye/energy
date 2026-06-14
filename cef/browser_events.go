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
	"github.com/energye/lcl/lcl"
)

func (m *TBrowser) initDefaultEvent() {
	m.chromium.SetOnProcessMessageReceived(m.onProcessMessageReceived)
	m.chromium.SetOnBeforeResourceLoad(m.onBeforeResourceLoad)
	m.chromium.SetOnBeforeDownload(m.onBeforeDownload)
	m.chromium.SetOnLoadStart(m.onLoadStart)
	m.chromium.SetOnGetResourceHandler(m.onGetResourceHandler)
	m.chromium.SetOnBeforeContextMenu(m.onBeforeContextMenu)
	m.chromium.SetOnContextMenuCommand(m.onContextMenuCommand)
	m.chromium.SetOnAfterCreated(m.onAfterCreated)
	m.chromium.SetOnKeyEvent(m.onKeyEvent)
	m.chromium.SetOnBeforeBrowse(m.onBeforeBrowse)
	m.chromium.SetOnTitleChange(m.onTitleChange)
	m.chromium.SetOnDragEnter(m.onDragEnter)
	m.chromium.SetOnDraggableRegionsChanged(m.onDraggableRegionsChanged)
	m.chromium.SetOnBeforeClose(m.onBeforeClose)
	m.chromium.SetOnClose(m.onClose)
	m.chromium.SetOnOpenUrlFromTab(m.onOpenUrlFromTab)
	m.chromium.SetOnBeforePopup(m.onBeforePopup)
}

func (m *TBrowser) onProcessMessageReceived(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, outResult *bool) {

}

func (m *TBrowser) onBeforeResourceLoad(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	callback cef.ICefCallback, outResult *cefTypes.TCefReturnValue) {

}

func (m *TBrowser) onAdapterBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {

}

func (m *TBrowser) onLoadStart(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, transitionType cefTypes.TCefTransitionType) {

}

func (m *TBrowser) onGetResourceHandler(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	resourceHandler *cef.IEngResourceHandler) {

}

func (m *TBrowser) onBeforeContextMenu(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, model cef.ICefMenuModel) {

}

func (m *TBrowser) onContextMenuCommand(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams,
	commandId int32, eventFlags cefTypes.TCefEventFlags, outResult *bool) {

}

func (m *TBrowser) onAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {

}

func (m *TBrowser) onKeyEvent(sender lcl.IObject, browser cef.ICefBrowser, event cef.TCefKeyEvent, osEvent cefTypes.TCefEventHandle, outResult *bool) {

}

func (m *TBrowser) onBeforeBrowse(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, userGesture bool,
	isRedirect bool, outResult *bool) {

}

func (m *TBrowser) onTitleChange(sender lcl.IObject, browser cef.ICefBrowser, title string) {

}

func (m *TBrowser) onDragEnter(sender lcl.IObject, browser cef.ICefBrowser, dragData cef.ICefDragData, mask cefTypes.TCefDragOperations, outResult *bool) {

}

func (m *TBrowser) onDraggableRegionsChanged(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, regionsCount cefTypes.NativeUInt,
	regions cef.ICefDraggableRegionArray) {

}

func (m *TBrowser) onBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {

}

func (m *TBrowser) onClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {

}

func (m *TBrowser) onOpenUrlFromTab(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, targetUrl string,
	targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, outResult *bool) {

}

func (m *TBrowser) onAdapterBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, popupId int32, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {

}
