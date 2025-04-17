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
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type ICefBaseRefCounted interface {
	Instance() uintptr
	IsValid() bool
	BaseRefCounted(instance uintptr) *TCefBaseRefCounted
}

func NewBaseRefCounted(instance uintptr) *TCefBaseRefCounted {
	return &TCefBaseRefCounted{instance: unsafe.Pointer(instance)}
}

// BaseRefCounted
// 将实例转换为 TCefBaseRefCounted
func (m *TCefBaseRefCounted) BaseRefCounted(instance uintptr) *TCefBaseRefCounted {
	m.instance = unsafe.Pointer(instance)
	return m
}

// Wrap 指针引用包裹
// 调用以增加对象的引用计数。应该为指向给定对象的指针的每个新副本调用。
func (m *TCefBaseRefCounted) Wrap(data uintptr) unsafe.Pointer {
	var result uintptr
	imports.Proc(def.CefBaseRefCounted_Wrap).Call(data, uintptr(unsafe.Pointer(&result)))
	return unsafe.Pointer(result)
}

// Free 释放底层指针
func (m *TCefBaseRefCounted) Free(data uintptr) {
	imports.Proc(def.CefBaseRefCounted_Free).Call(uintptr(unsafe.Pointer(&data)))
	m.instance = nil
}

// SameAs
// 将aData指针与当前实例的FData字段进行比较。
func (m *TCefBaseRefCounted) SameAs(data uintptr) bool {
	r1, _, _ := imports.Proc(def.CefBaseRefCounted_SameAs).Call(m.Instance(), data)
	return api.GoBool(r1)
}

// HasOneRef
// 如果当前引用计数为1，则返回true（1）。
func (m *TCefBaseRefCounted) HasOneRef() bool {
	r1, _, _ := imports.Proc(def.CefBaseRefCounted_HasOneRef).Call(m.Instance())
	return api.GoBool(r1)
}

// HasAtLeastOneRef
// 如果当前引用计数至少为1，则返回true（1）。
func (m *TCefBaseRefCounted) HasAtLeastOneRef() bool {
	r1, _, _ := imports.Proc(def.CefBaseRefCounted_HasAtLeastOneRef).Call(m.Instance())
	return api.GoBool(r1)
}

// DestroyOtherRefs
// 释放所有其他实例。
func (m *TCefBaseRefCounted) DestroyOtherRefs() {
	imports.Proc(def.CefBaseRefCounted_DestroyOtherRefs).Call(m.Instance())
}

// Instance 实例
func (m *TCefBaseRefCounted) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefBaseRefCounted) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
