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

// ICoreWebView2PrintToPdfStreamCompletedHandler Parent: IObject
//
//	Receives the result of the PrintToPdfStream method.
//	errorCode returns S_OK if the PrintToPdfStream operation succeeded.
//	The printable pdf data is returned in the pdfStream object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler">See the ICoreWebView2PrintToPdfStreamCompletedHandler article.</a>
type ICoreWebView2PrintToPdfStreamCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintToPdfStreamCompletedHandler Parent: TObject
//
//	Receives the result of the PrintToPdfStream method.
//	errorCode returns S_OK if the PrintToPdfStream operation succeeded.
//	The printable pdf data is returned in the pdfStream object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler">See the ICoreWebView2PrintToPdfStreamCompletedHandler article.</a>
type TCoreWebView2PrintToPdfStreamCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintToPdfStreamCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintToPdfStreamCompletedHandler {
	r1 := WV().SysCallN(521, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintToPdfStreamCompletedHandler(r1)
}

func CoreWebView2PrintToPdfStreamCompletedHandlerClass() TClass {
	ret := WV().SysCallN(520)
	return TClass(ret)
}
