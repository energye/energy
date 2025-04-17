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

// ICefResourceReadCallback
//
//	/include/capi/cef_resource_handler_capi.h (cef_resource_read_callback_t)
type ICefResourceReadCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefResourceReadCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResourceReadCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefResourceReadCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefResourceReadCallback) Cont(bytesSkipped int64) {
	if m == nil || m.instance == nil {
		return
	}
	imports.Proc(def.ResourceSkipCallback_Cont).Call(m.Instance(), uintptr(unsafe.Pointer(&bytesSkipped)))
}

// ICefResourceSkipCallback
//
//	/include/capi/cef_resource_handler_capi.h (cef_resource_skip_callback_t)
type ICefResourceSkipCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefResourceSkipCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResourceSkipCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefResourceSkipCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefResourceSkipCallback) Cont(bytesRead int64) {
	if m == nil || m.instance == nil {
		return
	}
	imports.Proc(def.ResourceReadCallback_Cont).Call(m.Instance(), uintptr(unsafe.Pointer(&bytesRead)))
}
