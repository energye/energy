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

// ICefDeleteCookiesCallback
type ICefDeleteCookiesCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// DeleteCookiesHandlerRef -> ICefDeleteCookiesCallback
var DeleteCookiesHandlerRef deleteCookiesHandler

type deleteCookiesHandler uintptr

func (*deleteCookiesHandler) NewForChromium(chromium IChromium) *ICefDeleteCookiesCallback {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDeleteCookiesCallbackRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDeleteCookiesCallback{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

func (*deleteCookiesHandler) New() *ICefDeleteCookiesCallback {
	var result uintptr
	imports.Proc(def.CefDeleteCookiesCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDeleteCookiesCallback{instance: unsafe.Pointer(result)}
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

func (m *ICefDeleteCookiesCallback) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefDeleteCookiesCallback) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

func (m *ICefDeleteCookiesCallback) SetOnComplete(fn deleteCookiesOnComplete) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.CefDeleteCookiesCallback_OnComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** ICefSetCookieCallback ************************** //

// ICefSetCookieCallback
type ICefSetCookieCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// SetCookieHandlerRef -> ICefSetCookieCallback
var SetCookieHandlerRef setCookieHandler

type setCookieHandler uintptr

func (*setCookieHandler) NewForChromium(chromium IChromium, id int32) *ICefSetCookieCallback {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefSetCookieCallbackRef_CreateForChromium).Call(chromium.Instance(), uintptr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSetCookieCallback{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

func (*setCookieHandler) New() *ICefSetCookieCallback {
	var result uintptr
	imports.Proc(def.CefSetCookieCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSetCookieCallback{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefSetCookieCallback) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefSetCookieCallback) IsOtherEvent() bool {
	return m.ct == consts.CtOther
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
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.CefSetCookieCallback_OnComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
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
