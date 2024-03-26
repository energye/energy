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

// ICefPrintJobCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print job requests.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_job_callback_t))
type ICefPrintJobCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Indicate completion of the print job.
	Cont() // procedure
}

// TCefPrintJobCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print job requests.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_job_callback_t))
type TCefPrintJobCallback struct {
	TCefBaseRefCounted
}

// PrintJobCallbackRef -> ICefPrintJobCallback
var PrintJobCallbackRef printJobCallback

// printJobCallback TCefPrintJobCallback Ref
type printJobCallback uintptr

func (m *printJobCallback) UnWrap(data uintptr) ICefPrintJobCallback {
	var resultCefPrintJobCallback uintptr
	CEF().SysCallN(1218, uintptr(data), uintptr(unsafePointer(&resultCefPrintJobCallback)))
	return AsCefPrintJobCallback(resultCefPrintJobCallback)
}

func (m *TCefPrintJobCallback) Cont() {
	CEF().SysCallN(1217, m.Instance())
}
