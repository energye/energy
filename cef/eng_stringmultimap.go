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

// ICefStringMultimap Parent: IObject
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type ICefStringMultimap interface {
	IObject
	GetHandle() TCefStringMultimapHandle // function
	// GetSize
	//  Return the number of elements in the string multimap.
	GetSize() NativeUInt // function
	// FindCount
	//  Return the number of values with the specified key.
	FindCount(key string) NativeUInt // function
	// GetEnumerate
	//  Return the value_index-th value with the specified key.
	GetEnumerate(key string, valueIndex NativeUInt) string // function
	// GetKey
	//  Return the key at the specified zero-based string multimap index.
	GetKey(index NativeUInt) string // function
	// GetValue
	//  Return the value at the specified zero-based string multimap index.
	GetValue(index NativeUInt) string // function
	// Append
	//  Append a new key/value pair at the end of the string multimap.
	Append(key, value string) bool // function
	// Clear
	//  Clear the string multimap.
	Clear() // procedure
}

// TCefStringMultimap Parent: TObject
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type TCefStringMultimap struct {
	TObject
}

func NewCefStringMultimap() ICefStringMultimap {
	r1 := CEF().SysCallN(1407)
	return AsCefStringMultimap(r1)
}

func (m *TCefStringMultimap) GetHandle() TCefStringMultimapHandle {
	r1 := CEF().SysCallN(1410, m.Instance())
	return TCefStringMultimapHandle(r1)
}

func (m *TCefStringMultimap) GetSize() NativeUInt {
	r1 := CEF().SysCallN(1412, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefStringMultimap) FindCount(key string) NativeUInt {
	r1 := CEF().SysCallN(1408, m.Instance(), PascalStr(key))
	return NativeUInt(r1)
}

func (m *TCefStringMultimap) GetEnumerate(key string, valueIndex NativeUInt) string {
	r1 := CEF().SysCallN(1409, m.Instance(), PascalStr(key), uintptr(valueIndex))
	return GoStr(r1)
}

func (m *TCefStringMultimap) GetKey(index NativeUInt) string {
	r1 := CEF().SysCallN(1411, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringMultimap) GetValue(index NativeUInt) string {
	r1 := CEF().SysCallN(1413, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringMultimap) Append(key, value string) bool {
	r1 := CEF().SysCallN(1404, m.Instance(), PascalStr(key), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefStringMultimap) Clear() {
	CEF().SysCallN(1405, m.Instance())
}
