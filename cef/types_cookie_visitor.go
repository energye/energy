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

func (m *ICefCookieVisitor) SetOnVisit(fn cookieVisitorOnCookieOnVisit) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefCookieVisitor_OnVisit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type cookieVisitorOnCookieOnVisit func(cookie *TCefCookie, deleteCookie, result *bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case cookieVisitorOnCookieOnVisit:
			cookiePtr := (*tCefCookiePtr)(getPtr(0))
			cookie := cookiePtr.convert()
			deleteCookiePtr := (*bool)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			fn.(cookieVisitorOnCookieOnVisit)(cookie, deleteCookiePtr, resultPtr)
		default:
			return false
		}
		return true
	})
}
