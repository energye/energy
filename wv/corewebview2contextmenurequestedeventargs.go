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

// ICoreWebView2ContextMenuRequestedEventArgs Parent: IObject
//
//	Event args for the ContextMenuRequested event. Will contain the selection information
//	and a collection of all of the default context menu items that the WebView
//	would show. Allows the app to draw its own context menu or add/remove
//	from the default context menu.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
type ICoreWebView2ContextMenuRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuRequestedEventArgs // property
	// MenuItems
	//  Gets the collection of `ContextMenuItem` objects.
	//  See `ICoreWebView2ContextMenuItemCollection` for more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_menuitems">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	MenuItems() ICoreWebView2ContextMenuItemCollection // property
	// ContextMenuTarget
	//  Gets the target information associated with the requested context menu.
	//  See `ICoreWebView2ContextMenuTarget` for more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_contextmenutarget">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	ContextMenuTarget() ICoreWebView2ContextMenuTarget // property
	// Location
	//  Gets the coordinates where the context menu request occurred in relation to the upper
	//  left corner of the WebView bounds.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_location">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	Location() (resultPoint TPoint) // property
	// SelectedCommandId
	//  Gets the selected CommandId.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_selectedcommandid">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	SelectedCommandId() int32 // property
	// SetSelectedCommandId Set SelectedCommandId
	SetSelectedCommandId(AValue int32) // property
	// Handled
	//  Gets whether the `ContextMenuRequested` event is handled by host.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_handled">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event when the custom context menu is closed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#getdeferral">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2ContextMenuRequestedEventArgs Parent: TObject
//
//	Event args for the ContextMenuRequested event. Will contain the selection information
//	and a collection of all of the default context menu items that the WebView
//	would show. Allows the app to draw its own context menu or add/remove
//	from the default context menu.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs">See the ICoreWebView2ContextMenuRequestedEventArgs article.</a>
type TCoreWebView2ContextMenuRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2ContextMenuRequestedEventArgs(aArgs ICoreWebView2ContextMenuRequestedEventArgs) ICoreWebView2ContextMenuRequestedEventArgs {
	r1 := WV().SysCallN(151, GetObjectUintptr(aArgs))
	return AsCoreWebView2ContextMenuRequestedEventArgs(r1)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(154, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) BaseIntf() ICoreWebView2ContextMenuRequestedEventArgs {
	var resultCoreWebView2ContextMenuRequestedEventArgs uintptr
	WV().SysCallN(148, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuRequestedEventArgs)))
	return AsCoreWebView2ContextMenuRequestedEventArgs(resultCoreWebView2ContextMenuRequestedEventArgs)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) MenuItems() ICoreWebView2ContextMenuItemCollection {
	var resultCoreWebView2ContextMenuItemCollection uintptr
	WV().SysCallN(156, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuItemCollection)))
	return AsCoreWebView2ContextMenuItemCollection(resultCoreWebView2ContextMenuItemCollection)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) ContextMenuTarget() ICoreWebView2ContextMenuTarget {
	var resultCoreWebView2ContextMenuTarget uintptr
	WV().SysCallN(150, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuTarget)))
	return AsCoreWebView2ContextMenuTarget(resultCoreWebView2ContextMenuTarget)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Location() (resultPoint TPoint) {
	WV().SysCallN(155, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SelectedCommandId() int32 {
	r1 := WV().SysCallN(157, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SetSelectedCommandId(AValue int32) {
	WV().SysCallN(157, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Handled() bool {
	r1 := WV().SysCallN(153, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SetHandled(AValue bool) {
	WV().SysCallN(153, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(152, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func CoreWebView2ContextMenuRequestedEventArgsClass() TClass {
	ret := WV().SysCallN(149)
	return TClass(ret)
}
