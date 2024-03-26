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

// ICefSslInfo Parent: ICefBaseRefCounted
//
//	Interface representing SSL information.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_info_capi.h">CEF source file: /include/capi/cef_ssl_info_capi.h (cef_sslinfo_t))
type ICefSslInfo interface {
	ICefBaseRefCounted
	// GetCertStatus
	//  Returns a bitmask containing any and all problems verifying the server certificate.
	GetCertStatus() TCefCertStatus // function
	// GetX509Certificate
	//  Returns the X.509 certificate.
	GetX509Certificate() ICefX509Certificate // function
}

// TCefSslInfo Parent: TCefBaseRefCounted
//
//	Interface representing SSL information.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_info_capi.h">CEF source file: /include/capi/cef_ssl_info_capi.h (cef_sslinfo_t))
type TCefSslInfo struct {
	TCefBaseRefCounted
}

// SslInfoRef -> ICefSslInfo
var SslInfoRef sslInfo

// sslInfo TCefSslInfo Ref
type sslInfo uintptr

func (m *sslInfo) UnWrap(data uintptr) ICefSslInfo {
	var resultCefSslInfo uintptr
	CEF().SysCallN(1376, uintptr(data), uintptr(unsafePointer(&resultCefSslInfo)))
	return AsCefSslInfo(resultCefSslInfo)
}

func (m *TCefSslInfo) GetCertStatus() TCefCertStatus {
	r1 := CEF().SysCallN(1374, m.Instance())
	return TCefCertStatus(r1)
}

func (m *TCefSslInfo) GetX509Certificate() ICefX509Certificate {
	var resultCefX509Certificate uintptr
	CEF().SysCallN(1375, m.Instance(), uintptr(unsafePointer(&resultCefX509Certificate)))
	return AsCefX509Certificate(resultCefX509Certificate)
}
