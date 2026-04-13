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

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "Cocoa/Cocoa.h"
#include "ns_toolbar.h"

*/
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"unsafe"
)

func NewToolBar(window INSWindow, delegate INSWindowDelegate, config ToolbarConfiguration) {
	if window == nil {
		return
	}
	cConfig := ToolbarConfigurationToOC(config)
	C.CreateToolbar(unsafe.Pointer(window.Instance()), unsafe.Pointer(delegate.Instance()), cConfig)
}
