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

// ICoreWebView2SourceChangedEventArgs Parent: IObject
//
//	Event args for the SourceChanged event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs">See the ICoreWebView2SourceChangedEventArgs article.</a>
type ICoreWebView2SourceChangedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SourceChangedEventArgs // property
	// IsNewDocument
	//  `TRUE` if the page being navigated to is a new document.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs#get_isnewdocument">See the ICoreWebView2SourceChangedEventArgs article.</a>
	IsNewDocument() bool // property
}

// TCoreWebView2SourceChangedEventArgs Parent: TObject
//
//	Event args for the SourceChanged event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs">See the ICoreWebView2SourceChangedEventArgs article.</a>
type TCoreWebView2SourceChangedEventArgs struct {
	TObject
}

func NewCoreWebView2SourceChangedEventArgs(aArgs ICoreWebView2SourceChangedEventArgs) ICoreWebView2SourceChangedEventArgs {
	r1 := WV().SysCallN(645, GetObjectUintptr(aArgs))
	return AsCoreWebView2SourceChangedEventArgs(r1)
}

func (m *TCoreWebView2SourceChangedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(646, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2SourceChangedEventArgs) BaseIntf() ICoreWebView2SourceChangedEventArgs {
	var resultCoreWebView2SourceChangedEventArgs uintptr
	WV().SysCallN(643, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2SourceChangedEventArgs)))
	return AsCoreWebView2SourceChangedEventArgs(resultCoreWebView2SourceChangedEventArgs)
}

func (m *TCoreWebView2SourceChangedEventArgs) IsNewDocument() bool {
	r1 := WV().SysCallN(647, m.Instance())
	return GoBool(r1)
}

func CoreWebView2SourceChangedEventArgsClass() TClass {
	ret := WV().SysCallN(644)
	return TClass(ret)
}
