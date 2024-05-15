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

// ICefSelectClientCertificateCallback Parent: ICefBaseRefCounted
//
//	Callback interface used to select a client certificate for authentication.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h (cef_select_client_certificate_callback_t))</a>
type ICefSelectClientCertificateCallback interface {
	ICefBaseRefCounted
	// Select
	//  Chooses the specified certificate for client certificate authentication. NULL value means that no client certificate should be used.
	Select(cert ICefX509Certificate) // procedure
}

// TCefSelectClientCertificateCallback Parent: TCefBaseRefCounted
//
//	Callback interface used to select a client certificate for authentication.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h (cef_select_client_certificate_callback_t))</a>
type TCefSelectClientCertificateCallback struct {
	TCefBaseRefCounted
}

// SelectClientCertificateCallbackRef -> ICefSelectClientCertificateCallback
var SelectClientCertificateCallbackRef selectClientCertificateCallback

// selectClientCertificateCallback TCefSelectClientCertificateCallback Ref
type selectClientCertificateCallback uintptr

func (m *selectClientCertificateCallback) UnWrap(data uintptr) ICefSelectClientCertificateCallback {
	var resultCefSelectClientCertificateCallback uintptr
	CEF().SysCallN(1355, uintptr(data), uintptr(unsafePointer(&resultCefSelectClientCertificateCallback)))
	return AsCefSelectClientCertificateCallback(resultCefSelectClientCertificateCallback)
}

func (m *TCefSelectClientCertificateCallback) Select(cert ICefX509Certificate) {
	CEF().SysCallN(1354, m.Instance(), GetObjectUintptr(cert))
}
