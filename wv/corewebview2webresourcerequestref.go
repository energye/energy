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

// ICoreWebView2WebResourceRequestRef Parent: IObject
//
//	An HTTP request used with the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest">See the ICoreWebView2WebResourceRequest article.</a>
type ICoreWebView2WebResourceRequestRef interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceRequestRef // property
	// URI
	//  The request URI.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_uri">See the ICoreWebView2WebResourceRequest article.</a>
	URI() string // property
	// SetURI Set URI
	SetURI(AValue string) // property
	// Method
	//  The HTTP request method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_method">See the ICoreWebView2WebResourceRequest article.</a>
	Method() string // property
	// SetMethod Set Method
	SetMethod(AValue string) // property
	// Content
	//  The HTTP request message body as stream. POST data should be here. If a
	//  stream is set, which overrides the message body, the stream must have
	//  all the content data available by the time the `WebResourceRequested`
	//  event deferral of this response is completed. Stream should be agile or
	//  be created from a background STA to prevent performance impact to the UI
	//  thread. `Null` means no content data. `IStream` semantics apply
	// (return `S_OK` to `Read` runs until all data is exhausted).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_content">See the ICoreWebView2WebResourceRequest article.</a>
	Content() IStream // property
	// SetContent Set Content
	SetContent(AValue IStream) // property
	// Headers
	//  The mutable HTTP request headers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_headers">See the ICoreWebView2WebResourceRequest article.</a>
	Headers() ICoreWebView2HttpRequestHeaders // property
}

// TCoreWebView2WebResourceRequestRef Parent: TObject
//
//	An HTTP request used with the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest">See the ICoreWebView2WebResourceRequest article.</a>
type TCoreWebView2WebResourceRequestRef struct {
	TObject
}

func NewCoreWebView2WebResourceRequestRef(aBaseIntf ICoreWebView2WebResourceRequestRef) ICoreWebView2WebResourceRequestRef {
	r1 := WV().SysCallN(667, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2WebResourceRequestRef(r1)
}

func (m *TCoreWebView2WebResourceRequestRef) Initialized() bool {
	r1 := WV().SysCallN(669, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebResourceRequestRef) BaseIntf() ICoreWebView2WebResourceRequestRef {
	var resultCoreWebView2WebResourceRequest uintptr
	WV().SysCallN(664, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceRequest)))
	return AsCoreWebView2WebResourceRequest(resultCoreWebView2WebResourceRequest)
}

func (m *TCoreWebView2WebResourceRequestRef) URI() string {
	r1 := WV().SysCallN(671, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2WebResourceRequestRef) SetURI(AValue string) {
	WV().SysCallN(671, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2WebResourceRequestRef) Method() string {
	r1 := WV().SysCallN(670, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2WebResourceRequestRef) SetMethod(AValue string) {
	WV().SysCallN(670, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2WebResourceRequestRef) Content() IStream {
	var resultStream uintptr
	WV().SysCallN(666, 0, m.Instance(), 0, uintptr(unsafePointer(&resultStream)))
	return AsStream(resultStream)
}

func (m *TCoreWebView2WebResourceRequestRef) SetContent(AValue IStream) {
	WV().SysCallN(666, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2WebResourceRequestRef) Headers() ICoreWebView2HttpRequestHeaders {
	var resultCoreWebView2HttpRequestHeaders uintptr
	WV().SysCallN(668, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpRequestHeaders)))
	return AsCoreWebView2HttpRequestHeaders(resultCoreWebView2HttpRequestHeaders)
}

func CoreWebView2WebResourceRequestRefClass() TClass {
	ret := WV().SysCallN(665)
	return TClass(ret)
}
