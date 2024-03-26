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

// ICoreWebView2PermissionSettingCollectionView Parent: IObject
//
//	Read-only collection of PermissionSettings (origin, kind, and state). Used to list
//	the nondefault permission settings on the profile that are persisted across
//	sessions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview">See the ICoreWebView2PermissionSettingCollectionView article.</a>
type ICoreWebView2PermissionSettingCollectionView interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionSettingCollectionView // property
	// ValueAtIndex
	//  Gets the `ICoreWebView2PermissionSetting` at the specified index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview#getvalueatindex">See the ICoreWebView2PermissionSettingCollectionView article.</a>
	ValueAtIndex(idx uint32) ICoreWebView2PermissionSetting // property
	// Count
	//  The number of `ICoreWebView2PermissionSetting`s in the collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview#get_count">See the ICoreWebView2PermissionSettingCollectionView article.</a>
	Count() uint32 // property
}

// TCoreWebView2PermissionSettingCollectionView Parent: TObject
//
//	Read-only collection of PermissionSettings (origin, kind, and state). Used to list
//	the nondefault permission settings on the profile that are persisted across
//	sessions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview">See the ICoreWebView2PermissionSettingCollectionView article.</a>
type TCoreWebView2PermissionSettingCollectionView struct {
	TObject
}

func NewCoreWebView2PermissionSettingCollectionView(aBaseIntf ICoreWebView2PermissionSettingCollectionView) ICoreWebView2PermissionSettingCollectionView {
	r1 := WV().SysCallN(449, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PermissionSettingCollectionView(r1)
}

func (m *TCoreWebView2PermissionSettingCollectionView) Initialized() bool {
	r1 := WV().SysCallN(450, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionSettingCollectionView) BaseIntf() ICoreWebView2PermissionSettingCollectionView {
	var resultCoreWebView2PermissionSettingCollectionView uintptr
	WV().SysCallN(446, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PermissionSettingCollectionView)))
	return AsCoreWebView2PermissionSettingCollectionView(resultCoreWebView2PermissionSettingCollectionView)
}

func (m *TCoreWebView2PermissionSettingCollectionView) ValueAtIndex(idx uint32) ICoreWebView2PermissionSetting {
	var resultCoreWebView2PermissionSetting uintptr
	WV().SysCallN(451, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2PermissionSetting)))
	return AsCoreWebView2PermissionSetting(resultCoreWebView2PermissionSetting)
}

func (m *TCoreWebView2PermissionSettingCollectionView) Count() uint32 {
	r1 := WV().SysCallN(448, m.Instance())
	return uint32(r1)
}

func CoreWebView2PermissionSettingCollectionViewClass() TClass {
	ret := WV().SysCallN(447)
	return TClass(ret)
}
