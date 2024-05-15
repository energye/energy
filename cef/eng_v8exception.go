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

// ICefV8Exception Parent: ICefBaseRefCounted
//
//	Interface representing a V8 exception. The functions of this interface may be called on any render process thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8exception_t))</a>
type ICefV8Exception interface {
	ICefBaseRefCounted
	// GetMessage
	//  Returns the exception message.
	GetMessage() string // function
	// GetSourceLine
	//  Returns the line of source code that the exception occurred within.
	GetSourceLine() string // function
	// GetScriptResourceName
	//  Returns the resource name for the script from where the function causing the error originates.
	GetScriptResourceName() string // function
	// GetLineNumber
	//  Returns the 1-based number of the line where the error occurred or 0 if the line number is unknown.
	GetLineNumber() int32 // function
	// GetStartPosition
	//  Returns the index within the script of the first character where the error occurred.
	GetStartPosition() int32 // function
	// GetEndPosition
	//  Returns the index within the script of the last character where the error occurred.
	GetEndPosition() int32 // function
	// GetStartColumn
	//  Returns the index within the line of the first character where the error occurred.
	GetStartColumn() int32 // function
	// GetEndColumn
	//  Returns the index within the line of the last character where the error occurred.
	GetEndColumn() int32 // function
}

// TCefV8Exception Parent: TCefBaseRefCounted
//
//	Interface representing a V8 exception. The functions of this interface may be called on any render process thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8exception_t))</a>
type TCefV8Exception struct {
	TCefBaseRefCounted
}

// V8ExceptionRef -> ICefV8Exception
var V8ExceptionRef v8Exception

// v8Exception TCefV8Exception Ref
type v8Exception uintptr

func (m *v8Exception) UnWrap(data uintptr) ICefV8Exception {
	var resultCefV8Exception uintptr
	CEF().SysCallN(1482, uintptr(data), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefV8Exception) GetMessage() string {
	r1 := CEF().SysCallN(1477, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetSourceLine() string {
	r1 := CEF().SysCallN(1479, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetScriptResourceName() string {
	r1 := CEF().SysCallN(1478, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetLineNumber() int32 {
	r1 := CEF().SysCallN(1476, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetStartPosition() int32 {
	r1 := CEF().SysCallN(1481, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetEndPosition() int32 {
	r1 := CEF().SysCallN(1475, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetStartColumn() int32 {
	r1 := CEF().SysCallN(1480, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetEndColumn() int32 {
	r1 := CEF().SysCallN(1474, m.Instance())
	return int32(r1)
}
