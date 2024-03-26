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

// ICoreWebView2PermissionRequestedEventArgs Parent: IObject
//
//	Event args for the PermissionRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
type ICoreWebView2PermissionRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionRequestedEventArgs // property
	// URI
	//  The origin of the web content that requests the permission.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_uri">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
	URI() string // property
	// State
	//  The status of a permission request,(for example is the request is granted).
	//  The default value is `COREWEBVIEW2_PERMISSION_STATE_DEFAULT`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_state">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
	State() TWVPermissionState // property
	// SetState Set State
	SetState(AValue TWVPermissionState) // property
	// PermissionKind
	//  The type of the permission that is requested.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_permissionkind">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
	PermissionKind() TWVPermissionKind // property
	// IsUserInitiated
	//  `TRUE` when the permission request was initiated through a user gesture.
	//  NOTE: Being initiated through a user gesture does not mean that user intended
	//  to access the associated resource.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_isuserinitiated">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
	IsUserInitiated() bool // property
	// Deferral
	//  Gets an `ICoreWebView2Deferral` object. Use the deferral object to make
	//  the permission decision at a later time. The deferral only applies to the
	//  current request, and does not prevent the `OnPermissionRequested` event from
	//  getting raised for new requests. However, for some permission kinds the
	//  WebView will avoid creating a new request if there is a pending request of
	//  the same kind.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#getdeferral">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
	// Handled
	//  By default, both the `OnPermissionRequested` event handlers on the
	//  `CoreWebView2Frame` and the `CoreWebView2` will be invoked, with the
	//  `CoreWebView2Frame` event handlers invoked first. The host may
	//  set this flag to `TRUE` within the `CoreWebView2Frame` event handlers
	//  to prevent the remaining `CoreWebView2` event handlers from being invoked.
	//  If a deferral is taken on the event args, then you must synchronously
	//  set `Handled` to TRUE prior to taking your deferral to prevent the
	//  `CoreWebView2`s event handlers from being invoked.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs2#get_handled">See the ICoreWebView2PermissionRequestedEventArgs2 article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// SavesInProfile
	//  The permission state set from the `PermissionRequested` event is saved in
	//  the profile by default; it persists across sessions and becomes the new
	//  default behavior for future `PermissionRequested` events. Browser
	//  heuristics can affect whether the event continues to be raised when the
	//  state is saved in the profile. Set the `SavesInProfile` property to
	//  `FALSE` to not persist the state beyond the current request, and to
	//  continue to receive `PermissionRequested`
	//  events for this origin and permission kind.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3#get_savesinprofile">See the ICoreWebView2PermissionRequestedEventArgs3 article.</a>
	SavesInProfile() bool // property
	// SetSavesInProfile Set SavesInProfile
	SetSavesInProfile(AValue bool) // property
}

// TCoreWebView2PermissionRequestedEventArgs Parent: TObject
//
//	Event args for the PermissionRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs">See the ICoreWebView2PermissionRequestedEventArgs article.</a>
type TCoreWebView2PermissionRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2PermissionRequestedEventArgs(aArgs ICoreWebView2PermissionRequestedEventArgs) ICoreWebView2PermissionRequestedEventArgs {
	r1 := WV().SysCallN(437, GetObjectUintptr(aArgs))
	return AsCoreWebView2PermissionRequestedEventArgs(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(440, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) BaseIntf() ICoreWebView2PermissionRequestedEventArgs {
	var resultCoreWebView2PermissionRequestedEventArgs uintptr
	WV().SysCallN(435, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PermissionRequestedEventArgs)))
	return AsCoreWebView2PermissionRequestedEventArgs(resultCoreWebView2PermissionRequestedEventArgs)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) URI() string {
	r1 := WV().SysCallN(445, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) State() TWVPermissionState {
	r1 := WV().SysCallN(444, 0, m.Instance(), 0)
	return TWVPermissionState(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetState(AValue TWVPermissionState) {
	WV().SysCallN(444, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PermissionRequestedEventArgs) PermissionKind() TWVPermissionKind {
	r1 := WV().SysCallN(442, m.Instance())
	return TWVPermissionKind(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) IsUserInitiated() bool {
	r1 := WV().SysCallN(441, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(438, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Handled() bool {
	r1 := WV().SysCallN(439, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetHandled(AValue bool) {
	WV().SysCallN(439, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SavesInProfile() bool {
	r1 := WV().SysCallN(443, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetSavesInProfile(AValue bool) {
	WV().SysCallN(443, 1, m.Instance(), PascalBool(AValue))
}

func CoreWebView2PermissionRequestedEventArgsClass() TClass {
	ret := WV().SysCallN(436)
	return TClass(ret)
}
