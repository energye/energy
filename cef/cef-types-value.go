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
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
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

func (m *ICefValue) GetIBinary() ipc.IBinaryValue {
	return m.GetBinary()
}

func (m *ICefValue) GetDictionary() *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefValue_GetDictionary).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIObject() ipc.IObjectValue {
	return m.GetDictionary()
}

func (m *ICefValue) GetList() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefValue_GetList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefValue) GetIArray() ipc.IArrayValue {
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

func (m *ICefValue) SetBinary(value *ICefBinaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetBinary).Call(m.Instance(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) SetDictionary(value *ICefDictionaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetDictionary).Call(m.Instance(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) SetList(value *ICefListValue) bool {
	r1, _, _ := imports.Proc(internale_CefValue_SetList).Call(m.Instance(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefValue) Free() {
	m.instance = nil
}
