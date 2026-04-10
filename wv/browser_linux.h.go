//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package wv

/*
#cgo pkg-config: webkit2gtk-4.0

#include <webkit2/webkit2.h>
#include "browser_linux.h"

void WebkitSetBackgroundColor(WebKitWebView *webview, gdouble r, gdouble g, gdouble b, gdouble a) {
	if (webview != NULL && WEBKIT_IS_WEB_VIEW(webview))
    {
		GdkRGBA colour = {r, g, b, a};
        webkit_web_view_set_background_color(WEBKIT_WEB_VIEW(webview), &colour);
    }
}

*/
import "C"
import (
	"github.com/energye/energy/v3/pkgs/linux/gtk3"
	"github.com/energye/energy/v3/pkgs/linux/gtk3/callback"
	gtk3Types "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

type TGtkWebview struct {
	wkWebview *C.WebKitWebView
	wsWebview gtk3Types.IWidget
}

func (m *TWebview) getCWkWebview() *C.WebKitWebView {
	if m.gtkWebview == nil {
		m.gtkWebview = &TGtkWebview{}
	}
	if m.gtkWebview.wkWebview == nil {
		webview := m.browser.WebView()
		m.gtkWebview.wkWebview = (*C.WebKitWebView)(unsafe.Pointer(webview))
	}
	return m.gtkWebview.wkWebview
}

func (m *TWebview) getGtkWebview() gtk3Types.IWidget {
	if m.gtkWebview == nil {
		m.gtkWebview = &TGtkWebview{}
	}
	if m.gtkWebview.wsWebview == nil {
		webview := m.browser.WebView()
		m.gtkWebview.wsWebview = gtk3.AsWidget(unsafe.Pointer(webview))
	}
	return m.gtkWebview.wsWebview
}

func (m *TWebview) mustGtkWebview() {
	m.getCWkWebview()
	m.getGtkWebview()
}

func (m *TWebview) SetBackgroundColor(color *colors.TARGB) {
	if color == nil {
		return
	}
	cWebview := m.getCWkWebview()
	cR := C.gdouble(float64(color.R) / 255.0)
	cG := C.gdouble(float64(color.G) / 255.0)
	cB := C.gdouble(float64(color.B) / 255.0)
	cA := C.gdouble(float64(color.A) / 255.0)
	C.WebkitSetBackgroundColor(cWebview, cR, cG, cB, cA)
}

func (m *TWebview) AddWindowWebview(iWindow window.IWindow) {
	if m.window == nil {
		m.window = iWindow.(window.ILinuxWindow)
	}
	m.isAddWindowSubview = true
	var (
		webviewBounds = m.BoundsRect()
		x, y, w, h    = webviewBounds.Left, webviewBounds.Top, webviewBounds.Width(), webviewBounds.Height()
	)

	windowLayout := m.window.GTKWindowLayout()

	m.gtkScrolledWindow.SetSizeRequest(int(w), int(h))
	m.gtkScrolledWindow.Add(m.getGtkWebview())

	windowLayout.Add(m.gtkScrolledWindow)
	windowLayout.Move(m.gtkScrolledWindow, int(x), int(y))

}

func (m *TWebview) UpdateBounds() {
	if m.isAddWindowSubview {
		allocation := m.window.GTKWindowScrolledWindow().GetAllocation()
		swx := int32(allocation.GetX())
		swy := int32(allocation.GetY())
		var (
			webviewAlign     = m.Align()
			webviewAnchors   = m.Anchors()
			windowBoundsRect = m.window.BoundsRect()
			webviewBounds    = m.BoundsRect()
			x, y, w, h       = webviewBounds.Left, webviewBounds.Top, webviewBounds.Width(), webviewBounds.Height()
		)
		// 真实的客户区大小, 当有菜单栏时
		windowBoundsRect.SetSize(windowBoundsRect.Width()-swx, windowBoundsRect.Height()-swy)

		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			x, y, w, h = webviewBounds.Left, webviewBounds.Top, webviewBounds.Width(), webviewBounds.Height()
		case types.AlClient:
			x, y, w, h = 0, 0, windowBoundsRect.Width(), windowBoundsRect.Height()
		case types.AlLeft, types.AlTop, types.AlRight, types.AlBottom:
			switch webviewAlign {
			case types.AlLeft:
				x, y, w, h = 0, 0, webviewBounds.Width(), windowBoundsRect.Height()
			case types.AlTop:
				x, y, w, h = 0, 0, windowBoundsRect.Width(), webviewBounds.Height()
			case types.AlRight:
				x, y, w, h = webviewBounds.Left, 0, webviewBounds.Width()+(windowBoundsRect.Width()-webviewBounds.Width()), windowBoundsRect.Height()
			case types.AlBottom:
				x, y, w, h = 0, windowBoundsRect.Height()-webviewBounds.Height(), windowBoundsRect.Width(), webviewBounds.Height()
			}
		}
		switch webviewAlign {
		case types.AlNone, types.AlCustom:
			//akLeft := webviewAnchors.In(types.AkLeft)
			//akTop := webviewAnchors.In(types.AkTop)
			akRight := webviewAnchors.In(types.AkRight)
			akBottom := webviewAnchors.In(types.AkBottom)
			//fmt.Println("webviewAlign:", webviewAlign, "webviewAnchors:", webviewAnchors, "akRight:", akRight, "akBottom:", akBottom)
			if akRight {
				if ow := m.oldBounds.Width(); ow > 0 {
					w += windowBoundsRect.Width() - ow
				}
			}
			if akBottom {
				if oh := m.oldBounds.Height(); oh > 0 {
					h += windowBoundsRect.Height() - oh
				}
			}
		}
		m.UpdateWebviewBounds(x, y, w, h)
		m.oldBounds = windowBoundsRect
	}
}

func (m *TWebview) UpdateWebviewBounds(x, y, width, height int32) {
	m.SetBounds(x, y, width, height)
	m.window.GTKWindowLayout().Move(m.gtkScrolledWindow, int(x), int(y))
	m.gtkScrolledWindow.SetSizeRequest(int(width), int(height))
}

func (m *TGtkWebview) SetOnDragDataReceived(fn gtk3Types.TDragDataReceivedEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragDataReceivedEvent,
		"c_trampoline_8_void_drag_data_received", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragDrop(fn gtk3Types.TDragDropEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragDropEvent,
		"c_trampoline_6_gboolean_drag_drop_motion", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragMotion(fn gtk3Types.TDragMotionEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragMotionEvent,
		"c_trampoline_6_gboolean_drag_drop_motion", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragLeave(fn gtk3Types.TDragLeaveEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragLeaveEvent,
		"c_trampoline_4_void", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragDataDelete(fn gtk3Types.TDragDataDeleteOrBeginOrEndEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragDataDeleteEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragBegin(fn gtk3Types.TDragDataDeleteOrBeginOrEndEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragBeginEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}

func (m *TGtkWebview) SetOnDragEnd(fn gtk3Types.TDragDataDeleteOrBeginOrEndEvent) gtk3Types.ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.wsWebview.Instance()), gtk3Types.EsnDragEndEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}
