//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 注册windows下CompMsg事件
func (m *TCEFChromiumBrowser) registerWindowsCompMsgEvent() {
	if m.window != nil {
		window := m.window.AsLCLBrowserWindow().BrowserWindow()
		var bwEvent = BrowserWindow.browserEvent
		m.Chromium().SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
			if bwEvent.onRenderCompMsg != nil {
				bwEvent.onRenderCompMsg(sender, message, lResult, aHandled)
			}
			if !*aHandled {
				window.doOnRenderCompMsg(m, cmtCEF, message, lResult, aHandled)
			}
		})
		// TODO 暂时不使用
		//m.SetOnWndProc(func(msg *types.TMessage) {
		//	var (
		//		tmpHandled bool
		//		lResult    types.LRESULT
		//	)
		//	m.doOnRenderCompMsg(cmtLCL, msg, &lResult, &tmpHandled)
		//	if tmpHandled {
		//		msg.Result = lResult
		//	}
		//})

		if window.WindowProperty().EnableWebkitAppRegion && window.WindowProperty().EnableWebkitAppRegionDClk {
			window.windowResize = func(sender lcl.IObject) bool {
				if window.WindowState() == types.WsMaximized && (window.WindowProperty().EnableHideCaption || window.BorderStyle() == types.BsNone || window.BorderStyle() == types.BsSingle) {
					var monitor = window.Monitor().WorkareaRect()
					window.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left, monitor.Bottom-monitor.Top)
					window.SetWindowState(types.WsMaximized)
				}
				return false
			}
		}
	}
}
