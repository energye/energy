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
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

var RequestContextRef requestContext

type requestContext uintptr

func (*requestContext) Global() *ICefRequestContext {
	var result uintptr
	imports.Proc(def.RequestContextRef_Global).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

func (*requestContext) New(requestContextSettings *TCefRequestContextSettings, handler *ICefRequestContextHandler) *ICefRequestContext {
	if requestContextSettings == nil {
		requestContextSettings = &TCefRequestContextSettings{}
	}
	requestContextSettingsPtr := requestContextSettings.ToPtr()
	var result uintptr
	imports.Proc(def.RequestContextRef_New).Call(uintptr(unsafe.Pointer(requestContextSettingsPtr)), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

func (*requestContext) NewTwo(cache, acceptLanguageList, cookieableSchemesList string, cookieableSchemesExcludeDefaults, persistSessionCookies, persistUserPreferences bool, handler *ICefRequestContextHandler) *ICefRequestContext {
	var result uintptr
	imports.Proc(def.RequestContextRef_NewTwo).Call(api.PascalStr(cache), api.PascalStr(acceptLanguageList), api.PascalStr(cookieableSchemesList), api.PascalBool(cookieableSchemesExcludeDefaults), api.PascalBool(persistSessionCookies), api.PascalBool(persistUserPreferences), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

func (*requestContext) Shared(other *ICefRequestContext, handler *ICefRequestContextHandler) *ICefRequestContext {
	var result uintptr
	imports.Proc(def.RequestContextRef_Shared).Call(other.Instance(), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

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
	r1, _, _ := imports.Proc(def.RequestContext_HasPreference).Call(m.Instance(), api.PascalStr(name))
	return api.GoBool(r1)
}

func (m *ICefRequestContext) GetPreference(name string) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContext_GetPreference).Call(m.Instance(), api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
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
	imports.Proc(def.RequestContext_GetAllPreferences).Call(m.Instance(), api.PascalBool(includeDefaults), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDictionaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) CanSetPreference(name string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.RequestContext_CanSetPreference).Call(m.Instance(), api.PascalStr(name))
	return api.GoBool(r1)
}

func (m *ICefRequestContext) SetPreference(name string, value *ICefValue) (error string, ok bool) {
	if !m.IsValid() {
		return "", false
	}
	var errorResult uintptr
	r1, _, _ := imports.Proc(def.RequestContext_SetPreference).Call(m.Instance(), api.PascalStr(name), value.Instance(), uintptr(unsafe.Pointer(&errorResult)))
	return api.GoStr(errorResult), api.GoBool(r1)
}

func (m *ICefRequestContext) IsSame(other *ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.RequestContext_IsSame).Call(m.Instance(), other.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) IsSharingWith(other *ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.RequestContext_IsSharingWith).Call(m.Instance(), other.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) IsGlobal() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.RequestContext_IsGlobal).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) GetHandler() *ICefRequestContextHandler {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContext_GetHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContextHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) GetCachePath() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.RequestContext_GetCachePath).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefRequestContext) GetCookieManager(callback *ICefCompletionCallback) *ICefCookieManager {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContext_GetCookieManager).Call(m.Instance(), callback.Instance(), uintptr(unsafe.Pointer(&result)))
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
	r1, _, _ := imports.Proc(def.RequestContext_RegisterSchemeHandlerFactory).Call(m.Instance(), api.PascalStr(schemeName), api.PascalStr(domainName), factory.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) ClearSchemeHandlerFactories() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.RequestContext_ClearSchemeHandlerFactories).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequestContext) ClearCertificateExceptions(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_ClearCertificateExceptions).Call(m.Instance(), callback.Instance())
}

func (m *ICefRequestContext) ClearHttpAuthCredentials(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_ClearHttpAuthCredentials).Call(m.Instance(), callback.Instance())
}

func (m *ICefRequestContext) CloseAllConnections(callback *ICefCompletionCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_CloseAllConnections).Call(m.Instance(), callback.Instance())
}

// ResolveHost TODO ICefResolveCallback
func (m *ICefRequestContext) ResolveHost(origin string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_ResolveHost).Call(m.Instance(), api.PascalStr(origin), uintptr(0))
}

func (m *ICefRequestContext) LoadExtension(rootDirectory string, manifest *ICefDictionaryValue, handler *ICefExtensionHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_LoadExtension).Call(m.Instance(), api.PascalStr(rootDirectory), manifest.Instance(), handler.Instance())
}

func (m *ICefRequestContext) DidLoadExtension(extensionId string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_DidLoadExtension).Call(m.Instance(), api.PascalStr(extensionId))
}

func (m *ICefRequestContext) HasExtension(extensionId string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_HasExtension).Call(m.Instance(), api.PascalStr(extensionId))
}

func (m *ICefRequestContext) GetExtensions() (result []string, ok bool) {
	if !m.IsValid() {
		return nil, false
	}
	extensionIds := lcl.NewStringList()
	defer extensionIds.Free()
	r1, _, _ := imports.Proc(def.RequestContext_GetExtensions).Call(m.Instance(), extensionIds.Instance())
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
	imports.Proc(def.RequestContext_GetExtension).Call(m.Instance(), api.PascalStr(extensionId), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefExtension{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) GetMediaRouter(callback *ICefCompletionCallback) *ICefMediaRouter {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContext_GetMediaRouter).Call(m.Instance(), callback.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRouter{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) GetWebsiteSetting(requestingUrl, topLevelUrl string, contentType consts.TCefContentSettingTypes) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContext_WebsiteSetting).Call(consts.GetValue, m.Instance(), api.PascalStr(requestingUrl),
		api.PascalStr(topLevelUrl), uintptr(contentType), 0, uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequestContext) SetWebsiteSetting(requestingUrl, topLevelUrl string, contentType consts.TCefContentSettingTypes, value *ICefValue) {
	if !m.IsValid() || !value.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_WebsiteSetting).Call(consts.SetValue, m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType), value.Instance(), 0)
}

func (m *ICefRequestContext) GetContentSetting(requestingUrl, topLevelUrl string, contentType consts.TCefContentSettingTypes) consts.TCefContentSettingValues {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.RequestContext_ContentSetting).Call(consts.GetValue, m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType), 0)
	return consts.TCefContentSettingValues(r1)
}

func (m *ICefRequestContext) SetContentSetting(requestingUrl, topLevelUrl string, contentType consts.TCefContentSettingTypes, value consts.TCefContentSettingValues) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RequestContext_ContentSetting).Call(consts.SetValue, m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType), uintptr(value))
}
