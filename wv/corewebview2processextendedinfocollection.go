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

// ICoreWebView2ProcessExtendedInfoCollection Parent: IObject
//
//	A list containing processInfo and associated extended information.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
type ICoreWebView2ProcessExtendedInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessExtendedInfoCollection // property
	// Count
	//  The number of process contained in the `ICoreWebView2ProcessExtendedInfoCollection`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#get_count">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the `ICoreWebView2ProcessExtendedInfo` located in the
	//  `ICoreWebView2ProcessExtendedInfoCollection` at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#getvalueatindex">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Items(idx uint32) ICoreWebView2ProcessExtendedInfo // property
}

// TCoreWebView2ProcessExtendedInfoCollection Parent: TObject
//
//	A list containing processInfo and associated extended information.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
type TCoreWebView2ProcessExtendedInfoCollection struct {
	TObject
}

func NewCoreWebView2ProcessExtendedInfoCollection(aBaseIntf ICoreWebView2ProcessExtendedInfoCollection) ICoreWebView2ProcessExtendedInfoCollection {
	r1 := WV().SysCallN(525, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessExtendedInfoCollection(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Initialized() bool {
	r1 := WV().SysCallN(526, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) BaseIntf() ICoreWebView2ProcessExtendedInfoCollection {
	var resultCoreWebView2ProcessExtendedInfoCollection uintptr
	WV().SysCallN(522, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessExtendedInfoCollection)))
	return AsCoreWebView2ProcessExtendedInfoCollection(resultCoreWebView2ProcessExtendedInfoCollection)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Count() uint32 {
	r1 := WV().SysCallN(524, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Items(idx uint32) ICoreWebView2ProcessExtendedInfo {
	var resultCoreWebView2ProcessExtendedInfo uintptr
	WV().SysCallN(527, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2ProcessExtendedInfo)))
	return AsCoreWebView2ProcessExtendedInfo(resultCoreWebView2ProcessExtendedInfo)
}

func CoreWebView2ProcessExtendedInfoCollectionClass() TClass {
	ret := WV().SysCallN(523)
	return TClass(ret)
}
