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
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ************************** ICefDeleteCookiesCallback ************************** //

// DeleteCookiesHandlerRef -> ICefDeleteCookiesCallback
var DeleteCookiesHandlerRef deleteCookiesHandler

type deleteCookiesHandler uintptr

func (*deleteCookiesHandler) NewForChromium(chromium IChromium) *ICefDeleteCookiesCallback {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefDeleteCookiesCallbackRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDeleteCookiesCallback{instance: unsafe.Pointer(result), ct: consts.CtChromium}
	}
	return nil
}

func (*deleteCookiesHandler) New() *ICefDeleteCookiesCallback {
	var result uintptr
	imports.Proc(internale_CefDeleteCookiesCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDeleteCookiesCallback{instance: unsafe.Pointer(result), ct: consts.CtTClient}
	}
	return nil
}

// Instance 实例
func (m *ICefDeleteCookiesCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDeleteCookiesCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
func (m *ICefDeleteCookiesCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDeleteCookiesCallback) IsTClientEvent() bool {
	return m.ct == consts.CtTClient
}

func (m *ICefDeleteCookiesCallback) IsChromiumEvent() bool {
	return m.ct == consts.CtChromium
}

func (m *ICefDeleteCookiesCallback) SetOnComplete(fn deleteCookiesOnComplete) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_CefDeleteCookiesCallback_OnComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** ICefSetCookieCallback ************************** //

// SetCookieHandlerRef -> ICefSetCookieCallback
var SetCookieHandlerRef setCookieHandler

type setCookieHandler uintptr

func (*setCookieHandler) NewForChromium(chromium IChromium, id int32) *ICefSetCookieCallback {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefSetCookieCallbackRef_CreateForChromium).Call(chromium.Instance(), uintptr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSetCookieCallback{instance: unsafe.Pointer(result), ct: consts.CtChromium}
	}
	return nil
}

func (*setCookieHandler) New() *ICefSetCookieCallback {
	var result uintptr
	imports.Proc(internale_CefSetCookieCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSetCookieCallback{instance: unsafe.Pointer(result), ct: consts.CtTClient}
	}
	return nil
}

func (m *ICefSetCookieCallback) IsTClientEvent() bool {
	return m.ct == consts.CtTClient
}

func (m *ICefSetCookieCallback) IsChromiumEvent() bool {
	return m.ct == consts.CtChromium
}

// Instance 实例
func (m *ICefSetCookieCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefSetCookieCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefSetCookieCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefSetCookieCallback) SetOnComplete(fn setCookieOnComplete) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_CefSetCookieCallback_OnComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** event ************************** //

type deleteCookiesOnComplete func(success bool)
type setCookieOnComplete func(numDeleted int32)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case deleteCookiesOnComplete:
			fn.(deleteCookiesOnComplete)(api.GoBool(getVal(0)))
		case setCookieOnComplete:
			fn.(setCookieOnComplete)(int32(getVal(0)))
		default:
			return false
		}
		return true
	})
}
