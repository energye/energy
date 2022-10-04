//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import "unsafe"

//该对象不做实现
type ICefV8Context struct {
	instance uintptr
	ptr      unsafe.Pointer
	Browser  *ICefBrowser
	Frame    *ICefFrame
	Global   *ICEFv8Value
}
