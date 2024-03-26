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

// ICoreWebView2WebResourceResponseViewGetContentCompletedHandler Parent: IObject
//
//	Receives the result of the ICoreWebView2WebResourceResponseView.GetContent method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler">See the ICoreWebView2WebResourceResponseViewGetContentCompletedHandler article.</a>
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler interface {
	IObject
	ResourceID() int32 // property
}

// TCoreWebView2WebResourceResponseViewGetContentCompletedHandler Parent: TObject
//
//	Receives the result of the ICoreWebView2WebResourceResponseView.GetContent method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler">See the ICoreWebView2WebResourceResponseViewGetContentCompletedHandler article.</a>
type TCoreWebView2WebResourceResponseViewGetContentCompletedHandler struct {
	TObject
}

func NewCoreWebView2WebResourceResponseViewGetContentCompletedHandler(aEvents IWVBrowserEvents, aResourceID int32) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	r1 := WV().SysCallN(687, GetObjectUintptr(aEvents), uintptr(aResourceID))
	return AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler(r1)
}

func (m *TCoreWebView2WebResourceResponseViewGetContentCompletedHandler) ResourceID() int32 {
	r1 := WV().SysCallN(688, m.Instance())
	return int32(r1)
}

func CoreWebView2WebResourceResponseViewGetContentCompletedHandlerClass() TClass {
	ret := WV().SysCallN(686)
	return TClass(ret)
}
