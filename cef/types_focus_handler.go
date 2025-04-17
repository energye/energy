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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefFocusHandler
type ICefFocusHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// FocusHandlerRef -> ICefFocusHandler
var FocusHandlerRef focusHandler

type focusHandler uintptr

func (*focusHandler) New() *ICefFocusHandler {
	var result uintptr
	imports.Proc(def.CefFocusHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFocusHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefFocusHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefFocusHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefFocusHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefFocusHandler) SetOnTakeFocus(fn onTakeFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFocusHandler_OnTakeFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefFocusHandler) SetOnSetFocus(fn onSetFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFocusHandler_OnSetFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefFocusHandler) SetOnGotFocus(fn onGotFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFocusHandler_OnGotFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onTakeFocus func(browser *ICefBrowser, next bool)
type onSetFocus func(browser *ICefBrowser, source consts.TCefFocusSource) bool
type onGotFocus func(browser *ICefBrowser)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onTakeFocus:
			browse := &ICefBrowser{instance: getPtr(0)}
			next := api.GoBool(getVal(1))
			fn.(onTakeFocus)(browse, next)
		case onSetFocus:
			browse := &ICefBrowser{instance: getPtr(0)}
			source := consts.TCefFocusSource(getVal(1))
			result := (*bool)(getPtr(2))
			*result = fn.(onSetFocus)(browse, source)
		case onGotFocus:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onGotFocus)(browse)
		default:
			return false
		}
		return true
	})
}
