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

// ICoreWebView2ClientCertificate Parent: IObject
//
//	Provides access to the client certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate">See the ICoreWebView2ClientCertificate article.</a>
type ICoreWebView2ClientCertificate interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificate // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ClientCertificate) // property
	// Subject
	//  Subject of the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_subject">See the ICoreWebView2ClientCertificate article.</a>
	Subject() string // property
	// Issuer
	//  Name of the certificate authority that issued the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_issuer">See the ICoreWebView2ClientCertificate article.</a>
	Issuer() string // property
	// ValidFrom
	//  The valid start date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validfrom">See the ICoreWebView2ClientCertificate article.</a>
	ValidFrom() TDateTime // property
	// ValidTo
	//  The valid expiration date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validto">See the ICoreWebView2ClientCertificate article.</a>
	ValidTo() TDateTime // property
	// DerEncodedSerialNumber
	//  Base64 encoding of DER encoded serial number of the certificate.
	//  Read more about DER at [RFC 7468 DER]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_derencodedserialnumber">See the ICoreWebView2ClientCertificate article.</a>
	DerEncodedSerialNumber() string // property
	// DisplayName
	//  Display name for a certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_displayname">See the ICoreWebView2ClientCertificate article.</a>
	DisplayName() string // property
	// PemEncodedIssuerCertificateChain
	//  Collection of PEM encoded client certificate issuer chain.
	//  In this collection first element is the current certificate followed by
	//  intermediate1, intermediate2...intermediateN-1. Root certificate is the
	//  last element in collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_pemencodedissuercertificatechain">See the ICoreWebView2ClientCertificate article.</a>
	PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection // property
	// Kind
	//  Kind of a certificate(eg., smart card, pin, other).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_kind">See the ICoreWebView2ClientCertificate article.</a>
	Kind() TWVClientCertificateKind // property
	// ToPemEncoding
	//  PEM encoded data for the certificate.
	//  Returns Base64 encoding of DER encoded certificate.
	//  Read more about PEM at [RFC 1421 Privacy Enhanced Mail]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#topemencoding">See the ICoreWebView2ClientCertificate article.</a>
	ToPemEncoding() string // function
}

// TCoreWebView2ClientCertificate Parent: TObject
//
//	Provides access to the client certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate">See the ICoreWebView2ClientCertificate article.</a>
type TCoreWebView2ClientCertificate struct {
	TObject
}

func NewCoreWebView2ClientCertificate(aBaseIntf ICoreWebView2ClientCertificate) ICoreWebView2ClientCertificate {
	r1 := WV().SysCallN(90, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ClientCertificate(r1)
}

func (m *TCoreWebView2ClientCertificate) Initialized() bool {
	r1 := WV().SysCallN(93, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificate) BaseIntf() ICoreWebView2ClientCertificate {
	var resultCoreWebView2ClientCertificate uintptr
	WV().SysCallN(88, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ClientCertificate)))
	return AsCoreWebView2ClientCertificate(resultCoreWebView2ClientCertificate)
}

func (m *TCoreWebView2ClientCertificate) SetBaseIntf(AValue ICoreWebView2ClientCertificate) {
	WV().SysCallN(88, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ClientCertificate) Subject() string {
	r1 := WV().SysCallN(97, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) Issuer() string {
	r1 := WV().SysCallN(94, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) ValidFrom() TDateTime {
	r1 := WV().SysCallN(99, m.Instance())
	return TDateTime(r1)
}

func (m *TCoreWebView2ClientCertificate) ValidTo() TDateTime {
	r1 := WV().SysCallN(100, m.Instance())
	return TDateTime(r1)
}

func (m *TCoreWebView2ClientCertificate) DerEncodedSerialNumber() string {
	r1 := WV().SysCallN(91, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) DisplayName() string {
	r1 := WV().SysCallN(92, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	WV().SysCallN(96, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2ClientCertificate) Kind() TWVClientCertificateKind {
	r1 := WV().SysCallN(95, m.Instance())
	return TWVClientCertificateKind(r1)
}

func (m *TCoreWebView2ClientCertificate) ToPemEncoding() string {
	r1 := WV().SysCallN(98, m.Instance())
	return GoStr(r1)
}

func CoreWebView2ClientCertificateClass() TClass {
	ret := WV().SysCallN(89)
	return TClass(ret)
}
