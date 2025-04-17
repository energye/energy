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

// ICefCookieAccessFilter
//
//	/include/capi/cef_resource_request_handler_capi.h (cef_cookie_access_filter_t)
type ICefCookieAccessFilter struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// CookieAccessFilterRef -> ICefCookieAccessFilter
var CookieAccessFilterRef cookieAccessFilter

type cookieAccessFilter uintptr

func (*cookieAccessFilter) New() *ICefCookieAccessFilter {
	var result uintptr
	imports.Proc(def.CookieAccessFilterRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCookieAccessFilter{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*cookieAccessFilter) NewForChromium(chromium IChromium) *ICefCookieAccessFilter {
	var result uintptr
	imports.Proc(def.CookieAccessFilterRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCookieAccessFilter{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

// Instance 实例
func (m *ICefCookieAccessFilter) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCookieAccessFilter) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCookieAccessFilter) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefCookieAccessFilter) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefCookieAccessFilter) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

func (m *ICefCookieAccessFilter) SetOnCanSendCookie(fn cookieAccessOnCanSendCookie) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.CookieAccessFilter_CanSendCookie).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefCookieAccessFilter) SetOnCanSaveCookie(fn cookieAccessOnCanSaveCookie) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.CookieAccessFilter_CanSaveCookie).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type cookieAccessOnCanSendCookie func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, cookie *TCefCookie) bool
type cookieAccessOnCanSaveCookie func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, cookie *TCefCookie) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case cookieAccessOnCanSendCookie:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			cookiePtr := (*tCefCookiePtr)(getInstance(getVal(3)))
			cookie := cookiePtr.convert()
			result := (*bool)(getInstance(getVal(4)))
			*result = fn.(cookieAccessOnCanSendCookie)(browse, frame, request, cookie)
		case cookieAccessOnCanSaveCookie:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			response := &ICefResponse{instance: getPtr(3)}
			cookiePtr := (*tCefCookiePtr)(getInstance(getVal(4)))
			cookie := cookiePtr.convert()
			result := (*bool)(getInstance(getVal(5)))
			*result = fn.(cookieAccessOnCanSaveCookie)(browse, frame, request, response, cookie)
		default:
			return false
		}
		return true
	})
}
