//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type ICefResponse struct {
	instance   unsafe.Pointer
	Status     int32
	StatusText string
	MimeType   string
	Charset    string
	Error      TCefErrorCode
	URL        string
}

type iCefResponse struct {
	Instance   uintptr
	Status     uintptr //int32
	StatusText uintptr //string
	MimeType   uintptr //string
	Charset    uintptr //string
	Error      uintptr //int32
	URL        uintptr //string
}

func (m *ICefResponse) IsReadOnly() bool {
	return api.GoBool(cefResponse_IsReadOnly(uintptr(m.instance)))
}

func (m *ICefResponse) SetError(error TCefErrorCode) {
	cefResponse_SetError(uintptr(m.instance), error)
}
func (m *ICefResponse) SetStatus(status int32) {
	cefResponse_SetStatus(uintptr(m.instance), status)
}
func (m *ICefResponse) SetStatusText(statusText string) {
	cefResponse_SetStatusText(uintptr(m.instance), statusText)
}
func (m *ICefResponse) SetMimeType(mimetype string) {
	cefResponse_SetMimeType(uintptr(m.instance), mimetype)
}
func (m *ICefResponse) SetCharset(charset string) {
	cefResponse_SetCharset(uintptr(m.instance), charset)
}

func (m *ICefResponse) GetHeaderByName(name string) string {
	return api.GoStr(cefResponse_GetHeaderByName(uintptr(m.instance), name))
}

func (m *ICefResponse) SetHeaderByName(name, value string, overwrite bool) {
	cefResponse_SetHeaderByName(uintptr(m.instance), name, value, overwrite)
}

func (m *ICefResponse) SetURL(url string) {
	cefResponse_SetURL(uintptr(m.instance), url)
}

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
