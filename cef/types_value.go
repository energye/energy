//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Value 所有值类型 ValueRef.New

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefValue -> ArgumentList
type ICefValue struct {
	base            TCefBaseRefCounted
	instance        unsafe.Pointer
	binaryValue     *ICefBinaryValue
	dictionaryValue *ICefDictionaryValue
	listValue       *ICefListValue
}

// ValueRef -> ICefValue
var ValueRef cefValue

// cefValue
type cefValue uintptr

func (*cefValue) New() *ICefValue {
	var result uintptr
	imports.Proc(def.CefValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (*cefValue) UnWrap(data *ICefValue) *ICefValue {
	var result uintptr
	imports.Proc(def.CefValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) Copy() *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetType() consts.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefValue_GetType).Call(m.Instance())
	return consts.TCefValueType(r1)
}

func (m *ICefValue) GetBool() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_GetBool).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) GetInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefValue_GetInt).Call(m.Instance())
	return int32(r1)
}

func (m *ICefValue) GetDouble() (result float64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CefValue_GetDouble).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefValue) GetString() (value string) {
	if !m.IsValid() {
		return ""
	}
	val := NewTString()
	imports.Proc(def.CefValue_GetString).Call(m.Instance(), val.Instance())
	value = val.Value()
	val.Free()
	return
}

func (m *ICefValue) GetBinary() *ICefBinaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefValue_GetBinary).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIBinary() types.IBinaryValue {
	return m.GetBinary()
}

func (m *ICefValue) GetDictionary() *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefValue_GetDictionary).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIObject() types.IObjectValue {
	if !m.IsValid() {
		return nil
	}
	return m.GetDictionary()
}

func (m *ICefValue) GetList() *ICefListValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefValue_GetList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIArray() types.IArrayValue {
	return m.GetList()
}

func (m *ICefValue) SetNull() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetNull).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) SetBool(value bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetBool).Call(m.Instance(), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetInt(value int32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetInt).Call(m.Instance(), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetDouble(value float64) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetDouble).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefValue) SetString(value string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetString).Call(m.Instance(), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetBinary(value *ICefBinaryValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetBinary).Call(m.Instance(), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.binaryValue != nil && m.binaryValue.instance != nil {
			m.binaryValue.Free()
		}
		m.binaryValue = value
	}
	return
}

func (m *ICefValue) SetDictionary(value *ICefDictionaryValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetDictionary).Call(m.Instance(), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.dictionaryValue != nil && m.dictionaryValue.instance != nil {
			m.dictionaryValue.Free()
		}
		m.dictionaryValue = value
	}
	return
}

func (m *ICefValue) SetList(value *ICefListValue) (result bool) {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefValue_SetList).Call(m.Instance(), value.Instance())
	result = api.GoBool(r1)
	if result {
		if m.listValue != nil && m.listValue.instance != nil {
			m.listValue.Free()
		}
		m.listValue = value
	}
	return
}

func (m *ICefValue) Free() {
	if m.instance != nil {
		if m.binaryValue != nil {
			m.binaryValue.Free()
			m.binaryValue = nil
		}
		if m.dictionaryValue != nil {
			m.dictionaryValue.Free()
			m.dictionaryValue = nil
		}
		if m.listValue != nil {
			m.listValue.Free()
			m.listValue = nil
		}
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
