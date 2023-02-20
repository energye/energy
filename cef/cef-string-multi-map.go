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
	instance uintptr
	ptr      unsafe.Pointer
}

// GetSize 大小
func (m *ICefStringMultiMap) GetSize() int {
	return cefHeaderMap_GetSize(m.instance)
}

// FindCount key值数量
func (m *ICefStringMultiMap) FindCount(key string) int {
	return cefHeaderMap_FindCount(m.instance, key)
}

// GetEnumerate 根据key & index获取枚举
func (m *ICefStringMultiMap) GetEnumerate(key string, valueIndex int) string {
	return api.GoStr(cefHeaderMap_GetEnumerate(m.instance, key, valueIndex))
}

// GetKey 根据 index 获取key
func (m *ICefStringMultiMap) GetKey(index int) string {
	return api.GoStr(cefHeaderMap_GetKey(m.instance, index))
}

// GetValue 根据 index 获取value
func (m *ICefStringMultiMap) GetValue(index int) string {
	return api.GoStr(cefHeaderMap_GetValue(m.instance, index))
}

// Append 给key追加值
func (m *ICefStringMultiMap) Append(key, value string) bool {
	return api.GoBool(cefHeaderMap_Append(m.instance, key, value))
}

// Clear 清空
func (m *ICefStringMultiMap) Clear() {
	cefHeaderMap_Clear(m.instance)
}

func cefHeaderMap_GetSize(instance uintptr) int {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_GetSize).Call(instance)
	return int(r1)
}
func cefHeaderMap_FindCount(instance uintptr, key string) int {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_FindCount).Call(instance, api.PascalStr(key))
	return int(r1)
}
func cefHeaderMap_GetEnumerate(instance uintptr, key string, valueIndex int) uintptr {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_GetEnumerate).Call(instance, api.PascalStr(key), uintptr(valueIndex))
	return r1
}
func cefHeaderMap_GetKey(instance uintptr, index int) uintptr {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_GetKey).Call(instance, uintptr(index))
	return r1
}
func cefHeaderMap_GetValue(instance uintptr, index int) uintptr {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_GetValue).Call(instance, uintptr(index))
	return r1
}
func cefHeaderMap_Append(instance uintptr, key, value string) uintptr {
	r1, _, _ := imports.Proc(internale_cefHeaderMap_Append).Call(instance, api.PascalStr(key), api.PascalStr(value))
	return r1
}
func cefHeaderMap_Clear(instance uintptr) {
	imports.Proc(internale_cefHeaderMap_Clear).Call(instance)
}
