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

// ICoreWebView2WebMessageReceivedEventArgs Parent: IObject
//
//	Event args for the WebMessageReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs">See the ICoreWebView2WebMessageReceivedEventArgs article.</a>
type ICoreWebView2WebMessageReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebMessageReceivedEventArgs // property
	// Source
	//  The URI of the document that sent this web message.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#get_source">See the ICoreWebView2WebMessageReceivedEventArgs article.</a>
	Source() string // property
	// WebMessageAsJson
	//  The message posted from the WebView content to the host converted to a
	//  JSON string. Run this operation to communicate using JavaScript objects.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#get_webmessageasjson">See the ICoreWebView2WebMessageReceivedEventArgs article.</a>
	WebMessageAsJson() string // property
	// WebMessageAsString
	//  If the message posted from the WebView content to the host is a string
	//  type, this method returns the value of that string. If the message
	//  posted is some other kind of JavaScript type this method fails with the
	//  following error.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#trygetwebmessageasstring">See the ICoreWebView2WebMessageReceivedEventArgs article.</a>
	WebMessageAsString() string // property
	// AdditionalObjects
	//  Additional received WebMessage objects. To pass `additionalObjects` via
	//  WebMessage to the app, use the
	//  `chrome.webview.postMessageWithAdditionalObjects` content API.
	//  Any DOM object type that can be natively representable that has been
	//  passed in to `additionalObjects` parameter will be accessible here.
	//  Currently a WebMessage object can be the `ICoreWebView2File` type.
	//  Entries in the collection can be `nullptr` if `null` or `undefined` was
	//  passed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2#get_additionalobjects">See the ICoreWebView2WebMessageReceivedEventArgs2 article.</a>
	AdditionalObjects() ICoreWebView2ObjectCollectionView // property
}

// TCoreWebView2WebMessageReceivedEventArgs Parent: TObject
//
//	Event args for the WebMessageReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs">See the ICoreWebView2WebMessageReceivedEventArgs article.</a>
type TCoreWebView2WebMessageReceivedEventArgs struct {
	TObject
}

func NewCoreWebView2WebMessageReceivedEventArgs(aArgs ICoreWebView2WebMessageReceivedEventArgs) ICoreWebView2WebMessageReceivedEventArgs {
	r1 := WV().SysCallN(659, GetObjectUintptr(aArgs))
	return AsCoreWebView2WebMessageReceivedEventArgs(r1)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(660, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) BaseIntf() ICoreWebView2WebMessageReceivedEventArgs {
	var resultCoreWebView2WebMessageReceivedEventArgs uintptr
	WV().SysCallN(657, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebMessageReceivedEventArgs)))
	return AsCoreWebView2WebMessageReceivedEventArgs(resultCoreWebView2WebMessageReceivedEventArgs)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) Source() string {
	r1 := WV().SysCallN(661, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) WebMessageAsJson() string {
	r1 := WV().SysCallN(662, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) WebMessageAsString() string {
	r1 := WV().SysCallN(663, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) AdditionalObjects() ICoreWebView2ObjectCollectionView {
	var resultCoreWebView2ObjectCollectionView uintptr
	WV().SysCallN(656, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ObjectCollectionView)))
	return AsCoreWebView2ObjectCollectionView(resultCoreWebView2ObjectCollectionView)
}

func CoreWebView2WebMessageReceivedEventArgsClass() TClass {
	ret := WV().SysCallN(658)
	return TClass(ret)
}
