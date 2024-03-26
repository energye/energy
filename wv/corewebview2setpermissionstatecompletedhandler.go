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

// ICoreWebView2SetPermissionStateCompletedHandler Parent: IObject
//
//	The caller implements this interface to handle the result of
//	SetPermissionState.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler">See the ICoreWebView2SetPermissionStateCompletedHandler article.</a>
type ICoreWebView2SetPermissionStateCompletedHandler interface {
	IObject
}

// TCoreWebView2SetPermissionStateCompletedHandler Parent: TObject
//
//	The caller implements this interface to handle the result of
//	SetPermissionState.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler">See the ICoreWebView2SetPermissionStateCompletedHandler article.</a>
type TCoreWebView2SetPermissionStateCompletedHandler struct {
	TObject
}

func NewCoreWebView2SetPermissionStateCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2SetPermissionStateCompletedHandler {
	r1 := WV().SysCallN(612, GetObjectUintptr(aEvents))
	return AsCoreWebView2SetPermissionStateCompletedHandler(r1)
}

func CoreWebView2SetPermissionStateCompletedHandlerClass() TClass {
	ret := WV().SysCallN(611)
	return TClass(ret)
}
