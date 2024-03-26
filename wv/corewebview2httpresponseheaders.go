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

// ICoreWebView2HttpResponseHeaders Parent: IObject
//
//	HTTP response headers.  Used to construct a WebResourceResponse for the
//	WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
type ICoreWebView2HttpResponseHeaders interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpResponseHeaders // property
	// Iterator
	//  Gets an iterator over the collection of entire response headers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getiterator">See the ICoreWebView2HttpResponseHeaders article.</a>
	Iterator() ICoreWebView2HttpHeadersCollectionIterator // property
	// GetHeader
	//  Gets the first header value in the collection matching the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheader">See the ICoreWebView2HttpResponseHeaders article.</a>
	GetHeader(aName string) string // function
	// GetHeaders
	//  Gets the header values matching the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
	GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool // function
	// Contains
	//  Verifies that the headers contain entries that match the header name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#contains">See the ICoreWebView2HttpResponseHeaders article.</a>
	Contains(aName string) bool // function
	// AppendHeader
	//  Appends header line with name and value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#appendheader">See the ICoreWebView2HttpResponseHeaders article.</a>
	AppendHeader(aName, aValue string) bool // function
}

// TCoreWebView2HttpResponseHeaders Parent: TObject
//
//	HTTP response headers.  Used to construct a WebResourceResponse for the
//	WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
type TCoreWebView2HttpResponseHeaders struct {
	TObject
}

func NewCoreWebView2HttpResponseHeaders(aBaseIntf ICoreWebView2HttpResponseHeaders) ICoreWebView2HttpResponseHeaders {
	r1 := WV().SysCallN(377, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2HttpResponseHeaders(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) Initialized() bool {
	r1 := WV().SysCallN(380, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) BaseIntf() ICoreWebView2HttpResponseHeaders {
	var resultCoreWebView2HttpResponseHeaders uintptr
	WV().SysCallN(374, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpResponseHeaders)))
	return AsCoreWebView2HttpResponseHeaders(resultCoreWebView2HttpResponseHeaders)
}

func (m *TCoreWebView2HttpResponseHeaders) Iterator() ICoreWebView2HttpHeadersCollectionIterator {
	var resultCoreWebView2HttpHeadersCollectionIterator uintptr
	WV().SysCallN(381, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpHeadersCollectionIterator)))
	return AsCoreWebView2HttpHeadersCollectionIterator(resultCoreWebView2HttpHeadersCollectionIterator)
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeader(aName string) string {
	r1 := WV().SysCallN(378, m.Instance(), PascalStr(aName))
	return GoStr(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool {
	var result1 uintptr
	r1 := WV().SysCallN(379, m.Instance(), PascalStr(aName), uintptr(unsafePointer(&result1)))
	*aIterator = AsCoreWebView2HttpHeadersCollectionIterator(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) Contains(aName string) bool {
	r1 := WV().SysCallN(376, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) AppendHeader(aName, aValue string) bool {
	r1 := WV().SysCallN(373, m.Instance(), PascalStr(aName), PascalStr(aValue))
	return GoBool(r1)
}

func CoreWebView2HttpResponseHeadersClass() TClass {
	ret := WV().SysCallN(375)
	return TClass(ret)
}
