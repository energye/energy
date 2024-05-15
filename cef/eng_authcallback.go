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

// ICefAuthCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of authentication requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_auth_callback_capi.h">CEF source file: /include/capi/cef_auth_callback_capi.h (cef_auth_callback_t))</a>
type ICefAuthCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the authentication request.
	Cont(username, password string) // procedure
	// Cancel
	//  Cancel the authentication request.
	Cancel() // procedure
}

// TCefAuthCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of authentication requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_auth_callback_capi.h">CEF source file: /include/capi/cef_auth_callback_capi.h (cef_auth_callback_t))</a>
type TCefAuthCallback struct {
	TCefBaseRefCounted
}

// AuthCallbackRef -> ICefAuthCallback
var AuthCallbackRef authCallback

// authCallback TCefAuthCallback Ref
type authCallback uintptr

func (m *authCallback) UnWrap(data uintptr) ICefAuthCallback {
	var resultCefAuthCallback uintptr
	CEF().SysCallN(586, uintptr(data), uintptr(unsafePointer(&resultCefAuthCallback)))
	return AsCefAuthCallback(resultCefAuthCallback)
}

func (m *TCefAuthCallback) Cont(username, password string) {
	CEF().SysCallN(585, m.Instance(), PascalStr(username), PascalStr(password))
}

func (m *TCefAuthCallback) Cancel() {
	CEF().SysCallN(584, m.Instance())
}
