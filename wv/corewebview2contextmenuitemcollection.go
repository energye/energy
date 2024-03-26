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

// ICoreWebView2ContextMenuItemCollection Parent: IObject
//
//	Represents a collection of ContextMenuItem objects. Used to get, remove and add
//	ContextMenuItem objects at the specified index.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection">See the ICoreWebView2ContextMenuItemCollection article.</a>
type ICoreWebView2ContextMenuItemCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuItemCollection // property
	// Count
	//  Gets the number of `ContextMenuItem` objects contained in the `ContextMenuItemCollection`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#get_count">See the ICoreWebView2ContextMenuItemCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the `ContextMenuItem` at the specified index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#getvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	Items(idx uint32) ICoreWebView2ContextMenuItem // property
	// RemoveValueAtIndex
	//  Removes the `ContextMenuItem` at the specified index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	RemoveValueAtIndex(index uint32) bool // function
	// InsertValueAtIndex
	//  Inserts the `ContextMenuItem` at the specified index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#insertvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	InsertValueAtIndex(index uint32, aValue ICoreWebView2ContextMenuItem) bool // function
	// AppendValue
	//  Appends the aValue item at the end of the collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#insertvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	AppendValue(aValue ICoreWebView2ContextMenuItem) bool // function
	// RemoveAllMenuItems
	//  Removes all items from the collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	RemoveAllMenuItems() // procedure
	// RemoveMenuItem
	//  Removes the item with the commandId value specified in the paramaters.
	//  <param name="aCommandId">The commandId value of the item that has to be removed.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	RemoveMenuItem(aCommandId int32) // procedure
	// RemoveMenuItem1
	//  Removes the item with the label value specified in the paramaters.
	//  <param name="aLabel">The label value of the item that has to be removed.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</a>
	RemoveMenuItem1(aLabel string) // procedure
}

// TCoreWebView2ContextMenuItemCollection Parent: TObject
//
//	Represents a collection of ContextMenuItem objects. Used to get, remove and add
//	ContextMenuItem objects at the specified index.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection">See the ICoreWebView2ContextMenuItemCollection article.</a>
type TCoreWebView2ContextMenuItemCollection struct {
	TObject
}

func NewCoreWebView2ContextMenuItemCollection(aBaseIntf ICoreWebView2ContextMenuItemCollection) ICoreWebView2ContextMenuItemCollection {
	r1 := WV().SysCallN(126, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ContextMenuItemCollection(r1)
}

func (m *TCoreWebView2ContextMenuItemCollection) Initialized() bool {
	r1 := WV().SysCallN(127, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItemCollection) BaseIntf() ICoreWebView2ContextMenuItemCollection {
	var resultCoreWebView2ContextMenuItemCollection uintptr
	WV().SysCallN(123, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuItemCollection)))
	return AsCoreWebView2ContextMenuItemCollection(resultCoreWebView2ContextMenuItemCollection)
}

func (m *TCoreWebView2ContextMenuItemCollection) Count() uint32 {
	r1 := WV().SysCallN(125, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ContextMenuItemCollection) Items(idx uint32) ICoreWebView2ContextMenuItem {
	var resultCoreWebView2ContextMenuItem uintptr
	WV().SysCallN(129, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2ContextMenuItem)))
	return AsCoreWebView2ContextMenuItem(resultCoreWebView2ContextMenuItem)
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) bool {
	r1 := WV().SysCallN(133, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, aValue ICoreWebView2ContextMenuItem) bool {
	r1 := WV().SysCallN(128, m.Instance(), uintptr(index), GetObjectUintptr(aValue))
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItemCollection) AppendValue(aValue ICoreWebView2ContextMenuItem) bool {
	r1 := WV().SysCallN(122, m.Instance(), GetObjectUintptr(aValue))
	return GoBool(r1)
}

func CoreWebView2ContextMenuItemCollectionClass() TClass {
	ret := WV().SysCallN(124)
	return TClass(ret)
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveAllMenuItems() {
	WV().SysCallN(130, m.Instance())
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveMenuItem(aCommandId int32) {
	WV().SysCallN(131, m.Instance(), uintptr(aCommandId))
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveMenuItem1(aLabel string) {
	WV().SysCallN(132, m.Instance(), PascalStr(aLabel))
}
