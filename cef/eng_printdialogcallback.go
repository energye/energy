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

// ICefPrintDialogCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print dialog requests.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_dialog_callback_t))
type ICefPrintDialogCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue printing with the specified |settings|.
	Cont(settings ICefPrintSettings) // procedure
	// Cancel
	//  Cancel the printing.
	Cancel() // procedure
}

// TCefPrintDialogCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print dialog requests.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_dialog_callback_t))
type TCefPrintDialogCallback struct {
	TCefBaseRefCounted
}

// PrintDialogCallbackRef -> ICefPrintDialogCallback
var PrintDialogCallbackRef printDialogCallback

// printDialogCallback TCefPrintDialogCallback Ref
type printDialogCallback uintptr

func (m *printDialogCallback) UnWrap(data uintptr) ICefPrintDialogCallback {
	var resultCefPrintDialogCallback uintptr
	CEF().SysCallN(1216, uintptr(data), uintptr(unsafePointer(&resultCefPrintDialogCallback)))
	return AsCefPrintDialogCallback(resultCefPrintDialogCallback)
}

func (m *TCefPrintDialogCallback) Cont(settings ICefPrintSettings) {
	CEF().SysCallN(1215, m.Instance(), GetObjectUintptr(settings))
}

func (m *TCefPrintDialogCallback) Cancel() {
	CEF().SysCallN(1214, m.Instance())
}
