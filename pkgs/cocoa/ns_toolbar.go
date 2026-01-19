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

*/
import "C"
import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
)

// NSToolBar 绑定到指定窗口
// 具有 toolbar delegate 实例
type NSToolBar struct {
	owner        lcl.IForm
	toolbar      Pointer
	delegate     Pointer
	config       *ToolbarConfiguration
	windowResize NotifyEvent
	items        tool.ArrayMap[string, IView]
}
