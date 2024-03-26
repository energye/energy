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

// ICoreWebView2PrintToPdfCompletedHandler Parent: IObject
//
//	Receives the result of the PrintToPdf method. If the print to PDF
//	operation succeeds, isSuccessful is true. Otherwise, if the operation
//	failed, isSuccessful is set to false. An invalid path returns
//	E_INVALIDARG.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfcompletedhandler">See the ICoreWebView2PrintToPdfCompletedHandler article.</a>
type ICoreWebView2PrintToPdfCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintToPdfCompletedHandler Parent: TObject
//
//	Receives the result of the PrintToPdf method. If the print to PDF
//	operation succeeds, isSuccessful is true. Otherwise, if the operation
//	failed, isSuccessful is set to false. An invalid path returns
//	E_INVALIDARG.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfcompletedhandler">See the ICoreWebView2PrintToPdfCompletedHandler article.</a>
type TCoreWebView2PrintToPdfCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintToPdfCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintToPdfCompletedHandler {
	r1 := WV().SysCallN(519, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintToPdfCompletedHandler(r1)
}

func CoreWebView2PrintToPdfCompletedHandlerClass() TClass {
	ret := WV().SysCallN(518)
	return TClass(ret)
}
