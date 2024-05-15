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

// ICefPostDataElement Parent: ICefBaseRefCounted
//
//	Interface used to represent a single element in the request post data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_post_data_element_t))</a>
type ICefPostDataElement interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetType
	//  Return the type of this post data element.
	GetType() TCefPostDataElementType // function
	// GetFile
	//  Return the file name.
	GetFile() string // function
	// GetBytesCount
	//  Return the number of bytes.
	GetBytesCount() NativeUInt // function
	// GetBytes
	//  Read up to |size| bytes into |bytes| and return the number of bytes actually read.
	GetBytes(size NativeUInt, bytes uintptr) NativeUInt // function
	// SetToEmpty
	//  Remove all contents from the post data element.
	SetToEmpty() // procedure
	// SetToFile
	//  The post data element will represent a file.
	SetToFile(fileName string) // procedure
	// SetToBytes
	//  The post data element will represent bytes. The bytes passed in will be copied.
	SetToBytes(size NativeUInt, bytes uintptr) // procedure
}

// TCefPostDataElement Parent: TCefBaseRefCounted
//
//	Interface used to represent a single element in the request post data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_capi.h">CEF source file: /include/capi/cef_request_capi.h (cef_post_data_element_t))</a>
type TCefPostDataElement struct {
	TCefBaseRefCounted
}

// PostDataElementRef -> ICefPostDataElement
var PostDataElementRef postDataElement

// postDataElement TCefPostDataElement Ref
type postDataElement uintptr

func (m *postDataElement) UnWrap(data uintptr) ICefPostDataElement {
	var resultCefPostDataElement uintptr
	CEF().SysCallN(1197, uintptr(data), uintptr(unsafePointer(&resultCefPostDataElement)))
	return AsCefPostDataElement(resultCefPostDataElement)
}

func (m *postDataElement) New() ICefPostDataElement {
	var resultCefPostDataElement uintptr
	CEF().SysCallN(1193, uintptr(unsafePointer(&resultCefPostDataElement)))
	return AsCefPostDataElement(resultCefPostDataElement)
}

func (m *TCefPostDataElement) IsReadOnly() bool {
	r1 := CEF().SysCallN(1192, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostDataElement) GetType() TCefPostDataElementType {
	r1 := CEF().SysCallN(1191, m.Instance())
	return TCefPostDataElementType(r1)
}

func (m *TCefPostDataElement) GetFile() string {
	r1 := CEF().SysCallN(1190, m.Instance())
	return GoStr(r1)
}

func (m *TCefPostDataElement) GetBytesCount() NativeUInt {
	r1 := CEF().SysCallN(1189, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPostDataElement) GetBytes(size NativeUInt, bytes uintptr) NativeUInt {
	r1 := CEF().SysCallN(1188, m.Instance(), uintptr(size), uintptr(bytes))
	return NativeUInt(r1)
}

func (m *TCefPostDataElement) SetToEmpty() {
	CEF().SysCallN(1195, m.Instance())
}

func (m *TCefPostDataElement) SetToFile(fileName string) {
	CEF().SysCallN(1196, m.Instance(), PascalStr(fileName))
}

func (m *TCefPostDataElement) SetToBytes(size NativeUInt, bytes uintptr) {
	CEF().SysCallN(1194, m.Instance(), uintptr(size), uintptr(bytes))
}
