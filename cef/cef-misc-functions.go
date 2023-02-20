//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF的复杂函数导出
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
)

// BrowserWindow 设置到指定窗口做为子窗口
func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsChild).Call(windowInfo, windowHandle, api.PascalStr(windowName))
}

// BrowserWindow 设置到做为弹出窗口
func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsPopUp).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

// BrowserWindow 设置到做为无窗口
func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsWindowless).Call(windowInfo, windowParent, api.PascalStr(windowName))
}
