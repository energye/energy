//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 快捷键事件

package cef

import (
	. "github.com/cyber-xxm/energy/v2/consts"
)

type KeyEventCallback func(browse *ICefBrowser, event *TCefKeyEvent, result *bool)

type AcceleratorCustom struct {
	Accelerator string
	KeyType     int
	Callback    KeyEventCallback
}

func acceleratorCode(shift, ctrl, alt bool, keyCode rune) string {
	var accelerator string
	if (shift && ctrl && alt) || (shift && ctrl && !alt) {
		accelerator = MA_Ctrl + "+" + MA_Shift + "+" + string(keyCode)
	} else if shift && alt {
		accelerator = MA_Alt + "+" + MA_Shift + "+" + string(keyCode)
	} else if ctrl || alt {
		accelerator = MA_Ctrl + "+" + string(keyCode)
	} else {
		accelerator = string(keyCode)
	}
	return accelerator
}
