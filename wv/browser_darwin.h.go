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
	"fmt"
	"github.com/energye/lcl/lcl"
	"unsafe"
)

func (m *TWebview) SetWebviewTransparent(isTransparent bool) {
	handle := unsafe.Pointer(m.browser.Data())
	isTransparent = !isTransparent
	v := _BoolToCInt(isTransparent)
	C.SetWebviewTransparent(handle, v)
}

func (m *TWebview) AddSubviewWebview() {
	fmt.Println("AddSubview")
	//CGRect init = { 0,0,0,0 };
	//[self.webview initWithFrame:init configuration:config];
	//[contentView addSubview:self.webview];
	//[self.webview setAutoresizingMask: NSViewWidthSizable|NSViewHeightSizable];
	//CGRect contentViewBounds = [contentView bounds];
	//[self.webview setFrame:contentViewBounds];

	nsWindow := unsafe.Pointer(lcl.PlatformWindow(m.window.Instance()))
	nsWebview := unsafe.Pointer(m.browser.Data())
	nsWebview = unsafe.Pointer(m.Handle())
	C.AddSubviewWebview(nsWindow, nsWebview)

	//m.SetBounds()
}

func _BoolToCInt(value bool) C.int {
	if value {
		return C.int(1)
	}
	return C.int(0)
}
