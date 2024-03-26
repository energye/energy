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

// ICoreWebView2ContentLoadingEventArgs Parent: IObject
//
//	Receives ContentLoading events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs">See the ICoreWebView2ContentLoadingEventArgs article.</a>
type ICoreWebView2ContentLoadingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContentLoadingEventArgs // property
	// IsErrorPage
	//  `TRUE` if the loaded content is an error page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_iserrorpage">See the ICoreWebView2ContentLoadingEventArgs article.</a>
	IsErrorPage() bool // property
	// NavigationId
	//  The ID of the navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_navigationid">See the ICoreWebView2ContentLoadingEventArgs article.</a>
	NavigationId() uint64 // property
}

// TCoreWebView2ContentLoadingEventArgs Parent: TObject
//
//	Receives ContentLoading events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs">See the ICoreWebView2ContentLoadingEventArgs article.</a>
type TCoreWebView2ContentLoadingEventArgs struct {
	TObject
}

func NewCoreWebView2ContentLoadingEventArgs(aArgs ICoreWebView2ContentLoadingEventArgs) ICoreWebView2ContentLoadingEventArgs {
	r1 := WV().SysCallN(118, GetObjectUintptr(aArgs))
	return AsCoreWebView2ContentLoadingEventArgs(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) Initialized() bool {
	r1 := WV().SysCallN(119, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) BaseIntf() ICoreWebView2ContentLoadingEventArgs {
	var resultCoreWebView2ContentLoadingEventArgs uintptr
	WV().SysCallN(116, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContentLoadingEventArgs)))
	return AsCoreWebView2ContentLoadingEventArgs(resultCoreWebView2ContentLoadingEventArgs)
}

func (m *TCoreWebView2ContentLoadingEventArgs) IsErrorPage() bool {
	r1 := WV().SysCallN(120, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) NavigationId() uint64 {
	r1 := WV().SysCallN(121, m.Instance())
	return uint64(r1)
}

func CoreWebView2ContentLoadingEventArgsClass() TClass {
	ret := WV().SysCallN(117)
	return TClass(ret)
}
