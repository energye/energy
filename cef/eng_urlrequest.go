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

// ICefUrlRequest Parent: ICefBaseRefCounted
//
//	Interface used to make a URL request. URL requests are not associated with a browser instance so no ICefClient callbacks will be executed. URL requests can be created on any valid CEF thread in either the browser or render process. Once created the functions of the URL request object must be accessed on the same thread that created it.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h (cef_urlrequest_t))</a>
type ICefUrlRequest interface {
	ICefBaseRefCounted
	// GetRequest
	//  Returns the request object used to create this URL request. The returned object is read-only and should not be modified.
	GetRequest() ICefRequest // function
	// GetClient
	//  Returns the client.
	GetClient() ICefUrlRequestClient // function
	// GetRequestStatus
	//  Returns the request status.
	GetRequestStatus() TCefUrlRequestStatus // function
	// GetRequestError
	//  Returns the request error if status is UR_CANCELED or UR_FAILED, or 0 otherwise.
	GetRequestError() int32 // function
	// GetResponse
	//  Returns the response, or NULL if no response information is available. Response information will only be available after the upload has completed. The returned object is read-only and should not be modified.
	GetResponse() ICefResponse // function
	// GetResponseWasCached
	//  Returns true (1) if the response body was served from the cache. This includes responses for which revalidation was required.
	GetResponseWasCached() bool // function
	// Cancel
	//  Cancel the request.
	Cancel() // procedure
}

// TCefUrlRequest Parent: TCefBaseRefCounted
//
//	Interface used to make a URL request. URL requests are not associated with a browser instance so no ICefClient callbacks will be executed. URL requests can be created on any valid CEF thread in either the browser or render process. Once created the functions of the URL request object must be accessed on the same thread that created it.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h (cef_urlrequest_t))</a>
type TCefUrlRequest struct {
	TCefBaseRefCounted
}

// UrlRequestRef -> ICefUrlRequest
var UrlRequestRef urlRequest

// urlRequest TCefUrlRequest Ref
type urlRequest uintptr

func (m *urlRequest) UnWrap(data uintptr) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	CEF().SysCallN(1473, uintptr(data), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *urlRequest) New(request ICefRequest, client ICefUrlRequestClient, requestContext ICefRequestContext) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	CEF().SysCallN(1472, GetObjectUintptr(request), GetObjectUintptr(client), GetObjectUintptr(requestContext), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *TCefUrlRequest) GetRequest() ICefRequest {
	var resultCefRequest uintptr
	CEF().SysCallN(1467, m.Instance(), uintptr(unsafePointer(&resultCefRequest)))
	return AsCefRequest(resultCefRequest)
}

func (m *TCefUrlRequest) GetClient() ICefUrlRequestClient {
	var resultCefUrlRequestClient uintptr
	CEF().SysCallN(1466, m.Instance(), uintptr(unsafePointer(&resultCefUrlRequestClient)))
	return AsCefUrlRequestClient(resultCefUrlRequestClient)
}

func (m *TCefUrlRequest) GetRequestStatus() TCefUrlRequestStatus {
	r1 := CEF().SysCallN(1469, m.Instance())
	return TCefUrlRequestStatus(r1)
}

func (m *TCefUrlRequest) GetRequestError() int32 {
	r1 := CEF().SysCallN(1468, m.Instance())
	return int32(r1)
}

func (m *TCefUrlRequest) GetResponse() ICefResponse {
	var resultCefResponse uintptr
	CEF().SysCallN(1470, m.Instance(), uintptr(unsafePointer(&resultCefResponse)))
	return AsCefResponse(resultCefResponse)
}

func (m *TCefUrlRequest) GetResponseWasCached() bool {
	r1 := CEF().SysCallN(1471, m.Instance())
	return GoBool(r1)
}

func (m *TCefUrlRequest) Cancel() {
	CEF().SysCallN(1465, m.Instance())
}
