//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefMenuButtonPressedLock Parent: ICefBaseRefCounted
//
//	MenuButton pressed lock is released when this object is destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_pressed_lock_t)</a>
type ICefMenuButtonPressedLock interface {
	ICefBaseRefCounted
}

// TCefMenuButtonPressedLock Parent: TCefBaseRefCounted
//
//	MenuButton pressed lock is released when this object is destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_pressed_lock_t)</a>
type TCefMenuButtonPressedLock struct {
	TCefBaseRefCounted
}

// MenuButtonPressedLockRef -> ICefMenuButtonPressedLock
var MenuButtonPressedLockRef menuButtonPressedLock

// menuButtonPressedLock TCefMenuButtonPressedLock Ref
type menuButtonPressedLock uintptr

// UnWrap
//
//	Returns a ICefMenuButtonPressedLock instance using a PCefMenuButtonPressedLock data pointer.
func (m *menuButtonPressedLock) UnWrap(data uintptr) ICefMenuButtonPressedLock {
	var resultCefMenuButtonPressedLock uintptr
	CEF().SysCallN(1078, uintptr(data), uintptr(unsafePointer(&resultCefMenuButtonPressedLock)))
	return AsCefMenuButtonPressedLock(resultCefMenuButtonPressedLock)
}
