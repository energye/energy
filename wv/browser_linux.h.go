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
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

func (m *TWebview) SetBackgroundColor(color *colors.TARGB) {
	if color == nil {
		return
	}
	webview := m.browser.WebView()
	cWebview := (*C.WebKitWebView)(unsafe.Pointer(webview))
	cR := C.gdouble(float64(color.R) / 255.0)
	cG := C.gdouble(float64(color.G) / 255.0)
	cB := C.gdouble(float64(color.B) / 255.0)
	cA := C.gdouble(float64(color.A) / 255.0)
	C.WebkitSetBackgroundColor(cWebview, cR, cG, cB, cA)
}
