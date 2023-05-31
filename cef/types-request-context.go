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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefRequestContext) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRequestContext) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRequestContext) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRequestContext) HasPreference(name string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_HasPreference).Call(m.Instance(), api.PascalStr(name))
	return api.GoBool(r1)
}

func (m *ICefRequestContext) GetPreference(name string) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_RequestContext_GetPreference).Call(m.Instance(), api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) GetAllPreferences(includeDefaults bool) *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_RequestContext_GetAllPreferences).Call(m.Instance(), api.PascalBool(includeDefaults), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDictionaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) CanSetPreference(name string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_CanSetPreference).Call(m.Instance(), api.PascalStr(name))
	return api.GoBool(r1)
}

func (m *ICefRequestContext) SetPreference(name string, value *ICefValue) (error string, ok bool) {
	if !m.IsValid() {
		return "", false
	}
	var errorResult uintptr
	r1, _, _ := imports.Proc(internale_RequestContext_SetPreference).Call(m.Instance(), api.PascalStr(name), value.Instance(), uintptr(unsafe.Pointer(&errorResult)))
	return api.GoStr(errorResult), api.GoBool(r1)
}

func (m *ICefRequestContext) IsSame(other *ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_IsSame).Call(m.Instance(), other.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) IsSharingWith(other *ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_IsSharingWith).Call(m.Instance(), other.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) IsGlobal() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_IsGlobal).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) GetHandler() *ICefRequestContextHandler {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_RequestContext_GetHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContextHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) GetCachePath() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_RequestContext_GetCachePath).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefRequestContext) GetCookieManager(callback *ICefCompletionCallback) *ICefCookieManager {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_RequestContext_GetCookieManager).Call(m.Instance(), callback.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCookieManager{instance: unsafe.Pointer(result)}
	}
	return nil
}

// RegisterSchemeHandlerFactory
func (m *ICefRequestContext) RegisterSchemeHandlerFactory(schemeName, domainName string, factory *ICefSchemeHandlerFactory) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_RegisterSchemeHandlerFactory).Call(m.Instance(), api.PascalStr(schemeName), api.PascalStr(domainName), factory.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) ClearSchemeHandlerFactories() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_RequestContext_ClearSchemeHandlerFactories).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) ClearCertificateExceptions(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_ClearCertificateExceptions).Call(m.Instance(), callback.Instance())
}

func (m *ICefRequestContext) ClearHttpAuthCredentials(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_ClearHttpAuthCredentials).Call(m.Instance(), callback.Instance())
}

func (m *ICefRequestContext) CloseAllConnections(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_CloseAllConnections).Call(m.Instance(), callback.Instance())
}

// ResolveHost TODO ICefResolveCallback
func (m *ICefRequestContext) ResolveHost(origin string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_ResolveHost).Call(m.Instance(), api.PascalStr(origin), uintptr(0))
}

func (m *ICefRequestContext) LoadExtension(rootDirectory string, manifest *ICefDictionaryValue, handler *ICefExtensionHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_LoadExtension).Call(m.Instance(), api.PascalStr(rootDirectory), manifest.Instance(), handler.Instance())
}

func (m *ICefRequestContext) DidLoadExtension(extensionId string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_DidLoadExtension).Call(m.Instance(), api.PascalStr(extensionId))
}

func (m *ICefRequestContext) HasExtension(extensionId string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RequestContext_HasExtension).Call(m.Instance(), api.PascalStr(extensionId))
}

func (m *ICefRequestContext) GetExtensions() (result []string, ok bool) {
	if !m.IsValid() {
		return nil, false
	}
	extensionIds := lcl.NewStringList()
	defer extensionIds.Free()
	r1, _, _ := imports.Proc(internale_RequestContext_GetExtensions).Call(m.Instance(), extensionIds.Instance())
	count := extensionIds.Count()
	if api.GoBool(r1) && count > 0 {
		result = make([]string, count, count)
		for i := 0; i < int(count); i++ {
			result[i] = extensionIds.Strings(int32(i))
		}
		return result, true
	}
	return nil, false
}

func (m *ICefRequestContext) GetExtension(extensionId string) *ICefExtension {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_RequestContext_GetExtension).Call(m.Instance(), api.PascalStr(extensionId), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefExtension{instance: unsafe.Pointer(result)}
	}
	return nil
}
