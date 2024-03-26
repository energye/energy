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

// ICoreWebView2BrowserExtensionList Parent: IObject
//
//	Provides a set of properties for managing browser Extension Lists from user profile. This
//	includes the number of browser Extensions in the list, and the ability to get an browser
//	Extension from the list at a particular index.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist">See the ICoreWebView2BrowserExtensionList article.</a>
type ICoreWebView2BrowserExtensionList interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserExtensionList // property
	// Count
	//  The number of browser Extensions in the list.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist#get_count">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the browser Extension located in the browser Extension List at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist#getvalueatindex">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Items(idx uint32) ICoreWebView2BrowserExtension // property
}

// TCoreWebView2BrowserExtensionList Parent: TObject
//
//	Provides a set of properties for managing browser Extension Lists from user profile. This
//	includes the number of browser Extensions in the list, and the ability to get an browser
//	Extension from the list at a particular index.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist">See the ICoreWebView2BrowserExtensionList article.</a>
type TCoreWebView2BrowserExtensionList struct {
	TObject
}

func NewCoreWebView2BrowserExtensionList(aBaseIntf ICoreWebView2BrowserExtensionList) ICoreWebView2BrowserExtensionList {
	r1 := WV().SysCallN(35, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2BrowserExtensionList(r1)
}

func (m *TCoreWebView2BrowserExtensionList) Initialized() bool {
	r1 := WV().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserExtensionList) BaseIntf() ICoreWebView2BrowserExtensionList {
	var resultCoreWebView2BrowserExtensionList uintptr
	WV().SysCallN(32, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2BrowserExtensionList)))
	return AsCoreWebView2BrowserExtensionList(resultCoreWebView2BrowserExtensionList)
}

func (m *TCoreWebView2BrowserExtensionList) Count() uint32 {
	r1 := WV().SysCallN(34, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2BrowserExtensionList) Items(idx uint32) ICoreWebView2BrowserExtension {
	var resultCoreWebView2BrowserExtension uintptr
	WV().SysCallN(37, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2BrowserExtension)))
	return AsCoreWebView2BrowserExtension(resultCoreWebView2BrowserExtension)
}

func CoreWebView2BrowserExtensionListClass() TClass {
	ret := WV().SysCallN(33)
	return TClass(ret)
}
