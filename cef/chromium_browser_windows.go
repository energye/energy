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
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 每一次拖拽区域改变都需要重新设置
func (m *TCEFChromiumBrowser) setDraggableRegions() {
	var scp float32
	// Windows 10 版本 1607 [仅限桌面应用]
	// Windows Server 2016 [仅限桌面应用]
	// 可动态调整
	dpi, err := winapi.GetDpiForWindow(et.HWND(m.window.Handle()))
	if err == nil {
		scp = float32(dpi) / 96.0
	} else {
		// 使用默认的，但不能动态调整
		scp = winapi.ScalePercent()
	}
	//在主线程中运行
	RunOnMainThread(func() {
		if m.rgn == nil {
			//第一次时创建RGN
			m.rgn = winapi.CreateRectRgn(0, 0, 0, 0)
		} else {
			//每次重置RGN
			winapi.SetRectRgn(m.rgn, 0, 0, 0, 0)
		}
		// 重新根据缩放比计算新的区域位置
		for i := 0; i < m.regions.RegionsCount(); i++ {
			region := m.regions.Region(i)
			x := int32(float32(region.Bounds.X) * scp)
			y := int32(float32(region.Bounds.Y) * scp)
			w := int32(float32(region.Bounds.Width) * scp)
			h := int32(float32(region.Bounds.Height) * scp)
			creRGN := winapi.CreateRectRgn(x, y, x+w, y+h)
			if region.Draggable {
				winapi.CombineRgn(m.rgn, m.rgn, creRGN, consts.RGN_OR)
			} else {
				winapi.CombineRgn(m.rgn, m.rgn, creRGN, consts.RGN_DIFF)
			}
			winapi.DeleteObject(creRGN)
		}
	})
}

// 注册windows下CompMsg事件
func (m *TCEFChromiumBrowser) registerWindowsCompMsgEvent() {
	if m.window != nil && m.window.IsLCL() {
		window := m.window.AsLCLBrowserWindow().BrowserWindow()
		var bwEvent = BrowserWindow.browserEvent
		m.Chromium().SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
			if bwEvent.onRenderCompMsg != nil {
				bwEvent.onRenderCompMsg(sender, message, lResult, aHandled)
			}
			if !*aHandled && window.cwcap != nil {
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

		//if window.WindowProperty().EnableWebkitAppRegion && window.WindowProperty().EnableWebkitAppRegionDClk {
		//	window.windowResize = func(sender lcl.IObject) bool {
		//		if window.WindowState() == types.WsMaximized && (window.WindowProperty().EnableHideCaption ||
		//			window.BorderStyle() == types.BsNone || window.BorderStyle() == types.BsSingle) {
		//			var monitor = window.Monitor().WorkareaRect()
		//			window.SetBounds(monitor.Left, monitor.Top, monitor.Right-monitor.Left, monitor.Bottom-monitor.Top)
		//			window.SetWindowState(types.WsMaximized)
		//		}
		//		return false
		//	}
		//}
	}
}
