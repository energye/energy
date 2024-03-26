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

// ICefBinaryValue Parent: ICefBaseRefCounted
//
//	Interface representing a binary value. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_binary_value_t))
type ICefBinaryValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. This object may become invalid if the underlying data is owned by another object (e.g. list or dictionary) and that other object is then modified or destroyed. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if this object is currently owned by another object.
	IsOwned() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying data.
	IsSame(that ICefBinaryValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefBinaryValue) bool // function
	// Copy
	//  Returns a copy of this object. The data in this object will also be copied.
	Copy() ICefBinaryValue // function
	// GetSize
	//  Returns the data size.
	GetSize() NativeUInt // function
	// GetData
	//  Read up to |buffer_size| number of bytes into |buffer|. Reading begins at the specified byte |data_offset|. Returns the number of bytes read.
	GetData(buffer uintptr, bufferSize, dataOffset NativeUInt) NativeUInt // function
}

// TCefBinaryValue Parent: TCefBaseRefCounted
//
//	Interface representing a binary value. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_binary_value_t))
type TCefBinaryValue struct {
	TCefBaseRefCounted
}

// BinaryValueRef -> ICefBinaryValue
var BinaryValueRef binaryValue

// binaryValue TCefBinaryValue Ref
type binaryValue uintptr

func (m *binaryValue) UnWrap(data uintptr) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(605, uintptr(data), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *binaryValue) New(data uintptr, dataSize NativeUInt) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(604, uintptr(data), uintptr(dataSize), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefBinaryValue) IsValid() bool {
	r1 := CEF().SysCallN(603, m.Instance())
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsOwned() bool {
	r1 := CEF().SysCallN(601, m.Instance())
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsSame(that ICefBinaryValue) bool {
	r1 := CEF().SysCallN(602, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsEqual(that ICefBinaryValue) bool {
	r1 := CEF().SysCallN(600, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBinaryValue) Copy() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(597, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefBinaryValue) GetSize() NativeUInt {
	r1 := CEF().SysCallN(599, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefBinaryValue) GetData(buffer uintptr, bufferSize, dataOffset NativeUInt) NativeUInt {
	r1 := CEF().SysCallN(598, m.Instance(), uintptr(buffer), uintptr(bufferSize), uintptr(dataOffset))
	return NativeUInt(r1)
}
