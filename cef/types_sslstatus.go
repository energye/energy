//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefSSLStatus
//
//	Interface representing the SSL information for a navigation entry.
//	<para><see cref="uCEFTypes|TCefSSLStatus">Implements TCefSSLStatus</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_status_capi.h">CEF source file: /include/capi/cef_ssl_status_capi.h (cef_sslstatus_t)</see></para>
type ICefSSLStatus struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// SSLStatusRef -> ICefSSLStatus
var SSLStatusRef sslStatus

type sslStatus uintptr

func (*sslStatus) UnWrap(data *ICefSSLStatus) *ICefSSLStatus {
	var result uintptr
	imports.Proc(def.SSLStatusRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSSLStatus{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefSSLStatus) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefSSLStatus) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefSSLStatus) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Returns true (1) if the status is related to a secure SSL/TLS connection.
func (m *ICefSSLStatus) IsSecureConnection() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.SSLStatus_IsSecureConnection).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns a bitmask containing any and all problems verifying the server certificate.
func (m *ICefSSLStatus) GetCertStatus() consts.TCefCertStatus {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.SSLStatus_GetCertStatus).Call(m.Instance())
	return consts.TCefCertStatus(r1)
}

// Returns the SSL version used for the SSL connection.
func (m *ICefSSLStatus) GetSSLVersion() consts.TCefSSLVersion {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.SSLStatus_GetSSLVersion).Call(m.Instance())
	return consts.TCefSSLVersion(r1)
}

// Returns a bitmask containing the page security content status.
func (m *ICefSSLStatus) GetContentStatus() consts.TCefSSLContentStatus {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.SSLStatus_GetContentStatus).Call(m.Instance())
	return consts.TCefSSLContentStatus(r1)
}

// Returns the X.509 certificate.
func (m *ICefSSLStatus) GetX509Certificate() *ICefX509Certificate {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.SSLStatus_GetX509Certificate).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefX509Certificate{instance: getInstance(result)}
	}
	return nil
}
