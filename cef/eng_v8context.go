//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefv8Context Parent: ICefBaseRefCounted
//
//	Interface representing a V8 context handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefV8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8context_t))
type ICefv8Context interface {
	ICefBaseRefCounted
	// GetTaskRunner
	//  Returns the task runner associated with this context. V8 handles can only be accessed from the thread on which they are created. This function can be called on any render process thread.
	GetTaskRunner() ICefTaskRunner // function
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetBrowser
	//  Returns the browser for this context. This function will return an NULL reference for WebWorker contexts.
	GetBrowser() ICefBrowser // function
	// GetFrame
	//  Returns the frame for this context. This function will return an NULL reference for WebWorker contexts.
	GetFrame() ICefFrame // function
	// GetGlobal
	//  Returns the global object for this context. The context must be entered before calling this function.
	GetGlobal() ICefv8Value // function
	// Enter
	//  Enter this context. A context must be explicitly entered before creating a V8 Object, Array, Function or Date asynchronously. exit() must be called the same number of times as enter() before releasing this context. V8 objects belong to the context in which they are created. Returns true (1) if the scope was entered successfully.
	Enter() bool // function
	// Exit
	//  Exit this context. Call this function only after calling enter(). Returns true (1) if the scope was exited successfully.
	Exit() bool // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefv8Context) bool // function
	// Eval
	//  Execute a string of JavaScript code in this V8 context. The |script_url| parameter is the URL where the script in question can be found, if any. The |start_line| parameter is the base line number to use for error reporting. On success |retval| will be set to the return value, if any, and the function will return true (1). On failure |exception| will be set to the exception, if any, and the function will return false (0).
	Eval(code string, scripturl string, startline int32, retval *ICefv8Value, exception *ICefV8Exception) bool // function
}

// TCefv8Context Parent: TCefBaseRefCounted
//
//	Interface representing a V8 context handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefV8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8context_t))
type TCefv8Context struct {
	TCefBaseRefCounted
}

// V8ContextRef -> ICefv8Context
var V8ContextRef v8Context

// v8Context TCefv8Context Ref
type v8Context uintptr

func (m *v8Context) UnWrap(data uintptr) ICefv8Context {
	var resultCefv8Context uintptr
	CEF().SysCallN(1642, uintptr(data), uintptr(unsafePointer(&resultCefv8Context)))
	return AsCefv8Context(resultCefv8Context)
}

func (m *v8Context) Current() ICefv8Context {
	var resultCefv8Context uintptr
	CEF().SysCallN(1631, uintptr(unsafePointer(&resultCefv8Context)))
	return AsCefv8Context(resultCefv8Context)
}

func (m *v8Context) Entered() ICefv8Context {
	var resultCefv8Context uintptr
	CEF().SysCallN(1633, uintptr(unsafePointer(&resultCefv8Context)))
	return AsCefv8Context(resultCefv8Context)
}

func (m *TCefv8Context) GetTaskRunner() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	CEF().SysCallN(1639, m.Instance(), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCefv8Context) IsValid() bool {
	r1 := CEF().SysCallN(1641, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Context) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(1636, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefv8Context) GetFrame() ICefFrame {
	var resultCefFrame uintptr
	CEF().SysCallN(1637, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefv8Context) GetGlobal() ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1638, m.Instance(), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Context) Enter() bool {
	r1 := CEF().SysCallN(1632, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Context) Exit() bool {
	r1 := CEF().SysCallN(1635, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Context) IsSame(that ICefv8Context) bool {
	r1 := CEF().SysCallN(1640, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefv8Context) Eval(code string, scripturl string, startline int32, retval *ICefv8Value, exception *ICefV8Exception) bool {
	var result3 uintptr
	var result4 uintptr
	r1 := CEF().SysCallN(1634, m.Instance(), PascalStr(code), PascalStr(scripturl), uintptr(startline), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*retval = AsCefv8Value(result3)
	*exception = AsCefV8Exception(result4)
	return GoBool(r1)
}
