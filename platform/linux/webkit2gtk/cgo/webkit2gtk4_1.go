//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo && webkit2_4_1

package cgo

/*
#cgo pkg-config: webkit2gtk-4.1

#include <webkit2/webkit2.h>
#include "webkit2gtk4_x.go.h"

void WebkitSetBackgroundColor(WebKitWebView *webview, gdouble r, gdouble g, gdouble b, gdouble a) {
	if (webview != NULL && WEBKIT_IS_WEB_VIEW(webview))
    {
		GdkRGBA colour = {r, g, b, a};
        webkit_web_view_set_background_color(WEBKIT_WEB_VIEW(webview), &colour);
    }
}

void WebkitOpenDevTools(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return;

    WebKitSettings *settings = webkit_web_view_get_settings(webview);

    if (!webkit_settings_get_enable_developer_extras(settings)) {
        return;
    }

    WebKitWebInspector *inspector = webkit_web_view_get_inspector(webview);
    webkit_web_inspector_show(inspector);
}


*/
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/gtk3/cgo"
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/types/colors"
	wvTypes "github.com/energye/wv/types/linux"
	"unsafe"
)

const Wkv = wvTypes.Wkv4_1

type Webkit2 struct {
	cgo.Widget
}

func (m *Webkit2) OpenDevTools() {
	webview := (*C.WebKitWebView)(unsafe.Pointer(m.Instance()))
	C.WebkitOpenDevTools(webview)
}

func (m *Webkit2) SetBackgroundColor(color *colors.TARGB) {
	if color == nil {
		return
	}
	cR := C.gdouble(float64(color.R) / 255.0)
	cG := C.gdouble(float64(color.G) / 255.0)
	cB := C.gdouble(float64(color.B) / 255.0)
	cA := C.gdouble(float64(color.A) / 255.0)
	C.WebkitSetBackgroundColor((*C.WebKitWebView)(unsafe.Pointer(m.Instance())), cR, cG, cB, cA)
}
