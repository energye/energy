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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"unsafe"
)

type ICefBaseRefCounted interface {
	Instance() uintptr
	IsValid() bool
}

func NewBaseRefCounted(instance uintptr) *TCefBaseRefCounted {
	return &TCefBaseRefCounted{instance: unsafe.Pointer(instance)}
}

// Wrap 指针引用包裹
func (m *TCefBaseRefCounted) Wrap(data uintptr) unsafe.Pointer {
	var result uintptr
	imports.Proc(def.CefBaseRefCounted_Wrap).Call(data, uintptr(unsafe.Pointer(&result)))
	return unsafe.Pointer(result)
}

// Free 释放底层指针
func (m *TCefBaseRefCounted) Free(data uintptr) {
	imports.Proc(def.CefBaseRefCounted_Free).Call(uintptr(unsafe.Pointer(&data)))
	m.instance = nil
}

// Instance 实例
func (m *TCefBaseRefCounted) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefBaseRefCounted) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
