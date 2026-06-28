//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build CEF147

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/lcl/lcl"
)

func (m *TBrowser) chromiumOnBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {

	m.chromiumOnAdapterBeforeDownload(sender, browser, downloadItem, suggestedName, callback, result)
}

func (m *TBrowser) chromiumOnBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, popupId int32, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {

	m.chromiumOnAdapterBeforePopup(sender, browser, frame, popupId, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, windowInfo, client,
		settings, extraInfo, noJavascriptAccess, result)
}

func (m *TViewsBrowser) initViewsBrowserTagsDefaultEvent() {
	if window147, ok := m.window.(cef.ICEFWindowComponent_147); ok {
		const CEF_ColorPrimaryBackground int32 = 225 // CEF_ColorPrimaryBackground, CEF 147
		window147.SetOnThemeColorsChanged(func(sender lcl.IObject, window cef.ICefWindow, chromeTheme int32) {
			backgroundColor := window147.GetThemeColor(CEF_ColorPrimaryBackground)
			isDark := ColorIsDark(backgroundColor)
			m.windowOnAdapterThemeChanged(sender, isDark)
		})
	}
}
