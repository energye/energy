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

// ICoreWebView2ProcessExtendedInfo Parent: IObject
//
//	Provides process with associated extended information in the `ICoreWebView2Environment`.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfo">See the ICoreWebView2ProcessExtendedInfo article.</a>
type ICoreWebView2ProcessExtendedInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessExtendedInfo // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ProcessExtendedInfo) // property
	// ProcessInfo
	//  The process info of the current process.
	ProcessInfo() ICoreWebView2ProcessInfo // property
	// AssociatedFrameInfos
	//  The collection of associated `FrameInfo`s which are actively running
	// (showing or hiding UI elements) in this renderer process. `AssociatedFrameInfos`
	//  will only be populated when this `ICoreWebView2ProcessExtendedInfo`
	//  corresponds to a renderer process. Non-renderer processes will always
	//  have an empty `AssociatedFrameInfos`. The `AssociatedFrameInfos` may
	//  also be empty for renderer processes that have no active frames.
	AssociatedFrameInfos() ICoreWebView2FrameInfoCollection // property
}

// TCoreWebView2ProcessExtendedInfo Parent: TObject
//
//	Provides process with associated extended information in the `ICoreWebView2Environment`.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfo">See the ICoreWebView2ProcessExtendedInfo article.</a>
type TCoreWebView2ProcessExtendedInfo struct {
	TObject
}

func NewCoreWebView2ProcessExtendedInfo(aBaseIntf ICoreWebView2ProcessExtendedInfo) ICoreWebView2ProcessExtendedInfo {
	r1 := WV().SysCallN(531, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessExtendedInfo(r1)
}

func (m *TCoreWebView2ProcessExtendedInfo) Initialized() bool {
	r1 := WV().SysCallN(532, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessExtendedInfo) BaseIntf() ICoreWebView2ProcessExtendedInfo {
	var resultCoreWebView2ProcessExtendedInfo uintptr
	WV().SysCallN(529, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ProcessExtendedInfo)))
	return AsCoreWebView2ProcessExtendedInfo(resultCoreWebView2ProcessExtendedInfo)
}

func (m *TCoreWebView2ProcessExtendedInfo) SetBaseIntf(AValue ICoreWebView2ProcessExtendedInfo) {
	WV().SysCallN(529, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ProcessExtendedInfo) ProcessInfo() ICoreWebView2ProcessInfo {
	var resultCoreWebView2ProcessInfo uintptr
	WV().SysCallN(533, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfo)))
	return AsCoreWebView2ProcessInfo(resultCoreWebView2ProcessInfo)
}

func (m *TCoreWebView2ProcessExtendedInfo) AssociatedFrameInfos() ICoreWebView2FrameInfoCollection {
	var resultCoreWebView2FrameInfoCollection uintptr
	WV().SysCallN(528, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollection)))
	return AsCoreWebView2FrameInfoCollection(resultCoreWebView2FrameInfoCollection)
}

func CoreWebView2ProcessExtendedInfoClass() TClass {
	ret := WV().SysCallN(530)
	return TClass(ret)
}
