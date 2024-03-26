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

// ICefValue Parent: ICefBaseRefCounted
//
//	Interface that wraps other data value types. Complex types (binary, dictionary and list) will be referenced but not owned by this object. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_value_t))
type ICefValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying data is valid. This will always be true (1) for simple types. For complex types (binary, dictionary and list) the underlying data may become invalid if owned by another object (e.g. list or dictionary) and that other object is then modified or destroyed. This value object can be re-used by calling Set*() even if the underlying data is invalid.
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if the underlying data is owned by another object.
	IsOwned() bool // function
	// IsReadOnly
	//  Returns true (1) if the underlying data is read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying data. If true (1) modifications to this object will also affect |that| object and vice-versa.
	IsSame(that ICefValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefValue) bool // function
	// Copy
	//  Returns a copy of this object. The underlying data will also be copied.
	Copy() ICefValue // function
	// GetType
	//  Returns the underlying value type.
	GetType() TCefValueType // function
	// GetBool
	//  Returns the underlying value as type bool.
	GetBool() bool // function
	// GetInt
	//  Returns the underlying value as type int.
	GetInt() int32 // function
	// GetDouble
	//  Returns the underlying value as type double.
	GetDouble() (resultFloat64 float64) // function
	// GetString
	//  Returns the underlying value as type string.
	GetString() string // function
	// GetBinary
	//  Returns the underlying value as type binary. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_binary().
	GetBinary() ICefBinaryValue // function
	// GetDictionary
	//  Returns the underlying value as type dictionary. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_dictionary().
	GetDictionary() ICefDictionaryValue // function
	// GetList
	//  Returns the underlying value as type list. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_list().
	GetList() ICefListValue // function
	// SetNull
	//  Sets the underlying value as type null. Returns true (1) if the value was set successfully.
	SetNull() bool // function
	// SetBool
	//  Sets the underlying value as type bool. Returns true (1) if the value was set successfully.
	SetBool(value bool) bool // function
	// SetInt
	//  Sets the underlying value as type int. Returns true (1) if the value was set successfully.
	SetInt(value int32) bool // function
	// SetDouble
	//  Sets the underlying value as type double. Returns true (1) if the value was set successfully.
	SetDouble(value float64) bool // function
	// SetString
	//  Sets the underlying value as type string. Returns true (1) if the value was set successfully.
	SetString(value string) bool // function
	// SetBinary
	//  Sets the underlying value as type binary. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetBinary(value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the underlying value as type dict. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetDictionary(value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the underlying value as type list. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetList(value ICefListValue) bool // function
}

// TCefValue Parent: TCefBaseRefCounted
//
//	Interface that wraps other data value types. Complex types (binary, dictionary and list) will be referenced but not owned by this object. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_value_t))
type TCefValue struct {
	TCefBaseRefCounted
}

// ValueRef -> ICefValue
var ValueRef value

// value TCefValue Ref
type value uintptr

func (m *value) UnWrap(data uintptr) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1520, uintptr(data), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *value) New() ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1511, uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefValue) IsValid() bool {
	r1 := CEF().SysCallN(1510, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsOwned() bool {
	r1 := CEF().SysCallN(1507, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsReadOnly() bool {
	r1 := CEF().SysCallN(1508, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsSame(that ICefValue) bool {
	r1 := CEF().SysCallN(1509, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefValue) IsEqual(that ICefValue) bool {
	r1 := CEF().SysCallN(1506, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefValue) Copy() ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1497, m.Instance(), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefValue) GetType() TCefValueType {
	r1 := CEF().SysCallN(1505, m.Instance())
	return TCefValueType(r1)
}

func (m *TCefValue) GetBool() bool {
	r1 := CEF().SysCallN(1499, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) GetInt() int32 {
	r1 := CEF().SysCallN(1502, m.Instance())
	return int32(r1)
}

func (m *TCefValue) GetDouble() (resultFloat64 float64) {
	CEF().SysCallN(1501, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefValue) GetString() string {
	r1 := CEF().SysCallN(1504, m.Instance())
	return GoStr(r1)
}

func (m *TCefValue) GetBinary() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(1498, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefValue) GetDictionary() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(1500, m.Instance(), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefValue) GetList() ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1503, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefValue) SetNull() bool {
	r1 := CEF().SysCallN(1518, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) SetBool(value bool) bool {
	r1 := CEF().SysCallN(1513, m.Instance(), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefValue) SetInt(value int32) bool {
	r1 := CEF().SysCallN(1516, m.Instance(), uintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetDouble(value float64) bool {
	r1 := CEF().SysCallN(1515, m.Instance(), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefValue) SetString(value string) bool {
	r1 := CEF().SysCallN(1519, m.Instance(), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetBinary(value ICefBinaryValue) bool {
	r1 := CEF().SysCallN(1512, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetDictionary(value ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(1514, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetList(value ICefListValue) bool {
	r1 := CEF().SysCallN(1517, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}
