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

// ICoreWebView2BasicAuthenticationResponse Parent: IObject
//
//	Represents a Basic HTTP authentication response that contains a user name
//	and a password as according to RFC7617 (https://tools.ietf.org/html/rfc7617)
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse">See the ICoreWebView2BasicAuthenticationResponse article.</a>
type ICoreWebView2BasicAuthenticationResponse interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BasicAuthenticationResponse // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2BasicAuthenticationResponse) // property
	// UserName
	//  User name provided for authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_username">See the ICoreWebView2BasicAuthenticationResponse article.</a>
	UserName() string // property
	// SetUserName Set UserName
	SetUserName(AValue string) // property
	// Password
	//  Password provided for authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_password">See the ICoreWebView2BasicAuthenticationResponse article.</a>
	Password() string // property
	// SetPassword Set Password
	SetPassword(AValue string) // property
}

// TCoreWebView2BasicAuthenticationResponse Parent: TObject
//
//	Represents a Basic HTTP authentication response that contains a user name
//	and a password as according to RFC7617 (https://tools.ietf.org/html/rfc7617)
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse">See the ICoreWebView2BasicAuthenticationResponse article.</a>
type TCoreWebView2BasicAuthenticationResponse struct {
	TObject
}

func NewCoreWebView2BasicAuthenticationResponse(aBaseIntf ICoreWebView2BasicAuthenticationResponse) ICoreWebView2BasicAuthenticationResponse {
	r1 := WV().SysCallN(26, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2BasicAuthenticationResponse(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) Initialized() bool {
	r1 := WV().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) BaseIntf() ICoreWebView2BasicAuthenticationResponse {
	var resultCoreWebView2BasicAuthenticationResponse uintptr
	WV().SysCallN(24, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2BasicAuthenticationResponse)))
	return AsCoreWebView2BasicAuthenticationResponse(resultCoreWebView2BasicAuthenticationResponse)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetBaseIntf(AValue ICoreWebView2BasicAuthenticationResponse) {
	WV().SysCallN(24, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2BasicAuthenticationResponse) UserName() string {
	r1 := WV().SysCallN(29, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetUserName(AValue string) {
	WV().SysCallN(29, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2BasicAuthenticationResponse) Password() string {
	r1 := WV().SysCallN(28, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetPassword(AValue string) {
	WV().SysCallN(28, 1, m.Instance(), PascalStr(AValue))
}

func CoreWebView2BasicAuthenticationResponseClass() TClass {
	ret := WV().SysCallN(25)
	return TClass(ret)
}
