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

// ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Parent: IObject
//
//	The caller implements this interface to handle the result of
//	GetNonDefaultPermissionSettings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler">See the ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler article.</a>
type ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler interface {
	IObject
}

// TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Parent: TObject
//
//	The caller implements this interface to handle the result of
//	GetNonDefaultPermissionSettings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler">See the ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler article.</a>
type TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler struct {
	TObject
}

func NewCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	r1 := WV().SysCallN(355, GetObjectUintptr(aEvents))
	return AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(r1)
}

func CoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerClass() TClass {
	ret := WV().SysCallN(354)
	return TClass(ret)
}
