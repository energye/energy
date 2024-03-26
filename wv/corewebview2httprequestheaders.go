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

// ICoreWebView2HttpRequestHeaders Parent: IObject
//
//	HTTP request headers.  Used to inspect the HTTP request on
//	WebResourceRequested event and NavigationStarting event.
//	It is possible to modify the HTTP request from a WebResourceRequested event, but not from a NavigationStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders">See the ICoreWebView2HttpRequestHeaders article.</a>
type ICoreWebView2HttpRequestHeaders interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpRequestHeaders // property
	// Iterator
	//  Gets an iterator over the collection of request headers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getiterator">See the ICoreWebView2HttpRequestHeaders article.</a>
	Iterator() ICoreWebView2HttpHeadersCollectionIterator // property
	// SetHeader
	//  Adds or updates header that matches the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#setheader">See the ICoreWebView2HttpRequestHeaders article.</a>
	SetHeader(aName, aValue string) bool // function
	// GetHeader
	//  Gets the header value matching the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getheader">See the ICoreWebView2HttpRequestHeaders article.</a>
	GetHeader(aName string) string // function
	// GetHeaders
	//  Gets the header value matching the name using an iterator.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getheaders">See the ICoreWebView2HttpRequestHeaders article.</a>
	GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool // function
	// Contains
	//  Verifies that the headers contain an entry that matches the header name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#contains">See the ICoreWebView2HttpRequestHeaders article.</a>
	Contains(aName string) bool // function
	// RemoveHeader
	//  Removes header that matches the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#removeheader">See the ICoreWebView2HttpRequestHeaders article.</a>
	RemoveHeader(aName string) bool // function
}

// TCoreWebView2HttpRequestHeaders Parent: TObject
//
//	HTTP request headers.  Used to inspect the HTTP request on
//	WebResourceRequested event and NavigationStarting event.
//	It is possible to modify the HTTP request from a WebResourceRequested event, but not from a NavigationStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders">See the ICoreWebView2HttpRequestHeaders article.</a>
type TCoreWebView2HttpRequestHeaders struct {
	TObject
}

func NewCoreWebView2HttpRequestHeaders(aBaseIntf ICoreWebView2HttpRequestHeaders) ICoreWebView2HttpRequestHeaders {
	r1 := WV().SysCallN(366, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2HttpRequestHeaders(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) Initialized() bool {
	r1 := WV().SysCallN(369, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) BaseIntf() ICoreWebView2HttpRequestHeaders {
	var resultCoreWebView2HttpRequestHeaders uintptr
	WV().SysCallN(363, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpRequestHeaders)))
	return AsCoreWebView2HttpRequestHeaders(resultCoreWebView2HttpRequestHeaders)
}

func (m *TCoreWebView2HttpRequestHeaders) Iterator() ICoreWebView2HttpHeadersCollectionIterator {
	var resultCoreWebView2HttpHeadersCollectionIterator uintptr
	WV().SysCallN(370, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpHeadersCollectionIterator)))
	return AsCoreWebView2HttpHeadersCollectionIterator(resultCoreWebView2HttpHeadersCollectionIterator)
}

func (m *TCoreWebView2HttpRequestHeaders) SetHeader(aName, aValue string) bool {
	r1 := WV().SysCallN(372, m.Instance(), PascalStr(aName), PascalStr(aValue))
	return GoBool(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) GetHeader(aName string) string {
	r1 := WV().SysCallN(367, m.Instance(), PascalStr(aName))
	return GoStr(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool {
	var result1 uintptr
	r1 := WV().SysCallN(368, m.Instance(), PascalStr(aName), uintptr(unsafePointer(&result1)))
	*aIterator = AsCoreWebView2HttpHeadersCollectionIterator(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) Contains(aName string) bool {
	r1 := WV().SysCallN(365, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TCoreWebView2HttpRequestHeaders) RemoveHeader(aName string) bool {
	r1 := WV().SysCallN(371, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func CoreWebView2HttpRequestHeadersClass() TClass {
	ret := WV().SysCallN(364)
	return TClass(ret)
}
