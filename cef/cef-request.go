//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Request
package cef

import (
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefRequest) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefRequest) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// IsReadOnly 是否只读
func (m *ICefRequest) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefRequest_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

// SetURL 设置URL
func (m *ICefRequest) SetURL(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetUrl).Call(m.Instance(), api.PascalStr(url))
}

// SetMethod 设置请求方式
func (m *ICefRequest) SetMethod(method string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetMethod).Call(m.Instance(), api.PascalStr(method))
}

// SetReferrer 设置来源策略
func (m *ICefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetReferrer).Call(m.Instance(), api.PascalStr(referrerUrl), policy.ToPtr())
}

// SetFlags 设置请求标记
func (m *ICefRequest) SetFlags(flags TCefUrlRequestFlags) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetFlags).Call(m.Instance(), flags.ToPtr())
}

// SetFirstPartyForCookies
func (m *ICefRequest) SetFirstPartyForCookies(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetFirstPartyForCookies).Call(m.Instance(), api.PascalStr(url))
}

// GetHeaderByName
func (m *ICefRequest) GetHeaderByName(name string) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefRequest_GetHeaderByName).Call(m.Instance(), api.PascalStr(name))
	return api.GoStr(r1)
}

// SetHeaderByName
func (m *ICefRequest) SetHeaderByName(name, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetHeaderByName).Call(m.Instance(), api.PascalStr(name), api.PascalStr(value), api.PascalBool(overwrite))
}

// GetHeaderMap
func (m *ICefRequest) GetHeaderMap() *ICefStringMultiMap {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefRequest_GetHeaderMap).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefStringMultiMap{instance: unsafe.Pointer(result)}
}

func (m *ICefRequest) SetHeaderMap(headerMap *ICefStringMultiMap) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetHeaderMap).Call(m.Instance(), headerMap.Instance())
}

func (m *ICefRequest) GetPostData() *ICefPostData {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefRequest_GetPostData).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefPostData{instance: unsafe.Pointer(result)}
}

func (m *ICefRequest) SetPostData(value *ICefPostData) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefRequest_SetPostData).Call(m.Instance(), value.Instance())
}
