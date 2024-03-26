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

// ICoreWebView2ContextMenuItem Parent: IObject
//
//	Represents a context menu item of a context menu displayed by WebView.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem">See the ICoreWebView2ContextMenuItem article.</a>
type ICoreWebView2ContextMenuItem interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuItem // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ContextMenuItem) // property
	// Name
	//  Gets the unlocalized name for the `ContextMenuItem`. Use this to
	//  distinguish between context menu item types. This will be the English
	//  label of the menu item in lower camel case. For example, the "Save as"
	//  menu item will be "saveAs". Extension menu items will be "extension",
	//  custom menu items will be "custom" and spellcheck items will be
	//  "spellCheck".
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_name">See the ICoreWebView2ContextMenuItem article.</a>
	Name() string // property
	// Label
	//  Gets the localized label for the `ContextMenuItem`. Will contain an
	//  ampersand for characters to be used as keyboard accelerator.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_label">See the ICoreWebView2ContextMenuItem article.</a>
	Label() string // property
	// CommandId
	//  Gets the Command ID for the `ContextMenuItem`. Use this to report the
	//  `SelectedCommandId` in `ContextMenuRequested` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_commandid">See the ICoreWebView2ContextMenuItem article.</a>
	CommandId() int32 // property
	// ShortcutKeyDescription
	//  Gets the localized keyboard shortcut for this ContextMenuItem. It will be
	//  the empty string if there is no keyboard shortcut. This is text intended
	//  to be displayed to the end user to show the keyboard shortcut. For example
	//  this property is Ctrl+Shift+I for the "Inspect" `ContextMenuItem`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_shortcutkeydescription">See the ICoreWebView2ContextMenuItem article.</a>
	ShortcutKeyDescription() string // property
	// Icon
	//  Gets the Icon for the `ContextMenuItem` in PNG, Bitmap or SVG formats in the form of an IStream.
	//  Stream will be rewound to the start of the image data.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_icon">See the ICoreWebView2ContextMenuItem article.</a>
	Icon() IStream // property
	// Kind
	//  Gets the `ContextMenuItem` kind.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_kind">See the ICoreWebView2ContextMenuItem article.</a>
	Kind() TWVMenuItemKind // property
	// IsEnabled
	//  Gets the enabled property of the `ContextMenuItem`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_isenabled">See the ICoreWebView2ContextMenuItem article.</a>
	IsEnabled() bool // property
	// SetIsEnabled Set IsEnabled
	SetIsEnabled(AValue bool) // property
	// IsChecked
	//  Gets the checked property of the `ContextMenuItem`, used if the kind is Check box or Radio.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_ischecked">See the ICoreWebView2ContextMenuItem article.</a>
	IsChecked() bool // property
	// SetIsChecked Set IsChecked
	SetIsChecked(AValue bool) // property
	// Children
	//  Gets the list of children menu items through a `ContextMenuItemCollection`
	//  if the kind is Submenu. If the kind is not submenu, will return null.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_children">See the ICoreWebView2ContextMenuItem article.</a>
	Children() ICoreWebView2ContextMenuItemCollection // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
}

// TCoreWebView2ContextMenuItem Parent: TObject
//
//	Represents a context menu item of a context menu displayed by WebView.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem">See the ICoreWebView2ContextMenuItem article.</a>
type TCoreWebView2ContextMenuItem struct {
	TObject
}

func NewCoreWebView2ContextMenuItem(aBaseIntf ICoreWebView2ContextMenuItem) ICoreWebView2ContextMenuItem {
	r1 := WV().SysCallN(139, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ContextMenuItem(r1)
}

func (m *TCoreWebView2ContextMenuItem) Initialized() bool {
	r1 := WV().SysCallN(141, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItem) BaseIntf() ICoreWebView2ContextMenuItem {
	var resultCoreWebView2ContextMenuItem uintptr
	WV().SysCallN(135, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ContextMenuItem)))
	return AsCoreWebView2ContextMenuItem(resultCoreWebView2ContextMenuItem)
}

func (m *TCoreWebView2ContextMenuItem) SetBaseIntf(AValue ICoreWebView2ContextMenuItem) {
	WV().SysCallN(135, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ContextMenuItem) Name() string {
	r1 := WV().SysCallN(146, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuItem) Label() string {
	r1 := WV().SysCallN(145, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuItem) CommandId() int32 {
	r1 := WV().SysCallN(138, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2ContextMenuItem) ShortcutKeyDescription() string {
	r1 := WV().SysCallN(147, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ContextMenuItem) Icon() IStream {
	var resultStream uintptr
	WV().SysCallN(140, m.Instance(), uintptr(unsafePointer(&resultStream)))
	return AsStream(resultStream)
}

func (m *TCoreWebView2ContextMenuItem) Kind() TWVMenuItemKind {
	r1 := WV().SysCallN(144, m.Instance())
	return TWVMenuItemKind(r1)
}

func (m *TCoreWebView2ContextMenuItem) IsEnabled() bool {
	r1 := WV().SysCallN(143, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItem) SetIsEnabled(AValue bool) {
	WV().SysCallN(143, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ContextMenuItem) IsChecked() bool {
	r1 := WV().SysCallN(142, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ContextMenuItem) SetIsChecked(AValue bool) {
	WV().SysCallN(142, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ContextMenuItem) Children() ICoreWebView2ContextMenuItemCollection {
	var resultCoreWebView2ContextMenuItemCollection uintptr
	WV().SysCallN(136, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContextMenuItemCollection)))
	return AsCoreWebView2ContextMenuItemCollection(resultCoreWebView2ContextMenuItemCollection)
}

func (m *TCoreWebView2ContextMenuItem) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(134, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func CoreWebView2ContextMenuItemClass() TClass {
	ret := WV().SysCallN(137)
	return TClass(ret)
}
