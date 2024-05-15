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

// ICefListValue Parent: ICefBaseRefCounted
//
//	Interface representing a list value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_list_value_t))</a>
type ICefListValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. This object may become invalid if the underlying data is owned by another object (e.g. list or dictionary) and that other object is then modified or destroyed. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if this object is currently owned by another object.
	IsOwned() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying data. If true (1) modifications to this object will also affect |that| object and vice-versa.
	IsSame(that ICefListValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefListValue) bool // function
	// Copy
	//  Returns a writable copy of this object.
	Copy() ICefListValue // function
	// SetSize
	//  Sets the number of values. If the number of values is expanded all new value slots will default to type null. Returns true (1) on success.
	SetSize(size NativeUInt) bool // function
	// GetSize
	//  Returns the number of values.
	GetSize() NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// Remove
	//  Removes the value at the specified index.
	Remove(index NativeUInt) bool // function
	// GetType
	//  Returns the value type at the specified index.
	GetType(index NativeUInt) TCefValueType // function
	// GetValue
	//  Returns the value at the specified index. For simple types the returned value will copy existing data and modifications to the value will not modify this object. For complex types (binary, dictionary and list) the returned value will reference existing data and modifications to the value will modify this object.
	GetValue(index NativeUInt) ICefValue // function
	// GetBool
	//  Returns the value at the specified index as type bool.
	GetBool(index NativeUInt) bool // function
	// GetInt
	//  Returns the value at the specified index as type int.
	GetInt(index NativeUInt) int32 // function
	// GetDouble
	//  Returns the value at the specified index as type double.
	GetDouble(index NativeUInt) (resultFloat64 float64) // function
	// GetString
	//  Returns the value at the specified index as type string.
	GetString(index NativeUInt) string // function
	// GetBinary
	//  Returns the value at the specified index as type binary. The returned value will reference existing data.
	GetBinary(index NativeUInt) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified index as type dictionary. The returned value will reference existing data and modifications to the value will modify this object.
	GetDictionary(index NativeUInt) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified index as type list. The returned value will reference existing data and modifications to the value will modify this object.
	GetList(index NativeUInt) ICefListValue // function
	// SetValue
	//  Sets the value at the specified index. Returns true (1) if the value was set successfully. If |value| represents simple data then the underlying data will be copied and modifications to |value| will not modify this object. If |value| represents complex data (binary, dictionary or list) then the underlying data will be referenced and modifications to |value| will modify this object.
	SetValue(index NativeUInt, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified index as type null. Returns true (1) if the value was set successfully.
	SetNull(index NativeUInt) bool // function
	// SetBool
	//  Sets the value at the specified index as type bool. Returns true (1) if the value was set successfully.
	SetBool(index NativeUInt, value bool) bool // function
	// SetInt
	//  Sets the value at the specified index as type int. Returns true (1) if the value was set successfully.
	SetInt(index NativeUInt, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified index as type double. Returns true (1) if the value was set successfully.
	SetDouble(index NativeUInt, value float64) bool // function
	// SetString
	//  Sets the value at the specified index as type string. Returns true (1) if the value was set successfully.
	SetString(index NativeUInt, value string) bool // function
	// SetBinary
	//  Sets the value at the specified index as type binary. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetBinary(index NativeUInt, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified index as type dict. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetDictionary(index NativeUInt, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified index as type list. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetList(index NativeUInt, value ICefListValue) bool // function
}

// TCefListValue Parent: TCefBaseRefCounted
//
//	Interface representing a list value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_list_value_t))</a>
type TCefListValue struct {
	TCefBaseRefCounted
}

// ListValueRef -> ICefListValue
var ListValueRef listValue

// listValue TCefListValue Ref
type listValue uintptr

func (m *listValue) UnWrap(data uintptr) ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1049, uintptr(data), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *listValue) New() ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1037, uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) IsValid() bool {
	r1 := CEF().SysCallN(1036, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsOwned() bool {
	r1 := CEF().SysCallN(1033, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsReadOnly() bool {
	r1 := CEF().SysCallN(1034, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsSame(that ICefListValue) bool {
	r1 := CEF().SysCallN(1035, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefListValue) IsEqual(that ICefListValue) bool {
	r1 := CEF().SysCallN(1032, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefListValue) Copy() ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1021, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) SetSize(size NativeUInt) bool {
	r1 := CEF().SysCallN(1046, m.Instance(), uintptr(size))
	return GoBool(r1)
}

func (m *TCefListValue) GetSize() NativeUInt {
	r1 := CEF().SysCallN(1028, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefListValue) Clear() bool {
	r1 := CEF().SysCallN(1020, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) Remove(index NativeUInt) bool {
	r1 := CEF().SysCallN(1038, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) GetType(index NativeUInt) TCefValueType {
	r1 := CEF().SysCallN(1030, m.Instance(), uintptr(index))
	return TCefValueType(r1)
}

func (m *TCefListValue) GetValue(index NativeUInt) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1031, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefListValue) GetBool(index NativeUInt) bool {
	r1 := CEF().SysCallN(1023, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) GetInt(index NativeUInt) int32 {
	r1 := CEF().SysCallN(1026, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefListValue) GetDouble(index NativeUInt) (resultFloat64 float64) {
	CEF().SysCallN(1025, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefListValue) GetString(index NativeUInt) string {
	r1 := CEF().SysCallN(1029, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefListValue) GetBinary(index NativeUInt) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(1022, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefListValue) GetDictionary(index NativeUInt) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(1024, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefListValue) GetList(index NativeUInt) ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1027, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) SetValue(index NativeUInt, value ICefValue) bool {
	r1 := CEF().SysCallN(1048, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetNull(index NativeUInt) bool {
	r1 := CEF().SysCallN(1045, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) SetBool(index NativeUInt, value bool) bool {
	r1 := CEF().SysCallN(1040, m.Instance(), uintptr(index), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetInt(index NativeUInt, value int32) bool {
	r1 := CEF().SysCallN(1043, m.Instance(), uintptr(index), uintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetDouble(index NativeUInt, value float64) bool {
	r1 := CEF().SysCallN(1042, m.Instance(), uintptr(index), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefListValue) SetString(index NativeUInt, value string) bool {
	r1 := CEF().SysCallN(1047, m.Instance(), uintptr(index), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetBinary(index NativeUInt, value ICefBinaryValue) bool {
	r1 := CEF().SysCallN(1039, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetDictionary(index NativeUInt, value ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(1041, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetList(index NativeUInt, value ICefListValue) bool {
	r1 := CEF().SysCallN(1044, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}
