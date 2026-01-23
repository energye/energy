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

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

#include "browser_darwin.h"
*/
import "C"
import (
	"github.com/energye/energy/v3/pkgs/cocoa"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"unsafe"
)

func (m *TWebview) SetWebviewTransparent(isTransparent bool) {
	handle := unsafe.Pointer(m.browser.Data())
	isTransparent = !isTransparent
	v := _BoolToCInt(isTransparent)
	C.SetWebviewTransparent(handle, v)
}

// AddWindowSubviewWebview 将webview作为子视图添加到窗口中
func (m *TWebview) AddWindowSubviewWebview(window window.IWindow) {
	if m.window == nil {
		m.window = window
	}
	m.nsWindow = lcl.PlatformWindow(window.Instance())
	m.isAddNSWindowSubview = true
	var (
		nsWindow           = unsafe.Pointer(m.nsWindow)
		nsWebview          = unsafe.Pointer(m.browser.Data())
		x, y, w, h float32 = 0, 0, 200, 200
	)
	cocoa.WindowAddSubview(nsWindow, nsWebview, x, y, w, h)
	m.UpdateBounds()
}

// UpdateBounds 更新WebView的边界矩形
// 当isAddNSWindowSubview为true时，根据当前的对齐方式和锚点设置来计算并更新WebView的位置和大小
func (m *TWebview) UpdateBounds() {
	if m.isAddNSWindowSubview {
		var (
			webviewAlign             = m.Align()
			windowBoundsRect         = m.window.BoundsRect()
			webviewBounds            = m.BoundsRect()
			x, y, w, h       float32 = 0, 0, 200, 200
			webviewAnchors           = m.Anchors()
		)
		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			x, y, w, h = float32(webviewBounds.Left), float32(webviewBounds.Top), float32(webviewBounds.Width()), float32(webviewBounds.Height())
		case types.AlClient:
			x, y, w, h = 0, 0, float32(windowBoundsRect.Width()), float32(windowBoundsRect.Height())
		case types.AlLeft, types.AlTop, types.AlRight, types.AlBottom:
			switch webviewAlign {
			case types.AlLeft:
				x, y, w, h = 0, 0, float32(webviewBounds.Width()), float32(windowBoundsRect.Height())
			case types.AlTop:
				x, y, w, h = 0, 0, float32(windowBoundsRect.Width()), float32(webviewBounds.Height())
			case types.AlRight:
				x, y, w, h = float32(windowBoundsRect.Width()-webviewBounds.Width()), 0, float32(webviewBounds.Width()), float32(windowBoundsRect.Height())
			case types.AlBottom:
				x, y, w, h = 0, float32(windowBoundsRect.Height()-webviewBounds.Height()), float32(windowBoundsRect.Width()), float32(webviewBounds.Height())
			}
		}
		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			//akLeft := webviewAnchors.In(types.AkLeft)
			//akTop := webviewAnchors.In(types.AkTop)
			akRight := webviewAnchors.In(types.AkRight)
			akBottom := webviewAnchors.In(types.AkBottom)
			if akRight {
				w += float32(windowBoundsRect.Width() - m.oldBounds.Width())
			}
			if akBottom {
				h += float32(windowBoundsRect.Height() - m.oldBounds.Height())
			}
		}
		m.UpdateWebviewBounds(x, y, w, h)
		m.oldBounds = windowBoundsRect
	}
}

// UpdateWebviewBounds 更新WebView组件的位置和尺寸
// 该方法将指定的坐标和尺寸参数转换为整数类型并设置组件边界，
// 同时通过C语言接口更新原生WebView的显示区域
func (m *TWebview) UpdateWebviewBounds(x, y, width, height float32) {
	m.SetBounds(int32(x), int32(y), int32(width), int32(height))
	nsWindow := unsafe.Pointer(m.nsWindow)
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.UpdateWebviewBounds(nsWindow, nsWebview, C.float(x), C.float(y), C.float(width), C.float(height))
}

// BecomeFirstResponder 使webview成为第一响应者，获取焦点并准备接收用户输入
// 该方法将当前webview设置为活动状态，使其能够响应键盘事件和其他用户交互
func (m *TWebview) BecomeFirstResponder() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewBecomeFirstResponder(nsWebview)
}

func (m *TWebview) Undo() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewUndo(nsWebview)
}

func (m *TWebview) Redo() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewRedo(nsWebview)
}

func (m *TWebview) Cut() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewCut(nsWebview)
}

func (m *TWebview) Copy() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewCopy(nsWebview)
}

func (m *TWebview) Paste() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewPaste(nsWebview)
}

func (m *TWebview) SelectAll() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewSelectAll(nsWebview)
}

// _WebViewRegisterPerformKeyMethod 注册WebView执行键盘方法的功能
// 该方法将当前浏览器实例与底层C库的键盘方法执行功能进行绑定
func (m *TWebview) _WebViewRegisterPerformKeyEquivalentMethod() {
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.WebViewRegisterPerformKeyEquivalentMethod(nsWebview)
}

func _BoolToCInt(value bool) C.int {
	if value {
		return C.int(1)
	}
	return C.int(0)
}
