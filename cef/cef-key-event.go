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
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
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

var defaultAcceleratorCustom = func() {
	//macos 下快捷键
	if common.IsDarwin() {
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+a",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().SelectAll()
				*result = true
			},
		})
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+x",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().Cut()
				*result = true
			},
		})
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+c",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().Copy()
				*result = true
			},
		})
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+v",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().Paste()
				*result = true
			},
		})
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+z",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().Undo()
				*result = true
			},
		})
		KeyAccelerator.AddAcceleratorCustom(&AcceleratorCustom{
			Accelerator: "ctrl+shift+z",
			Callback: func(browse *ICefBrowser, event *TCefKeyEvent, result *bool) {
				browse.GetFocusedFrame().Redo()
				*result = true
			},
		})
	}
}
