//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 多值MAP类型
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefStringMultiMap 实例
type ICefStringMultiMap struct {
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefStringMultiMap) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// GetSize 大小
func (m *ICefStringMultiMap) GetSize() int32 {
	r1, _, _ := imports.Proc(internale_StringMultimap_GetSize).Call(m.Instance())
	return int32(r1)
}

// FindCount key值数量
func (m *ICefStringMultiMap) FindCount(key string) int32 {
	r1, _, _ := imports.Proc(internale_StringMultimap_FindCount).Call(m.Instance(), api.PascalStr(key))
	return int32(r1)
}

// GetEnumerate 根据 key and index 获取
func (m *ICefStringMultiMap) GetEnumerate(key string, valueIndex int32) string {
	r1, _, _ := imports.Proc(internale_StringMultimap_GetEnumerate).Call(m.Instance(), api.PascalStr(key), uintptr(valueIndex))
	return api.GoStr(r1)
}

// GetKey 根据 index 获取key
func (m *ICefStringMultiMap) GetKey(index int32) string {
	r1, _, _ := imports.Proc(internale_StringMultimap_GetKey).Call(m.Instance(), uintptr(index))
	return api.GoStr(r1)
}

// GetValue 根据 index 获取value
func (m *ICefStringMultiMap) GetValue(index int32) string {
	r1, _, _ := imports.Proc(internale_StringMultimap_GetValue).Call(m.Instance(), uintptr(index))
	return api.GoStr(r1)
}

// Append 给key追加值
func (m *ICefStringMultiMap) Append(key, value string) bool {
	r1, _, _ := imports.Proc(internale_StringMultimap_Append).Call(m.Instance(), api.PascalStr(key), api.PascalStr(value))
	return api.GoBool(r1)
}

// Clear 清空
func (m *ICefStringMultiMap) Clear() {
	imports.Proc(internale_StringMultimap_Clear).Call(m.Instance())
}
