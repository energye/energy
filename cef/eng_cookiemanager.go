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

// ICefCookieManager Parent: ICefBaseRefCounted
//
//	Interface used for managing cookies. The functions of this interface may be called on any thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_cookie_manager_t))</a>
type ICefCookieManager interface {
	ICefBaseRefCounted
	// VisitAllCookies
	//  Visit all cookies on the UI thread. The returned cookies are ordered by longest path, then by earliest creation date. Returns false (0) if cookies cannot be accessed.
	VisitAllCookies(visitor ICefCookieVisitor) bool // function
	// VisitUrlCookies
	//  Visit a subset of cookies on the UI thread. The results are filtered by the given url scheme, host, domain and path. If |includeHttpOnly| is true (1) HTTP-only cookies will also be included in the results. The returned cookies are ordered by longest path, then by earliest creation date. Returns false (0) if cookies cannot be accessed.
	VisitUrlCookies(url string, includeHttpOnly bool, visitor ICefCookieVisitor) bool // function
	// SetCookie
	//  Sets a cookie given a valid URL and explicit user-provided cookie attributes. This function expects each attribute to be well-formed. It will check for disallowed characters (e.g. the ';' character is disallowed within the cookie value attribute) and fail without setting the cookie if such characters are found. If |callback| is non-NULL it will be executed asnychronously on the UI thread after the cookie has been set. Returns false (0) if an invalid URL is specified or if cookies cannot be accessed.
	SetCookie(url, name, value, domain, path string, secure, httponly, hasExpires bool, creation, lastAccess, expires TDateTime, samesite TCefCookieSameSite, priority TCefCookiePriority, callback ICefSetCookieCallback) bool // function
	// DeleteCookies
	//  Delete all cookies that match the specified parameters. If both |url| and |cookie_name| values are specified all host and domain cookies matching both will be deleted. If only |url| is specified all host cookies (but not domain cookies) irrespective of path will be deleted. If |url| is NULL all cookies for all hosts and domains will be deleted. If |callback| is non- NULL it will be executed asnychronously on the UI thread after the cookies have been deleted. Returns false (0) if a non-NULL invalid URL is specified or if cookies cannot be accessed. Cookies can alternately be deleted using the Visit*Cookies() functions.
	DeleteCookies(url, cookieName string, callback ICefDeleteCookiesCallback) bool // function
	// FlushStore
	//  Flush the backing store (if any) to disk. If |callback| is non-NULL it will be executed asnychronously on the UI thread after the flush is complete. Returns false (0) if cookies cannot be accessed.
	FlushStore(callback ICefCompletionCallback) bool // function
}

// TCefCookieManager Parent: TCefBaseRefCounted
//
//	Interface used for managing cookies. The functions of this interface may be called on any thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_cookie_manager_t))</a>
type TCefCookieManager struct {
	TCefBaseRefCounted
}

// CookieManagerRef -> ICefCookieManager
var CookieManagerRef cookieManager

// cookieManager TCefCookieManager Ref
type cookieManager uintptr

func (m *cookieManager) UnWrap(data uintptr) ICefCookieManager {
	var resultCefCookieManager uintptr
	CEF().SysCallN(771, uintptr(data), uintptr(unsafePointer(&resultCefCookieManager)))
	return AsCefCookieManager(resultCefCookieManager)
}

func (m *cookieManager) Global(callback ICefCompletionCallback) ICefCookieManager {
	var resultCefCookieManager uintptr
	CEF().SysCallN(769, GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefCookieManager)))
	return AsCefCookieManager(resultCefCookieManager)
}

func (m *TCefCookieManager) VisitAllCookies(visitor ICefCookieVisitor) bool {
	r1 := CEF().SysCallN(772, m.Instance(), GetObjectUintptr(visitor))
	return GoBool(r1)
}

func (m *TCefCookieManager) VisitUrlCookies(url string, includeHttpOnly bool, visitor ICefCookieVisitor) bool {
	r1 := CEF().SysCallN(773, m.Instance(), PascalStr(url), PascalBool(includeHttpOnly), GetObjectUintptr(visitor))
	return GoBool(r1)
}

func (m *TCefCookieManager) SetCookie(url, name, value, domain, path string, secure, httponly, hasExpires bool, creation, lastAccess, expires TDateTime, samesite TCefCookieSameSite, priority TCefCookiePriority, callback ICefSetCookieCallback) bool {
	r1 := CEF().SysCallN(770, m.Instance(), PascalStr(url), PascalStr(name), PascalStr(value), PascalStr(domain), PascalStr(path), PascalBool(secure), PascalBool(httponly), PascalBool(hasExpires), uintptr(unsafePointer(&creation)), uintptr(unsafePointer(&lastAccess)), uintptr(unsafePointer(&expires)), uintptr(samesite), uintptr(priority), GetObjectUintptr(callback))
	return GoBool(r1)
}

func (m *TCefCookieManager) DeleteCookies(url, cookieName string, callback ICefDeleteCookiesCallback) bool {
	r1 := CEF().SysCallN(767, m.Instance(), PascalStr(url), PascalStr(cookieName), GetObjectUintptr(callback))
	return GoBool(r1)
}

func (m *TCefCookieManager) FlushStore(callback ICefCompletionCallback) bool {
	r1 := CEF().SysCallN(768, m.Instance(), GetObjectUintptr(callback))
	return GoBool(r1)
}
