//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
)

func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsChild).Call(windowInfo, windowHandle, api.PascalStr(windowName))
}

func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsPopUp).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(internale_CEFWindowInfoAsWindowless).Call(windowInfo, windowParent, api.PascalStr(windowName))
}
