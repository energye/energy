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

// ICoreWebView2ProcessInfoCollection Parent: IObject
//
//	A list containing process id and corresponding process type.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection">See the ICoreWebView2ProcessInfoCollection article.</a>
type ICoreWebView2ProcessInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessInfoCollection // property
	// Count
	//  The number of process contained in the ICoreWebView2ProcessInfoCollection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection#get_count">See the ICoreWebView2ProcessInfoCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the `ICoreWebView2ProcessInfo` located in the `ICoreWebView2ProcessInfoCollection`
	//  at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection#getvalueatindex">See the ICoreWebView2ProcessInfoCollection article.</a>
	Items(idx uint32) ICoreWebView2ProcessInfo // property
}

// TCoreWebView2ProcessInfoCollection Parent: TObject
//
//	A list containing process id and corresponding process type.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection">See the ICoreWebView2ProcessInfoCollection article.</a>
type TCoreWebView2ProcessInfoCollection struct {
	TObject
}

func NewCoreWebView2ProcessInfoCollection(aBaseIntf ICoreWebView2ProcessInfoCollection) ICoreWebView2ProcessInfoCollection {
	r1 := WV().SysCallN(546, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessInfoCollection(r1)
}

func (m *TCoreWebView2ProcessInfoCollection) Initialized() bool {
	r1 := WV().SysCallN(547, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessInfoCollection) BaseIntf() ICoreWebView2ProcessInfoCollection {
	var resultCoreWebView2ProcessInfoCollection uintptr
	WV().SysCallN(543, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfoCollection)))
	return AsCoreWebView2ProcessInfoCollection(resultCoreWebView2ProcessInfoCollection)
}

func (m *TCoreWebView2ProcessInfoCollection) Count() uint32 {
	r1 := WV().SysCallN(545, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ProcessInfoCollection) Items(idx uint32) ICoreWebView2ProcessInfo {
	var resultCoreWebView2ProcessInfo uintptr
	WV().SysCallN(548, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2ProcessInfo)))
	return AsCoreWebView2ProcessInfo(resultCoreWebView2ProcessInfo)
}

func CoreWebView2ProcessInfoCollectionClass() TClass {
	ret := WV().SysCallN(544)
	return TClass(ret)
}
