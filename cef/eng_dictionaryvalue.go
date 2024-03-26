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

// ICefDictionaryValue Parent: ICefBaseRefCounted
//
//	Interface representing a dictionary value. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_dictionary_value_t))
type ICefDictionaryValue interface {
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
	IsSame(that ICefDictionaryValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefDictionaryValue) bool // function
	// Copy
	//  Returns a writable copy of this object. If |exclude_NULL_children| is true (1) any NULL dictionaries or lists will be excluded from the copy.
	Copy(excludeEmptyChildren bool) ICefDictionaryValue // function
	// GetSize
	//  Returns the number of values.
	GetSize() NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// HasKey
	//  Returns true (1) if the current dictionary has a value for the given key.
	HasKey(key string) bool // function
	// GetKeys
	//  Reads all keys for this dictionary into the specified vector.
	GetKeys(keys IStrings) bool // function
	// Remove
	//  Removes the value at the specified key. Returns true (1) is the value was removed successfully.
	Remove(key string) bool // function
	// GetType
	//  Returns the value type for the specified key.
	GetType(key string) TCefValueType // function
	// GetValue
	//  Returns the value at the specified key. For simple types the returned value will copy existing data and modifications to the value will not modify this object. For complex types (binary, dictionary and list) the returned value will reference existing data and modifications to the value will modify this object.
	GetValue(key string) ICefValue // function
	// GetBool
	//  Returns the value at the specified key as type bool.
	GetBool(key string) bool // function
	// GetInt
	//  Returns the value at the specified key as type int.
	GetInt(key string) int32 // function
	// GetDouble
	//  Returns the value at the specified key as type double.
	GetDouble(key string) (resultFloat64 float64) // function
	// GetString
	//  Returns the value at the specified key as type string.
	GetString(key string) string // function
	// GetBinary
	//  Returns the value at the specified key as type binary. The returned value will reference existing data.
	GetBinary(key string) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified key as type dictionary. The returned value will reference existing data and modifications to the value will modify this object.
	GetDictionary(key string) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified key as type list. The returned value will reference existing data and modifications to the value will modify this object.
	GetList(key string) ICefListValue // function
	// SetValue
	//  Sets the value at the specified key. Returns true (1) if the value was set successfully. If |value| represents simple data then the underlying data will be copied and modifications to |value| will not modify this object. If |value| represents complex data (binary, dictionary or list) then the underlying data will be referenced and modifications to |value| will modify this object.
	SetValue(key string, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified key as type null. Returns true (1) if the value was set successfully.
	SetNull(key string) bool // function
	// SetBool
	//  Sets the value at the specified key as type bool. Returns true (1) if the value was set successfully.
	SetBool(key string, value bool) bool // function
	// SetInt
	//  Sets the value at the specified key as type int. Returns true (1) if the value was set successfully.
	SetInt(key string, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified key as type double. Returns true (1) if the value was set successfully.
	SetDouble(key string, value float64) bool // function
	// SetString
	//  Sets the value at the specified key as type string. Returns true (1) if the value was set successfully.
	SetString(key, value string) bool // function
	// SetBinary
	//  Sets the value at the specified key as type binary. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetBinary(key string, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified key as type dict. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetDictionary(key string, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified key as type list. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetList(key string, value ICefListValue) bool // function
}

// TCefDictionaryValue Parent: TCefBaseRefCounted
//
//	Interface representing a dictionary value. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_dictionary_value_t))
type TCefDictionaryValue struct {
	TCefBaseRefCounted
}

// DictionaryValueRef -> ICefDictionaryValue
var DictionaryValueRef dictionaryValue

// dictionaryValue TCefDictionaryValue Ref
type dictionaryValue uintptr

func (m *dictionaryValue) UnWrap(data uintptr) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(819, uintptr(data), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *dictionaryValue) New() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(808, uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) IsValid() bool {
	r1 := CEF().SysCallN(807, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsOwned() bool {
	r1 := CEF().SysCallN(804, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsReadOnly() bool {
	r1 := CEF().SysCallN(805, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsSame(that ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(806, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsEqual(that ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(803, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) Copy(excludeEmptyChildren bool) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(790, m.Instance(), PascalBool(excludeEmptyChildren), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) GetSize() NativeUInt {
	r1 := CEF().SysCallN(798, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefDictionaryValue) Clear() bool {
	r1 := CEF().SysCallN(789, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) HasKey(key string) bool {
	r1 := CEF().SysCallN(802, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetKeys(keys IStrings) bool {
	r1 := CEF().SysCallN(796, m.Instance(), GetObjectUintptr(keys))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) Remove(key string) bool {
	r1 := CEF().SysCallN(809, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetType(key string) TCefValueType {
	r1 := CEF().SysCallN(800, m.Instance(), PascalStr(key))
	return TCefValueType(r1)
}

func (m *TCefDictionaryValue) GetValue(key string) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(801, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefDictionaryValue) GetBool(key string) bool {
	r1 := CEF().SysCallN(792, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetInt(key string) int32 {
	r1 := CEF().SysCallN(795, m.Instance(), PascalStr(key))
	return int32(r1)
}

func (m *TCefDictionaryValue) GetDouble(key string) (resultFloat64 float64) {
	CEF().SysCallN(794, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefDictionaryValue) GetString(key string) string {
	r1 := CEF().SysCallN(799, m.Instance(), PascalStr(key))
	return GoStr(r1)
}

func (m *TCefDictionaryValue) GetBinary(key string) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	CEF().SysCallN(791, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefDictionaryValue) GetDictionary(key string) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(793, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) GetList(key string) ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(797, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefDictionaryValue) SetValue(key string, value ICefValue) bool {
	r1 := CEF().SysCallN(818, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetNull(key string) bool {
	r1 := CEF().SysCallN(816, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetBool(key string, value bool) bool {
	r1 := CEF().SysCallN(811, m.Instance(), PascalStr(key), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetInt(key string, value int32) bool {
	r1 := CEF().SysCallN(814, m.Instance(), PascalStr(key), uintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetDouble(key string, value float64) bool {
	r1 := CEF().SysCallN(813, m.Instance(), PascalStr(key), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetString(key, value string) bool {
	r1 := CEF().SysCallN(817, m.Instance(), PascalStr(key), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetBinary(key string, value ICefBinaryValue) bool {
	r1 := CEF().SysCallN(810, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetDictionary(key string, value ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(812, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetList(key string, value ICefListValue) bool {
	r1 := CEF().SysCallN(815, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}
