package cgo

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

#include "wk_webview.h"

*/
import "C"
import (
	"github.com/energye/energy/v3/platform/darwin/cocoa/cgo"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type WkWebView struct {
	cgo.NSView
}

func AsWkWebView(ptr unsafe.Pointer) IWkWebView {
	if ptr == nil {
		return nil
	}
	m := &WkWebView{}
	m.SetInstance(ptr)
	return m
}

func (m *WkWebView) SetWebviewTransparent(isTransparent bool) {
	webview := unsafe.Pointer(m.Instance())
	isTransparent = !isTransparent
	v := _BoolToCInt(isTransparent)
	C.SetWebviewTransparent(webview, v)
}

// UpdateBounds 更新WebView组件的位置和尺寸
// 该方法将指定的坐标和尺寸参数转换为整数类型并设置组件边界，
// 同时通过C语言接口更新原生WebView的显示区域
func (m *WkWebView) UpdateBounds(window INSWindow, x, y, width, height float32) {
	nsWindow := unsafe.Pointer(window.Instance())
	webview := unsafe.Pointer(m.Instance())
	C.UpdateWebviewBounds(nsWindow, webview, C.float(x), C.float(y), C.float(width), C.float(height))
	//println("UpdateWebviewBounds:", int(x), int(y), int(width), int(height))
}

// BecomeFirstResponder 使webview成为第一响应者，获取焦点并准备接收用户输入
// 该方法将当前webview设置为活动状态，使其能够响应键盘事件和其他用户交互
func (m *WkWebView) BecomeFirstResponder() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewBecomeFirstResponder(webview)
}

func (m *WkWebView) Undo() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewUndo(webview)
}

func (m *WkWebView) Redo() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewRedo(webview)
}

func (m *WkWebView) Cut() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewCut(webview)
}

func (m *WkWebView) Copy() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewCopy(webview)
}

func (m *WkWebView) Paste() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewPaste(webview)
}

func (m *WkWebView) SelectAll() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewSelectAll(webview)
}

// RegisterPerformKeyEquivalentMethod 注册WebView执行键盘方法的功能
// 该方法将当前浏览器实例与底层C库的键盘方法执行功能进行绑定
func (m *WkWebView) RegisterPerformKeyEquivalentMethod() {
	webview := unsafe.Pointer(m.Instance())
	C.WebViewRegisterPerformKeyEquivalentMethod(webview)
}

func (m *WkWebView) ConvertPoint(inPoint types.TPoint) (point types.TPoint) {
	webview := unsafe.Pointer(m.Instance())
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
