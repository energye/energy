//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package wv

import (
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

// AddWindowSubviewWebview 将webview作为子视图添加到窗口中
func (m *TWebview) AddWindowSubviewWebview(iWindow window.IWindow) {
	if m.window == nil {
		m.window = iWindow.(window.IDarwinWindow)
	}
	m.nsWindow = lcl.PlatformWindow(iWindow.Instance())
	m.isAddNSWindowSubview = true
	var (
		webviewBounds = m.BoundsRect()
		x, y, w, h    = float32(webviewBounds.Left), float32(webviewBounds.Top), float32(webviewBounds.Width()), float32(webviewBounds.Height())
	)
	m.window.NSWindow().AddSubview(m.wkWebView, x, y, w, h)
}

// UpdateBounds 更新WebView的边界矩形
// 当isAddNSWindowSubview为true时，根据当前的对齐方式和锚点设置来计算并更新WebView的位置和大小
func (m *TWebview) UpdateBounds() {
	if m.isAddNSWindowSubview {
		var (
			webviewAlign   = m.Align()
			windowBounds   = m.window.BoundsRect()
			webviewBounds  = m.BoundsRect()
			x, y, w, h     = float32(webviewBounds.Left), float32(webviewBounds.Top), float32(webviewBounds.Width()), float32(webviewBounds.Height())
			webviewAnchors = m.Anchors()
		)
		//println("UpdateBounds-windowBounds:", windowBounds.Left, windowBounds.Top, windowBounds.Width(), windowBounds.Height())
		//println("UpdateBounds-webviewBounds:", webviewBounds.Left, webviewBounds.Top, webviewBounds.Width(), webviewBounds.Height())
		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			x, y, w, h = float32(webviewBounds.Left), float32(webviewBounds.Top), float32(webviewBounds.Width()), float32(webviewBounds.Height())
		case types.AlClient:
			x, y, w, h = 0, 0, float32(windowBounds.Width()), float32(windowBounds.Height())
		case types.AlLeft, types.AlTop, types.AlRight, types.AlBottom:
			switch webviewAlign {
			case types.AlLeft:
				x, y, w, h = 0, 0, float32(webviewBounds.Width()), float32(windowBounds.Height())
			case types.AlTop:
				x, y, w, h = 0, 0, float32(windowBounds.Width()), float32(webviewBounds.Height())
			case types.AlRight:
				x, y, w, h = float32(windowBounds.Width()-webviewBounds.Width()), 0, float32(webviewBounds.Width()), float32(windowBounds.Height())
			case types.AlBottom:
				x, y, w, h = 0, float32(windowBounds.Height()-webviewBounds.Height()), float32(windowBounds.Width()), float32(webviewBounds.Height())
			}
		}
		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			//akLeft := webviewAnchors.In(types.AkLeft)
			//akTop := webviewAnchors.In(types.AkTop)
			akRight := webviewAnchors.In(types.AkRight)
			akBottom := webviewAnchors.In(types.AkBottom)
			if akRight {
				//println("m.oldBounds.Width()", m.oldBounds.Width())
				if ow := m.oldBounds.Width(); ow > 0 {
					w += float32(windowBounds.Width() - ow)
				}
			}
			if akBottom {
				//println("m.oldBounds.Height()", m.oldBounds.Height())
				if oh := m.oldBounds.Height(); oh > 0 {
					h += float32(windowBounds.Height() - oh)
				}
			}
		}
		m.UpdateWebviewBounds(x, y, w, h)
		m.oldBounds = windowBounds
	}
}

// UpdateWebviewBounds 更新WebView组件的位置和尺寸
// 该方法将指定的坐标和尺寸参数转换为整数类型并设置组件边界，
// 同时通过C语言接口更新原生WebView的显示区域
func (m *TWebview) UpdateWebviewBounds(x, y, width, height float32) {
	m.SetBounds(int32(x), int32(y), int32(width), int32(height))
	m.wkWebView.UpdateBounds(m.window.NSWindow(), x, y, width, height)
}

func (m *TWebview) ExecuteScriptCallback(script string, callback TOnEvaluateScriptCallbackEvent) {
	m.wkWebView.ExecuteScriptCallback(script, func(result string, err string) {
		callback(result, err)
	})
}
