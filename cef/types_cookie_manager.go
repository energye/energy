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
	"time"
	"unsafe"
)

// ICefCookieManager
type ICefCookieManager struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefCookieManager) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCookieManager) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCookieManager) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefCookieManager) VisitAllCookies(visitor *ICefCookieVisitor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefCookieManager_VisitAllCookies).Call(m.Instance(), visitor.Instance())
}

func (m *ICefCookieManager) VisitUrlCookies(url string, includeHttpOnly bool, visitor *ICefCookieVisitor) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefCookieManager_VisitUrlCookies).Call(m.Instance(), api.PascalStr(url), api.PascalBool(includeHttpOnly), visitor.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool, creation, lastAccess, expires time.Time,
	sameSite consts.TCefCookieSameSite, priority consts.TCefCookiePriority, callback *ICefSetCookieCallback) bool {
	if !m.IsValid() {
		return false
	}
	cookie := &TCefCookie{
		Url:        url,
		Name:       name,
		Value:      value,
		Domain:     domain,
		Path:       path,
		Secure:     secure,
		Httponly:   httponly,
		HasExpires: hasExpires,
		Creation:   creation,
		LastAccess: lastAccess,
		Expires:    expires,
		SameSite:   sameSite,
		Priority:   priority,
	}
	cookiePtr := cookie.ToPtr()
	r1, _, _ := imports.Proc(def.CefCookieManager_SetCookie).Call(m.Instance(), uintptr(unsafe.Pointer(cookiePtr)), callback.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) DeleteCookies(url, cookieName string, callback *ICefDeleteCookiesCallback) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefCookieManager_DeleteCookies).Call(m.Instance(), api.PascalStr(url), api.PascalStr(cookieName), callback.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) FlushStore(callback *ICefCompletionCallback) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefCookieManager_FlushStore).Call(m.Instance(), callback.Instance())
	return api.GoBool(r1)
}
