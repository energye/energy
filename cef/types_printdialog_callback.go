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

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"unsafe"
)

// ICefPrintDialogCallback
type ICefPrintDialogCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefPrintDialogCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPrintDialogCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPrintDialogCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPrintDialogCallback) Cont(settings *ICefPrintSettings) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintDialogCallback_Cont).Call(m.Instance(), settings.Instance())
}

func (m *ICefPrintDialogCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintDialogCallback_Cancel).Call(m.Instance())
}
