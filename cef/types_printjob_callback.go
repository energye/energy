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

// ICefPrintJobCallback
type ICefPrintJobCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefPrintJobCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPrintJobCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPrintJobCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPrintJobCallback) Cont() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintJobCallback_Cont).Call(m.Instance())
}
