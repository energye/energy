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

// ICefPermissionHandler
type ICefPermissionHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// PermissionHandlerRef -> ICefPermissionHandler
var PermissionHandlerRef permissionHandler

type permissionHandler uintptr

func (*permissionHandler) New() *ICefPermissionHandler {
	var result uintptr
	imports.Proc(def.CefPermissionHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPermissionHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefPermissionHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPermissionHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPermissionHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPermissionHandler) OnRequestMediaAccessPermission(fn onRequestMediaAccessPermission) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPermissionHandler_OnRequestMediaAccessPermission).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPermissionHandler) OnShowPermissionPrompt(fn onShowPermissionPrompt) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPermissionHandler_OnShowPermissionPrompt).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPermissionHandler) OnDismissPermissionPrompt(fn onDismissPermissionPrompt) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPermissionHandler_OnDismissPermissionPrompt).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onRequestMediaAccessPermission func(browser *ICefBrowser, frame *ICefFrame, requestingOrigin string, requestedPermissions uint32, callback *ICefMediaAccessCallback) bool
type onShowPermissionPrompt func(browser *ICefBrowser, promptId uint64, requestingOrigin string, requestedPermissions uint32, callback *ICefPermissionPromptCallback) bool
type onDismissPermissionPrompt func(browser *ICefBrowser, promptId uint64, result consts.TCefPermissionRequestResult)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onRequestMediaAccessPermission:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			requestingOrigin := api.GoStr(getVal(2))
			requestedPermissions := uint32(getVal(3))
			callback := &ICefMediaAccessCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(onRequestMediaAccessPermission)(browse, frame, requestingOrigin, requestedPermissions, callback)
		case onShowPermissionPrompt:
			browse := &ICefBrowser{instance: getPtr(0)}
			promptId := *(*uint64)(getPtr(1))
			requestingOrigin := api.GoStr(getVal(2))
			requestedPermissions := uint32(getVal(3))
			callback := &ICefPermissionPromptCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(onShowPermissionPrompt)(browse, promptId, requestingOrigin, requestedPermissions, callback)
		case onDismissPermissionPrompt:
			browse := &ICefBrowser{instance: getPtr(0)}
			promptId := *(*uint64)(getPtr(1))
			result := consts.TCefPermissionRequestResult(getVal(2))
			fn.(onDismissPermissionPrompt)(browse, promptId, result)
		default:
			return false
		}
		return true
	})
}
