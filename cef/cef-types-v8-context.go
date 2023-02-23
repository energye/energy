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

// Instance 实例
func (m *ICefV8Context) Instance() uintptr {
	return uintptr(m.instance)
}
