//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type ICefRequest struct {
	instance             uintptr
	ptr                  unsafe.Pointer
	Url                  string
	Method               string
	ReferrerUrl          string
	ReferrerPolicy       TCefReferrerPolicy
	Flags                TCefUrlRequestFlags
	FirstPartyForCookies string
	ResourceType         TCefResourceType
	TransitionType       TCefTransitionType
	Identifier           uint64
}

type rICefRequest struct {
	Instance             uintptr
	Url                  uintptr //string
	Method               uintptr //string
	ReferrerUrl          uintptr //string
	ReferrerPolicy       uintptr //int32
	Flags                uintptr //int32
	FirstPartyForCookies uintptr //string
	ResourceType         uintptr //int32
	TransitionType       uintptr //int32
	Identifier           uintptr //uint64
}

//request
func (m *ICefRequest) IsReadOnly() bool {
	return api.GoBool(cefRequest_IsReadOnly(m.instance))
}

func (m *ICefRequest) SetUrl(url string) {
	cefRequest_SetUrl(m.instance, url)
}

func (m *ICefRequest) SetMethod(method string) {
	cefRequest_SetMethod(m.instance, method)
}

func (m *ICefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	cefRequest_SetReferrer(m.instance, referrerUrl, policy)
}

func (m *ICefRequest) SetFlags(flags TCefUrlRequestFlags) {
	cefRequest_SetFlags(m.instance, flags)
}

func (m *ICefRequest) SetFirstPartyForCookies(url string) {
	cefRequest_SetFirstPartyForCookies(m.instance, url)
}

func (m *ICefRequest) GetHeaderByName(name string) string {
	return api.GoStr(cefRequest_GetHeaderByName(m.instance, name))
}

func (m *ICefRequest) SetHeaderByName(name, value string, overwrite bool) {
	cefRequest_SetHeaderByName(m.instance, name, value, overwrite)
}

func (m *ICefRequest) GetHeaderMap() *ICefStringMultiMap {
	headerMap := &ICefStringMultiMap{}
	headerMap.instance = cefRequest_GetHeaderMap(m.instance)
	headerMap.ptr = unsafe.Pointer(headerMap.instance)
	return headerMap
}

func (m *ICefRequest) SetHeaderMap(headerMap *ICefStringMultiMap) {
	cefRequest_SetHeaderMap(m.instance, headerMap.instance)
}

//request
func cefRequest_IsReadOnly(instance uintptr) uintptr {
	r1, _, _ := Proc(internale_cefRequest_IsReadOnly).Call(instance)
	return r1
}

func cefRequest_SetUrl(instance uintptr, url string) {
	Proc(internale_cefRequest_SetUrl).Call(instance, api.PascalStr(url))
}

func cefRequest_SetMethod(instance uintptr, method string) {
	Proc(internale_cefRequest_SetMethod).Call(instance, api.PascalStr(method))
}

func cefRequest_SetReferrer(instance uintptr, referrerUrl string, policy TCefReferrerPolicy) {
	Proc(internale_cefRequest_SetReferrer).Call(instance, api.PascalStr(referrerUrl), uintptr(policy))
}

func cefRequest_SetFlags(instance uintptr, flags TCefUrlRequestFlags) {
	Proc(internale_cefRequest_SetFlags).Call(instance, uintptr(flags))
}

func cefRequest_SetFirstPartyForCookies(instance uintptr, url string) {
	Proc(internale_cefRequest_SetFirstPartyForCookies).Call(instance, api.PascalStr(url))
}

func cefRequest_GetHeaderByName(instance uintptr, name string) uintptr {
	r1, _, _ := Proc(internale_cefRequest_GetHeaderByName).Call(instance, api.PascalStr(name))
	return r1
}

func cefRequest_SetHeaderByName(instance uintptr, url, value string, overwrite bool) {
	Proc(internale_cefRequest_SetHeaderByName).Call(instance, api.PascalStr(url), api.PascalStr(value), api.PascalBool(overwrite))
}

func cefRequest_GetHeaderMap(instance uintptr) uintptr {
	r1, _, _ := Proc(internale_cefRequest_GetHeaderMap).Call(instance)
	return r1
}

func cefRequest_SetHeaderMap(instance, headerMap uintptr) uintptr {
	r1, _, _ := Proc(internale_cefRequest_SetHeaderMap).Call(instance, headerMap)
	return r1
}
