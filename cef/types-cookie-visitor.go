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
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// CookieVisitorRef -> ICefCookieVisitor
var CookieVisitorRef cookieVisitor

type cookieVisitor uintptr

func (*cookieVisitor) New() *ICefCookieVisitor {
	var result uintptr
	imports.Proc(internale_CefCookieVisitorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
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
	imports.Proc(internale_CefCookieVisitor_OnVisit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type cookieOnVisit func(cookie *ICefCookie) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case cookieOnVisit:
			cookie := *(*iCefCookiePtr)(getInstance(getVal(0)))
			creation := *(*float64)(common.GetParamPtr(cookie.creation, 0))
			lastAccess := *(*float64)(common.GetParamPtr(cookie.lastAccess, 0))
			expires := *(*float64)(common.GetParamPtr(cookie.expires, 0))
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
			*result = fn.(cookieOnVisit)(iCookie)
		default:
			return false
		}
		return true
	})
}
