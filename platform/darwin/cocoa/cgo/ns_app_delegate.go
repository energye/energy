//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "ns_app_delegate.h"

void InitAppDelegate(void);
*/
import "C"

//export doOnAppDelegateCallback
func doOnAppDelegateCallback(cContext *C.TCallbackContext) {
	ctx := CCallbackContextToGo(cContext)
	doDispatchEvent(ctx)
}

// InitAppDelegate 初始化 macOS 应用程序代理
func (m *NSApp) InitAppDelegate() {
	if m.initializationAppDelegate {
		return
	}
	m.initializationAppDelegate = true
	C.InitAppDelegate()
}
