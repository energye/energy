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

// ICoreWebView2PrintCompletedHandler Parent: IObject
//
//	Receives the result of the Print method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler">See the ICoreWebView2PrintCompletedHandler article.</a>
type ICoreWebView2PrintCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintCompletedHandler Parent: TObject
//
//	Receives the result of the Print method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler">See the ICoreWebView2PrintCompletedHandler article.</a>
type TCoreWebView2PrintCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintCompletedHandler {
	r1 := WV().SysCallN(492, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintCompletedHandler(r1)
}

func CoreWebView2PrintCompletedHandlerClass() TClass {
	ret := WV().SysCallN(491)
	return TClass(ret)
}
