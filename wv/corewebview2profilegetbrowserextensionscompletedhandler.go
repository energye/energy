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

// ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of
//	getting the browser Extensions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profilegetbrowserextensionscompletedhandler">See the ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler article.</a>
type ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler interface {
	IObject
}

// TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of
//	getting the browser Extensions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profilegetbrowserextensionscompletedhandler">See the ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler article.</a>
type TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler struct {
	TObject
}

func NewCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	r1 := WV().SysCallN(559, GetObjectUintptr(aEvents))
	return AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(r1)
}

func CoreWebView2ProfileGetBrowserExtensionsCompletedHandlerClass() TClass {
	ret := WV().SysCallN(558)
	return TClass(ret)
}
