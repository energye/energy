//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefNavigationEntry Parent: ICefBaseRefCounted
//
//	Interface used to represent an entry in navigation history.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_navigation_entry_capi.h">CEF source file: /include/capi/cef_navigation_entry_capi.h (cef_navigation_entry_t))</a>
type ICefNavigationEntry interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetUrl
	//  Returns the actual URL of the page. For some pages this may be data: URL or similar. Use get_display_url() to return a display-friendly version.
	GetUrl() string // function
	// GetDisplayUrl
	//  Returns a display-friendly version of the URL.
	GetDisplayUrl() string // function
	// GetOriginalUrl
	//  Returns the original URL that was entered by the user before any redirects.
	GetOriginalUrl() string // function
	// GetTitle
	//  Returns the title set by the page. This value may be NULL.
	GetTitle() string // function
	// GetTransitionType
	//  Returns the transition type which indicates what the user did to move to this page from the previous page.
	GetTransitionType() TCefTransitionType // function
	// HasPostData
	//  Returns true (1) if this navigation includes post data.
	HasPostData() bool // function
	// GetCompletionTime
	//  Returns the time for the last known successful navigation completion. A navigation may be completed more than once if the page is reloaded. May be 0 if the navigation has not yet completed.
	GetCompletionTime() (resultDateTime TDateTime) // function
	// GetHttpStatusCode
	//  Returns the HTTP status code for the last known successful navigation response. May be 0 if the response has not yet been received or if the navigation has not yet completed.
	GetHttpStatusCode() int32 // function
	// GetSSLStatus
	//  Returns the SSL information for this navigation entry.
	GetSSLStatus() ICefSSLStatus // function
}

// TCefNavigationEntry Parent: TCefBaseRefCounted
//
//	Interface used to represent an entry in navigation history.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_navigation_entry_capi.h">CEF source file: /include/capi/cef_navigation_entry_capi.h (cef_navigation_entry_t))</a>
type TCefNavigationEntry struct {
	TCefBaseRefCounted
}

// NavigationEntryRef -> ICefNavigationEntry
var NavigationEntryRef navigationEntry

// navigationEntry TCefNavigationEntry Ref
type navigationEntry uintptr

func (m *navigationEntry) UnWrap(data uintptr) ICefNavigationEntry {
	var resultCefNavigationEntry uintptr
	CEF().SysCallN(1151, uintptr(data), uintptr(unsafePointer(&resultCefNavigationEntry)))
	return AsCefNavigationEntry(resultCefNavigationEntry)
}

func (m *TCefNavigationEntry) IsValid() bool {
	r1 := CEF().SysCallN(1150, m.Instance())
	return GoBool(r1)
}

func (m *TCefNavigationEntry) GetUrl() string {
	r1 := CEF().SysCallN(1148, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetDisplayUrl() string {
	r1 := CEF().SysCallN(1142, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetOriginalUrl() string {
	r1 := CEF().SysCallN(1144, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetTitle() string {
	r1 := CEF().SysCallN(1146, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetTransitionType() TCefTransitionType {
	r1 := CEF().SysCallN(1147, m.Instance())
	return TCefTransitionType(r1)
}

func (m *TCefNavigationEntry) HasPostData() bool {
	r1 := CEF().SysCallN(1149, m.Instance())
	return GoBool(r1)
}

func (m *TCefNavigationEntry) GetCompletionTime() (resultDateTime TDateTime) {
	CEF().SysCallN(1141, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefNavigationEntry) GetHttpStatusCode() int32 {
	r1 := CEF().SysCallN(1143, m.Instance())
	return int32(r1)
}

func (m *TCefNavigationEntry) GetSSLStatus() ICefSSLStatus {
	var resultCefSSLStatus uintptr
	CEF().SysCallN(1145, m.Instance(), uintptr(unsafePointer(&resultCefSSLStatus)))
	return AsCefSSLStatus(resultCefSSLStatus)
}
