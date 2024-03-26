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

// ICoreWebView2ProfileAddBrowserExtensionCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result
//	of loading an browser Extension.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profileaddbrowserextensioncompletedhandler">See the ICoreWebView2ProfileAddBrowserExtensionCompletedHandler article.</a>
type ICoreWebView2ProfileAddBrowserExtensionCompletedHandler interface {
	IObject
}

// TCoreWebView2ProfileAddBrowserExtensionCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result
//	of loading an browser Extension.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profileaddbrowserextensioncompletedhandler">See the ICoreWebView2ProfileAddBrowserExtensionCompletedHandler article.</a>
type TCoreWebView2ProfileAddBrowserExtensionCompletedHandler struct {
	TObject
}

func NewCoreWebView2ProfileAddBrowserExtensionCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	r1 := WV().SysCallN(557, GetObjectUintptr(aEvents))
	return AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler(r1)
}

func CoreWebView2ProfileAddBrowserExtensionCompletedHandlerClass() TClass {
	ret := WV().SysCallN(556)
	return TClass(ret)
}
