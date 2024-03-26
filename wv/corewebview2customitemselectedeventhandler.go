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

// ICoreWebView2CustomItemSelectedEventHandler Parent: IObject
//
//	Raised to notify the host that the end user selected a custom
//	ContextMenuItem. CustomItemSelected event is raised on the specific
//	ContextMenuItem that the end user selected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2customitemselectedeventhandler">See the ICoreWebView2CustomItemSelectedEventHandler article.</a>
type ICoreWebView2CustomItemSelectedEventHandler interface {
	IObject
}

// TCoreWebView2CustomItemSelectedEventHandler Parent: TObject
//
//	Raised to notify the host that the end user selected a custom
//	ContextMenuItem. CustomItemSelected event is raised on the specific
//	ContextMenuItem that the end user selected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2customitemselectedeventhandler">See the ICoreWebView2CustomItemSelectedEventHandler article.</a>
type TCoreWebView2CustomItemSelectedEventHandler struct {
	TObject
}

func NewCoreWebView2CustomItemSelectedEventHandler(aEvents IWVBrowserEvents) ICoreWebView2CustomItemSelectedEventHandler {
	r1 := WV().SysCallN(234, GetObjectUintptr(aEvents))
	return AsCoreWebView2CustomItemSelectedEventHandler(r1)
}

func CoreWebView2CustomItemSelectedEventHandlerClass() TClass {
	ret := WV().SysCallN(233)
	return TClass(ret)
}
