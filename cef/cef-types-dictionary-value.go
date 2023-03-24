//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//  CEF 字典类型 DictionaryValueRef.New()
package cef

import (
	"github.com/energye/energy/cef/ipc"
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// DictionaryValueRef -> ICefDictionaryValue
var DictionaryValueRef cefDictionaryValue

//cefDictionaryValue
type cefDictionaryValue uintptr

// New 创建一个字典类型
func (*cefDictionaryValue) New() *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

// data
func (*cefDictionaryValue) UnWrap(data *ICefDictionaryValue) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
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
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) IsOwned() bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) IsReadOnly() bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) Copy(excludeEmptyChildren bool) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValue_Copy).Call(m.Instance(), api.PascalBool(excludeEmptyChildren), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) Size() uint32 {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefDictionaryValue) Clear() (result bool) {
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
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_Clear).Call(m.Instance())
	result = api.GoBool(r1)
	return
}

func (m *ICefDictionaryValue) HasKey(key string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_HasKey).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetKeys() *ICefV8ValueKeys {
	var result uintptr
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetKeys).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8ValueKeys{keys: lcl.AsStrings(result), count: int(int32(r1))}
}

func (m *ICefDictionaryValue) GetIKeys() ipc.IV8ValueKeys {
	return m.GetKeys()
}

func (m *ICefDictionaryValue) Remove(key string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_Remove).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetType(key string) consts.TCefValueType {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetType).Call(m.Instance(), api.PascalStr(key))
	return consts.TCefValueType(r1)
}

func (m *ICefDictionaryValue) GetValue(key string) *ICefValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValue_GetValue).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIValue(key string) ipc.IValue {
	return m.GetValue(key)
}

func (m *ICefDictionaryValue) GetBool(key string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetBool).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) GetInt(key string) int32 {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetInt).Call(m.Instance(), api.PascalStr(key))
	return int32(r1)
}

func (m *ICefDictionaryValue) GetDouble(key string) (result float64) {
	imports.Proc(internale_CefDictionaryValue_GetDouble).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefDictionaryValue) GetString(key string) string {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetString).Call(m.Instance(), api.PascalStr(key))
	return api.GoStr(r1)
}

func (m *ICefDictionaryValue) GetBinary(key string) *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValue_GetBinary).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIBinary(key string) ipc.IBinaryValue {
	return m.GetBinary(key)
}

func (m *ICefDictionaryValue) GetDictionary(key string) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValue_GetDictionary).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIObject(key string) ipc.IObjectValue {
	return m.GetDictionary(key)
}

func (m *ICefDictionaryValue) GetList(key string) *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefDictionaryValue_GetList).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefDictionaryValue) GetIArray(key string) ipc.IArrayValue {
	return m.GetList(key)
}

func (m *ICefDictionaryValue) SetNull(key string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetNull).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetBool(key string, value bool) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetBool).Call(m.Instance(), api.PascalStr(key), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetInt(key string, value int32) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetInt).Call(m.Instance(), api.PascalStr(key), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetDouble(key string, value float64) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetDouble).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetString(key string, value string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetString).Call(m.Instance(), api.PascalStr(key), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetValue(key string, value *ICefValue) (result bool) {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetValue).Call(m.Instance(), api.PascalStr(key), value.Instance())
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
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetBinary).Call(m.Instance(), api.PascalStr(key), value.Instance())
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
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetDictionary).Call(m.Instance(), api.PascalStr(key), value.Instance())
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
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetList).Call(m.Instance(), api.PascalStr(key), value.Instance())
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
