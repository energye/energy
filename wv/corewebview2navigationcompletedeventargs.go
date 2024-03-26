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

// ICoreWebView2NavigationCompletedEventArgs Parent: IObject
//
//	Event args for the NavigationCompleted event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs">See the ICoreWebView2NavigationCompletedEventArgs article.</a>
type ICoreWebView2NavigationCompletedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NavigationCompletedEventArgs // property
	// IsSuccess
	//  `TRUE` when the navigation is successful. `FALSE` for a navigation that
	//  ended up in an error page(failures due to no network, DNS lookup
	//  failure, HTTP server responds with 4xx), but may also be `FALSE` for
	//  additional scenarios such as `window.stop()` run on navigated page.
	//  Note that WebView2 will report the navigation as 'unsuccessful' if the load
	//  for the navigation did not reach the expected completion for any reason. Such
	//  reasons include potentially catastrophic issues such network and certificate
	//  issues, but can also be the result of intended actions such as the app canceling a navigation or
	//  navigating away before the original navigation completed. Applications should not
	//  just rely on this flag, but also consider the reported WebErrorStatus to
	//  determine whether the failure is indeed catastrophic in their context.
	//  WebErrorStatuses that may indicate a non-catastrophic failure include:
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_issuccess">See the ICoreWebView2NavigationCompletedEventArgs article.</a>
	IsSuccess() bool // property
	// WebErrorStatus
	//  The error code if the navigation failed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_weberrorstatus">See the ICoreWebView2NavigationCompletedEventArgs article.</a>
	WebErrorStatus() TWVWebErrorStatus // property
	// NavigationID
	//  The ID of the navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_navigationid">See the ICoreWebView2NavigationCompletedEventArgs article.</a>
	NavigationID() uint64 // property
	// HttpStatusCode
	//  The HTTP status code of the navigation if it involved an HTTP request.
	//  For instance, this will usually be 200 if the request was successful, 404
	//  if a page was not found, etc. See
	//  https://developer.mozilla.org/docs/Web/HTTP/Status for a list of
	//  common status codes.
	//  The `HttpStatusCode` property will be 0 in the following cases:
	//  * The navigation did not involve an HTTP request. For instance, if it was
	//  a navigation to a file:// URL, or if it was a same-document navigation.
	//  * The navigation failed before a response was received. For instance, if
	//  the hostname was not found, or if there was a network error.
	//  In those cases, you can get more information from the `IsSuccess` and
	//  `WebErrorStatus` properties.
	//  If the navigation receives a successful HTTP response, but the navigated
	//  page calls `window.stop()` before it finishes loading, then
	//  `HttpStatusCode` may contain a success code like 200, but `IsSuccess` will
	//  be FALSE and `WebErrorStatus` will be
	//  `COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED`.
	//  Since WebView2 handles HTTP continuations and redirects automatically, it
	//  is unlikely for `HttpStatusCode` to ever be in the 1xx or 3xx ranges.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs2#get_httpstatuscode">See the ICoreWebView2NavigationCompletedEventArgs2 article.</a>
	HttpStatusCode() int32 // property
}

// TCoreWebView2NavigationCompletedEventArgs Parent: TObject
//
//	Event args for the NavigationCompleted event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs">See the ICoreWebView2NavigationCompletedEventArgs article.</a>
type TCoreWebView2NavigationCompletedEventArgs struct {
	TObject
}

func NewCoreWebView2NavigationCompletedEventArgs(aArgs ICoreWebView2NavigationCompletedEventArgs) ICoreWebView2NavigationCompletedEventArgs {
	r1 := WV().SysCallN(399, GetObjectUintptr(aArgs))
	return AsCoreWebView2NavigationCompletedEventArgs(r1)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(401, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) BaseIntf() ICoreWebView2NavigationCompletedEventArgs {
	var resultCoreWebView2NavigationCompletedEventArgs uintptr
	WV().SysCallN(397, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2NavigationCompletedEventArgs)))
	return AsCoreWebView2NavigationCompletedEventArgs(resultCoreWebView2NavigationCompletedEventArgs)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) IsSuccess() bool {
	r1 := WV().SysCallN(402, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) WebErrorStatus() TWVWebErrorStatus {
	r1 := WV().SysCallN(404, m.Instance())
	return TWVWebErrorStatus(r1)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) NavigationID() uint64 {
	r1 := WV().SysCallN(403, m.Instance())
	return uint64(r1)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) HttpStatusCode() int32 {
	r1 := WV().SysCallN(400, m.Instance())
	return int32(r1)
}

func CoreWebView2NavigationCompletedEventArgsClass() TClass {
	ret := WV().SysCallN(398)
	return TClass(ret)
}
