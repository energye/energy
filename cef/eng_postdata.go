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

// ICefPostData Parent: ICefBaseRefCounted
//
//	Interface used to represent post data for a web request. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_post_data_t))</a>
type ICefPostData interface {
	ICefBaseRefCounted
	// GetElements
	//  Retrieve the post data elements.
	GetElements(elementsCount *NativeUInt, elements *ICefPostDataElementArray)
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// HasExcludedElements
	//  Returns true (1) if the underlying POST data includes elements that are not represented by this ICefPostData object (for example, multi-part file upload data). Modifying ICefPostData objects with excluded elements may result in the request failing.
	HasExcludedElements() bool // function
	// GetElementCount
	//  Returns the number of existing post data elements.
	GetElementCount() NativeUInt // function
	// RemoveElement
	//  Remove the specified post data element. Returns true (1) if the removal succeeds.
	RemoveElement(element ICefPostDataElement) bool // function
	// AddElement
	//  Add the specified post data element. Returns true (1) if the add succeeds.
	AddElement(element ICefPostDataElement) bool // function
	// RemoveElements
	//  Remove all existing post data elements.
	RemoveElements() // procedure
}

// TCefPostData Parent: TCefBaseRefCounted
//
//	Interface used to represent post data for a web request. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_post_data_t))</a>
type TCefPostData struct {
	TCefBaseRefCounted
}

// PostDataRef -> ICefPostData
var PostDataRef postData

// postData TCefPostData Ref
type postData uintptr

func (m *postData) UnWrap(data uintptr) ICefPostData {
	var resultCefPostData uintptr
	CEF().SysCallN(1205, uintptr(data), uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *postData) New() ICefPostData {
	var resultCefPostData uintptr
	CEF().SysCallN(1202, uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *TCefPostData) IsReadOnly() bool {
	r1 := CEF().SysCallN(1201, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostData) HasExcludedElements() bool {
	r1 := CEF().SysCallN(1200, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostData) GetElementCount() NativeUInt {
	r1 := CEF().SysCallN(1199, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPostData) RemoveElement(element ICefPostDataElement) bool {
	r1 := CEF().SysCallN(1203, m.Instance(), GetObjectUintptr(element))
	return GoBool(r1)
}

func (m *TCefPostData) AddElement(element ICefPostDataElement) bool {
	r1 := CEF().SysCallN(1198, m.Instance(), GetObjectUintptr(element))
	return GoBool(r1)
}

func (m *TCefPostData) RemoveElements() {
	CEF().SysCallN(1204, m.Instance())
}
