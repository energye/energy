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

import "unsafe"

// 该对象不做实现
type ICefV8Context struct {
	instance uintptr
	ptr      unsafe.Pointer
	Browser  *ICefBrowser
	Frame    *ICefFrame
	Global   *ICEFv8Value
}
