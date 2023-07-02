//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 右键菜单

package cef

import (
	"unsafe"
)

// MenuModelDelegateRef -> ICefMenuModelDelegate
var MenuModelDelegateRef menuModelDelegate

type menuModelDelegate uintptr

func (*menuModelDelegate) New() *ICefMenuModelDelegate {
	var result uintptr
	//imports.Proc(def.MenuModelRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefMenuModelDelegate{instance: unsafe.Pointer(result)}
}

// Instance 实例
func (m *ICefMenuModelDelegate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMenuModelDelegate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMenuModelDelegate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
