//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !CEF127 && !CEF147

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool/exec"
	"path/filepath"
)

func (m *TBrowser) chromiumOnAdapterBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {
	callback.Cont(filepath.Join(exec.AppDir(), suggestedName), true)
	*result = true
}

func (m *TBrowser) chromiumOnBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback) {

	var result bool
	m.chromiumOnAdapterBeforeDownload(sender, browser, downloadItem, suggestedName, callback, &result)
}

func (m *TBrowser) chromiumOnBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {
	m.chromiumOnAdapterBeforePopup(sender, browser, frame, -1, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, windowInfo, client,
		settings, extraInfo, noJavascriptAccess, result)
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
