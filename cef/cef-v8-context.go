//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF V8 上下文
package cef

import "unsafe"

// ICefV8Context
// 渲染进程创建时对V8Value创建
type ICefV8Context struct {
	instance uintptr
	ptr      unsafe.Pointer
	Browser  *ICefBrowser
	Frame    *ICefFrame
	Global   *V8Value
}
