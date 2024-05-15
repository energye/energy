//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefMediaAccessCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of media access permission requests.
//	This record is declared twice with almost identical parameters. "allowed_permissions" is defined as int and uint32.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_access_handler_capi.h">CEF source file: /include/capi/cef_media_access_handler_capi.h (cef_media_access_callback_t))</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_media_access_callback_t))</a>
type ICefMediaAccessCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Call to allow or deny media access. If this callback was initiated in response to a getUserMedia (indicated by CEF_MEDIA_PERMISSION_DEVICE_AUDIO_CAPTURE and/or CEF_MEDIA_PERMISSION_DEVICE_VIDEO_CAPTURE being set) then |allowed_permissions| must match |required_permissions| passed to OnRequestMediaAccessPermission.
	Cont(allowedpermissions TCefMediaAccessPermissionTypes) // procedure
	// Cancel
	//  Cancel the media access request.
	Cancel() // procedure
}

// TCefMediaAccessCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of media access permission requests.
//	This record is declared twice with almost identical parameters. "allowed_permissions" is defined as int and uint32.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_access_handler_capi.h">CEF source file: /include/capi/cef_media_access_handler_capi.h (cef_media_access_callback_t))</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_media_access_callback_t))</a>
type TCefMediaAccessCallback struct {
	TCefBaseRefCounted
}

// MediaAccessCallbackRef -> ICefMediaAccessCallback
var MediaAccessCallbackRef mediaAccessCallback

// mediaAccessCallback TCefMediaAccessCallback Ref
type mediaAccessCallback uintptr

func (m *mediaAccessCallback) UnWrap(data uintptr) ICefMediaAccessCallback {
	var resultCefMediaAccessCallback uintptr
	CEF().SysCallN(1052, uintptr(data), uintptr(unsafePointer(&resultCefMediaAccessCallback)))
	return AsCefMediaAccessCallback(resultCefMediaAccessCallback)
}

func (m *TCefMediaAccessCallback) Cont(allowedpermissions TCefMediaAccessPermissionTypes) {
	CEF().SysCallN(1051, m.Instance(), uintptr(allowedpermissions))
}

func (m *TCefMediaAccessCallback) Cancel() {
	CEF().SysCallN(1050, m.Instance())
}
