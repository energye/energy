//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef view
package cef

// Instance 实例
func (m *ICefView) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}
