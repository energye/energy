//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 多值MAP类型 StringMultiMapRef.New

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefStringMultiMap 实例
type ICefStringMultiMap struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// StringMultiMapRef -> ICefStringMultiMap
var StringMultiMapRef stringMultiMap

// stringMultiMap
type stringMultiMap uintptr

// New 创建一个新的 StringMultiMap
func (m *stringMultiMap) New() *ICefStringMultiMap {
	var result uintptr
	imports.Proc(def.StringMultimapRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefStringMultiMap{instance: unsafe.Pointer(result)}
}

// Instance 实例
func (m *ICefStringMultiMap) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefStringMultiMap) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// GetSize 大小
func (m *ICefStringMultiMap) GetSize() uint32 {
	r1, _, _ := imports.Proc(def.StringMultimap_GetSize).Call(m.Instance())
	return uint32(r1)
}

// FindCount key值数量
func (m *ICefStringMultiMap) FindCount(key string) uint32 {
	r1, _, _ := imports.Proc(def.StringMultimap_FindCount).Call(m.Instance(), api.PascalStr(key))
	return uint32(r1)
}

// GetEnumerate 根据 key and index 获取
func (m *ICefStringMultiMap) GetEnumerate(key string, index uint32) (r string) {
	value := NewTString()
	imports.Proc(def.StringMultimap_GetEnumerate).Call(m.Instance(), api.PascalStr(key), uintptr(index), value.Instance())
	r = value.Value()
	value.Free()
	return
}

// GetKey 根据 index 获取key
func (m *ICefStringMultiMap) GetKey(index uint32) (key string) {
	tKey := NewTString()
	imports.Proc(def.StringMultimap_GetKey).Call(m.Instance(), uintptr(index), tKey.Instance())
	key = tKey.Value()
	tKey.Free()
	return
}

// GetValue 根据 index 获取value
func (m *ICefStringMultiMap) GetValue(index uint32) (value string) {
	tValue := NewTString()
	imports.Proc(def.StringMultimap_GetValue).Call(m.Instance(), uintptr(index), tValue.Instance())
	value = tValue.Value()
	tValue.Free()
	return
}

// Append 给key追加值
func (m *ICefStringMultiMap) Append(key, value string) bool {
	r1, _, _ := imports.Proc(def.StringMultimap_Append).Call(m.Instance(), api.PascalStr(key), api.PascalStr(value))
	return api.GoBool(r1)
}

// Clear 清空
func (m *ICefStringMultiMap) Clear() {
	imports.Proc(def.StringMultimap_Clear).Call(m.Instance())
}

func (m *ICefStringMultiMap) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
