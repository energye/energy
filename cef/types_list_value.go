//----------------------------------------
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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefListValue -> ArgumentList
type ICefListValue struct {
	base             TCefBaseRefCounted
	instance         unsafe.Pointer
	values           map[int]*ICefValue
	binaryValues     map[int]*ICefBinaryValue
	dictionaryValues map[int]*ICefDictionaryValue
	listValues       map[int]*ICefListValue
}

// ListValueRef -> ICefListValue
var ListValueRef listValue

// cefListValue
type listValue uintptr

// New 创建一个集合类型
func (*listValue) New() *ICefListValue {
	var result uintptr
	imports.Proc(def.CefListValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (*listValue) UnWrap(data *ICefListValue) *ICefListValue {
	var result uintptr
	imports.Proc(def.CefListValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

func (m *ICefListValue) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefListValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) Copy() *ICefListValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefListValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) SetSize(size uint32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetSize).Call(m.Instance(), uintptr(size))
	return api.GoBool(r1)
}

func (m *ICefListValue) Size() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefListValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefListValue) Clear() (result bool) {
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
	if m.instance != nil {
		r1, _, _ := imports.Proc(def.CefListValue_Clear).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return
}

func (m *ICefListValue) Remove(index uint32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_Remove).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

func (m *ICefListValue) GetType(index uint32) consts.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefListValue_GetType).Call(m.Instance(), uintptr(index))
	return consts.TCefValueType(r1)
}

func (m *ICefListValue) GetValue(index uint32) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefListValue_GetValue).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIValue(index uint32) types.IValue {
	return m.GetValue(index)
}

func (m *ICefListValue) GetBool(index uint32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_GetBool).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

func (m *ICefListValue) GetInt(index uint32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefListValue_GetInt).Call(m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *ICefListValue) GetDouble(index uint32) (result float64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CefListValue_GetDouble).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefListValue) GetString(index uint32) (value string) {
	if !m.IsValid() {
		return ""
	}
	val := NewTString()
	imports.Proc(def.CefListValue_GetString).Call(m.Instance(), uintptr(index), val.Instance())
	value = val.Value()
	val.Free()
	return
}

func (m *ICefListValue) GetBinary(index uint32) *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefListValue_GetBinary).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIBinary(index uint32) types.IBinaryValue {
	return m.GetBinary(index)
}

func (m *ICefListValue) GetDictionary(index uint32) *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefListValue_GetDictionary).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIObject(index uint32) types.IObjectValue {
	return m.GetDictionary(index)
}

func (m *ICefListValue) GetList(index uint32) *ICefListValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefListValue_GetList).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIArray(index uint32) types.IArrayValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetList(index)
}

func (m *ICefListValue) SetNull(index uint32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetNull).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetBool(index uint32, value bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetBool).Call(m.Instance(), uintptr(index), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetInt(index uint32, value int32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetInt).Call(m.Instance(), uintptr(index), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetDouble(index uint32, value float64) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetDouble).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetString(index uint32, value string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetString).Call(m.Instance(), uintptr(index), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetValue(index uint32, value *ICefValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetValue).Call(m.Instance(), uintptr(index), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.values == nil {
			m.values = make(map[int]*ICefValue)
		}
		if v, ok := m.values[int(index)]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.values[int(index)] = value
	}
	return
}

func (m *ICefListValue) SetBinary(index uint32, value *ICefBinaryValue) (result bool) {
	if !m.IsValid() || value == nil || value.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetBinary).Call(m.Instance(), uintptr(index), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.binaryValues == nil {
			m.binaryValues = make(map[int]*ICefBinaryValue)
		}
		if v, ok := m.binaryValues[int(index)]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.binaryValues[int(index)] = value
	}
	return
}

func (m *ICefListValue) SetDictionary(index uint32, value *ICefDictionaryValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetDictionary).Call(m.Instance(), uintptr(index), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.dictionaryValues == nil {
			m.dictionaryValues = make(map[int]*ICefDictionaryValue)
		}
		if v, ok := m.dictionaryValues[int(index)]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.dictionaryValues[int(index)] = value
	}
	return
}

func (m *ICefListValue) SetList(index uint32, value *ICefListValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefListValue_SetList).Call(m.Instance(), uintptr(index), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.listValues == nil {
			m.listValues = make(map[int]*ICefListValue)
		}
		if v, ok := m.listValues[int(index)]; ok && v != nil && v.instance != nil {
			v.Free()
		}
		m.listValues[int(index)] = value
	}
	return
}

func (m *ICefListValue) Free() {
	if m != nil && m.instance != nil {
		m.Clear()
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
