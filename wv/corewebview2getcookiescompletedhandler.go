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

// ICoreWebView2GetCookiesCompletedHandler Parent: IObject
//
//	Receives the result of the GetCookies method.  The result is written to
//	the cookie list provided in the GetCookies method call.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getcookiescompletedhandler">See the ICoreWebView2GetCookiesCompletedHandler article.</a>
type ICoreWebView2GetCookiesCompletedHandler interface {
	IObject
}

// TCoreWebView2GetCookiesCompletedHandler Parent: TObject
//
//	Receives the result of the GetCookies method.  The result is written to
//	the cookie list provided in the GetCookies method call.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getcookiescompletedhandler">See the ICoreWebView2GetCookiesCompletedHandler article.</a>
type TCoreWebView2GetCookiesCompletedHandler struct {
	TObject
}

func NewCoreWebView2GetCookiesCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2GetCookiesCompletedHandler {
	r1 := WV().SysCallN(353, GetObjectUintptr(aEvents))
	return AsCoreWebView2GetCookiesCompletedHandler(r1)
}

func CoreWebView2GetCookiesCompletedHandlerClass() TClass {
	ret := WV().SysCallN(352)
	return TClass(ret)
}
