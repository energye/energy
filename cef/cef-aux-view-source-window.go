//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

const (
	view_source_name = "view-source"
)

func (m *ICefBrowser) createBrowserViewSource(frame *ICefFrame) {
	if currentWindowInfo := BrowserWindow.GetWindowInfo(m.Identifier()); currentWindowInfo != nil {
		if currentWindowInfo.IsLCL() {
			var viewSourceUrl = fmt.Sprintf("view-source:%s", frame.Url)
			QueueAsyncCall(func(id int) {
				wp := NewWindowProperty()
				wp.Url = viewSourceUrl
				wp.Title = fmt.Sprintf("%s - %s", view_source_name, frame.Url)
				wp.WindowType = WT_VIEW_SOURCE
				viewSourceWindow := NewLCLBrowserWindow(nil, wp)
				viewSourceWindow.SetWidth(800)
				viewSourceWindow.SetHeight(600)
				if common.IsDarwin() {
					viewSourceWindow.Chromium().SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
						viewSourceWindow.Chromium().LoadUrl(viewSourceUrl)
					})
				}
				if assetserve.AssetsServerHeaderKeyValue != "" {
					viewSourceWindow.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *TCefReturnValue) {
						request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
					})
				}
				viewSourceWindow.EnableDefaultCloseEvent()
				viewSourceWindow.Show()
			})
		} else if currentWindowInfo.IsViewsFramework() {
			frame.ViewSource()
		}
	}
}
