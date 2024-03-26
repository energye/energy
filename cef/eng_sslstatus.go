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

// ICefSSLStatus Parent: ICefBaseRefCounted
//
//	Interface representing the SSL information for a navigation entry.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_status_capi.h">CEF source file: /include/capi/cef_ssl_status_capi.h (cef_sslstatus_t))
type ICefSSLStatus interface {
	ICefBaseRefCounted
	// IsSecureConnection
	//  Returns true (1) if the status is related to a secure SSL/TLS connection.
	IsSecureConnection() bool // function
	// GetCertStatus
	//  Returns a bitmask containing any and all problems verifying the server certificate.
	GetCertStatus() TCefCertStatus // function
	// GetSSLVersion
	//  Returns the SSL version used for the SSL connection.
	GetSSLVersion() TCefSSLVersion // function
	// GetContentStatus
	//  Returns a bitmask containing the page security content status.
	GetContentStatus() TCefSSLContentStatus // function
	// GetX509Certificate
	//  Returns the X.509 certificate.
	GetX509Certificate() ICefX509Certificate // function
}

// TCefSSLStatus Parent: TCefBaseRefCounted
//
//	Interface representing the SSL information for a navigation entry.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_status_capi.h">CEF source file: /include/capi/cef_ssl_status_capi.h (cef_sslstatus_t))
type TCefSSLStatus struct {
	TCefBaseRefCounted
}

// SSLStatusRef -> ICefSSLStatus
var SSLStatusRef sSLStatus

// sSLStatus TCefSSLStatus Ref
type sSLStatus uintptr

func (m *sSLStatus) UnWrap(data uintptr) ICefSSLStatus {
	var resultCefSSLStatus uintptr
	CEF().SysCallN(1341, uintptr(data), uintptr(unsafePointer(&resultCefSSLStatus)))
	return AsCefSSLStatus(resultCefSSLStatus)
}

func (m *TCefSSLStatus) IsSecureConnection() bool {
	r1 := CEF().SysCallN(1340, m.Instance())
	return GoBool(r1)
}

func (m *TCefSSLStatus) GetCertStatus() TCefCertStatus {
	r1 := CEF().SysCallN(1336, m.Instance())
	return TCefCertStatus(r1)
}

func (m *TCefSSLStatus) GetSSLVersion() TCefSSLVersion {
	r1 := CEF().SysCallN(1338, m.Instance())
	return TCefSSLVersion(r1)
}

func (m *TCefSSLStatus) GetContentStatus() TCefSSLContentStatus {
	r1 := CEF().SysCallN(1337, m.Instance())
	return TCefSSLContentStatus(r1)
}

func (m *TCefSSLStatus) GetX509Certificate() ICefX509Certificate {
	var resultCefX509Certificate uintptr
	CEF().SysCallN(1339, m.Instance(), uintptr(unsafePointer(&resultCefX509Certificate)))
	return AsCefX509Certificate(resultCefX509Certificate)
}
