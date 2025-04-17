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

// ICefKeyboardHandler
type ICefKeyboardHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// KeyboardHandlerRef -> ICefKeyboardHandler
var KeyboardHandlerRef keyboardHandler

type keyboardHandler uintptr

func (*keyboardHandler) New() *ICefKeyboardHandler {
	var result uintptr
	imports.Proc(def.CefKeyboardHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefKeyboardHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefKeyboardHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefKeyboardHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefKeyboardHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefKeyboardHandler) SetOnPreKeyEvent(fn onPreKeyEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefKeyboardHandler_OnPreKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefKeyboardHandler) SetOnKeyEvent(fn onKeyEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefKeyboardHandler_OnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onPreKeyEvent func(browser *ICefBrowser, event *TCefKeyEvent, osEvent consts.TCefEventHandle) (isKeyboardShortcut, result bool)
type onKeyEvent func(browser *ICefBrowser, event *TCefKeyEvent, osEvent consts.TCefEventHandle) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onPreKeyEvent:
			browse := &ICefBrowser{instance: getPtr(0)}
			event := (*TCefKeyEvent)(getPtr(1))
			osEvent := consts.EventHandle(getVal(2))
			isKeyboardShortcut := (*bool)(getPtr(3))
			result := (*bool)(getPtr(4))
			*isKeyboardShortcut, *result = fn.(onPreKeyEvent)(browse, event, osEvent)
		case onKeyEvent:
			browse := &ICefBrowser{instance: getPtr(0)}
			event := (*TCefKeyEvent)(getPtr(1))
			osEvent := consts.EventHandle(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(onKeyEvent)(browse, event, osEvent)
		default:
			return false
		}
		return true
	})
}
