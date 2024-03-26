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

// ICoreWebView2DevToolsProtocolEventReceivedEventArgs Parent: IObject
//
//	Event args for the DevToolsProtocolEventReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
type ICoreWebView2DevToolsProtocolEventReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DevToolsProtocolEventReceivedEventArgs // property
	// ParameterObjectAsJson
	//  The parameter object of the corresponding `DevToolsProtocol` event
	//  represented as a JSON string.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs#get_parameterobjectasjson">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
	ParameterObjectAsJson() string // property
	// SessionId
	//  The sessionId of the target where the event originates from.
	//  Empty string is returned as sessionId if the event comes from the default
	//  session for the top page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs2#get_sessionid">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 article.</a>
	SessionId() string // property
}

// TCoreWebView2DevToolsProtocolEventReceivedEventArgs Parent: TObject
//
//	Event args for the DevToolsProtocolEventReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
type TCoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	TObject
}

func NewCoreWebView2DevToolsProtocolEventReceivedEventArgs(aArgs ICoreWebView2DevToolsProtocolEventReceivedEventArgs) ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	r1 := WV().SysCallN(247, GetObjectUintptr(aArgs))
	return AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(248, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) BaseIntf() ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	var resultCoreWebView2DevToolsProtocolEventReceivedEventArgs uintptr
	WV().SysCallN(245, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DevToolsProtocolEventReceivedEventArgs)))
	return AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(resultCoreWebView2DevToolsProtocolEventReceivedEventArgs)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) ParameterObjectAsJson() string {
	r1 := WV().SysCallN(249, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) SessionId() string {
	r1 := WV().SysCallN(250, m.Instance())
	return GoStr(r1)
}

func CoreWebView2DevToolsProtocolEventReceivedEventArgsClass() TClass {
	ret := WV().SysCallN(246)
	return TClass(ret)
}
