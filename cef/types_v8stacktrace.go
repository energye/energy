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

// ICefV8StackTrace
//
//	Interface representing a V8 stack trace handle. V8 handles can only be
//	accessed from the thread on which they are created. Valid threads for
//	creating a V8 handle include the render process main thread (TID_RENDERER)
//	and WebWorker threads. A task runner for posting tasks on the associated
//	thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<para><see cref="uCEFTypes|TCefV8StackTrace">Implements TCefV8StackTrace</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_trace_t)</see></para>
type ICefV8StackTrace struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// V8StackTraceRef -> ICefV8StackTrace
var V8StackTraceRef v8StackTrace

type v8StackTrace uintptr

func (*v8StackTrace) UnWrap(data *ICefV8StackTrace) *ICefV8StackTrace {
	var result uintptr
	imports.Proc(def.V8StackTraceRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefV8StackTrace{instance: getInstance(result)}
	}
	return nil
}

func (*v8StackTrace) Current(frameLimit int32) *ICefV8StackTrace {
	var result uintptr
	imports.Proc(def.V8StackTraceRef_Current).Call(uintptr(frameLimit), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefV8StackTrace{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefV8StackTrace) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8StackTrace) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefV8StackTrace) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.V8StackTrace_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns true (1) if the underlying handle is valid and it can be accessed
// on the current thread. Do not call any other functions if this function
// returns false (0).
// Returns the number of stack frames.
func (m *ICefV8StackTrace) GetFrameCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.V8StackTrace_GetFrameCount).Call(m.Instance())
	return int32(r1)
}

// Returns the stack frame at the specified 0-based index.
func (m *ICefV8StackTrace) GetFrame(index int32) *ICefV8StackFrame {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.V8StackTrace_GetFrame).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefV8StackFrame{instance: getInstance(result)}
	}
	return nil
}

// Interface representing a V8 stack frame handle. V8 handles can only be
// accessed from the thread on which they are created. Valid threads for
// creating a V8 handle include the render process main thread (TID_RENDERER)
// and WebWorker threads. A task runner for posting tasks on the associated
// thread can be retrieved via the ICefv8context.GetTaskRunner() function.
// <para><see cref="uCEFTypes|TCefV8StackFrame">Implements TCefV8StackFrame</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_frame_t)</see></para>
type ICefV8StackFrame struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// V8StackFrameRef -> ICefV8StackFrame
var V8StackFrameRef v8StackFrame

type v8StackFrame uintptr

func (*v8StackFrame) UnWrap(data *ICefV8StackFrame) *ICefV8StackFrame {
	var result uintptr
	imports.Proc(def.V8StackFrameRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefV8StackFrame{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefV8StackFrame) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8StackFrame) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefV8StackFrame) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns the name of the resource script that contains the function.
func (m *ICefV8StackFrame) GetScriptName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_GetScriptName).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the name of the resource script that contains the function or the
// sourceURL value if the script name is undefined and its source ends with a "//@ sourceURL=..." string.
func (m *ICefV8StackFrame) GetScriptNameOrSourceUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_GetScriptNameOrSourceUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the name of the function.
func (m *ICefV8StackFrame) GetFunctionName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_GetFunctionName).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the 1-based line number for the function call or 0 if unknown.
func (m *ICefV8StackFrame) GetLineNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_GetLineNumber).Call(m.Instance())
	return int32(r1)
}

// Returns the 1-based column offset on the line for the function call or 0 if unknown.
func (m *ICefV8StackFrame) GetColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_GetColumn).Call(m.Instance())
	return int32(r1)
}

// Returns true (1) if the function was compiled using eval().
func (m *ICefV8StackFrame) IsEval() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_IsEval).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns true (1) if the function was called as a constructor via "new".
func (m *ICefV8StackFrame) IsConstructor() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.V8StackFrame_IsConstructor).Call(m.Instance())
	return api.GoBool(r1)
}
