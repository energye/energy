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

func (m *ICefRequest) Instance() uintptr {
	return uintptr(m.instance)
}

// IsReadOnly 是否只读
func (m *ICefRequest) IsReadOnly() bool {
	return api.GoBool(cefRequest_IsReadOnly(m.Instance()))
}

// SetURL 设置URL
func (m *ICefRequest) SetURL(url string) {
	cefRequest_SetUrl(m.Instance(), url)
}

// SetMethod 设置请求方式
func (m *ICefRequest) SetMethod(method string) {
	cefRequest_SetMethod(m.Instance(), method)
}

// SetReferrer 设置来源策略
func (m *ICefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	cefRequest_SetReferrer(m.Instance(), referrerUrl, policy)
}

// SetFlags 设置请求标记
func (m *ICefRequest) SetFlags(flags TCefUrlRequestFlags) {
	cefRequest_SetFlags(m.Instance(), flags)
}

// SetFirstPartyForCookies
func (m *ICefRequest) SetFirstPartyForCookies(url string) {
	cefRequest_SetFirstPartyForCookies(m.Instance(), url)
}

// GetHeaderByName
func (m *ICefRequest) GetHeaderByName(name string) string {
	return api.GoStr(cefRequest_GetHeaderByName(m.Instance(), name))
}

// SetHeaderByName
func (m *ICefRequest) SetHeaderByName(name, value string, overwrite bool) {
	cefRequest_SetHeaderByName(m.Instance(), name, value, overwrite)
}

// GetHeaderMap
func (m *ICefRequest) GetHeaderMap() *ICefStringMultiMap {
	headerMap := &ICefStringMultiMap{}
	headerMap.instance = cefRequest_GetHeaderMap(m.Instance())
	headerMap.ptr = unsafe.Pointer(headerMap.instance)
	return headerMap
}

func (m *ICefRequest) SetHeaderMap(headerMap *ICefStringMultiMap) {
	cefRequest_SetHeaderMap(m.Instance(), headerMap.instance)
}

// request
func cefRequest_IsReadOnly(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_cefRequest_IsReadOnly).Call(instance)
	return r1
}

func cefRequest_SetUrl(instance uintptr, url string) {
	imports.Proc(internale_cefRequest_SetUrl).Call(instance, api.PascalStr(url))
}

func cefRequest_SetMethod(instance uintptr, method string) {
	imports.Proc(internale_cefRequest_SetMethod).Call(instance, api.PascalStr(method))
}

func cefRequest_SetReferrer(instance uintptr, referrerUrl string, policy TCefReferrerPolicy) {
	imports.Proc(internale_cefRequest_SetReferrer).Call(instance, api.PascalStr(referrerUrl), uintptr(policy))
}

func cefRequest_SetFlags(instance uintptr, flags TCefUrlRequestFlags) {
	imports.Proc(internale_cefRequest_SetFlags).Call(instance, uintptr(flags))
}

func cefRequest_SetFirstPartyForCookies(instance uintptr, url string) {
	imports.Proc(internale_cefRequest_SetFirstPartyForCookies).Call(instance, api.PascalStr(url))
}

func cefRequest_GetHeaderByName(instance uintptr, name string) uintptr {
	r1, _, _ := imports.Proc(internale_cefRequest_GetHeaderByName).Call(instance, api.PascalStr(name))
	return r1
}

func cefRequest_SetHeaderByName(instance uintptr, url, value string, overwrite bool) {
	imports.Proc(internale_cefRequest_SetHeaderByName).Call(instance, api.PascalStr(url), api.PascalStr(value), api.PascalBool(overwrite))
}

func cefRequest_GetHeaderMap(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_cefRequest_GetHeaderMap).Call(instance)
	return r1
}

func cefRequest_SetHeaderMap(instance, headerMap uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_cefRequest_SetHeaderMap).Call(instance, headerMap)
	return r1
}
