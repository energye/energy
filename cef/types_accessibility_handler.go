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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefAccessibilityHandler
//
//	/include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t)
type ICefAccessibilityHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// AccessibilityHandlerRef -> ICefAccessibilityHandler
var AccessibilityHandlerRef accessibilityHandler

type accessibilityHandler uintptr

func (*accessibilityHandler) New() *ICefAccessibilityHandler {
	var result uintptr
	imports.Proc(def.CefAccessibilityHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefAccessibilityHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefAccessibilityHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefAccessibilityHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefAccessibilityHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefAccessibilityHandler) SetOnAccessibilityTreeChange(fn onAccessibilityTreeChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAccessibilityHandler_OnAccessibilityTreeChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefAccessibilityHandler) SetOnAccessibilityLocationChange(fn onAccessibilityLocationChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAccessibilityHandler_OnAccessibilityLocationChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onAccessibilityTreeChange func(value *ICefValue)
type onAccessibilityLocationChange func(value *ICefValue)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onAccessibilityTreeChange:
			fn.(onAccessibilityTreeChange)(&ICefValue{instance: getInstance(getVal(0))})
		case onAccessibilityLocationChange:
			fn.(onAccessibilityLocationChange)(&ICefValue{instance: getInstance(getVal(0))})
		default:
			return false
		}
		return true
	})
}
