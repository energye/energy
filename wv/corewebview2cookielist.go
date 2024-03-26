//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2CookieList Parent: IObject
//
//	A list of cookie objects. See ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist">See the ICoreWebView2CookieList article.</a>
type ICoreWebView2CookieList interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CookieList // property
	// Count
	//  The number of cookies contained in the ICoreWebView2CookieList.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist#get_count">See the ICoreWebView2CookieList article.</a>
	Count() uint32 // property
	// Items
	//  Gets the cookie object at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist#getvalueatindex">See the ICoreWebView2CookieList article.</a>
	Items(idx uint32) ICoreWebView2Cookie // property
}

// TCoreWebView2CookieList Parent: TObject
//
//	A list of cookie objects. See ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist">See the ICoreWebView2CookieList article.</a>
type TCoreWebView2CookieList struct {
	TObject
}

func NewCoreWebView2CookieList(aBaseIntf ICoreWebView2CookieList) ICoreWebView2CookieList {
	r1 := WV().SysCallN(204, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2CookieList(r1)
}

func (m *TCoreWebView2CookieList) Initialized() bool {
	r1 := WV().SysCallN(205, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2CookieList) BaseIntf() ICoreWebView2CookieList {
	var resultCoreWebView2CookieList uintptr
	WV().SysCallN(201, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieList)))
	return AsCoreWebView2CookieList(resultCoreWebView2CookieList)
}

func (m *TCoreWebView2CookieList) Count() uint32 {
	r1 := WV().SysCallN(203, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2CookieList) Items(idx uint32) ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(206, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func CoreWebView2CookieListClass() TClass {
	ret := WV().SysCallN(202)
	return TClass(ret)
}
