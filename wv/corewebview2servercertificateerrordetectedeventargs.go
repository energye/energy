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

// ICoreWebView2ServerCertificateErrorDetectedEventArgs Parent: IObject
//
//	Event args for the ServerCertificateErrorDetected event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
type ICoreWebView2ServerCertificateErrorDetectedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ServerCertificateErrorDetectedEventArgs // property
	// ErrorStatus
	//  The TLS error code for the invalid certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_errorstatus">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
	ErrorStatus() TWVWebErrorStatus // property
	// RequestUri
	//  URI associated with the request for the invalid certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_requesturi">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
	RequestUri() string // property
	// ServerCertificate
	//  Returns the server certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_servercertificate">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
	ServerCertificate() ICoreWebView2Certificate // property
	// Action
	//  The action of the server certificate error detection.
	//  The default value is `COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_DEFAULT`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_action">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
	Action() TWVServerCertificateErrorAction // property
	// SetAction Set Action
	SetAction(AValue TWVServerCertificateErrorAction) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#getdeferral">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2ServerCertificateErrorDetectedEventArgs Parent: TObject
//
//	Event args for the ServerCertificateErrorDetected event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</a>
type TCoreWebView2ServerCertificateErrorDetectedEventArgs struct {
	TObject
}

func NewCoreWebView2ServerCertificateErrorDetectedEventArgs(aArgs ICoreWebView2ServerCertificateErrorDetectedEventArgs) ICoreWebView2ServerCertificateErrorDetectedEventArgs {
	r1 := WV().SysCallN(605, GetObjectUintptr(aArgs))
	return AsCoreWebView2ServerCertificateErrorDetectedEventArgs(r1)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(608, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) BaseIntf() ICoreWebView2ServerCertificateErrorDetectedEventArgs {
	var resultCoreWebView2ServerCertificateErrorDetectedEventArgs uintptr
	WV().SysCallN(603, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ServerCertificateErrorDetectedEventArgs)))
	return AsCoreWebView2ServerCertificateErrorDetectedEventArgs(resultCoreWebView2ServerCertificateErrorDetectedEventArgs)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) ErrorStatus() TWVWebErrorStatus {
	r1 := WV().SysCallN(607, m.Instance())
	return TWVWebErrorStatus(r1)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) RequestUri() string {
	r1 := WV().SysCallN(609, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) ServerCertificate() ICoreWebView2Certificate {
	var resultCoreWebView2Certificate uintptr
	WV().SysCallN(610, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Certificate)))
	return AsCoreWebView2Certificate(resultCoreWebView2Certificate)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Action() TWVServerCertificateErrorAction {
	r1 := WV().SysCallN(602, 0, m.Instance(), 0)
	return TWVServerCertificateErrorAction(r1)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) SetAction(AValue TWVServerCertificateErrorAction) {
	WV().SysCallN(602, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(606, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func CoreWebView2ServerCertificateErrorDetectedEventArgsClass() TClass {
	ret := WV().SysCallN(604)
	return TClass(ret)
}
