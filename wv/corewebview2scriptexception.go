//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2ScriptException Parent: IObject
//
//	This interface represents a JavaScript exception.
//	If the CoreWebView2.ExecuteScriptWithResult result has Succeeded as false,
//	you can use the result's Exception property to get the script exception.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptexception">See the ICoreWebView2ScriptException article.</a>
type ICoreWebView2ScriptException interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScriptException // property
	// LineNumber
	//  The line number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.lineNumber`.
	//  Note that this position starts at 0.
	LineNumber() uint32 // property
	// ColumnNumber
	//  The column number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.columnNumber`.
	//  Note that this position starts at 0.
	ColumnNumber() uint32 // property
	// Name
	//  The Name is the exception's class name.
	//  In the JSON it is `exceptionDetail.exception.className`.
	//  This is the empty string if the exception doesn't have a class name.
	//  This can happen if the script throws a non-Error object such as `throw "abc";`
	Name() string // property
	// Message
	//  The Message is the exception's message and potentially stack.
	//  In the JSON it is exceptionDetail.exception.description.
	//  This is the empty string if the exception doesn't have a description.
	//  This can happen if the script throws a non-Error object such as throw "abc";.
	Message() string // property
	// ToJson
	//  This will return all details of the exception as a JSON string.
	//  In the case that script has thrown a non-Error object such as `throw "abc";`
	//  or any other non-Error object, you can get object specific properties.
	ToJson() string // function
}

// TCoreWebView2ScriptException Parent: TObject
//
//	This interface represents a JavaScript exception.
//	If the CoreWebView2.ExecuteScriptWithResult result has Succeeded as false,
//	you can use the result's Exception property to get the script exception.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptexception">See the ICoreWebView2ScriptException article.</a>
type TCoreWebView2ScriptException struct {
	TObject
}

func NewCoreWebView2ScriptException(aBaseIntf ICoreWebView2ScriptException) ICoreWebView2ScriptException {
	r1 := WV().SysCallN(596, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ScriptException(r1)
}

func (m *TCoreWebView2ScriptException) Initialized() bool {
	r1 := WV().SysCallN(597, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ScriptException) BaseIntf() ICoreWebView2ScriptException {
	var resultCoreWebView2ScriptException uintptr
	WV().SysCallN(593, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ScriptException)))
	return AsCoreWebView2ScriptException(resultCoreWebView2ScriptException)
}

func (m *TCoreWebView2ScriptException) LineNumber() uint32 {
	r1 := WV().SysCallN(598, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ScriptException) ColumnNumber() uint32 {
	r1 := WV().SysCallN(595, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ScriptException) Name() string {
	r1 := WV().SysCallN(600, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptException) Message() string {
	r1 := WV().SysCallN(599, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptException) ToJson() string {
	r1 := WV().SysCallN(601, m.Instance())
	return GoStr(r1)
}

func CoreWebView2ScriptExceptionClass() TClass {
	ret := WV().SysCallN(594)
	return TClass(ret)
}
