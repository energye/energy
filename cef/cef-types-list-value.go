//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 集合类型 ListValueRef.New
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ListValueRef -> ICefListValue
var ListValueRef cefListValue

//cefListValue
type cefListValue uintptr

// New 创建一个集合类型
func (*cefListValue) New() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefListValue_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefListValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefListValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) IsOwned() bool {
	r1, _, _ := imports.Proc(internale_CefListValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) IsReadOnly() bool {
	r1, _, _ := imports.Proc(internale_CefListValue_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) Copy() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefListValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) SetSize(size types.NativeUInt) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetSize).Call(m.Instance(), size.ToPtr())
	return api.GoBool(r1)
}

func (m *ICefListValue) Size() uint32 {
	r1, _, _ := imports.Proc(internale_CefListValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefListValue) Clear() bool {
	r1, _, _ := imports.Proc(internale_CefListValue_Clear).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) Remove(index types.NativeUInt) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_Remove).Call(m.Instance(), index.ToPtr())
	return api.GoBool(r1)
}

func (m *ICefListValue) GetType(index types.NativeUInt) consts.TCefValueType {
	r1, _, _ := imports.Proc(internale_CefListValue_GetType).Call(m.Instance(), index.ToPtr())
	return consts.TCefValueType(r1)
}

func (m *ICefListValue) GetValue(index types.NativeUInt) *ICefValue {
	var result uintptr
	imports.Proc(internale_CefListValue_GetValue).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIValue(index types.NativeUInt) ipc.IValue {
	return m.GetValue(index)
}

func (m *ICefListValue) GetBool(index types.NativeUInt) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_GetBool).Call(m.Instance(), index.ToPtr())
	return api.GoBool(r1)
}

func (m *ICefListValue) GetInt(index types.NativeUInt) int32 {
	r1, _, _ := imports.Proc(internale_CefListValue_GetInt).Call(m.Instance(), index.ToPtr())
	return int32(r1)
}

func (m *ICefListValue) GetDouble(index types.NativeUInt) (result float64) {
	imports.Proc(internale_CefListValue_GetDouble).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefListValue) GetString(index types.NativeUInt) string {
	r1, _, _ := imports.Proc(internale_CefListValue_GetString).Call(m.Instance(), index.ToPtr())
	return api.GoStr(r1)
}

func (m *ICefListValue) GetBinary(index types.NativeUInt) *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefListValue_GetBinary).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIBinary(index types.NativeUInt) ipc.IBinaryValue {
	return m.GetBinary(index)
}

func (m *ICefListValue) GetDictionary(index types.NativeUInt) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(internale_CefListValue_GetDictionary).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIObject(index types.NativeUInt) ipc.IObjectValue {
	return m.GetDictionary(index)
}

func (m *ICefListValue) GetList(index types.NativeUInt) *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefListValue_GetList).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefListValue) GetIArray(index types.NativeUInt) ipc.IArrayValue {
	return m.GetList(index)
}

func (m *ICefListValue) SetValue(index types.NativeUInt, value *ICefValue) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetValue).Call(m.Instance(), index.ToPtr(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) SetNull(index types.NativeUInt) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetNull).Call(m.Instance(), index.ToPtr())
	return api.GoBool(r1)
}

func (m *ICefListValue) SetBool(index types.NativeUInt, value bool) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetBool).Call(m.Instance(), index.ToPtr(), api.PascalBool(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetInt(index types.NativeUInt, value int32) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetInt).Call(m.Instance(), index.ToPtr(), uintptr(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetDouble(index types.NativeUInt, value float64) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetDouble).Call(m.Instance(), index.ToPtr(), uintptr(unsafe.Pointer(&value)))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetString(index types.NativeUInt, value string) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetString).Call(m.Instance(), index.ToPtr(), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefListValue) SetBinary(index types.NativeUInt, value *ICefBinaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetBinary).Call(m.Instance(), index.ToPtr(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) SetDictionary(index types.NativeUInt, value *ICefDictionaryValue) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetDictionary).Call(m.Instance(), index.ToPtr(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) SetList(index types.NativeUInt, value *ICefListValue) bool {
	r1, _, _ := imports.Proc(internale_CefListValue_SetList).Call(m.Instance(), index.ToPtr(), value.Instance())
	return api.GoBool(r1)
}

func (m *ICefListValue) Free() {
	m.Clear()
	m.instance = nil
}
