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

// ICefFileDialogCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of file dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_file_dialog_callback_t))</a>
type ICefFileDialogCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the file selection. |file_paths| should be a single value or a list of values depending on the dialog mode. An NULL |file_paths| value is treated the same as calling cancel().
	Cont(filePaths IStrings) // procedure
	// Cancel
	//  Cancel the file selection.
	Cancel() // procedure
}

// TCefFileDialogCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of file dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_file_dialog_callback_t))</a>
type TCefFileDialogCallback struct {
	TCefBaseRefCounted
}

// FileDialogCallbackRef -> ICefFileDialogCallback
var FileDialogCallbackRef fileDialogCallback

// fileDialogCallback TCefFileDialogCallback Ref
type fileDialogCallback uintptr

func (m *fileDialogCallback) UnWrap(data uintptr) ICefFileDialogCallback {
	var resultCefFileDialogCallback uintptr
	CEF().SysCallN(953, uintptr(data), uintptr(unsafePointer(&resultCefFileDialogCallback)))
	return AsCefFileDialogCallback(resultCefFileDialogCallback)
}

func (m *TCefFileDialogCallback) Cont(filePaths IStrings) {
	CEF().SysCallN(952, m.Instance(), GetObjectUintptr(filePaths))
}

func (m *TCefFileDialogCallback) Cancel() {
	CEF().SysCallN(951, m.Instance())
}
