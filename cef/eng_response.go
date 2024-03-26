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

// ICefResponse Parent: ICefBaseRefCounted
//
//	Interface used to represent a web response. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_response_capi.h">CEF source file: /include/capi/cef_response_capi.h (cef_response_t))
type ICefResponse interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetError
	//  Get the response error code. Returns ERR_NONE if there was no error.
	GetError() TCefErrorCode // function
	// GetStatus
	//  Get the response status code.
	GetStatus() int32 // function
	// GetStatusText
	//  Get the response status text.
	GetStatusText() string // function
	// GetMimeType
	//  Get the response mime type.
	GetMimeType() string // function
	// GetCharset
	//  Get the response charset.
	GetCharset() string // function
	// GetHeaderByName
	//  Get the value for the specified response header field.
	GetHeaderByName(name string) string // function
	// GetURL
	//  Get the resolved URL after redirects or changed as a result of HSTS.
	GetURL() string // function
	// SetError
	//  Set the response error code. This can be used by custom scheme handlers to return errors during initial request processing.
	SetError(error TCefErrorCode) // procedure
	// SetStatus
	//  Set the response status code.
	SetStatus(status int32) // procedure
	// SetStatusText
	//  Set the response status text.
	SetStatusText(statusText string) // procedure
	// SetMimeType
	//  Set the response mime type.
	SetMimeType(mimetype string) // procedure
	// SetCharset
	//  Set the response charset.
	SetCharset(charset string) // procedure
	// SetHeaderByName
	//  Set the header |name| to |value|. If |overwrite| is true (1) any existing values will be replaced with the new value. If |overwrite| is false (0) any existing values will not be overwritten.
	SetHeaderByName(name, value string, overwrite bool) // procedure
	// GetHeaderMap
	//  Get all response header fields.
	GetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetHeaderMap
	//  Set all response header fields.
	SetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetURL
	//  Set the resolved URL after redirects or changed as a result of HSTS.
	SetURL(url string) // procedure
}

// TCefResponse Parent: TCefBaseRefCounted
//
//	Interface used to represent a web response. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_response_capi.h">CEF source file: /include/capi/cef_response_capi.h (cef_response_t))
type TCefResponse struct {
	TCefBaseRefCounted
}

// ResponseRef -> ICefResponse
var ResponseRef response

// response TCefResponse Ref
type response uintptr

func (m *response) UnWrap(data uintptr) ICefResponse {
	var resultCefResponse uintptr
	CEF().SysCallN(1329, uintptr(data), uintptr(unsafePointer(&resultCefResponse)))
	return AsCefResponse(resultCefResponse)
}

func (m *response) New() ICefResponse {
	var resultCefResponse uintptr
	CEF().SysCallN(1320, uintptr(unsafePointer(&resultCefResponse)))
	return AsCefResponse(resultCefResponse)
}

func (m *TCefResponse) IsReadOnly() bool {
	r1 := CEF().SysCallN(1319, m.Instance())
	return GoBool(r1)
}

func (m *TCefResponse) GetError() TCefErrorCode {
	r1 := CEF().SysCallN(1312, m.Instance())
	return TCefErrorCode(r1)
}

func (m *TCefResponse) GetStatus() int32 {
	r1 := CEF().SysCallN(1316, m.Instance())
	return int32(r1)
}

func (m *TCefResponse) GetStatusText() string {
	r1 := CEF().SysCallN(1317, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetMimeType() string {
	r1 := CEF().SysCallN(1315, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetCharset() string {
	r1 := CEF().SysCallN(1311, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetHeaderByName(name string) string {
	r1 := CEF().SysCallN(1313, m.Instance(), PascalStr(name))
	return GoStr(r1)
}

func (m *TCefResponse) GetURL() string {
	r1 := CEF().SysCallN(1318, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) SetError(error TCefErrorCode) {
	CEF().SysCallN(1322, m.Instance(), uintptr(error))
}

func (m *TCefResponse) SetStatus(status int32) {
	CEF().SysCallN(1326, m.Instance(), uintptr(status))
}

func (m *TCefResponse) SetStatusText(statusText string) {
	CEF().SysCallN(1327, m.Instance(), PascalStr(statusText))
}

func (m *TCefResponse) SetMimeType(mimetype string) {
	CEF().SysCallN(1325, m.Instance(), PascalStr(mimetype))
}

func (m *TCefResponse) SetCharset(charset string) {
	CEF().SysCallN(1321, m.Instance(), PascalStr(charset))
}

func (m *TCefResponse) SetHeaderByName(name, value string, overwrite bool) {
	CEF().SysCallN(1323, m.Instance(), PascalStr(name), PascalStr(value), PascalBool(overwrite))
}

func (m *TCefResponse) GetHeaderMap(headerMap ICefStringMultimap) {
	CEF().SysCallN(1314, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefResponse) SetHeaderMap(headerMap ICefStringMultimap) {
	CEF().SysCallN(1324, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefResponse) SetURL(url string) {
	CEF().SysCallN(1328, m.Instance(), PascalStr(url))
}
