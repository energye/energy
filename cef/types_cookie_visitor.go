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
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefCookieVisitor
type ICefCookieVisitor struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// CookieVisitorRef -> ICefCookieVisitor
var CookieVisitorRef cookieVisitor

type cookieVisitor uintptr

func (*cookieVisitor) New() *ICefCookieVisitor {
	var result uintptr
	imports.Proc(def.CefCookieVisitorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCookieVisitor{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefCookieVisitor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCookieVisitor) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCookieVisitor) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefCookieVisitor) SetOnVisit(fn cookieOnVisit) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefCookieVisitor_OnVisit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type cookieOnVisit func(cookie *ICefCookie) (deleteCookie, result bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case cookieOnVisit:
			cookie := *(*iCefCookiePtr)(getInstance(getVal(0)))
			creation := *(*float64)(common.GetParamPtr(cookie.creation, 0))
			lastAccess := *(*float64)(common.GetParamPtr(cookie.lastAccess, 0))
			expires := *(*float64)(common.GetParamPtr(cookie.expires, 0))
			deleteCookiePtr := (*bool)(common.GetParamPtr(cookie.aDeleteCookie, 0))
			var deleteCookie bool
			iCookie := &ICefCookie{
				Url:          api.GoStr(cookie.url),
				Name:         api.GoStr(cookie.name),
				Value:        api.GoStr(cookie.value),
				Domain:       api.GoStr(cookie.domain),
				Path:         api.GoStr(cookie.path),
				Secure:       *(*bool)(common.GetParamPtr(cookie.secure, 0)),
				Httponly:     *(*bool)(common.GetParamPtr(cookie.httponly, 0)),
				HasExpires:   *(*bool)(common.GetParamPtr(cookie.hasExpires, 0)),
				Creation:     common.DDateTimeToGoDateTime(creation),
				LastAccess:   common.DDateTimeToGoDateTime(lastAccess),
				Expires:      common.DDateTimeToGoDateTime(expires),
				Count:        int32(cookie.count),
				Total:        int32(cookie.total),
				SameSite:     consts.TCefCookieSameSite(cookie.sameSite),
				Priority:     consts.TCefCookiePriority(cookie.priority),
				DeleteCookie: *(*bool)(common.GetParamPtr(cookie.aDeleteCookie, 0)),
			}
			result := (*bool)(getInstance(getVal(1)))
			deleteCookie, *result = fn.(cookieOnVisit)(iCookie)
			*deleteCookiePtr = deleteCookie
		default:
			return false
		}
		return true
	})
}
