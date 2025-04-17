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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefV8Context
//
// Interface representing a V8 context handle. V8 handles can only be accessed
// from the thread on which they are created. Valid threads for creating a V8
// handle include the render process main thread (TID_RENDERER) and WebWorker
// threads. A task runner for posting tasks on the associated thread can be
// retrieved via the ICefV8context.GetTaskRunner() function.
// <para><see cref="uCEFTypes|TCefV8Context">Implements TCefV8Context</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8context_t)</see></para>
type ICefV8Context struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	browser  *ICefBrowser
	frame    *ICefFrame
	global   *ICefV8Value
}

// V8ContextRef -> ICefV8Context
var V8ContextRef *cefV8ContextRef

// cefV8ContextRef
type cefV8ContextRef uintptr

// Instance 实例
func (m *ICefV8Context) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefV8Context) Eval(code, scriptUrl string, startLine int32) (value *ICefV8Value, exception *ICefV8Exception, ok bool) {
	if !m.IsValid() {
		return
	}
	var returnValuePtr uintptr
	var returnExceptionPtr uintptr
	r1, _, _ := imports.Proc(def.CefV8Context_Eval).Call(m.Instance(), api.PascalStr(code), api.PascalStr(scriptUrl), uintptr(startLine), uintptr(unsafe.Pointer(&returnValuePtr)), uintptr(unsafe.Pointer(&returnExceptionPtr)))
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
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Context_Enter).Call(m.Instance())
	return api.GoBool(r1)
}

// Exit 退出上下文
func (m *ICefV8Context) Exit() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Context_Exit).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Context) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefV8Context) IsSame(that *ICefV8Context) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Context_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Context) Browser() *ICefBrowser {
	if !m.IsValid() {
		return nil
	}
	if m.browser == nil {
		var result uintptr
		imports.Proc(def.CefV8Context_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.browser = &ICefBrowser{
			instance: unsafe.Pointer(result),
		}
	}
	return m.browser
}

func (m *ICefV8Context) Frame() *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	if m.frame == nil {
		var result uintptr
		imports.Proc(def.CefV8Context_Frame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.frame = &ICefFrame{
			instance: unsafe.Pointer(result),
		}
	}
	return m.frame
}

func (m *ICefV8Context) Global() *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	if m.global == nil {
		var result uintptr
		imports.Proc(def.CefV8Context_Global).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.global = &ICefV8Value{
			instance: unsafe.Pointer(result),
		}
	}
	return m.global
}

func (m *ICefV8Context) Free() {
	if m.instance != nil {
		if m.frame != nil {
			m.frame.Free()
			m.frame = nil
		}
		if m.browser != nil {
			m.browser.Free()
			m.browser = nil
		}
		if m.global != nil {
			m.global.Free()
			m.global = nil
		}
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *cefV8ContextRef) Current() *ICefV8Context {
	var result uintptr
	imports.Proc(def.CefV8ContextRef_Current).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{
		instance: unsafe.Pointer(result),
	}
}

func (m *cefV8ContextRef) Entered() *ICefV8Context {
	var result uintptr
	imports.Proc(def.CefV8ContextRef_Entered).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{
		instance: unsafe.Pointer(result),
	}
}

func (m *cefV8ContextRef) UnWrap(data *ICefV8Context) *ICefV8Context {
	var result uintptr
	imports.Proc(def.CefV8ContextRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}
