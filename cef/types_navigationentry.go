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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefNavigationEntry
// Interface used to represent an entry in navigation history.
// <para><see cref="uCEFTypes|TCefNavigationEntry">Implements TCefNavigationEntry</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_navigation_entry_capi.h">CEF source file: /include/capi/cef_navigation_entry_capi.h (cef_navigation_entry_t)</see></para>
type ICefNavigationEntry struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// NavigationEntryRef -> ICefNavigationEntry
var NavigationEntryRef navigationEntry

type navigationEntry uintptr

func (*navigationEntry) UnWrap(data *ICefNavigationEntry) *ICefNavigationEntry {
	var result uintptr
	imports.Proc(def.NavigationEntryRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefNavigationEntry{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefNavigationEntry) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefNavigationEntry) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefNavigationEntry) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns the actual URL of the page. For some pages this may be data: URL
// or similar. Use get_display_url() to return a display-friendly version.
func (m *ICefNavigationEntry) GetUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns a display-friendly version of the URL.
func (m *ICefNavigationEntry) GetDisplayUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetDisplayUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the original URL that was entered by the user before any
// redirects.
func (m *ICefNavigationEntry) GetOriginalUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetOriginalUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the title set by the page. This value may be NULL.
func (m *ICefNavigationEntry) GetTitle() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetTitle).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the transition type which indicates what the user did to move to
// this page from the previous page.
func (m *ICefNavigationEntry) GetTransitionType() consts.TCefTransitionType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetTransitionType).Call(m.Instance())
	return consts.TCefTransitionType(r1)
}

// Returns true (1) if this navigation includes post data.
func (m *ICefNavigationEntry) HasPostData() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_HasPostData).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns the time for the last known successful navigation completion. A
// navigation may be completed more than once if the page is reloaded. May be
// 0 if the navigation has not yet completed.
func (m *ICefNavigationEntry) GetCompletionTime() (time consts.TDateTime) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.NavigationEntry_GetCompletionTime).Call(m.Instance(), uintptr(unsafe.Pointer(&time)))
	return
}

// Returns the HTTP status code for the last known successful navigation
// response. May be 0 if the response has not yet been received or if the
// navigation has not yet completed.
func (m *ICefNavigationEntry) GetHttpStatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.NavigationEntry_GetHttpStatusCode).Call(m.Instance())
	return int32(r1)
}

// Returns the actual URL of the page. For some pages this may be data: URL
// or similar. Use get_display_url() to return a display-friendly version.
func (m *ICefNavigationEntry) GetSSLStatus() *ICefSSLStatus {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.NavigationEntry_GetSSLStatus).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSSLStatus{instance: getInstance(result)}
	}
	return nil
}
