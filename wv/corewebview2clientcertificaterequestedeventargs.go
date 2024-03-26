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

// ICoreWebView2ClientCertificateRequestedEventArgs Parent: IObject
//
//	Event args for the ClientCertificateRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
type ICoreWebView2ClientCertificateRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificateRequestedEventArgs // property
	// Host
	//  Host name of the server that requested client certificate authentication.
	//  Normalization rules applied to the hostname are:
	//  * Convert to lowercase characters for ascii characters.
	//  * Punycode is used for representing non ascii characters.
	//  * Strip square brackets for IPV6 address.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_host">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Host() string // property
	// Port
	//  Port of the server that requested client certificate authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_port">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Port() int32 // property
	// IsProxy
	//  Returns true if the server that issued this request is an http proxy.
	//  Returns false if the server is the origin server.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_isproxy">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	IsProxy() bool // property
	// AllowedCertificateAuthorities
	//  Returns the `ICoreWebView2StringCollection`.
	//  The collection contains Base64 encoding of DER encoded distinguished names of
	//  certificate authorities allowed by the server.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_allowedcertificateauthorities">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	AllowedCertificateAuthorities() ICoreWebView2StringCollection // property
	// MutuallyTrustedCertificates
	//  Returns the `ICoreWebView2ClientCertificateCollection` when client
	//  certificate authentication is requested. The collection contains mutually
	//  trusted CA certificates.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_mutuallytrustedcertificates">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	MutuallyTrustedCertificates() ICoreWebView2ClientCertificateCollection // property
	// SelectedCertificate
	//  Returns the selected certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_selectedcertificate">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	SelectedCertificate() ICoreWebView2ClientCertificate // property
	// SetSelectedCertificate Set SelectedCertificate
	SetSelectedCertificate(AValue ICoreWebView2ClientCertificate) // property
	// Cancel
	//  You may set this flag to cancel the certificate selection. If canceled,
	//  the request is aborted regardless of the `Handled` property. By default the
	//  value is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_cancel">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// Handled
	//  You may set this flag to `TRUE` to respond to the server with or
	//  without a certificate. If this flag is `TRUE` with a `SelectedCertificate`
	//  it responds to the server with the selected certificate otherwise respond to the
	//  server without a certificate. By default the value of `Handled` and `Cancel` are `FALSE`
	//  and display default client certificate selection dialog prompt to allow the user to
	//  choose a certificate. The `SelectedCertificate` is ignored unless `Handled` is set `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_handled">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#getdeferral">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2ClientCertificateRequestedEventArgs Parent: TObject
//
//	Event args for the ClientCertificateRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
type TCoreWebView2ClientCertificateRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2ClientCertificateRequestedEventArgs(aArgs ICoreWebView2ClientCertificateRequestedEventArgs) ICoreWebView2ClientCertificateRequestedEventArgs {
	r1 := WV().SysCallN(79, GetObjectUintptr(aArgs))
	return AsCoreWebView2ClientCertificateRequestedEventArgs(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(83, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) BaseIntf() ICoreWebView2ClientCertificateRequestedEventArgs {
	var resultCoreWebView2ClientCertificateRequestedEventArgs uintptr
	WV().SysCallN(76, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ClientCertificateRequestedEventArgs)))
	return AsCoreWebView2ClientCertificateRequestedEventArgs(resultCoreWebView2ClientCertificateRequestedEventArgs)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Host() string {
	r1 := WV().SysCallN(82, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Port() int32 {
	r1 := WV().SysCallN(86, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) IsProxy() bool {
	r1 := WV().SysCallN(84, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) AllowedCertificateAuthorities() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	WV().SysCallN(75, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) MutuallyTrustedCertificates() ICoreWebView2ClientCertificateCollection {
	var resultCoreWebView2ClientCertificateCollection uintptr
	WV().SysCallN(85, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ClientCertificateCollection)))
	return AsCoreWebView2ClientCertificateCollection(resultCoreWebView2ClientCertificateCollection)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SelectedCertificate() ICoreWebView2ClientCertificate {
	var resultCoreWebView2ClientCertificate uintptr
	WV().SysCallN(87, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ClientCertificate)))
	return AsCoreWebView2ClientCertificate(resultCoreWebView2ClientCertificate)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetSelectedCertificate(AValue ICoreWebView2ClientCertificate) {
	WV().SysCallN(87, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Cancel() bool {
	r1 := WV().SysCallN(77, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetCancel(AValue bool) {
	WV().SysCallN(77, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Handled() bool {
	r1 := WV().SysCallN(81, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetHandled(AValue bool) {
	WV().SysCallN(81, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(80, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func CoreWebView2ClientCertificateRequestedEventArgsClass() TClass {
	ret := WV().SysCallN(78)
	return TClass(ret)
}
