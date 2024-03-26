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

// ICoreWebView2ClearBrowsingDataCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the ClearBrowsingData result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clearbrowsingdatacompletedhandler">See the ICoreWebView2ClearBrowsingDataCompletedHandler article.</a>
type ICoreWebView2ClearBrowsingDataCompletedHandler interface {
	IObject
}

// TCoreWebView2ClearBrowsingDataCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the ClearBrowsingData result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clearbrowsingdatacompletedhandler">See the ICoreWebView2ClearBrowsingDataCompletedHandler article.</a>
type TCoreWebView2ClearBrowsingDataCompletedHandler struct {
	TObject
}

func NewCoreWebView2ClearBrowsingDataCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ClearBrowsingDataCompletedHandler {
	r1 := WV().SysCallN(68, GetObjectUintptr(aEvents))
	return AsCoreWebView2ClearBrowsingDataCompletedHandler(r1)
}

func CoreWebView2ClearBrowsingDataCompletedHandlerClass() TClass {
	ret := WV().SysCallN(67)
	return TClass(ret)
}
