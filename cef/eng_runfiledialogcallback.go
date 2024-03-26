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

// IRunFileDialogCallback Parent: ICefRunFileDialogCallback
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of
//	this interface will be called on the browser process UI thread.
//	<a cref="uCEFTypes|TCefRunFileDialogCallback">Implements TCefRunFileDialogCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t)</a>
type IRunFileDialogCallback interface {
	ICefRunFileDialogCallback
	// SetOnRunFileDialogDismissed
	//  Called asynchronously after the file dialog is dismissed. |file_paths|
	//  will be a single value or a list of values depending on the dialog mode.
	//  If the selection was cancelled |file_paths| will be NULL.
	SetOnRunFileDialogDismissed(fn TOnRunFileDialogDismissed) // property event
}

// TRunFileDialogCallback Parent: TCefRunFileDialogCallback
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of
//	this interface will be called on the browser process UI thread.
//	<a cref="uCEFTypes|TCefRunFileDialogCallback">Implements TCefRunFileDialogCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t)</a>
type TRunFileDialogCallback struct {
	TCefRunFileDialogCallback
	runFileDialogDismissedPtr uintptr
}

func NewRunFileDialogCallback() IRunFileDialogCallback {
	r1 := CEF().SysCallN(2224)
	return AsRunFileDialogCallback(r1)
}

func RunFileDialogCallbackClass() TClass {
	ret := CEF().SysCallN(2223)
	return TClass(ret)
}

func (m *TRunFileDialogCallback) SetOnRunFileDialogDismissed(fn TOnRunFileDialogDismissed) {
	if m.runFileDialogDismissedPtr != 0 {
		RemoveEventElement(m.runFileDialogDismissedPtr)
	}
	m.runFileDialogDismissedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2225, m.Instance(), m.runFileDialogDismissedPtr)
}
