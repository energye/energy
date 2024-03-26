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

// ICoreWebView2NavigationStartingEventArgs Parent: IObject
//
//	Event args for the NavigationStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs">See the ICoreWebView2NavigationStartingEventArgs article.</a>
type ICoreWebView2NavigationStartingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NavigationStartingEventArgs // property
	// URI
	//  The uri of the requested navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_uri">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	URI() string // property
	// IsUserInitiated
	//  `TRUE` when the navigation was initiated through a user gesture as
	//  opposed to programmatic navigation by page script. Navigations initiated
	//  via WebView2 APIs are considered as user initiated.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_isuserinitiated">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	IsUserInitiated() bool // property
	// IsRedirected
	//  `TRUE` when the navigation is redirected.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_isredirected">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	IsRedirected() bool // property
	// Cancel
	//  The host may set this flag to cancel the navigation. If set, the
	//  navigation is not longer present and the content of the current page is
	//  intact. For performance reasons, `GET` HTTP requests may happen, while
	//  the host is responding. You may set cookies and use part of a request
	//  for the navigation. Navigations to about schemes are cancellable, unless
	//  `msWebView2CancellableAboutNavigations` feature flag is disabled.
	//  Cancellation of frame navigation to `srcdoc` is not supported and
	//  wil be ignored. A cancelled navigation will fire a `NavigationCompleted`
	//  event with a `WebErrorStatus` of
	//  `COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_cancel">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// NavigationID
	//  The ID of the navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_navigationid">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	NavigationID() uint64 // property
	// RequestHeaders
	//  The HTTP request headers for the navigation.
	//  NOTE: You are not able to modify the HTTP request headers in a
	//  `NavigationStarting` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_requestheaders">See the ICoreWebView2NavigationStartingEventArgs article.</a>
	RequestHeaders() ICoreWebView2HttpRequestHeaders // property
	// AdditionalAllowedFrameAncestors
	//  Get additional allowed frame ancestors set by the host app.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs2#get_additionalallowedframeancestors">See the ICoreWebView2NavigationStartingEventArgs3 article.</a>
	AdditionalAllowedFrameAncestors() string // property
	// SetAdditionalAllowedFrameAncestors Set AdditionalAllowedFrameAncestors
	SetAdditionalAllowedFrameAncestors(AValue string) // property
	// NavigationKind
	//  Get the navigation kind of this navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3#get_navigationkind">See the ICoreWebView2NavigationStartingEventArgs3 article.</a>
	NavigationKind() TWVNavigationKind // property
}

// TCoreWebView2NavigationStartingEventArgs Parent: TObject
//
//	Event args for the NavigationStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs">See the ICoreWebView2NavigationStartingEventArgs article.</a>
type TCoreWebView2NavigationStartingEventArgs struct {
	TObject
}

func NewCoreWebView2NavigationStartingEventArgs(aArgs ICoreWebView2NavigationStartingEventArgs) ICoreWebView2NavigationStartingEventArgs {
	r1 := WV().SysCallN(409, GetObjectUintptr(aArgs))
	return AsCoreWebView2NavigationStartingEventArgs(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) Initialized() bool {
	r1 := WV().SysCallN(410, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) BaseIntf() ICoreWebView2NavigationStartingEventArgs {
	var resultCoreWebView2NavigationStartingEventArgs uintptr
	WV().SysCallN(406, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2NavigationStartingEventArgs)))
	return AsCoreWebView2NavigationStartingEventArgs(resultCoreWebView2NavigationStartingEventArgs)
}

func (m *TCoreWebView2NavigationStartingEventArgs) URI() string {
	r1 := WV().SysCallN(416, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) IsUserInitiated() bool {
	r1 := WV().SysCallN(412, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) IsRedirected() bool {
	r1 := WV().SysCallN(411, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) Cancel() bool {
	r1 := WV().SysCallN(407, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) SetCancel(AValue bool) {
	WV().SysCallN(407, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2NavigationStartingEventArgs) NavigationID() uint64 {
	r1 := WV().SysCallN(413, m.Instance())
	return uint64(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) RequestHeaders() ICoreWebView2HttpRequestHeaders {
	var resultCoreWebView2HttpRequestHeaders uintptr
	WV().SysCallN(415, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpRequestHeaders)))
	return AsCoreWebView2HttpRequestHeaders(resultCoreWebView2HttpRequestHeaders)
}

func (m *TCoreWebView2NavigationStartingEventArgs) AdditionalAllowedFrameAncestors() string {
	r1 := WV().SysCallN(405, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2NavigationStartingEventArgs) SetAdditionalAllowedFrameAncestors(AValue string) {
	WV().SysCallN(405, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2NavigationStartingEventArgs) NavigationKind() TWVNavigationKind {
	r1 := WV().SysCallN(414, m.Instance())
	return TWVNavigationKind(r1)
}

func CoreWebView2NavigationStartingEventArgsClass() TClass {
	ret := WV().SysCallN(408)
	return TClass(ret)
}
