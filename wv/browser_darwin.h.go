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
	"github.com/energye/lcl/lcl"
	"unsafe"
)

func (m *TWebview) SetWebviewTransparent(isTransparent bool) {
	handle := unsafe.Pointer(m.browser.Data())
	isTransparent = !isTransparent
	v := _BoolToCInt(isTransparent)
	C.SetWebviewTransparent(handle, v)
}

func (m *TWebview) AddFormSubviewWebview(form lcl.IEngForm) {
	var (
		nsWindow  unsafe.Pointer
		nsWebview unsafe.Pointer
	)
	//webviewBounds := m.BoundsRect()
	//webviewAlign := m.Align()
	//webviewAnchors := m.Anchors()

	m.nsWindow = lcl.PlatformWindow(form.Instance())
	nsWindow = unsafe.Pointer(m.nsWindow)
	nsWebview = unsafe.Pointer(m.browser.Data())
	cocoa.WindowAddSubview(nsWindow, nsWebview)
}

func (m *TWebview) updateBounds() {
	nsWindow := unsafe.Pointer(m.nsWindow)
	nsWebview := unsafe.Pointer(m.browser.Data())
	C.UpdateWebviewBounds(nsWindow, nsWebview)
}

func _BoolToCInt(value bool) C.int {
	if value {
		return C.int(1)
	}
	return C.int(0)
}
