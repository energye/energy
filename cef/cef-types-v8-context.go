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

// V8ContextRef -> ICefV8Context
var V8ContextRef *cefV8ContextRef

// cefV8ContextRef
type cefV8ContextRef uintptr

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

func (m *ICefV8Context) Browser() *ICefBrowser {
	if m.browser == nil {
		var result uintptr
		imports.Proc(internale_CefV8Context_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.browser = &ICefBrowser{
			instance: unsafe.Pointer(result),
		}
	}
	return m.browser
}

func (m *ICefV8Context) Frame() *ICefFrame {
	if m.frame == nil {
		var result uintptr
		imports.Proc(internale_CefV8Context_Frame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.frame = &ICefFrame{
			instance: unsafe.Pointer(result),
		}
	}
	return m.frame
}

func (m *ICefV8Context) Global() *ICefV8Value {
	if m.global == nil {
		var result uintptr
		imports.Proc(internale_CefV8Context_Global).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.global = &ICefV8Value{
			instance: unsafe.Pointer(result),
		}
	}
	return m.global
}

func (m *ICefV8Context) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *cefV8ContextRef) Current() *ICefV8Context {
	var result uintptr
	imports.Proc(internale_CefV8ContextRef_Current).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{
		instance: unsafe.Pointer(result),
	}
}

func (m *cefV8ContextRef) Entered() *ICefV8Context {
	var result uintptr
	imports.Proc(internale_CefV8ContextRef_Entered).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{
		instance: unsafe.Pointer(result),
	}
}

func (m *cefV8ContextRef) UnWrap(data *ICefV8Context) *ICefV8Context {
	var result uintptr
	imports.Proc(internale_CefV8ContextRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}
