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
	"github.com/energye/energy/cef/ipc/types"
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

//ValueRef -> ICefValue
var ValueRef cefValue

//cefValue
type cefValue uintptr

func (*cefValue) New() *ICefValue {
	var result uintptr
	imports.Proc(internale_CefValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (*cefValue) UnWrap(data *ICefValue) *ICefValue {
	var result uintptr
	imports.Proc(internale_CefValueRef_UnWrap).Call(uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
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
	r1, _, _ := imports.Proc(internale_CefValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) IsOwned() bool {
	r1, _, _ := imports.Proc(internale_CefValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) IsReadOnly() bool {
	r1, _, _ := imports.Proc(internale_CefValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) Copy() *ICefValue {
	var result uintptr
	imports.Proc(internale_CefValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetType() consts.TCefValueType {
	r1, _, _ := imports.Proc(internale_CefValue_GetType).Call(m.Instance())
	return consts.TCefValueType(r1)
}

func (m *ICefValue) GetBool() bool {
	r1, _, _ := imports.Proc(internale_CefValue_GetBool).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) GetInt() int32 {
	r1, _, _ := imports.Proc(internale_CefValue_GetInt).Call(m.Instance())
	return int32(r1)
}

func (m *ICefValue) GetDouble() (result float64) {
	imports.Proc(internale_CefValue_GetDouble).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefValue) GetString() string {
	r1, _, _ := imports.Proc(internale_CefValue_GetString).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefValue) GetBinary() *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefValue_GetBinary).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIBinary() types.IBinaryValue {
	return m.GetBinary()
}

func (m *ICefValue) GetDictionary() *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefValue_GetDictionary).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIObject() types.IObjectValue {
	return m.GetDictionary()
}

func (m *ICefValue) GetList() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefValue_GetList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIArray() types.IArrayValue {
	return m.GetList()
}

func (m *ICefValue) SetNull() bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetNull).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) SetBool(value bool) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetBool).Call(m.Instance(), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetInt(value int32) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetInt).Call(m.Instance(), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetDouble(value float64) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetDouble).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefValue) SetString(value string) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetString).Call(m.Instance(), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefValue) SetBinary(value *ICefBinaryValue) (result bool) {
	r1, _, _ := imports.Proc(internale_CefValue_SetBinary).Call(m.Instance(), value.Instance())
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
	r1, _, _ := imports.Proc(internale_CefValue_SetDictionary).Call(m.Instance(), value.Instance())
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
	r1, _, _ := imports.Proc(internale_CefValue_SetList).Call(m.Instance(), value.Instance())
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
