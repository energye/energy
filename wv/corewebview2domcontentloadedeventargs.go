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

// ICoreWebView2DOMContentLoadedEventArgs Parent: IObject
//
//	Event args for the DOMContentLoaded event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
type ICoreWebView2DOMContentLoadedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DOMContentLoadedEventArgs // property
	// NavigationId
	//  The ID of the navigation which corresponds to other navigation ID properties on other navigation events.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs#get_navigationid">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
	NavigationId() uint64 // property
}

// TCoreWebView2DOMContentLoadedEventArgs Parent: TObject
//
//	Event args for the DOMContentLoaded event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
type TCoreWebView2DOMContentLoadedEventArgs struct {
	TObject
}

func NewCoreWebView2DOMContentLoadedEventArgs(aArgs ICoreWebView2DOMContentLoadedEventArgs) ICoreWebView2DOMContentLoadedEventArgs {
	r1 := WV().SysCallN(237, GetObjectUintptr(aArgs))
	return AsCoreWebView2DOMContentLoadedEventArgs(r1)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(238, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) BaseIntf() ICoreWebView2DOMContentLoadedEventArgs {
	var resultCoreWebView2DOMContentLoadedEventArgs uintptr
	WV().SysCallN(235, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DOMContentLoadedEventArgs)))
	return AsCoreWebView2DOMContentLoadedEventArgs(resultCoreWebView2DOMContentLoadedEventArgs)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) NavigationId() uint64 {
	r1 := WV().SysCallN(239, m.Instance())
	return uint64(r1)
}

func CoreWebView2DOMContentLoadedEventArgsClass() TClass {
	ret := WV().SysCallN(236)
	return TClass(ret)
}
