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

// ICoreWebView2WebResourceResponse Parent: IObject
//
//	An HTTP response used with the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse">See the ICoreWebView2WebResourceResponse article.</a>
type ICoreWebView2WebResourceResponse interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponse // property
	// StatusCode
	//  The HTTP response status code.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_statuscode">See the ICoreWebView2WebResourceResponse article.</a>
	StatusCode() int32 // property
	// SetStatusCode Set StatusCode
	SetStatusCode(AValue int32) // property
	// ReasonPhrase
	//  The HTTP response reason phrase.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_reasonphrase">See the ICoreWebView2WebResourceResponse article.</a>
	ReasonPhrase() string // property
	// SetReasonPhrase Set ReasonPhrase
	SetReasonPhrase(AValue string) // property
	// Content
	//  HTTP response content as stream. Stream must have all the content data
	//  available by the time the `WebResourceRequested` event deferral of this
	//  response is completed. Stream should be agile or be created from a
	//  background thread to prevent performance impact to the UI thread. `Null`
	//  means no content data. `IStream` semantics apply(return `S_OK` to
	//  `Read` runs until all data is exhausted).
	//  When providing the response data, you should consider relevant HTTP
	//  request headers just like an HTTP server would do. For example, if the
	//  request was for a video resource in a HTML video element, the request may
	//  contain the [Range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range)
	//  header to request only a part of the video that is streaming. In this
	//  case, your response stream should be only the portion of the video
	//  specified by the range HTTP request headers and you should set the
	//  appropriate
	//  [Content-Range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range)
	//  header in the response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_content">See the ICoreWebView2WebResourceResponse article.</a>
	Content() IStream // property
	// SetContent Set Content
	SetContent(AValue IStream) // property
	// Headers
	//  Overridden HTTP response headers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_headers">See the ICoreWebView2WebResourceResponse article.</a>
	Headers() ICoreWebView2HttpResponseHeaders // property
}

// TCoreWebView2WebResourceResponse Parent: TObject
//
//	An HTTP response used with the WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse">See the ICoreWebView2WebResourceResponse article.</a>
type TCoreWebView2WebResourceResponse struct {
	TObject
}

func NewCoreWebView2WebResourceResponse(aBaseIntf ICoreWebView2WebResourceResponse) ICoreWebView2WebResourceResponse {
	r1 := WV().SysCallN(700, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2WebResourceResponse(r1)
}

func (m *TCoreWebView2WebResourceResponse) Initialized() bool {
	r1 := WV().SysCallN(702, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebResourceResponse) BaseIntf() ICoreWebView2WebResourceResponse {
	var resultCoreWebView2WebResourceResponse uintptr
	WV().SysCallN(697, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceResponse)))
	return AsCoreWebView2WebResourceResponse(resultCoreWebView2WebResourceResponse)
}

func (m *TCoreWebView2WebResourceResponse) StatusCode() int32 {
	r1 := WV().SysCallN(704, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2WebResourceResponse) SetStatusCode(AValue int32) {
	WV().SysCallN(704, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2WebResourceResponse) ReasonPhrase() string {
	r1 := WV().SysCallN(703, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2WebResourceResponse) SetReasonPhrase(AValue string) {
	WV().SysCallN(703, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2WebResourceResponse) Content() IStream {
	var resultStream uintptr
	WV().SysCallN(699, 0, m.Instance(), 0, uintptr(unsafePointer(&resultStream)))
	return AsStream(resultStream)
}

func (m *TCoreWebView2WebResourceResponse) SetContent(AValue IStream) {
	WV().SysCallN(699, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2WebResourceResponse) Headers() ICoreWebView2HttpResponseHeaders {
	var resultCoreWebView2HttpResponseHeaders uintptr
	WV().SysCallN(701, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpResponseHeaders)))
	return AsCoreWebView2HttpResponseHeaders(resultCoreWebView2HttpResponseHeaders)
}

func CoreWebView2WebResourceResponseClass() TClass {
	ret := WV().SysCallN(698)
	return TClass(ret)
}
