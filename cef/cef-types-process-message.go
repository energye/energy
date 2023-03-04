//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 进程消息 ProcessMessageRef.New()
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ProcessMessageRef -> ICefProcessMessage
var ProcessMessageRef processMessage

// processMessage
type processMessage uintptr

// new 创建一个进程消息类型 - internal
//
// 参数: name 消息名
func (*processMessage) new(name string) *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessageRef_New).Call(api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		name:     name,
		instance: unsafe.Pointer(result),
	}
}

// New 创建一个进程消息类型 - export
//
// 参数: name 消息名
func (m *processMessage) New(name string) *ICefProcessMessage {
	if isInternalKey(name) {
		return nil
	}
	return m.new(name)
}

func (m *processMessage) UnWrap(data *ICefProcessMessage) *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessageRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}

// Instance 实例
func (m *ICefProcessMessage) Instance() uintptr {
	return uintptr(m.instance)
}

// ArgumentList 参数列表
func (m *ICefProcessMessage) ArgumentList() *ICefListValue {
	if m.name == "" {
		return nil
	}
	if m.argumentList == nil {
		var result uintptr
		imports.Proc(internale_CefProcessMessage_ArgumentList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.argumentList = &ICefListValue{
			instance: unsafe.Pointer(result),
		}
	}
	return m.argumentList
}

func (m *ICefProcessMessage) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefProcessMessage_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefProcessMessage) Copy() *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessage_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefProcessMessage) Name() string {
	if m.name == "" {
		r1, _, _ := imports.Proc(internale_CefProcessMessage_Name).Call(m.Instance())
		m.name = api.GoStr(r1)
	}
	return m.name
}

func (m *ICefProcessMessage) Free() {
	m.ArgumentList().Free()
	m.instance = nil
}
