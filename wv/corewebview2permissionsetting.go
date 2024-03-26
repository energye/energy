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

// ICoreWebView2PermissionSetting Parent: IObject
//
//	Provides a set of properties for a permission setting.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting">See the ICoreWebView2PermissionSetting article.</a>
type ICoreWebView2PermissionSetting interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionSetting // property
	// PermissionKind
	//  The kind of the permission setting. See `COREWEBVIEW2_PERMISSION_KIND` for
	//  more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionkind">See the ICoreWebView2PermissionSetting article.</a>
	PermissionKind() TWVPermissionKind // property
	// PermissionOrigin
	//  The origin of the permission setting.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionorigin">See the ICoreWebView2PermissionSetting article.</a>
	PermissionOrigin() string // property
	// PermissionState
	//  The state of the permission setting.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionstate">See the ICoreWebView2PermissionSetting article.</a>
	PermissionState() TWVPermissionState // property
}

// TCoreWebView2PermissionSetting Parent: TObject
//
//	Provides a set of properties for a permission setting.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting">See the ICoreWebView2PermissionSetting article.</a>
type TCoreWebView2PermissionSetting struct {
	TObject
}

func NewCoreWebView2PermissionSetting(aBaseIntf ICoreWebView2PermissionSetting) ICoreWebView2PermissionSetting {
	r1 := WV().SysCallN(454, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PermissionSetting(r1)
}

func (m *TCoreWebView2PermissionSetting) Initialized() bool {
	r1 := WV().SysCallN(455, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionSetting) BaseIntf() ICoreWebView2PermissionSetting {
	var resultCoreWebView2PermissionSetting uintptr
	WV().SysCallN(452, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PermissionSetting)))
	return AsCoreWebView2PermissionSetting(resultCoreWebView2PermissionSetting)
}

func (m *TCoreWebView2PermissionSetting) PermissionKind() TWVPermissionKind {
	r1 := WV().SysCallN(456, m.Instance())
	return TWVPermissionKind(r1)
}

func (m *TCoreWebView2PermissionSetting) PermissionOrigin() string {
	r1 := WV().SysCallN(457, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2PermissionSetting) PermissionState() TWVPermissionState {
	r1 := WV().SysCallN(458, m.Instance())
	return TWVPermissionState(r1)
}

func CoreWebView2PermissionSettingClass() TClass {
	ret := WV().SysCallN(453)
	return TClass(ret)
}
