//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/golcl/lcl/api"
)

func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	Proc(internale_CEFWindowInfoAsChild).Call(windowInfo, windowHandle, api.PascalStr(windowName))
}

func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	Proc(internale_CEFWindowInfoAsPopUp).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	Proc(internale_CEFWindowInfoAsWindowless).Call(windowInfo, windowParent, api.PascalStr(windowName))
}
