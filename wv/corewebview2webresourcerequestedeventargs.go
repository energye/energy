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

// ICoreWebView2WebResourceRequestedEventArgs Parent: IObject
//
//	Event args for the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
type ICoreWebView2WebResourceRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceRequestedEventArgs // property
	// Request
	//  The Web resource request. The request object may be missing some headers
	//  that are added by network stack at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_request">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
	Request() ICoreWebView2WebResourceRequestRef // property
	// Response
	//  A placeholder for the web resource response object. If this object is
	//  set, the web resource request is completed with the specified response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_response">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
	Response() ICoreWebView2WebResourceResponse // property
	// SetResponse Set Response
	SetResponse(AValue ICoreWebView2WebResourceResponse) // property
	// Deferral
	//  Obtain an `ICoreWebView2Deferral` object and put the event into a
	//  deferred state. Use the `ICoreWebView2Deferral` object to complete the
	//  request at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#getdeferral">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
	// ResourceContext
	//  The web resource request context.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_resourcecontext">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
	ResourceContext() TWVWebResourceContext // property
}

// TCoreWebView2WebResourceRequestedEventArgs Parent: TObject
//
//	Event args for the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs">See the ICoreWebView2WebResourceRequestedEventArgs article.</a>
type TCoreWebView2WebResourceRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2WebResourceRequestedEventArgs(aArgs ICoreWebView2WebResourceRequestedEventArgs) ICoreWebView2WebResourceRequestedEventArgs {
	r1 := WV().SysCallN(674, GetObjectUintptr(aArgs))
	return AsCoreWebView2WebResourceRequestedEventArgs(r1)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(676, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) BaseIntf() ICoreWebView2WebResourceRequestedEventArgs {
	var resultCoreWebView2WebResourceRequestedEventArgs uintptr
	WV().SysCallN(672, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceRequestedEventArgs)))
	return AsCoreWebView2WebResourceRequestedEventArgs(resultCoreWebView2WebResourceRequestedEventArgs)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Request() ICoreWebView2WebResourceRequestRef {
	var resultCoreWebView2WebResourceRequest uintptr
	WV().SysCallN(677, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceRequest)))
	return AsCoreWebView2WebResourceRequest(resultCoreWebView2WebResourceRequest)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Response() ICoreWebView2WebResourceResponse {
	var resultCoreWebView2WebResourceResponse uintptr
	WV().SysCallN(679, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2WebResourceResponse)))
	return AsCoreWebView2WebResourceResponse(resultCoreWebView2WebResourceResponse)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) SetResponse(AValue ICoreWebView2WebResourceResponse) {
	WV().SysCallN(679, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(675, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) ResourceContext() TWVWebResourceContext {
	r1 := WV().SysCallN(678, m.Instance())
	return TWVWebResourceContext(r1)
}

func CoreWebView2WebResourceRequestedEventArgsClass() TClass {
	ret := WV().SysCallN(673)
	return TClass(ret)
}
