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

// ICoreWebView2ExecuteScriptResult Parent: IObject
//
//	This is the result for ExecuteScriptWithResult.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2executescriptresult">See the ICoreWebView2ExecuteScriptResult article.</a>
type ICoreWebView2ExecuteScriptResult interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ExecuteScriptResult // property
	// Succeeded
	//  This property is true if ExecuteScriptWithResult successfully executed script with
	//  no unhandled exceptions and the result is available in the ResultAsJson property
	//  or via the TryGetResultAsString method.
	//  If it is false then the script execution had an unhandled exception which you
	//  can get via the Exception property.
	Succeeded() bool // property
	// ResultAsJson
	//  A function that has no explicit return value returns undefined. If the
	//  script that was run throws an unhandled exception, then the result is
	//  also "null". This method is applied asynchronously. If the method is
	//  run before `ContentLoading`, the script will not be executed
	//  and the string "null" will be returned.
	//  The return value description is as follows
	//  1. S_OK: Execution succeeds.
	//  2. E_POINTER: When the `jsonResult` is nullptr.
	ResultAsJson() string // property
	// Exception
	//  If Succeeded is false, you can use this property to get the unhandled exception thrown by script execution
	//  Note that due to the compatibility of the WinRT/.NET interface,
	//  S_OK will be returned even if the acquisition fails.
	//  We can determine whether the acquisition is successful by judging whether the `exception` is nullptr.
	Exception() ICoreWebView2ScriptException // property
	// TryGetResultAsString
	//  If Succeeded is true and the result of script execution is a string, this method provides the value of the string result,
	//  and we will get the `FALSE` var value when the js result is not string type.
	//  The return value description is as follows
	//  1. S_OK: Execution succeeds.
	//  2. E_POINTER: When the `stringResult` or `value` is nullptr.
	//  NOTE: If the `value` returns `FALSE`, the `stringResult` will be set to a empty string.
	TryGetResultAsString(stringResult *string, value *bool) bool // function
}

// TCoreWebView2ExecuteScriptResult Parent: TObject
//
//	This is the result for ExecuteScriptWithResult.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2executescriptresult">See the ICoreWebView2ExecuteScriptResult article.</a>
type TCoreWebView2ExecuteScriptResult struct {
	TObject
}

func NewCoreWebView2ExecuteScriptResult(aBaseIntf ICoreWebView2ExecuteScriptResult) ICoreWebView2ExecuteScriptResult {
	r1 := WV().SysCallN(306, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ExecuteScriptResult(r1)
}

func (m *TCoreWebView2ExecuteScriptResult) Initialized() bool {
	r1 := WV().SysCallN(308, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ExecuteScriptResult) BaseIntf() ICoreWebView2ExecuteScriptResult {
	var resultCoreWebView2ExecuteScriptResult uintptr
	WV().SysCallN(304, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ExecuteScriptResult)))
	return AsCoreWebView2ExecuteScriptResult(resultCoreWebView2ExecuteScriptResult)
}

func (m *TCoreWebView2ExecuteScriptResult) Succeeded() bool {
	r1 := WV().SysCallN(310, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ExecuteScriptResult) ResultAsJson() string {
	r1 := WV().SysCallN(309, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ExecuteScriptResult) Exception() ICoreWebView2ScriptException {
	var resultCoreWebView2ScriptException uintptr
	WV().SysCallN(307, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ScriptException)))
	return AsCoreWebView2ScriptException(resultCoreWebView2ScriptException)
}

func (m *TCoreWebView2ExecuteScriptResult) TryGetResultAsString(stringResult *string, value *bool) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := WV().SysCallN(311, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*stringResult = GoStr(result0)
	*value = GoBool(result1)
	return GoBool(r1)
}

func CoreWebView2ExecuteScriptResultClass() TClass {
	ret := WV().SysCallN(305)
	return TClass(ret)
}
