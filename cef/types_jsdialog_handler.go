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

// ICefJsDialogHandler
type ICefJsDialogHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// JsDialogHandlerRef -> ICefJsDialogHandler
var JsDialogHandlerRef jsDialogHandler

type jsDialogHandler uintptr

func (*jsDialogHandler) New() *ICefJsDialogHandler {
	var result uintptr
	imports.Proc(def.CefJsDialogHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefJsDialogHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefJsDialogHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefJsDialogHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefJsDialogHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefJsDialogHandler) OnJsDialog(fn onJsDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefJsDialogHandler_OnJsdialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefJsDialogHandler) OnBeforeUnloadDialog(fn onBeforeUnloadDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefJsDialogHandler_OnBeforeUnloadDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefJsDialogHandler) OnResetDialogState(fn onResetDialogState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefJsDialogHandler_OnResetDialogState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefJsDialogHandler) OnDialogClosed(fn onDialogClosed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefJsDialogHandler_OnDialogClosed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onJsDialog func(browser *ICefBrowser, originUrl string, dialogType consts.TCefJsDialogType, messageText, defaultPromptText string, callback *ICefJsDialogCallback) (suppressMessage bool, result bool)
type onBeforeUnloadDialog func(browser *ICefBrowser, messageText string, isReload bool, callback *ICefJsDialogCallback) bool
type onResetDialogState func(browser *ICefBrowser)
type onDialogClosed func(browser *ICefBrowser)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onJsDialog:
			browse := &ICefBrowser{instance: getPtr(0)}
			originUrl := api.GoStr(getVal(1))
			dialogType := consts.TCefJsDialogType(getVal(2))
			messageText, defaultPromptText := api.GoStr(getVal(3)), api.GoStr(getVal(4))
			callback := &ICefJsDialogCallback{instance: getPtr(5)}
			suppressMessage := (*bool)(getPtr(6))
			result := (*bool)(getPtr(7))
			*suppressMessage, *result = fn.(onJsDialog)(browse, originUrl, dialogType, messageText, defaultPromptText, callback)
		case onBeforeUnloadDialog:
			browse := &ICefBrowser{instance: getPtr(0)}
			messageText := api.GoStr(getVal(1))
			isReload := api.GoBool(getVal(2))
			callback := &ICefJsDialogCallback{instance: getPtr(3)}
			result := (*bool)(getPtr(4))
			*result = fn.(onBeforeUnloadDialog)(browse, messageText, isReload, callback)
		case onResetDialogState:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onResetDialogState)(browse)
		case onDialogClosed:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onDialogClosed)(browse)
		default:
			return false
		}
		return true
	})
}
