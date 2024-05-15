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

// IResourceHandler Parent: ICefResourceHandler
//
//	Interface used to implement a custom request handler interface. The
//	functions of this interface will be called on the IO thread unless otherwise
//	indicated.
//	<a cref="uCEFTypes|TCefResourceHandler">Implements TCefResourceHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_handler_t)</a>
type IResourceHandler interface {
	ICefResourceHandler
	// SetOnOpen
	//  Open the response stream. To handle the request immediately set
	//  |handle_request| to true(1) and return true(1). To decide at a later
	//  time set |handle_request| to false(0), return true(1), and execute
	//  |callback| to continue or cancel the request. To cancel the request
	//  immediately set |handle_request| to true(1) and return false(0). This
	//  function will be called in sequence but not from a dedicated thread. For
	//  backwards compatibility set |handle_request| to false(0) and return false
	//  (0) and the ProcessRequest function will be called.
	SetOnOpen(fn TOnResourceHandlerOpen) // property event
	// SetOnProcessRequest
	//  Begin processing the request. To handle the request return true(1) and
	//  call ICefCallback.cont() once the response header information is
	//  available(ICefCallback.cont() can also be called from inside this
	//  function if header information is available immediately). To cancel the
	//  request return false(0).
	//  WARNING: This function is deprecated. Use Open instead.
	SetOnProcessRequest(fn TOnResourceHandlerProcessRequest) // property event
	// SetOnGetResponseHeaders
	//  Retrieve response header information. If the response length is not known
	//  set |response_length| to -1 and read_response() will be called until it
	//  returns false(0). If the response length is known set |response_length|
	//  to a positive value and read_response() will be called until it returns
	//  false(0) or the specified number of bytes have been read. Use the
	//  |response| object to set the mime type, http status code and other
	//  optional header values. To redirect the request to a new URL set
	//  |redirectUrl| to the new URL. |redirectUrl| can be either a relative or
	//  fully qualified URL. It is also possible to set |response| to a redirect
	//  http status code and pass the new URL via a Location header. Likewise with
	//  |redirectUrl| it is valid to set a relative or fully qualified URL as the
	//  Location header value. If an error occured while setting up the request
	//  you can call set_error() on |response| to indicate the error condition.
	SetOnGetResponseHeaders(fn TOnResourceHandlerGetResponseHeaders) // property event
	// SetOnSkip
	//  Skip response data when requested by a Range header. Skip over and discard
	//  |bytes_to_skip| bytes of response data. If data is available immediately
	//  set |bytes_skipped| to the number of bytes skipped and return true(1). To
	//  read the data at a later time set |bytes_skipped| to 0, return true(1)
	//  and execute |callback| when the data is available. To indicate failure set
	//  |bytes_skipped| to < 0(e.g. -2 for ERR_FAILED) and return false(0). This
	//  function will be called in sequence but not from a dedicated thread.
	SetOnSkip(fn TOnResourceHandlerSkip) // property event
	// SetOnRead
	//  Read response data. If data is available immediately copy up to
	//  |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
	//  bytes copied, and return true(1). To read the data at a later time keep a
	//  pointer to |data_out|, set |bytes_read| to 0, return true(1) and execute
	//  |callback| when the data is available(|data_out| will remain valid until
	//  the callback is executed). To indicate response completion set
	//  |bytes_read| to 0 and return false(0). To indicate failure set
	//  |bytes_read| to < 0(e.g. -2 for ERR_FAILED) and return false(0). This
	//  function will be called in sequence but not from a dedicated thread. For
	//  backwards compatibility set |bytes_read| to -1 and return false(0) and
	//  the ReadResponse function will be called.
	SetOnRead(fn TOnResourceHandlerRead) // property event
	// SetOnReadResponse
	//  Read response data. If data is available immediately copy up to
	//  |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
	//  bytes copied, and return true(1). To read the data at a later time set
	//  |bytes_read| to 0, return true(1) and call ICefCallback.cont() when
	//  the data is available. To indicate response completion return false(0).
	//  WARNING: This function is deprecated. Use Skip and Read instead.
	SetOnReadResponse(fn TOnResourceHandlerReadResponse) // property event
	// SetOnCancel
	//  Request processing has been canceled.
	SetOnCancel(fn TOnResourceHandlerCancel) // property event
}

// TResourceHandler Parent: TCefResourceHandler
//
//	Interface used to implement a custom request handler interface. The
//	functions of this interface will be called on the IO thread unless otherwise
//	indicated.
//	<a cref="uCEFTypes|TCefResourceHandler">Implements TCefResourceHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_handler_t)</a>
type TResourceHandler struct {
	TCefResourceHandler
	openPtr               uintptr
	processRequestPtr     uintptr
	getResponseHeadersPtr uintptr
	skipPtr               uintptr
	readPtr               uintptr
	readResponsePtr       uintptr
	cancelPtr             uintptr
}

func NewResourceHandler(browser ICefBrowser, frame ICefFrame, schemeName string, request ICefRequest) IResourceHandler {
	r1 := CEF().SysCallN(2215, GetObjectUintptr(browser), GetObjectUintptr(frame), PascalStr(schemeName), GetObjectUintptr(request))
	return AsResourceHandler(r1)
}

func ResourceHandlerClass() TClass {
	ret := CEF().SysCallN(2214)
	return TClass(ret)
}

func (m *TResourceHandler) SetOnOpen(fn TOnResourceHandlerOpen) {
	if m.openPtr != 0 {
		RemoveEventElement(m.openPtr)
	}
	m.openPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2218, m.Instance(), m.openPtr)
}

func (m *TResourceHandler) SetOnProcessRequest(fn TOnResourceHandlerProcessRequest) {
	if m.processRequestPtr != 0 {
		RemoveEventElement(m.processRequestPtr)
	}
	m.processRequestPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2219, m.Instance(), m.processRequestPtr)
}

func (m *TResourceHandler) SetOnGetResponseHeaders(fn TOnResourceHandlerGetResponseHeaders) {
	if m.getResponseHeadersPtr != 0 {
		RemoveEventElement(m.getResponseHeadersPtr)
	}
	m.getResponseHeadersPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2217, m.Instance(), m.getResponseHeadersPtr)
}

func (m *TResourceHandler) SetOnSkip(fn TOnResourceHandlerSkip) {
	if m.skipPtr != 0 {
		RemoveEventElement(m.skipPtr)
	}
	m.skipPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2222, m.Instance(), m.skipPtr)
}

func (m *TResourceHandler) SetOnRead(fn TOnResourceHandlerRead) {
	if m.readPtr != 0 {
		RemoveEventElement(m.readPtr)
	}
	m.readPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2220, m.Instance(), m.readPtr)
}

func (m *TResourceHandler) SetOnReadResponse(fn TOnResourceHandlerReadResponse) {
	if m.readResponsePtr != 0 {
		RemoveEventElement(m.readResponsePtr)
	}
	m.readResponsePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2221, m.Instance(), m.readResponsePtr)
}

func (m *TResourceHandler) SetOnCancel(fn TOnResourceHandlerCancel) {
	if m.cancelPtr != 0 {
		RemoveEventElement(m.cancelPtr)
	}
	m.cancelPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2216, m.Instance(), m.cancelPtr)
}
