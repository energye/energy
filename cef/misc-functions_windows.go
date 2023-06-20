//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

// 函数工具 - windows

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/types"
)

func CefIsKeyDown(wparam types.WPARAM) bool {
	return winapi.WinGetKeyState(types.Int32(wparam)) < 0
}

func CefIsKeyToggled(wparam types.WPARAM) bool {
	return (int16(winapi.WinGetKeyState(types.Int32(wparam))) & 0x1) != 0
}

func GetCefMouseModifiersByWPARAM(wparam types.WPARAM) (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if (wparam & winapi.MK_CONTROL) != 0 {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if (wparam & winapi.MK_SHIFT) != 0 {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if (wparam & winapi.MK_LBUTTON) != 0 {
		result = result | consts.EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if (wparam & winapi.MK_MBUTTON) != 0 {
		result = result | consts.EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if (wparam & winapi.MK_RBUTTON) != 0 {
		result = result | consts.EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(winapi.VK_MENU) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(winapi.VK_NUMLOCK) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(winapi.VK_CAPITAL) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefMouseModifiers() (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if CefIsKeyDown(winapi.MK_CONTROL) {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(winapi.MK_SHIFT) {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(winapi.MK_LBUTTON) {
		result = result | consts.EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if CefIsKeyDown(winapi.MK_MBUTTON) {
		result = result | consts.EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if CefIsKeyDown(winapi.MK_RBUTTON) {
		result = result | consts.EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(winapi.VK_MENU) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(winapi.VK_NUMLOCK) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(winapi.VK_CAPITAL) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefKeyboardModifiers(aWparam types.WPARAM, aLparam types.LPARAM) (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if CefIsKeyDown(winapi.VK_SHIFT) {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(winapi.VK_CONTROL) {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(winapi.VK_MENU) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(winapi.VK_NUMLOCK) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(winapi.VK_CAPITAL) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	switch aWparam {
	case winapi.VK_RETURN:
		if ((aLparam >> 16) & winapi.KF_EXTENDED) != 0 {
			result = result | consts.EVENTFLAG_IS_KEY_PAD
		}
	case winapi.VK_INSERT, winapi.VK_DELETE, winapi.VK_HOME, winapi.VK_END, winapi.VK_PRIOR, winapi.VK_NEXT, winapi.VK_UP, winapi.VK_DOWN, winapi.VK_LEFT, winapi.VK_RIGHT:
		if ((aLparam >> 16) & winapi.KF_EXTENDED) == 0 {
			result = result | consts.EVENTFLAG_IS_KEY_PAD
		}
	case winapi.VK_NUMLOCK, winapi.VK_NUMPAD0, winapi.VK_NUMPAD1, winapi.VK_NUMPAD2, winapi.VK_NUMPAD3, winapi.VK_NUMPAD4, winapi.VK_NUMPAD5, winapi.VK_NUMPAD6, winapi.VK_NUMPAD7,
		winapi.VK_NUMPAD8, winapi.VK_NUMPAD9, winapi.VK_DIVIDE, winapi.VK_MULTIPLY, winapi.VK_SUBTRACT, winapi.VK_ADD, winapi.VK_DECIMAL, winapi.VK_CLEAR:
		result = result | consts.EVENTFLAG_IS_KEY_PAD
	case winapi.VK_SHIFT:
		if CefIsKeyDown(winapi.VK_LSHIFT) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(winapi.VK_RSHIFT) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case winapi.VK_CONTROL:
		if CefIsKeyDown(winapi.VK_LCONTROL) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(winapi.VK_RCONTROL) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case winapi.VK_MENU:
		if CefIsKeyDown(winapi.VK_LMENU) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(winapi.VK_RMENU) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case winapi.VK_LWIN:
		result = result | consts.EVENTFLAG_IS_LEFT
	case winapi.VK_RWIN:
		result = result | consts.EVENTFLAG_IS_RIGHT
	}
	return
}
