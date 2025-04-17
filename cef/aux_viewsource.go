//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 辅助工具-显示网页源代码

package cef

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/common"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

const (
	viewSourceName = "view-source"
)

func (m *ICefBrowser) createBrowserViewSource(currentWindow IBrowserWindow) {
	if currentWindow != nil {
		var frame = m.MainFrame()
		if currentWindow.IsLCL() {
			var viewSourceUrl = fmt.Sprintf("view-source:%s", frame.Url())
			wp := NewWindowProperty()
			wp.Url = viewSourceUrl
			wp.Title = fmt.Sprintf("%s - %s", viewSourceName, frame.Url())
			wp.WindowType = WT_VIEW_SOURCE
			viewSourceWindow := NewLCLBrowserWindow(nil, wp, nil)
			viewSourceWindow.SetWidth(800)
			viewSourceWindow.SetHeight(600)
			if common.IsDarwin() || common.IsLinux() {
				viewSourceWindow.Chromium().SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
					viewSourceWindow.Chromium().LoadUrl(viewSourceUrl)
				})
			}
			if assetserve.AssetsServerHeaderKeyValue != "" {
				viewSourceWindow.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *TCefReturnValue) {
					if application.Is49() {
						headerMap := request.GetHeaderMap()
						headerMap.Append(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue)
						request.SetHeaderMap(headerMap)
						headerMap.Free()
					} else {
						request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
					}
				})
			}
			viewSourceWindow.Chromium().SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, client *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool {
				wp := NewWindowProperty()
				wp.Url = beforePopupInfo.TargetUrl
				wp.Title = beforePopupInfo.TargetUrl
				wp.WindowType = WT_VIEW_SOURCE
				bw := NewLCLBrowserWindow(nil, wp, nil)
				bw.SetWidth(800)
				bw.SetHeight(600)
				bw.EnableDefaultCloseEvent()
				QueueAsyncCall(func(id int) { //main thread run
					bw.Show()
				})
				return true
			})
			viewSourceWindow.EnableDefaultCloseEvent()
			QueueAsyncCall(func(id int) { //main thread run
				viewSourceWindow.Show()
			})
		} else if currentWindow.IsViewsFramework() {
			frame.ViewSource()
		}
	}
}
