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

// ICoreWebView2ContextMenuTarget Parent: IObject
//
//	Represents the information regarding the context menu target.
//	Includes the context selected and the appropriate data used for the actions of a context menu.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget">See the ICoreWebView2ContextMenuTarget article.</a>
type ICoreWebView2ContextMenuTarget interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuTarget // property
	// Kind
	//  Gets the kind of context that the user selected.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_kind">See the ICoreWebView2ContextMenuTarget article.</a>
	Kind() TWVMenuTargetKind // property
	// IsEditable
	//  Returns TRUE if the context menu is requested on an editable component.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_iseditable">See the ICoreWebView2ContextMenuTarget article.</a>
	IsEditable() bool // property
	// IsRequestedForMainFrame
	//  Returns TRUE if the context menu was requested on the main frame and
	//  FALSE if invoked on another frame.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_isrequestedformainframe">See the ICoreWebView2ContextMenuTarget article.</a>
	IsRequestedForMainFrame() bool // property
	// PageUri
	//  Gets the uri of the page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_pageuri">See the ICoreWebView2ContextMenuTarget article.</a>
	PageUri() string // property
	// FrameUri
	//  Gets the uri of the frame. Will match the PageUri if `IsRequestedForMainFrame` is TRUE.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_frameuri">See the ICoreWebView2ContextMenuTarget article.</a>
	FrameUri() string // property
	// HasLinkUri
	//  Returns TRUE if the context menu is requested on HTML containing an anchor tag.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_haslinkuri">See the ICoreWebView2ContextMenuTarget article.</a>
	HasLinkUri() bool // property
	// LinkUri
	//  Gets the uri of the link(if `HasLinkUri` is TRUE, null otherwise).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_linkuri">See the ICoreWebView2ContextMenuTarget article.</a>
	LinkUri() string // property
	// HasLinkText
	//  Returns TRUE if the context menu is requested on text element that contains an anchor tag.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_haslinktext">See the ICoreWebView2ContextMenuTarget article.</a>
	HasLinkText() bool // property
	// LinkText
	//  Gets the text of the link(if `HasLinkText` is TRUE, null otherwise).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_linktext">See the ICoreWebView2ContextMenuTarget article.</a>
	LinkText() string // property
	// HasSourceUri
	//  Returns TRUE if the context menu is requested on HTML containing a source uri.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_hassourceuri">See the ICoreWebView2ContextMenuTarget article.</a>
	HasSourceUri() bool // property
	// SourceUri
	//  Gets the active source uri of element(if `HasSourceUri` is TRUE, null otherwise).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_sourceuri">See the ICoreWebView2ContextMenuTarget article.</a>
	SourceUri() string // property
	// HasSelection
	//  Returns TRUE if the context menu is requested on a selection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_hasselection">See the ICoreWebView2ContextMenuTarget article.</a>
	HasSelection() bool // property
	// SelectionText
	//  Gets the selected text(if `HasSelection` is TRUE, null otherwise).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_selectiontext">See the ICoreWebView2ContextMenuTarget article.</a>
	SelectionText() string // property
}

// TCoreWebView2ContextMenuTarget Parent: TObject
//
//	Represents the information regarding the context menu target.
//	Includes the context selected and the appropriate data used for the actions of a context menu.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget">See the ICoreWebView2ContextMenuTarget article.</a>
type TCoreWebView2ContextMenuTarget struct {
	TObject
}

func NewCoreWebView2ContextMenuTarget(aBaseIntf ICoreWebView2ContextMenuTarget) ICoreWebView2ContextMenuTarget {
	r1 := WV().SysCallN(160, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ContextMenuTarget(r1)
}

func (m *TCoreWebView2ContextMenuTarget) Initialized() bool {
	r1 := WV().SysCallN(166, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) BaseIntf() ICoreWebView2ContextMenuTarget {
	var resultCoreWebView2ContextMenuTarget uintptr
	WV().SysCallN(158, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuTarget)))
	return AsCoreWebView2ContextMenuTarget(resultCoreWebView2ContextMenuTarget)
}

func (m *TCoreWebView2ContextMenuTarget) Kind() TWVMenuTargetKind {
	r1 := WV().SysCallN(169, m.Instance())
	return TWVMenuTargetKind(r1)
}

func (m *TCoreWebView2ContextMenuTarget) IsEditable() bool {
	r1 := WV().SysCallN(167, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) IsRequestedForMainFrame() bool {
	r1 := WV().SysCallN(168, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) PageUri() string {
	r1 := WV().SysCallN(172, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuTarget) FrameUri() string {
	r1 := WV().SysCallN(161, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuTarget) HasLinkUri() bool {
	r1 := WV().SysCallN(163, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) LinkUri() string {
	r1 := WV().SysCallN(171, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuTarget) HasLinkText() bool {
	r1 := WV().SysCallN(162, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) LinkText() string {
	r1 := WV().SysCallN(170, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuTarget) HasSourceUri() bool {
	r1 := WV().SysCallN(165, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) SourceUri() string {
	r1 := WV().SysCallN(174, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuTarget) HasSelection() bool {
	r1 := WV().SysCallN(164, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuTarget) SelectionText() string {
	r1 := WV().SysCallN(173, m.Instance())
	return GoStr(r1)
}

func CoreWebView2ContextMenuTargetClass() TClass {
	ret := WV().SysCallN(159)
	return TClass(ret)
}
