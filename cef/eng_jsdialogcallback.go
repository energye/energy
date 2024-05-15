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

// ICefJsDialogCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of JavaScript dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h (cef_jsdialog_callback_t))</a>
type ICefJsDialogCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the JS dialog request. Set |success| to true (1) if the OK button was pressed. The |user_input| value should be specified for prompt dialogs.
	Cont(success bool, userInput string) // procedure
}

// TCefJsDialogCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of JavaScript dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h (cef_jsdialog_callback_t))</a>
type TCefJsDialogCallback struct {
	TCefBaseRefCounted
}

// JsDialogCallbackRef -> ICefJsDialogCallback
var JsDialogCallbackRef jsDialogCallback

// jsDialogCallback TCefJsDialogCallback Ref
type jsDialogCallback uintptr

func (m *jsDialogCallback) UnWrap(data uintptr) ICefJsDialogCallback {
	var resultCefJsDialogCallback uintptr
	CEF().SysCallN(1002, uintptr(data), uintptr(unsafePointer(&resultCefJsDialogCallback)))
	return AsCefJsDialogCallback(resultCefJsDialogCallback)
}

func (m *TCefJsDialogCallback) Cont(success bool, userInput string) {
	CEF().SysCallN(1001, m.Instance(), PascalBool(success), PascalStr(userInput))
}
