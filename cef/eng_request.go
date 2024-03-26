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

// ICefRequest Parent: ICefBaseRefCounted
//
//	Interface used to represent a web request. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_request_t))
type ICefRequest interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetUrl
	//  Get the fully qualified URL.
	GetUrl() string // function
	// GetMethod
	//  Get the request function type. The value will default to POST if post data is provided and GET otherwise.
	GetMethod() string // function
	// GetPostData
	//  Get the post data.
	GetPostData() ICefPostData // function
	// GetReferrerUrl
	//  Get the referrer URL.
	GetReferrerUrl() string // function
	// GetReferrerPolicy
	//  Get the referrer policy.
	GetReferrerPolicy() TCefReferrerPolicy // function
	// GetHeaderByName
	//  Returns the first header value for |name| or an NULL string if not found. Will not return the Referer value if any. Use GetHeaderMap instead if |name| might have multiple values.
	GetHeaderByName(name string) string // function
	// GetFlags
	//  Get the flags used in combination with ICefUrlRequest. See TCefUrlRequestFlags for supported values.
	GetFlags() TCefUrlRequestFlags // function
	// GetFirstPartyForCookies
	//  Get the URL to the first party for cookies used in combination with ICefUrlRequest.
	GetFirstPartyForCookies() string // function
	// GetResourceType
	//  Get the resource type for this request. Only available in the browser process.
	GetResourceType() TCefResourceType // function
	// GetTransitionType
	//  Get the transition type for this request. Only available in the browser process and only applies to requests that represent a main frame or sub- frame navigation.
	GetTransitionType() TCefTransitionType // function
	// GetIdentifier
	//  Returns the globally unique identifier for this request or 0 if not specified. Can be used by ICefResourceRequestHandler implementations in the browser process to track a single request across multiple callbacks.
	GetIdentifier() (resultUint64 uint64) // function
	// GetHeaderMap
	//  Get the header values. Will not include the Referer value if any.
	GetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetUrl
	//  Set the fully qualified URL.
	SetUrl(value string) // procedure
	// SetMethod
	//  Set the request function type.
	SetMethod(value string) // procedure
	// SetReferrer
	//  Set the referrer URL and policy. If non-NULL the referrer URL must be fully qualified with an HTTP or HTTPS scheme component. Any username, password or ref component will be removed.
	SetReferrer(referrerUrl string, policy TCefReferrerPolicy) // procedure
	// SetPostData
	//  Set the post data.
	SetPostData(value ICefPostData) // procedure
	// SetHeaderMap
	//  Set the header values. If a Referer value exists in the header map it will be removed and ignored.
	SetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetHeaderByName
	//  Set the header |name| to |value|. If |overwrite| is true (1) any existing values will be replaced with the new value. If |overwrite| is false (0) any existing values will not be overwritten. The Referer value cannot be set using this function.
	SetHeaderByName(name, value string, overwrite bool) // procedure
	// SetFlags
	//  Set the flags used in combination with ICefUrlRequest. See TCefUrlRequestFlags for supported values.
	SetFlags(flags TCefUrlRequestFlags) // procedure
	// SetFirstPartyForCookies
	//  Set the URL to the first party for cookies used in combination with ICefUrlRequest.
	SetFirstPartyForCookies(url string) // procedure
	// Assign
	//  Set all values at one time. This method corresponds to TCefRequest.set_ and cef_request_t.set
	Assign(url, method string, postData ICefPostData, headerMap ICefStringMultimap) // procedure
}

// TCefRequest Parent: TCefBaseRefCounted
//
//	Interface used to represent a web request. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_request_t))
type TCefRequest struct {
	TCefBaseRefCounted
}

// RequestRef -> ICefRequest
var RequestRef request

// request TCefRequest Ref
type request uintptr

func (m *request) UnWrap(data uintptr) ICefRequest {
	var resultCefRequest uintptr
	CEF().SysCallN(1306, uintptr(data), uintptr(unsafePointer(&resultCefRequest)))
	return AsCefRequest(resultCefRequest)
}

func (m *request) New() ICefRequest {
	var resultCefRequest uintptr
	CEF().SysCallN(1297, uintptr(unsafePointer(&resultCefRequest)))
	return AsCefRequest(resultCefRequest)
}

func (m *TCefRequest) IsReadOnly() bool {
	r1 := CEF().SysCallN(1296, m.Instance())
	return GoBool(r1)
}

func (m *TCefRequest) GetUrl() string {
	r1 := CEF().SysCallN(1295, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetMethod() string {
	r1 := CEF().SysCallN(1289, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetPostData() ICefPostData {
	var resultCefPostData uintptr
	CEF().SysCallN(1290, m.Instance(), uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *TCefRequest) GetReferrerUrl() string {
	r1 := CEF().SysCallN(1292, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetReferrerPolicy() TCefReferrerPolicy {
	r1 := CEF().SysCallN(1291, m.Instance())
	return TCefReferrerPolicy(r1)
}

func (m *TCefRequest) GetHeaderByName(name string) string {
	r1 := CEF().SysCallN(1286, m.Instance(), PascalStr(name))
	return GoStr(r1)
}

func (m *TCefRequest) GetFlags() TCefUrlRequestFlags {
	r1 := CEF().SysCallN(1285, m.Instance())
	return TCefUrlRequestFlags(r1)
}

func (m *TCefRequest) GetFirstPartyForCookies() string {
	r1 := CEF().SysCallN(1284, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetResourceType() TCefResourceType {
	r1 := CEF().SysCallN(1293, m.Instance())
	return TCefResourceType(r1)
}

func (m *TCefRequest) GetTransitionType() TCefTransitionType {
	r1 := CEF().SysCallN(1294, m.Instance())
	return TCefTransitionType(r1)
}

func (m *TCefRequest) GetIdentifier() (resultUint64 uint64) {
	CEF().SysCallN(1288, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCefRequest) GetHeaderMap(headerMap ICefStringMultimap) {
	CEF().SysCallN(1287, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefRequest) SetUrl(value string) {
	CEF().SysCallN(1305, m.Instance(), PascalStr(value))
}

func (m *TCefRequest) SetMethod(value string) {
	CEF().SysCallN(1302, m.Instance(), PascalStr(value))
}

func (m *TCefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	CEF().SysCallN(1304, m.Instance(), PascalStr(referrerUrl), uintptr(policy))
}

func (m *TCefRequest) SetPostData(value ICefPostData) {
	CEF().SysCallN(1303, m.Instance(), GetObjectUintptr(value))
}

func (m *TCefRequest) SetHeaderMap(headerMap ICefStringMultimap) {
	CEF().SysCallN(1301, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefRequest) SetHeaderByName(name, value string, overwrite bool) {
	CEF().SysCallN(1300, m.Instance(), PascalStr(name), PascalStr(value), PascalBool(overwrite))
}

func (m *TCefRequest) SetFlags(flags TCefUrlRequestFlags) {
	CEF().SysCallN(1299, m.Instance(), uintptr(flags))
}

func (m *TCefRequest) SetFirstPartyForCookies(url string) {
	CEF().SysCallN(1298, m.Instance(), PascalStr(url))
}

func (m *TCefRequest) Assign(url, method string, postData ICefPostData, headerMap ICefStringMultimap) {
	CEF().SysCallN(1283, m.Instance(), PascalStr(url), PascalStr(method), GetObjectUintptr(postData), GetObjectUintptr(headerMap))
}
