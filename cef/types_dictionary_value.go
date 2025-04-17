//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefDictionaryValue -> ArgumentList
type ICefDictionaryValue struct {
	base             TCefBaseRefCounted
	instance         unsafe.Pointer
	values           map[string]*ICefValue
	binaryValues     map[string]*ICefBinaryValue
	dictionaryValues map[string]*ICefDictionaryValue
	listValues       map[string]*ICefListValue
}

// DictionaryValueRef -> ICefDictionaryValue
var DictionaryValueRef cefDictionaryValue

// cefDictionaryValue
type cefDictionaryValue uintptr

// New 创建一个字典类型
func (*cefDictionaryValue) New() *ICefDictionaryValue {
	var result uintptr
	imports.Proc(def.CefDictionaryValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

// UnWrap
func (*cefDictionaryValue) UnWrap(data *ICefDictionaryValue) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(def.CefDictionaryValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefDictionaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDictionaryValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) Copy(excludeEmptyChildren bool) *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDictionaryValue_Copy).Call(m.Instance(), api.PascalBool(excludeEmptyChildren), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) Size() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefDictionaryValue) Clear() (result bool) {
	if !m.IsValid() {
		return false
	}
	if m.values != nil {
		for _, v := range m.values {
			if v != nil && v.instance != nil {
				v.Free()
			}
		}
		m.values = nil
	}
	if m.binaryValues != nil {
		for _, v := range m.binaryValues {
			if v != nil && v.instance != nil {
				v.Free()
			}
		}
		m.binaryValues = nil
	}
	if m.dictionaryValues != nil {
		for _, v := range m.dictionaryValues {
			if v != nil && v.instance != nil {
				v.Free()
			}
		}
		m.dictionaryValues = nil
	}
	if m.listValues != nil {
		for _, v := range m.listValues {
			if v != nil && v.instance != nil {
				v.Free()
			}
		}
		m.listValues = nil
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_Clear).Call(m.Instance())
	result = api.GoBool(r1)
	return
}

func (m *ICefDictionaryValue) HasKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_HasKey).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetKeys() *ICefV8ValueKeys {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	r1, _, _ := imports.Proc(def.CefDictionaryValue_GetKeys).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8ValueKeys{keys: lcl.AsStrings(result), count: int(int32(r1))}
}

func (m *ICefDictionaryValue) GetIKeys() types.IV8ValueKeys {
	if !m.IsValid() {
		return nil
	}
	return m.GetKeys()
}

func (m *ICefDictionaryValue) Remove(key string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_Remove).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetType(key string) consts.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_GetType).Call(m.Instance(), api.PascalStr(key))
	return consts.TCefValueType(r1)
}

func (m *ICefDictionaryValue) GetValue(key string) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDictionaryValue_GetValue).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIValue(key string) types.IValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetValue(key)
}

func (m *ICefDictionaryValue) GetBool(key string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_GetBool).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetInt(key string) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_GetInt).Call(m.Instance(), api.PascalStr(key))
	return int32(r1)
}

func (m *ICefDictionaryValue) GetDouble(key string) (result float64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CefDictionaryValue_GetDouble).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefDictionaryValue) GetString(key string) (value string) {
	if !m.IsValid() {
		return ""
	}
	val := NewTString()
	imports.Proc(def.CefDictionaryValue_GetString).Call(m.Instance(), api.PascalStr(key), val.Instance())
	value = val.Value()
	val.Free()
	return
}

func (m *ICefDictionaryValue) GetBinary(key string) *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDictionaryValue_GetBinary).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIBinary(key string) types.IBinaryValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetBinary(key)
}

func (m *ICefDictionaryValue) GetDictionary(key string) *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDictionaryValue_GetDictionary).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIObject(key string) types.IObjectValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetDictionary(key)
}

func (m *ICefDictionaryValue) GetList(key string) *ICefListValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDictionaryValue_GetList).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIArray(key string) types.IArrayValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetList(key)
}

func (m *ICefDictionaryValue) SetNull(key string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetNull).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetBool(key string, value bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetBool).Call(m.Instance(), api.PascalStr(key), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetInt(key string, value int32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetInt).Call(m.Instance(), api.PascalStr(key), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetDouble(key string, value float64) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetDouble).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetString(key string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetString).Call(m.Instance(), api.PascalStr(key), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetValue(key string, value *ICefValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetValue).Call(m.Instance(), api.PascalStr(key), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.values == nil {
			m.values = make(map[string]*ICefValue)
		}
		if v, ok := m.values[key]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.values[key] = value
	}
	return
}

func (m *ICefDictionaryValue) SetBinary(key string, value *ICefBinaryValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetBinary).Call(m.Instance(), api.PascalStr(key), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.binaryValues == nil {
			m.binaryValues = make(map[string]*ICefBinaryValue)
		}
		if v, ok := m.binaryValues[key]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.binaryValues[key] = value
	}
	return
}

func (m *ICefDictionaryValue) SetDictionary(key string, value *ICefDictionaryValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetDictionary).Call(m.Instance(), api.PascalStr(key), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.dictionaryValues == nil {
			m.dictionaryValues = make(map[string]*ICefDictionaryValue)
		}
		if v, ok := m.dictionaryValues[key]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.dictionaryValues[key] = value
	}
	return
}

func (m *ICefDictionaryValue) SetList(key string, value *ICefListValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDictionaryValue_SetList).Call(m.Instance(), api.PascalStr(key), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.listValues == nil {
			m.listValues = make(map[string]*ICefListValue)
		}
		if v, ok := m.listValues[key]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.listValues[key] = value
	}
	return
}

func (m *ICefDictionaryValue) Free() {
	if m.instance != nil {
		m.Clear()
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
