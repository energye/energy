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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefRequestContext
type ICefRequestContext struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// RequestContextRef -> ICefRequestContext
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

// Creates a new context object with the specified |settings| and optional |handler|.
// <param name="settings">Pointer to TCefRequestContextSettings.</param>
// <param name="handler">Optional handler for the request context.</param>
func (*requestContext) New(requestContextSettings TCefRequestContextSettings, handler *ICefRequestContextHandler) *ICefRequestContext {
	requestContextSettingsPtr := requestContextSettings.ToPtr()
	var result uintptr
	imports.Proc(def.RequestContextRef_New).Call(uintptr(unsafe.Pointer(requestContextSettingsPtr)), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

// Creates a new context object with the specified settings and optional |handler|.
//
// <param name="aCache">The directory where cache data for this request context will be stored on disk. See TCefRequestContextSettings.cache_path for more information.</param>
// <param name="aAcceptLanguageList">Comma delimited ordered list of language codes without any whitespace that will be used in the "Accept-Language" HTTP header. See TCefRequestContextSettings.accept_language_list for more information.</param>
// <param name="aCookieableSchemesList">Comma delimited list of schemes supported by the associated ICefCookieManager. See TCefRequestContextSettings.cookieable_schemes_list for more information.</param>
// <param name="aCookieableSchemesExcludeDefaults">Setting this parameter to true will disable all loading and saving of cookies. See TCefRequestContextSettings.cookieable_schemes_list for more information.</param>
// <param name="aPersistSessionCookies">To persist session cookies (cookies without an expiry date or validity interval) by default when using the global cookie manager set this value to true. See TCefRequestContextSettings.persist_session_cookies for more information.</param>
// <param name="handler">Optional handler for the request context.</param>
func (*requestContext) NewTwo(cache, acceptLanguageList, cookieableSchemesList string, cookieableSchemesExcludeDefaults, persistSessionCookies bool,
	handler *ICefRequestContextHandler) *ICefRequestContext {
	var result uintptr
	imports.Proc(def.RequestContextRef_NewTwo).Call(api.PascalStr(cache), api.PascalStr(acceptLanguageList), api.PascalStr(cookieableSchemesList), api.PascalBool(cookieableSchemesExcludeDefaults), api.PascalBool(persistSessionCookies), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: getInstance(result)}
	}
	return nil
}

// Creates a new context object that shares storage with |other| and uses an optional |handler|.
// <param name="other">Another ICefRequestContext instance that will share storage with the new ICefRequestContext instance.</param>
// <param name="handler">Optional handler for the request context.</param>
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
