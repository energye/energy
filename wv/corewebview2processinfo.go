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

// ICoreWebView2ProcessInfo Parent: IObject
//
//	Provides a set of properties for a process in the ICoreWebView2Environment.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo">See the ICoreWebView2ProcessInfo article.</a>
type ICoreWebView2ProcessInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessInfo // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ProcessInfo) // property
	// Kind
	//  The kind of the process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</a>
	Kind() TWVProcessKind // property
	// KindStr
	//  The kind of the process in string format.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</a>
	KindStr() string // property
	// ProcessId
	//  The process id of the process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_processid">See the ICoreWebView2ProcessInfo article.</a>
	ProcessId() int32 // property
}

// TCoreWebView2ProcessInfo Parent: TObject
//
//	Provides a set of properties for a process in the ICoreWebView2Environment.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo">See the ICoreWebView2ProcessInfo article.</a>
type TCoreWebView2ProcessInfo struct {
	TObject
}

func NewCoreWebView2ProcessInfo(aBaseIntf ICoreWebView2ProcessInfo) ICoreWebView2ProcessInfo {
	r1 := WV().SysCallN(551, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessInfo(r1)
}

func (m *TCoreWebView2ProcessInfo) Initialized() bool {
	r1 := WV().SysCallN(552, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessInfo) BaseIntf() ICoreWebView2ProcessInfo {
	var resultCoreWebView2ProcessInfo uintptr
	WV().SysCallN(549, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ProcessInfo)))
	return AsCoreWebView2ProcessInfo(resultCoreWebView2ProcessInfo)
}

func (m *TCoreWebView2ProcessInfo) SetBaseIntf(AValue ICoreWebView2ProcessInfo) {
	WV().SysCallN(549, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ProcessInfo) Kind() TWVProcessKind {
	r1 := WV().SysCallN(553, m.Instance())
	return TWVProcessKind(r1)
}

func (m *TCoreWebView2ProcessInfo) KindStr() string {
	r1 := WV().SysCallN(554, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ProcessInfo) ProcessId() int32 {
	r1 := WV().SysCallN(555, m.Instance())
	return int32(r1)
}

func CoreWebView2ProcessInfoClass() TClass {
	ret := WV().SysCallN(550)
	return TClass(ret)
}
