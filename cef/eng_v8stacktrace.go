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

// ICefV8StackTrace Parent: ICefBaseRefCounted
//
//	Interface representing a V8 stack trace handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_trace_t))
type ICefV8StackTrace interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetFrameCount
	//  Returns the number of stack frames.
	GetFrameCount() int32 // function
	// GetFrame
	//  Returns the stack frame at the specified 0-based index.
	GetFrame(index int32) ICefV8StackFrame // function
}

// TCefV8StackTrace Parent: TCefBaseRefCounted
//
//	Interface representing a V8 stack trace handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_trace_t))
type TCefV8StackTrace struct {
	TCefBaseRefCounted
}

// V8StackTraceRef -> ICefV8StackTrace
var V8StackTraceRef v8StackTrace

// v8StackTrace TCefV8StackTrace Ref
type v8StackTrace uintptr

func (m *v8StackTrace) UnWrap(data uintptr) ICefV8StackTrace {
	var resultCefV8StackTrace uintptr
	CEF().SysCallN(1496, uintptr(data), uintptr(unsafePointer(&resultCefV8StackTrace)))
	return AsCefV8StackTrace(resultCefV8StackTrace)
}

func (m *v8StackTrace) Current(frameLimit int32) ICefV8StackTrace {
	var resultCefV8StackTrace uintptr
	CEF().SysCallN(1492, uintptr(frameLimit), uintptr(unsafePointer(&resultCefV8StackTrace)))
	return AsCefV8StackTrace(resultCefV8StackTrace)
}

func (m *TCefV8StackTrace) IsValid() bool {
	r1 := CEF().SysCallN(1495, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackTrace) GetFrameCount() int32 {
	r1 := CEF().SysCallN(1494, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackTrace) GetFrame(index int32) ICefV8StackFrame {
	var resultCefV8StackFrame uintptr
	CEF().SysCallN(1493, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefV8StackFrame)))
	return AsCefV8StackFrame(resultCefV8StackFrame)
}
