//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	"github.com/energye/golcl/lcl/api"
)

func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	Proc("CEFWindowInfoAsChild").Call(windowInfo, windowHandle, api.GoStrToDStr(windowName))
}

func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	Proc("CEFWindowInfoAsPopUp").Call(windowInfo, windowParent, api.GoStrToDStr(windowName))
}

func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	Proc("CEFWindowInfoAsWindowless").Call(windowInfo, windowParent, api.GoStrToDStr(windowName))
}
