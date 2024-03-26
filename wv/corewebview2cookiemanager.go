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

// ICoreWebView2CookieManager Parent: IObject
//
//	Creates, adds or updates, gets, or or view the cookies. The changes would
//	apply to the context of the user profile. That is, other WebViews under the
//	same user profile could be affected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager">See the ICoreWebView2CookieManager article.</a>
type ICoreWebView2CookieManager interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CookieManager // property
	// CreateCookie
	//  Create a cookie object with a specified name, value, domain, and path.
	//  One can set other optional properties after cookie creation.
	//  This only creates a cookie object and it is not added to the cookie
	//  manager until you call AddOrUpdateCookie.
	//  Leading or trailing whitespace(s), empty string, and special characters
	//  are not allowed for name.
	//  See ICoreWebView2Cookie for more details.
	CreateCookie(aName, aValue, aDomain, aPath string) ICoreWebView2Cookie // function
	// CopyCookie
	//  Creates a cookie whose params matches those of the specified cookie.
	CopyCookie(aCookie ICoreWebView2Cookie) ICoreWebView2Cookie // function
	// GetCookies
	//  Gets a list of cookies matching the specific URI.
	//  If uri is empty string or null, all cookies under the same profile are
	//  returned.
	//  You can modify the cookie objects by calling
	//  ICoreWebView2CookieManager::AddOrUpdateCookie, and the changes
	//  will be applied to the webview.
	GetCookies(aURI string, aHandler ICoreWebView2GetCookiesCompletedHandler) bool // function
	// AddOrUpdateCookie
	//  Adds or updates a cookie with the given cookie data; may overwrite
	//  cookies with matching name, domain, and path if they exist.
	//  This method will fail if the domain of the given cookie is not specified.
	AddOrUpdateCookie(aCookie ICoreWebView2Cookie) bool // function
	// DeleteCookie
	//  Deletes a cookie whose name and domain/path pair
	//  match those of the specified cookie.
	DeleteCookie(aCookie ICoreWebView2Cookie) bool // function
	// DeleteCookies
	//  Deletes cookies with matching name and uri.
	//  Cookie name is required.
	//  All cookies with the given name where domain
	//  and path match provided URI are deleted.
	DeleteCookies(aName, aURI string) bool // function
	// DeleteCookiesWithDomainAndPath
	//  Deletes cookies with matching name and domain/path pair.
	//  Cookie name is required.
	//  If domain is specified, deletes only cookies with the exact domain.
	//  If path is specified, deletes only cookies with the exact path.
	DeleteCookiesWithDomainAndPath(aName, aDomain, aPath string) bool // function
	// DeleteAllCookies
	//  Deletes all cookies under the same profile.
	//  This could affect other WebViews under the same user profile.
	DeleteAllCookies() bool // function
}

// TCoreWebView2CookieManager Parent: TObject
//
//	Creates, adds or updates, gets, or or view the cookies. The changes would
//	apply to the context of the user profile. That is, other WebViews under the
//	same user profile could be affected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager">See the ICoreWebView2CookieManager article.</a>
type TCoreWebView2CookieManager struct {
	TObject
}

func NewCoreWebView2CookieManager(aBaseIntf ICoreWebView2CookieManager) ICoreWebView2CookieManager {
	r1 := WV().SysCallN(211, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2CookieManager(r1)
}

func (m *TCoreWebView2CookieManager) Initialized() bool {
	r1 := WV().SysCallN(218, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) BaseIntf() ICoreWebView2CookieManager {
	var resultCoreWebView2CookieManager uintptr
	WV().SysCallN(208, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieManager)))
	return AsCoreWebView2CookieManager(resultCoreWebView2CookieManager)
}

func (m *TCoreWebView2CookieManager) CreateCookie(aName, aValue, aDomain, aPath string) ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(212, m.Instance(), PascalStr(aName), PascalStr(aValue), PascalStr(aDomain), PascalStr(aPath), uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TCoreWebView2CookieManager) CopyCookie(aCookie ICoreWebView2Cookie) ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(210, m.Instance(), GetObjectUintptr(aCookie), uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TCoreWebView2CookieManager) GetCookies(aURI string, aHandler ICoreWebView2GetCookiesCompletedHandler) bool {
	r1 := WV().SysCallN(217, m.Instance(), PascalStr(aURI), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) AddOrUpdateCookie(aCookie ICoreWebView2Cookie) bool {
	r1 := WV().SysCallN(207, m.Instance(), GetObjectUintptr(aCookie))
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) DeleteCookie(aCookie ICoreWebView2Cookie) bool {
	r1 := WV().SysCallN(214, m.Instance(), GetObjectUintptr(aCookie))
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) DeleteCookies(aName, aURI string) bool {
	r1 := WV().SysCallN(215, m.Instance(), PascalStr(aName), PascalStr(aURI))
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) DeleteCookiesWithDomainAndPath(aName, aDomain, aPath string) bool {
	r1 := WV().SysCallN(216, m.Instance(), PascalStr(aName), PascalStr(aDomain), PascalStr(aPath))
	return GoBool(r1)
}

func (m *TCoreWebView2CookieManager) DeleteAllCookies() bool {
	r1 := WV().SysCallN(213, m.Instance())
	return GoBool(r1)
}

func CoreWebView2CookieManagerClass() TClass {
	ret := WV().SysCallN(209)
	return TClass(ret)
}
