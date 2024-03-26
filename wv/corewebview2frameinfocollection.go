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

// ICoreWebView2FrameInfoCollection Parent: IObject
//
//	Collection of FrameInfos (name and source). Used to list the affected
//	frames' info when a frame-only render process failure occurs in the
//	ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection">See the ICoreWebView2FrameInfoCollection article.</a>
type ICoreWebView2FrameInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfoCollection // property
	// Iterator
	//  Gets an iterator over the collection of `FrameInfo`s.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection#getiterator">See the ICoreWebView2FrameInfoCollection article.</a>
	Iterator() ICoreWebView2FrameInfoCollectionIterator // property
}

// TCoreWebView2FrameInfoCollection Parent: TObject
//
//	Collection of FrameInfos (name and source). Used to list the affected
//	frames' info when a frame-only render process failure occurs in the
//	ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection">See the ICoreWebView2FrameInfoCollection article.</a>
type TCoreWebView2FrameInfoCollection struct {
	TObject
}

func NewCoreWebView2FrameInfoCollection(aBaseIntf ICoreWebView2FrameInfoCollection) ICoreWebView2FrameInfoCollection {
	r1 := WV().SysCallN(326, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2FrameInfoCollection(r1)
}

func (m *TCoreWebView2FrameInfoCollection) Initialized() bool {
	r1 := WV().SysCallN(327, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameInfoCollection) BaseIntf() ICoreWebView2FrameInfoCollection {
	var resultCoreWebView2FrameInfoCollection uintptr
	WV().SysCallN(324, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollection)))
	return AsCoreWebView2FrameInfoCollection(resultCoreWebView2FrameInfoCollection)
}

func (m *TCoreWebView2FrameInfoCollection) Iterator() ICoreWebView2FrameInfoCollectionIterator {
	var resultCoreWebView2FrameInfoCollectionIterator uintptr
	WV().SysCallN(328, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollectionIterator)))
	return AsCoreWebView2FrameInfoCollectionIterator(resultCoreWebView2FrameInfoCollectionIterator)
}

func CoreWebView2FrameInfoCollectionClass() TClass {
	ret := WV().SysCallN(325)
	return TClass(ret)
}
