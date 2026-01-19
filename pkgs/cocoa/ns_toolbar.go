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

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "Cocoa/Cocoa.h"
#include "ns_toolbar.h"

*/
import "C"
import "github.com/energye/lcl/lcl"

func NewToolBar(form lcl.IForm, config ToolbarConfiguration) {
	nsWindow := uintptr(lcl.PlatformWindow(form.Instance()))
	if nsWindow == 0 {
		return
	}
	cConfig := ToolbarConfigurationToOC(config)
	C.CreateToolbar(C.ulong(nsWindow), cConfig)
}
