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

// ICoreWebView2TrySuspendCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the TrySuspend result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2trysuspendcompletedhandler">See the ICoreWebView2TrySuspendCompletedHandler article.</a>
type ICoreWebView2TrySuspendCompletedHandler interface {
	IObject
}

// TCoreWebView2TrySuspendCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the TrySuspend result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2trysuspendcompletedhandler">See the ICoreWebView2TrySuspendCompletedHandler article.</a>
type TCoreWebView2TrySuspendCompletedHandler struct {
	TObject
}

func NewCoreWebView2TrySuspendCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2TrySuspendCompletedHandler {
	r1 := WV().SysCallN(655, GetObjectUintptr(aEvents))
	return AsCoreWebView2TrySuspendCompletedHandler(r1)
}

func CoreWebView2TrySuspendCompletedHandlerClass() TClass {
	ret := WV().SysCallN(654)
	return TClass(ret)
}
