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

// ICoreWebView2Cookie Parent: IObject
//
//	Provides a set of properties that are used to manage an
//	ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie">See the ICoreWebView2Cookie article.</a>
type ICoreWebView2Cookie interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Cookie // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2Cookie) // property
	// Name
	//  Cookie name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_name">See the ICoreWebView2Cookie article.</a>
	Name() string // property
	// Value
	//  Cookie value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_value">See the ICoreWebView2Cookie article.</a>
	Value() string // property
	// SetValue Set Value
	SetValue(AValue string) // property
	// Domain
	//  The domain for which the cookie is valid.
	//  The default is the host that this cookie has been received from.
	//  Note that, for instance, ".bing.com", "bing.com", and "www.bing.com" are
	//  considered different domains.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_domain">See the ICoreWebView2Cookie article.</a>
	Domain() string // property
	// Path
	//  The path for which the cookie is valid. The default is "/", which means
	//  this cookie will be sent to all pages on the Domain.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_path">See the ICoreWebView2Cookie article.</a>
	Path() string // property
	// Expires
	//  The expiration date and time for the cookie as the number of seconds since the UNIX epoch.
	//  The default is -1.0, which means cookies are session cookies by default.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</a>
	Expires() (resultDouble float64) // property
	// SetExpires Set Expires
	SetExpires(AValue float64) // property
	// ExpiresDate
	//  The expiration date and time for the cookie in TDateTime format.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</a>
	ExpiresDate() TDateTime // property
	// SetExpiresDate Set ExpiresDate
	SetExpiresDate(AValue TDateTime) // property
	// IsHttpOnly
	//  Whether this cookie is http-only.
	//  True if a page script or other active content cannot access this
	//  cookie. The default is false.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_ishttponly">See the ICoreWebView2Cookie article.</a>
	IsHttpOnly() bool // property
	// SetIsHttpOnly Set IsHttpOnly
	SetIsHttpOnly(AValue bool) // property
	// SameSite
	//  SameSite status of the cookie which represents the enforcement mode of the cookie.
	//  The default is COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_samesite">See the ICoreWebView2Cookie article.</a>
	SameSite() TWVCookieSameSiteKind // property
	// SetSameSite Set SameSite
	SetSameSite(AValue TWVCookieSameSiteKind) // property
	// IsSecure
	//  The security level of this cookie. True if the client is only to return
	//  the cookie in subsequent requests if those requests use HTTPS.
	//  The default is false.
	//  Note that cookie that requests COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE but
	//  is not marked Secure will be rejected.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issecure">See the ICoreWebView2Cookie article.</a>
	IsSecure() bool // property
	// SetIsSecure Set IsSecure
	SetIsSecure(AValue bool) // property
	// IsSession
	//  Whether this is a session cookie. The default is false.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issession">See the ICoreWebView2Cookie article.</a>
	IsSession() bool // property
}

// TCoreWebView2Cookie Parent: TObject
//
//	Provides a set of properties that are used to manage an
//	ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie">See the ICoreWebView2Cookie article.</a>
type TCoreWebView2Cookie struct {
	TObject
}

func NewCoreWebView2Cookie(aBaseIntf ICoreWebView2Cookie) ICoreWebView2Cookie {
	r1 := WV().SysCallN(221, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Cookie(r1)
}

func (m *TCoreWebView2Cookie) Initialized() bool {
	r1 := WV().SysCallN(225, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) BaseIntf() ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(219, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TCoreWebView2Cookie) SetBaseIntf(AValue ICoreWebView2Cookie) {
	WV().SysCallN(219, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2Cookie) Name() string {
	r1 := WV().SysCallN(229, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Value() string {
	r1 := WV().SysCallN(232, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) SetValue(AValue string) {
	WV().SysCallN(232, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2Cookie) Domain() string {
	r1 := WV().SysCallN(222, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Path() string {
	r1 := WV().SysCallN(230, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Expires() (resultDouble float64) {
	WV().SysCallN(223, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2Cookie) SetExpires(AValue float64) {
	WV().SysCallN(223, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2Cookie) ExpiresDate() TDateTime {
	r1 := WV().SysCallN(224, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCoreWebView2Cookie) SetExpiresDate(AValue TDateTime) {
	WV().SysCallN(224, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Cookie) IsHttpOnly() bool {
	r1 := WV().SysCallN(226, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) SetIsHttpOnly(AValue bool) {
	WV().SysCallN(226, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Cookie) SameSite() TWVCookieSameSiteKind {
	r1 := WV().SysCallN(231, 0, m.Instance(), 0)
	return TWVCookieSameSiteKind(r1)
}

func (m *TCoreWebView2Cookie) SetSameSite(AValue TWVCookieSameSiteKind) {
	WV().SysCallN(231, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Cookie) IsSecure() bool {
	r1 := WV().SysCallN(227, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) SetIsSecure(AValue bool) {
	WV().SysCallN(227, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Cookie) IsSession() bool {
	r1 := WV().SysCallN(228, m.Instance())
	return GoBool(r1)
}

func CoreWebView2CookieClass() TClass {
	ret := WV().SysCallN(220)
	return TClass(ret)
}
