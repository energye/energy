//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF v8 value V8所有类型
//
// ICefV8Value 创建和使用一搬在 v8context 上下文中使用

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"unsafe"
)

// ICefCustomUserData
type ICefCustomUserData struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

func NewCustomUserData(userDataType, userData uintptr) *ICefCustomUserData {
	var result uintptr
	imports.Proc(def.CustomUserData_Create).Call(userDataType, userData, uintptr(unsafe.Pointer(&result)))
	return &ICefCustomUserData{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefCustomUserData) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCustomUserData) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefCustomUserData) UserDataType() uintptr {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CustomUserData_UserDataType).Call(m.Instance())
	return r1
}

func (m *ICefCustomUserData) UserData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CustomUserData_UserData).Call(m.Instance())
	return r1
}

func (m *ICefCustomUserData) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
