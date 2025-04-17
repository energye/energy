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

// ICefProcessMessage
type ICefProcessMessage struct {
	base         TCefBaseRefCounted
	instance     unsafe.Pointer
	argumentList *ICefListValue
	name         string
}

// ProcessMessageRef -> ICefProcessMessage
var ProcessMessageRef processMessage

// processMessage
type processMessage uintptr

// new 创建一个进程消息类型 - internal
//
// 参数: name 消息名
func (*processMessage) new(name string) *ICefProcessMessage {
	var result uintptr
	imports.Proc(def.CefProcessMessageRef_New).Call(api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		name:     name,
		instance: unsafe.Pointer(result),
	}
}

// New 创建一个进程消息类型 - export
//
// 参数: name 消息名
func (m *processMessage) New(name string) *ICefProcessMessage {
	if isIPCInternalKey(name) {
		return nil
	}
	return m.new(name)
}

func (m *processMessage) UnWrap(data *ICefProcessMessage) *ICefProcessMessage {
	var result uintptr
	imports.Proc(def.CefProcessMessageRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefProcessMessage) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefProcessMessage) Free() {
	if m != nil && m.instance != nil {
		m.ArgumentList().Free()
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefProcessMessage) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefProcessMessage_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// ArgumentList 参数列表
func (m *ICefProcessMessage) ArgumentList() *ICefListValue {
	if m.Name() == "" || !m.IsValid() {
		return nil
	}
	if m.argumentList == nil {
		var result uintptr
		imports.Proc(def.CefProcessMessage_ArgumentList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.argumentList = &ICefListValue{
			instance: unsafe.Pointer(result),
		}
	}
	return m.argumentList
}

func (m *ICefProcessMessage) Copy() *ICefProcessMessage {
	var result uintptr
	imports.Proc(def.CefProcessMessage_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefProcessMessage) Name() string {
	if m.name == "" {
		r1, _, _ := imports.Proc(def.CefProcessMessage_Name).Call(m.Instance())
		m.name = api.GoStr(r1)
	}
	return m.name
}
