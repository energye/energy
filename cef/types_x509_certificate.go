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
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

// ICefX509Certificate
//
//	/include/capi/cef_x509_certificate_capi.h (cef_x509certificate_t)
type ICefX509Certificate struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefX509Certificate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefX509Certificate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefX509Certificate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefX509Certificate) GetSubject() *ICefX509CertPrincipal {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetSubject).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefX509CertPrincipal{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetIssuer() *ICefX509CertPrincipal {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetIssuer).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefX509CertPrincipal{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetSerialNumber() *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetSerialNumber).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBinaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetValidStart() (result int64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefX509Certificate_GetValidStart).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefX509Certificate) GetValidExpiry() (result int64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefX509Certificate_GetValidExpiry).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefX509Certificate) GetValidStartAsDateTime() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	var result float64
	imports.Proc(def.CefX509Certificate_GetValidStartAsDateTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return common.DDateTimeToGoDateTime(result)
	}
	return time.Time{}
}

func (m *ICefX509Certificate) GetValidExpiryAsDateTime() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	var result float64
	imports.Proc(def.CefX509Certificate_GetValidExpiryAsDateTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return common.DDateTimeToGoDateTime(result)
	}
	return time.Time{}
}

func (m *ICefX509Certificate) GetDerEncoded() *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetDerEncoded).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBinaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetPemEncoded() *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetPemEncoded).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBinaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetIssuerChainSize() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefX509Certificate_GetIssuerChainSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefX509Certificate) GetDEREncodedIssuerChain(chainCount uint32) *TCefBinaryValueArray {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetDEREncodedIssuerChain).Call(m.Instance(), uintptr(chainCount), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCefBinaryValueArray{count: chainCount, instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefX509Certificate) GetPEMEncodedIssuerChain(chainCount uint32) *TCefBinaryValueArray {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefX509Certificate_GetPEMEncodedIssuerChain).Call(m.Instance(), uintptr(chainCount), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCefBinaryValueArray{count: chainCount, instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** TCefX509CertificateArray ************************** //

// TCefX509CertificateArray
// []ICefX509Certificate
type TCefX509CertificateArray struct {
	instance     unsafe.Pointer
	certificates []*ICefX509Certificate
	count        uint32
}

func (m *TCefX509CertificateArray) Get(index uint32) *ICefX509Certificate {
	if index < m.count {
		if m.certificates == nil {
			m.certificates = make([]*ICefX509Certificate, m.count, m.count)
		}
		var result uintptr
		imports.Proc(def.CefX509CertificateArray_Get).Call(uintptr(m.instance), uintptr(index), uintptr(unsafe.Pointer(&result)))
		if result != 0 {
			m.certificates[index] = &ICefX509Certificate{instance: unsafe.Pointer(result)}
			return m.certificates[index]
		}
	}
	return nil
}

func (m *TCefX509CertificateArray) Count() uint32 {
	return m.count
}

func (m *TCefX509CertificateArray) Free() {
	if m.instance != nil {
		if m.certificates != nil {
			for _, cert := range m.certificates {
				cert.Free()
			}
			m.certificates = nil
		}
		m.instance = nil
		m.count = 0
	}
}

// ************************** ICefX509CertPrincipal ************************** //

// ICefX509CertPrincipal
// include/capi/cef_x509_certificate_capi.h (cef_x509cert_principal_t)
type ICefX509CertPrincipal struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefX509CertPrincipal) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefX509CertPrincipal) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefX509CertPrincipal) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefX509CertPrincipal) GetDisplayName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefX509CertPrincipal_GetDisplayName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefX509CertPrincipal) GetCommonName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefX509CertPrincipal_GetCommonName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefX509CertPrincipal) GetLocalityName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefX509CertPrincipal_GetLocalityName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefX509CertPrincipal) GetStateOrProvinceName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefX509CertPrincipal_GetStateOrProvinceName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefX509CertPrincipal) GetCountryName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefX509CertPrincipal_GetCountryName).Call(m.Instance())
	return api.GoStr(r1)
}

// GetStreetAddresses
//
//	CEF 115 Remove
func (m *ICefX509CertPrincipal) GetStreetAddresses() *lcl.TStrings {
	if !m.IsValid() {
		return nil
	}
	addresses := lcl.NewStrings()
	imports.Proc(def.CefX509CertPrincipal_GetStreetAddresses).Call(m.Instance(), addresses.Instance())
	return addresses
}

func (m *ICefX509CertPrincipal) GetOrganizationNames() *lcl.TStrings {
	if !m.IsValid() {
		return nil
	}
	addresses := lcl.NewStrings()
	imports.Proc(def.CefX509CertPrincipal_GetOrganizationNames).Call(m.Instance(), addresses.Instance())
	return addresses
}

func (m *ICefX509CertPrincipal) GetOrganizationUnitNames() *lcl.TStrings {
	if !m.IsValid() {
		return nil
	}
	addresses := lcl.NewStrings()
	imports.Proc(def.CefX509CertPrincipal_GetOrganizationUnitNames).Call(m.Instance(), addresses.Instance())
	return addresses
}

// GetDomainComponents
//
//	CEF 115 Remove
func (m *ICefX509CertPrincipal) GetDomainComponents() *lcl.TStrings {
	if !m.IsValid() {
		return nil
	}
	addresses := lcl.NewStrings()
	imports.Proc(def.CefX509CertPrincipal_GetDomainComponents).Call(m.Instance(), addresses.Instance())
	return addresses
}

// ************************** ICefX509CertPrincipal ************************** //

// ICefSelectClientCertificateCallback
//
//	/include/capi/cef_request_handler_capi.h (cef_select_client_certificate_callback_t)
type ICefSelectClientCertificateCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefSelectClientCertificateCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefSelectClientCertificateCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefSelectClientCertificateCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefSelectClientCertificateCallback) Select(cert *ICefX509Certificate) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefSelectClientCertificateCallback_Select).Call(m.Instance(), cert.Instance())
}

// ************************** ICefSslInfo ************************** //

// ICefSslInfo
//
//	/include/capi/cef_ssl_info_capi.h (cef_sslinfo_t)
type ICefSslInfo struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefSslInfo) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefSslInfo) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefSslInfo) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefSslInfo) GetCertStatus() consts.TCefCertStatus {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefSslInfo_GetCertStatus).Call(m.Instance())
	return consts.TCefCertStatus(r1)
}

func (m *ICefSslInfo) GetX509Certificate() *ICefX509Certificate {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefSslInfo_GetX509Certificate).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefX509Certificate{instance: unsafe.Pointer(result)}
	}
	return nil
}
