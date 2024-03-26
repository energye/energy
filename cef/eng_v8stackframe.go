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

// ICefV8StackFrame Parent: ICefBaseRefCounted
//
//	Interface representing a V8 stack frame handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_frame_t))
type ICefV8StackFrame interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetScriptName
	//  Returns the name of the resource script that contains the function.
	GetScriptName() string // function
	// GetScriptNameOrSourceUrl
	//  Returns the name of the resource script that contains the function or the sourceURL value if the script name is undefined and its source ends with a "//@ sourceURL=..." string.
	GetScriptNameOrSourceUrl() string // function
	// GetFunctionName
	//  Returns the name of the function.
	GetFunctionName() string // function
	// GetLineNumber
	//  Returns the 1-based line number for the function call or 0 if unknown.
	GetLineNumber() int32 // function
	// GetColumn
	//  Returns the 1-based column offset on the line for the function call or 0 if unknown.
	GetColumn() int32 // function
	// IsEval
	//  Returns true (1) if the function was compiled using eval().
	IsEval() bool // function
	// IsConstructor
	//  Returns true (1) if the function was called as a constructor via "new".
	IsConstructor() bool // function
}

// TCefV8StackFrame Parent: TCefBaseRefCounted
//
//	Interface representing a V8 stack frame handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_frame_t))
type TCefV8StackFrame struct {
	TCefBaseRefCounted
}

// V8StackFrameRef -> ICefV8StackFrame
var V8StackFrameRef v8StackFrame

// v8StackFrame TCefV8StackFrame Ref
type v8StackFrame uintptr

func (m *v8StackFrame) UnWrap(data uintptr) ICefV8StackFrame {
	var resultCefV8StackFrame uintptr
	CEF().SysCallN(1491, uintptr(data), uintptr(unsafePointer(&resultCefV8StackFrame)))
	return AsCefV8StackFrame(resultCefV8StackFrame)
}

func (m *TCefV8StackFrame) IsValid() bool {
	r1 := CEF().SysCallN(1490, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackFrame) GetScriptName() string {
	r1 := CEF().SysCallN(1486, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetScriptNameOrSourceUrl() string {
	r1 := CEF().SysCallN(1487, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetFunctionName() string {
	r1 := CEF().SysCallN(1484, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetLineNumber() int32 {
	r1 := CEF().SysCallN(1485, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackFrame) GetColumn() int32 {
	r1 := CEF().SysCallN(1483, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackFrame) IsEval() bool {
	r1 := CEF().SysCallN(1489, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackFrame) IsConstructor() bool {
	r1 := CEF().SysCallN(1488, m.Instance())
	return GoBool(r1)
}
