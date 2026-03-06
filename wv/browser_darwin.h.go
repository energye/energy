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

//export evaluateScriptCallback
func evaluateScriptCallback(cCallbackID C.int, resC *C.char, errC *C.char) {
	var (
		result, err string
		callbackID  int
	)
	callbackID = int(cCallbackID)
	if resC != nil {
		result = C.GoString(resC)
	}
	if errC != nil {
		err = C.GoString(errC)
	}
	if callback, ok := gEvaluateScriptEventCallback.Load(callbackID); ok {
		gEvaluateScriptEventCallback.Delete(callbackID)
		callback.(TOnEvaluateScriptCallbackEvent)(result, err)
	}
}

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
		nsWindow      = unsafe.Pointer(m.nsWindow)
		webview       = unsafe.Pointer(m.browser.Data())
		webviewBounds = m.BoundsRect()
		x, y, w, h    = float32(webviewBounds.Left), float32(webviewBounds.Top), float32(webviewBounds.Width()), float32(webviewBounds.Height())
	)

	cocoa.WindowAddSubview(nsWindow, webview, x, y, w, h)
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
	nsWindow := unsafe.Pointer(m.nsWindow)
	webview := unsafe.Pointer(m.browser.Data())
	C.UpdateWebviewBounds(nsWindow, webview, C.float(x), C.float(y), C.float(width), C.float(height))
	//println("UpdateWebviewBounds:", int(x), int(y), int(width), int(height))
}

// BecomeFirstResponder 使webview成为第一响应者，获取焦点并准备接收用户输入
// 该方法将当前webview设置为活动状态，使其能够响应键盘事件和其他用户交互
func (m *TWebview) BecomeFirstResponder() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewBecomeFirstResponder(webview)
}

func (m *TWebview) Undo() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewUndo(webview)
}

func (m *TWebview) Redo() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewRedo(webview)
}

func (m *TWebview) Cut() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewCut(webview)
}

func (m *TWebview) Copy() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewCopy(webview)
}

func (m *TWebview) Paste() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewPaste(webview)
}

func (m *TWebview) SelectAll() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewSelectAll(webview)
}

func (m *TWebview) ExecuteScriptCallback(script string, callback TOnEvaluateScriptCallbackEvent) {
	if script == "" || callback == nil {
		return
	}
	webview := unsafe.Pointer(m.browser.Data())
	cScript := C.CString(script)
	defer C.free(unsafe.Pointer(cScript))
	eventID := gNextEvaluateScriptEventID()
	cEventID := C.int(eventID)
	cCallback := (C.CGoEvaluateScriptCallback)(C.evaluateScriptCallback)
	gEvaluateScriptEventCallback.Store(eventID, callback)
	C.WebViewEvaluateScriptCallback(webview, cEventID, cScript, cCallback)
}

// _WebViewRegisterPerformKeyMethod 注册WebView执行键盘方法的功能
// 该方法将当前浏览器实例与底层C库的键盘方法执行功能进行绑定
func (m *TWebview) _WebViewRegisterPerformKeyEquivalentMethod() {
	webview := unsafe.Pointer(m.browser.Data())
	C.WebViewRegisterPerformKeyEquivalentMethod(webview)
}

func (m *TWebview) ConvertPoint(inPoint types.TPoint) (point types.TPoint) {
	webview := unsafe.Pointer(m.browser.Data())
	outPoint := C.ConvertPoint(webview, C.float(inPoint.X), C.float(inPoint.Y))
	point.X = int32(outPoint.x)
	point.Y = int32(outPoint.y)
	return
}

func _BoolToCInt(value bool) C.int {
	if value {
		return C.int(1)
	}
	return C.int(0)
}
