//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF V8 上下文
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefV8Context) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefV8Context) Eval(code, scriptUrl string, startLine int32) (value *ICefV8Value, exception *ICefV8Exception, ok bool) {
	var returnValuePtr uintptr
	var returnExceptionPtr uintptr
	r1, _, _ := imports.Proc(internale_CefV8Context_Eval).Call(m.Instance(), api.PascalStr(code), api.PascalStr(scriptUrl), uintptr(startLine), uintptr(unsafe.Pointer(&returnValuePtr)), uintptr(unsafe.Pointer(&returnExceptionPtr)))
	ok = api.GoBool(r1)
	if ok {
		value = &ICefV8Value{instance: unsafe.Pointer(returnValuePtr)}
	} else {
		exception = &ICefV8Exception{instance: unsafe.Pointer(returnExceptionPtr)}
	}
	return value, exception, api.GoBool(r1)
}

// Enter 进入上下文
func (m *ICefV8Context) Enter() bool {
	r1, _, _ := imports.Proc(internale_CefV8Context_Enter).Call(m.Instance())
	return api.GoBool(r1)
}

// Exit 退出上下文
func (m *ICefV8Context) Exit() bool {
	r1, _, _ := imports.Proc(internale_CefV8Context_Exit).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Context) IsSame(that *ICefV8Context) bool {
	r1, _, _ := imports.Proc(internale_CefV8Context_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}
