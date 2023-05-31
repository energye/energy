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
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

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
	imports.Proc(internale_CefCookieManager_VisitAllCookies).Call(m.Instance(), visitor.Instance())
}

func (m *ICefCookieManager) VisitUrlCookies(url string, includeHttpOnly bool, visitor *ICefCookieVisitor) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefCookieManager_VisitUrlCookies).Call(m.Instance(), api.PascalStr(url), api.PascalBool(includeHttpOnly), visitor.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool, creation, lastAccess, expires time.Time,
	sameSite consts.TCefCookieSameSite, priority consts.TCefCookiePriority, callback *ICefSetCookieCallback) bool {
	if !m.IsValid() {
		return false
	}
	creationPtr := common.GoDateTimeToDDateTime(creation)
	lastAccessPtr := common.GoDateTimeToDDateTime(lastAccess)
	expiresPtr := common.GoDateTimeToDDateTime(expires)
	cCookie := &iCefCookiePtr{
		url:             api.PascalStr(url),
		name:            api.PascalStr(name),
		value:           api.PascalStr(value),
		domain:          api.PascalStr(domain),
		path:            api.PascalStr(path),
		secure:          api.PascalBool(secure),
		httponly:        api.PascalBool(httponly),
		hasExpires:      api.PascalBool(hasExpires),
		creation:        uintptr(unsafe.Pointer(&creationPtr)),
		lastAccess:      uintptr(unsafe.Pointer(&lastAccessPtr)),
		expires:         uintptr(unsafe.Pointer(&expiresPtr)),
		sameSite:        uintptr(sameSite),
		priority:        uintptr(priority),
		aSetImmediately: uintptr(0),
		aID:             uintptr(0),
		aDeleteCookie:   uintptr(0),
		aResult:         uintptr(0),
		count:           uintptr(0),
		total:           uintptr(0),
	}
	r1, _, _ := imports.Proc(internale_CefCookieManager_SetCookie).Call(m.Instance(), uintptr(unsafe.Pointer(cCookie)), callback.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) DeleteCookies(url, cookieName string, callback *ICefDeleteCookiesCallback) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefCookieManager_DeleteCookies).Call(m.Instance(), api.PascalStr(url), api.PascalStr(cookieName), callback.Instance())
	return api.GoBool(r1)
}

func (m *ICefCookieManager) FlushStore(callback *ICefCompletionCallback) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefCookieManager_FlushStore).Call(m.Instance(), callback.Instance())
	return api.GoBool(r1)
}
