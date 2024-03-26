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

// ICoreWebView2BrowserExtensionEnableCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of setting the
//	browser Extension as enabled or disabled. If enabled, the browser Extension is
//	running in WebView instances. If disabled, the browser Extension is not running in WebView instances.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionenablecompletedhandler">See the ICoreWebView2BrowserExtensionEnableCompletedHandler article.</a>
type ICoreWebView2BrowserExtensionEnableCompletedHandler interface {
	IObject
}

// TCoreWebView2BrowserExtensionEnableCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of setting the
//	browser Extension as enabled or disabled. If enabled, the browser Extension is
//	running in WebView instances. If disabled, the browser Extension is not running in WebView instances.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionenablecompletedhandler">See the ICoreWebView2BrowserExtensionEnableCompletedHandler article.</a>
type TCoreWebView2BrowserExtensionEnableCompletedHandler struct {
	TObject
}

func NewCoreWebView2BrowserExtensionEnableCompletedHandler(aEvents IWVBrowserEvents, aExtensionID string) ICoreWebView2BrowserExtensionEnableCompletedHandler {
	r1 := WV().SysCallN(31, GetObjectUintptr(aEvents), PascalStr(aExtensionID))
	return AsCoreWebView2BrowserExtensionEnableCompletedHandler(r1)
}

func CoreWebView2BrowserExtensionEnableCompletedHandlerClass() TClass {
	ret := WV().SysCallN(30)
	return TClass(ret)
}
