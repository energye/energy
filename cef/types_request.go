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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefRequest
type ICefRequest struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// RequestRef -> ICefRequest
var RequestRef request

// request
type request uintptr

func (*request) New() *ICefRequest {
	var result uintptr
	imports.Proc(def.CefRequestRef_New).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequest{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*request) UnWrap(data *ICefRequest) *ICefRequest {
	var result uintptr
	imports.Proc(def.CefRequestRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

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
	r1, _, _ := imports.Proc(def.CefRequest_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefRequest) URL() (r string) {
	if !m.IsValid() {
		return ""
	}
	value := NewTString()
	imports.Proc(def.CefRequest_GetUrl).Call(m.Instance(), value.Instance())
	r = value.Value()
	value.Free()
	return
}

// SetURL 设置URL
func (m *ICefRequest) SetURL(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetUrl).Call(m.Instance(), api.PascalStr(url))
}

func (m *ICefRequest) Method() (r string) {
	if !m.IsValid() {
		return ""
	}
	value := NewTString()
	imports.Proc(def.CefRequest_GetMethod).Call(m.Instance(), value.Instance())
	r = value.Value()
	value.Free()
	return
}

// SetMethod 设置请求方式
func (m *ICefRequest) SetMethod(method string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetMethod).Call(m.Instance(), api.PascalStr(method))
}

func (m *ICefRequest) ReferrerUrl() (r string) {
	if !m.IsValid() {
		return ""
	}
	value := NewTString()
	imports.Proc(def.CefRequest_GetReferrerUrl).Call(m.Instance(), value.Instance())
	r = value.Value()
	value.Free()
	return r
}

// SetReferrer 设置来源策略
func (m *ICefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetReferrer).Call(m.Instance(), api.PascalStr(referrerUrl), policy.ToPtr())
}

func (m *ICefRequest) ReferrerPolicy() TCefReferrerPolicy {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefRequest_GetReferrerPolicy).Call(m.Instance())
	return TCefReferrerPolicy(r1)
}

func (m *ICefRequest) Flags() TCefUrlRequestFlags {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefRequest_GetFlags).Call(m.Instance())
	return TCefUrlRequestFlags(r1)
}

// SetFlags 设置请求标记
func (m *ICefRequest) SetFlags(flags TCefUrlRequestFlags) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetFlags).Call(m.Instance(), flags.ToPtr())
}

func (m *ICefRequest) GetFirstPartyForCookies() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefRequest_GetFirstPartyForCookies).Call(m.Instance())
	return api.GoStr(r1)
}

// SetFirstPartyForCookies
func (m *ICefRequest) SetFirstPartyForCookies(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetFirstPartyForCookies).Call(m.Instance(), api.PascalStr(url))
}

// GetHeaderByName
func (m *ICefRequest) GetHeaderByName(name string) (r string) {
	if !m.IsValid() {
		return ""
	}
	value := NewTString()
	imports.Proc(def.CefRequest_GetHeaderByName).Call(m.Instance(), api.PascalStr(name), value.Instance())
	r = value.Value()
	value.Free()
	return r
}

// SetHeaderByName
func (m *ICefRequest) SetHeaderByName(name, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetHeaderByName).Call(m.Instance(), api.PascalStr(name), api.PascalStr(value), api.PascalBool(overwrite))
}

// GetHeaderMap
func (m *ICefRequest) GetHeaderMap() *ICefStringMultiMap {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefRequest_GetHeaderMap).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefStringMultiMap{instance: unsafe.Pointer(result)}
}

func (m *ICefRequest) SetHeaderMap(headerMap *ICefStringMultiMap) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetHeaderMap).Call(m.Instance(), headerMap.Instance())
}

func (m *ICefRequest) GetPostData() *ICefPostData {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefRequest_GetPostData).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPostData{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefRequest) SetPostData(value *ICefPostData) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequest_SetPostData).Call(m.Instance(), value.Instance())
}

func (m *ICefRequest) ResourceType() TCefResourceType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefRequest_GetResourceType).Call(m.Instance())
	return TCefResourceType(r1)
}

func (m *ICefRequest) TransitionType() TCefTransitionType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefRequest_GetTransitionType).Call(m.Instance())
	return TCefTransitionType(r1)
}

func (m *ICefRequest) Identifier() (result uint64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CefRequest_GetIdentifier).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefRequest) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
