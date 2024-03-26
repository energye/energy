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

// ICoreWebView2BrowserExtensionRemoveCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of removing
//	the browser Extension from the Profile.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionremovecompletedhandler">See the ICoreWebView2BrowserExtensionRemoveCompletedHandler article.</a>
type ICoreWebView2BrowserExtensionRemoveCompletedHandler interface {
	IObject
}

// TCoreWebView2BrowserExtensionRemoveCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of removing
//	the browser Extension from the Profile.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionremovecompletedhandler">See the ICoreWebView2BrowserExtensionRemoveCompletedHandler article.</a>
type TCoreWebView2BrowserExtensionRemoveCompletedHandler struct {
	TObject
}

func NewCoreWebView2BrowserExtensionRemoveCompletedHandler(aEvents IWVBrowserEvents, aExtensionID string) ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	r1 := WV().SysCallN(39, GetObjectUintptr(aEvents), PascalStr(aExtensionID))
	return AsCoreWebView2BrowserExtensionRemoveCompletedHandler(r1)
}

func CoreWebView2BrowserExtensionRemoveCompletedHandlerClass() TClass {
	ret := WV().SysCallN(38)
	return TClass(ret)
}
