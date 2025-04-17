//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 回调事件实现

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"unsafe"
)

// ICefCallback
type ICefCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Cont 继续执行
func (m *ICefCallback) Cont() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFCallback_Cont).Call(uintptr(m.instance))
}

// Cancel 取消执行
func (m *ICefCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFCallback_Cancel).Call(uintptr(m.instance))
}
