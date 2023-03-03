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
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
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

func (m *ICefDictionaryValue) GetSize() uint32 {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefDictionaryValue) Clear() bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_Clear).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) HasKey(key string) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_HasKey).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

//func (m *ICefDictionaryValue) GetKeys() bool {
//
//}

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

func (m *ICefDictionaryValue) SetValue(key string, value *ICefValue) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetValue).Call(m.Instance(), api.PascalStr(key), value.Instance())
	return api.GoBool(r1)
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

func (m *ICefDictionaryValue) SetBinary(key string, value *ICefBinaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetBinary).Call(m.Instance(), api.PascalStr(key), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetDictionary(key string, value *ICefDictionaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetDictionary).Call(m.Instance(), api.PascalStr(key), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) SetList(key string, value *ICefListValue) bool {
	r1, _, _ := imports.Proc(internale_CefDictionaryValue_SetList).Call(m.Instance(), api.PascalStr(key), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefDictionaryValue) Free() {
	m.Clear()
	m.instance = nil
}
