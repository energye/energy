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
// +build windows

// 函数工具 - windows

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/consts/messages"
	"github.com/cyber-xxm/energy/v2/types"
)

func CefIsKeyDown(wparam types.WPARAM) bool {
	return winapi.GetKeyState(types.Int32(wparam)) < 0
}

func CefIsKeyToggled(wparam types.WPARAM) bool {
	return (int16(winapi.GetKeyState(types.Int32(wparam))) & 0x1) != 0
}

func GetCefMouseModifiersByWPARAM(wparam types.WPARAM) (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if (wparam & messages.MK_CONTROL) != 0 {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if (wparam & messages.MK_SHIFT) != 0 {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if (wparam & messages.MK_LBUTTON) != 0 {
		result = result | consts.EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if (wparam & messages.MK_MBUTTON) != 0 {
		result = result | consts.EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if (wparam & messages.MK_RBUTTON) != 0 {
		result = result | consts.EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(consts.VkMenu) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(consts.VkNumLock) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(consts.VkCapital) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefMouseModifiers() (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if CefIsKeyDown(messages.MK_CONTROL) {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(messages.MK_SHIFT) {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(messages.MK_LBUTTON) {
		result = result | consts.EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if CefIsKeyDown(messages.MK_MBUTTON) {
		result = result | consts.EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if CefIsKeyDown(messages.MK_RBUTTON) {
		result = result | consts.EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(consts.VkMenu) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(consts.VkNumLock) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(consts.VkCapital) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefKeyboardModifiers(aWparam types.WPARAM, aLparam types.LPARAM) (result consts.TCefEventFlags) {
	result = consts.EVENTFLAG_NONE
	if CefIsKeyDown(consts.VkShift) {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(consts.VkControl) {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(consts.VkMenu) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(consts.VkNumLock) {
		result = result | consts.EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(consts.VkCapital) {
		result = result | consts.EVENTFLAG_CAPS_LOCK_ON
	}
	switch aWparam {
	case consts.VkReturn:
		if ((aLparam >> 16) & consts.KF_EXTENDED) != 0 {
			result = result | consts.EVENTFLAG_IS_KEY_PAD
		}
	case consts.VkInsert, consts.VkDelete, consts.VkHome, consts.VkEnd, consts.VkPrior, consts.VkNext, consts.VkUp, consts.VkDown, consts.VkLeft, consts.VkRight:
		if ((aLparam >> 16) & winapi.KF_EXTENDED) == 0 {
			result = result | consts.EVENTFLAG_IS_KEY_PAD
		}
	case consts.VkNumLock, consts.VkNumpad0, consts.VkNumpad1, consts.VkNumpad2, consts.VkNumpad3, consts.VkNumpad4, consts.VkNumpad5, consts.VkNumpad6, consts.VkNumpad7,
		consts.VkNumpad8, consts.VkNumpad9, consts.VkDivide, consts.VkMultiply, consts.VkSubtract, consts.VkAdd, consts.VkDecimal, consts.VkClear:
		result = result | consts.EVENTFLAG_IS_KEY_PAD
	case consts.VkShift:
		if CefIsKeyDown(consts.VkLShift) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(consts.VkRShift) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case consts.VkControl:
		if CefIsKeyDown(consts.VkLControl) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(consts.VkRControl) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case consts.VkMenu:
		if CefIsKeyDown(consts.VkLMenu) {
			result = result | consts.EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(consts.VkRMenu) {
			result = result | consts.EVENTFLAG_IS_RIGHT
		}
	case consts.VkLWin:
		result = result | consts.EVENTFLAG_IS_LEFT
	case consts.VkRWin:
		result = result | consts.EVENTFLAG_IS_RIGHT
	}
	return
}
