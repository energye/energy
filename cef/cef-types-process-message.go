//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 进程消息 ProcessMessage.New()
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ProcessMessage -> ICefProcessMessage
var ProcessMessage processMessageRef

//processMessageRef
type processMessageRef uintptr

// New 创建一个进程消息类型
//
// 参数: name 消息名
func (*processMessageRef) New(name string) *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessageRef_New).Call(api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		instance: unsafe.Pointer(result),
	}
}
func (m *ICefProcessMessage) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefProcessMessage) ArgumentList() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefProcessMessage_ArgumentList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
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
	r1, _, _ := imports.Proc(internale_CefProcessMessage_Name).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefProcessMessage) Free() {
	m.ArgumentList().Free()
	m.instance = nil
}
