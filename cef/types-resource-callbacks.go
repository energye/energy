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
	"github.com/energye/energy/common/imports"
	"unsafe"
)

// ************************** ICefResourceReadCallback ************************** //

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
	imports.Proc(internale_ResourceSkipCallback_Cont).Call(m.Instance(), uintptr(unsafe.Pointer(&bytesSkipped)))
}

// ************************** ICefResourceSkipCallback ************************** //

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
	imports.Proc(internale_ResourceReadCallback_Cont).Call(m.Instance(), uintptr(unsafe.Pointer(&bytesRead)))
}
