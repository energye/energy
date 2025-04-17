//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Response

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefResponse
type ICefResponse struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// ResponseRef -> ICefResponse
var ResponseRef response

// request
type response uintptr

func (*response) New() *ICefResponse {
	var result uintptr
	imports.Proc(def.CefResponseRef_New).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefResponse{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*response) UnWrap(data *ICefResponse) *ICefResponse {
	var result uintptr
	imports.Proc(def.CefResponseRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefResponse) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResponse) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// IsReadOnly 是否只读
func (m *ICefResponse) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefResponse_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefResponse) URL() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetURL).Call(m.Instance())
	return api.GoStr(r1)
}

// SetURL 设置URL
func (m *ICefResponse) SetURL(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetURL).Call(m.Instance(), api.PascalStr(url))
}

func (m *ICefResponse) Error() TCefErrorCode {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetError).Call(m.Instance())
	return TCefErrorCode(r1)
}

// SetError 设置错误码
func (m *ICefResponse) SetError(error TCefErrorCode) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetError).Call(m.Instance(), error.ToPtr())
}

func (m *ICefResponse) Status() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetStatus).Call(m.Instance())
	return int32(r1)
}

// SetStatus 设置状态码
func (m *ICefResponse) SetStatus(status int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetStatus).Call(m.Instance(), uintptr(status))
}

func (m *ICefResponse) StatusText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetStatusText).Call(m.Instance())
	return api.GoStr(r1)
}

// SetStatusText 设置状态文本
func (m *ICefResponse) SetStatusText(statusText string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetStatusText).Call(m.Instance(), api.PascalStr(statusText))
}

func (m *ICefResponse) MimeType() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetMimeType).Call(m.Instance())
	return api.GoStr(r1)
}

// SetMimeType mime类型
func (m *ICefResponse) SetMimeType(mimetype string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetMimeType).Call(m.Instance(), api.PascalStr(mimetype))
}

func (m *ICefResponse) Charset() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetCharset).Call(m.Instance())
	return api.GoStr(r1)
}

// SetCharset 设置编码
func (m *ICefResponse) SetCharset(charset string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetCharset).Call(m.Instance(), api.PascalStr(charset))
}

// GetHeaderByName
func (m *ICefResponse) GetHeaderByName(name string) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefResponse_GetHeaderByName).Call(m.Instance(), api.PascalStr(name))
	return api.GoStr(r1)
}

// SetHeaderByName
func (m *ICefResponse) SetHeaderByName(name, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetHeaderByName).Call(m.Instance(), api.PascalStr(name), api.PascalStr(value), api.PascalBool(overwrite))
}

// GetHeaderMap
func (m *ICefResponse) GetHeaderMap() *ICefStringMultiMap {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefResponse_GetHeaderMap).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefStringMultiMap{instance: unsafe.Pointer(result)}
}

// GetHeaderMap
func (m *ICefResponse) SetHeaderMap(headerMap *ICefStringMultiMap) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefResponse_SetHeaderMap).Call(m.Instance(), headerMap.Instance())
}

func (m *ICefResponse) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
