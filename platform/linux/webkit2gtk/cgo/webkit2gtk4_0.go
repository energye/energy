//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package cgo

/*
#cgo pkg-config: webkit2gtk-4.0

#include <webkit2/webkit2.h>
#include "webkit2gtk4_0.go.h"

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
	"github.com/energye/energy/v3/platform/linux/callback"
	"github.com/energye/energy/v3/platform/linux/gtk3/cgo"
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

type Webkit2 struct {
	cgo.Widget
}

func AsWebkit2(ptr unsafe.Pointer) IWebkit2 {
	if ptr == nil {
		return nil
	}
	m := new(Webkit2)
	m.Object = cgo.ToGoObject(ptr)
	return m
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

func (m *Webkit2) SetOnDragDataReceived(fn TDragDataReceivedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragDataReceivedEvent,
		"c_trampoline_8_void_drag_data_received", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDrop(fn TDragDropEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragDropEvent,
		"c_trampoline_6_gboolean_drag_drop_motion", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragMotion(fn TDragMotionEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragMotionEvent,
		"c_trampoline_6_gboolean_drag_drop_motion", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragLeave(fn TDragLeaveEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragLeaveEvent,
		"c_trampoline_4_void", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDataDelete(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragDataDeleteEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragBegin(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragBeginEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragEnd(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(m.Instance()), EsnDragEndEvent,
		"c_trampoline_3_void", fn, nil)
	return signalHandlerID
}
