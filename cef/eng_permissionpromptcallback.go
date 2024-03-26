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

// ICefPermissionPromptCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of permission prompts.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_prompt_callback_t))
type ICefPermissionPromptCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Complete the permissions request with the specified |result|.
	Cont(result TCefPermissionRequestResult) // procedure
}

// TCefPermissionPromptCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of permission prompts.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_prompt_callback_t))
type TCefPermissionPromptCallback struct {
	TCefBaseRefCounted
}

// PermissionPromptCallbackRef -> ICefPermissionPromptCallback
var PermissionPromptCallbackRef permissionPromptCallback

// permissionPromptCallback TCefPermissionPromptCallback Ref
type permissionPromptCallback uintptr

func (m *permissionPromptCallback) UnWrap(data uintptr) ICefPermissionPromptCallback {
	var resultCefPermissionPromptCallback uintptr
	CEF().SysCallN(1187, uintptr(data), uintptr(unsafePointer(&resultCefPermissionPromptCallback)))
	return AsCefPermissionPromptCallback(resultCefPermissionPromptCallback)
}

func (m *TCefPermissionPromptCallback) Cont(result TCefPermissionRequestResult) {
	CEF().SysCallN(1186, m.Instance(), uintptr(result))
}
