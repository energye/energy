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
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefResponse 实例
type ICefResponse struct {
	instance   unsafe.Pointer
	Status     int32
	StatusText string
	MimeType   string
	Charset    string
	Error      TCefErrorCode
	URL        string
}

// ICefResponse 指针实例
type iCefResponse struct {
	Instance   uintptr
	Status     uintptr //int32
	StatusText uintptr //string
	MimeType   uintptr //string
	Charset    uintptr //string
	Error      uintptr //int32
	URL        uintptr //string
}

// IsReadOnly 是否只读
func (m *ICefResponse) IsReadOnly() bool {
	return api.GoBool(cefResponse_IsReadOnly(uintptr(m.instance)))
}

// SetURL 设置URL
func (m *ICefResponse) SetURL(url string) {
	cefResponse_SetURL(uintptr(m.instance), url)
}

// SetError 设置错误码
func (m *ICefResponse) SetError(error TCefErrorCode) {
	cefResponse_SetError(uintptr(m.instance), error)
}

// SetStatus 设置状态码
func (m *ICefResponse) SetStatus(status int32) {
	cefResponse_SetStatus(uintptr(m.instance), status)
}

// SetStatusText 设置状态文本
func (m *ICefResponse) SetStatusText(statusText string) {
	cefResponse_SetStatusText(uintptr(m.instance), statusText)
}

// SetMimeType mime类型
func (m *ICefResponse) SetMimeType(mimetype string) {
	cefResponse_SetMimeType(uintptr(m.instance), mimetype)
}

// SetCharset 设置编码
func (m *ICefResponse) SetCharset(charset string) {
	cefResponse_SetCharset(uintptr(m.instance), charset)
}

// GetHeaderByName
func (m *ICefResponse) GetHeaderByName(name string) string {
	return api.GoStr(cefResponse_GetHeaderByName(uintptr(m.instance), name))
}

// SetHeaderByName
func (m *ICefResponse) SetHeaderByName(name, value string, overwrite bool) {
	cefResponse_SetHeaderByName(uintptr(m.instance), name, value, overwrite)
}

// GetHeaderMap
func (m *ICefResponse) GetHeaderMap() *ICefStringMultiMap {
	headerMap := &ICefStringMultiMap{}
	headerMap.instance = cefResponse_GetHeaderMap(uintptr(m.instance))
	headerMap.ptr = unsafe.Pointer(headerMap.instance)
	return headerMap
}

func cefResponse_IsReadOnly(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_cefResponse_IsReadOnly).Call(instance)
	return r1
}

func cefResponse_SetError(instance uintptr, error TCefErrorCode) {
	imports.Proc(internale_cefResponse_SetError).Call(instance, error.ToPtr())
}

func cefResponse_SetStatus(instance uintptr, error int32) {
	imports.Proc(internale_cefResponse_SetStatus).Call(instance, uintptr(error))
}

func cefResponse_SetStatusText(instance uintptr, statusText string) {
	imports.Proc(internale_cefResponse_SetStatusText).Call(instance, api.PascalStr(statusText))
}

func cefResponse_SetMimeType(instance uintptr, mimetype string) {
	imports.Proc(internale_cefResponse_SetMimeType).Call(instance, api.PascalStr(mimetype))
}

func cefResponse_SetCharset(instance uintptr, charset string) {
	imports.Proc(internale_cefResponse_SetCharset).Call(instance, api.PascalStr(charset))
}

func cefResponse_GetHeaderByName(instance uintptr, name string) uintptr {
	r1, _, _ := imports.Proc(internale_cefResponse_GetHeaderByName).Call(instance, api.PascalStr(name))
	return r1
}

func cefResponse_SetHeaderByName(instance uintptr, name, value string, overwrite bool) {
	imports.Proc(internale_cefResponse_SetHeaderByName).Call(instance, api.PascalStr(name), api.PascalStr(value), api.PascalBool(overwrite))
}

func cefResponse_SetURL(instance uintptr, url string) {
	imports.Proc(internale_cefResponse_SetURL).Call(instance, api.PascalStr(url))
}

func cefResponse_GetHeaderMap(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_cefResponse_GetHeaderMap).Call(instance)
	return r1
}
