//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import "unsafe"

// GMenuModel is a representation of GLib's GMenuModel.
type GMenuModel struct {
	Object
}

func AsGMenuModel(ptr unsafe.Pointer) *GMenuModel {
	if ptr == nil {
		return nil
	}
	m := new(GMenuModel)
	m.instance = ptr
	return m
}

// Native returns the underlying GMenuModel pointer as uintptr.
func (m *GMenuModel) Native() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// IsMutable is a wrapper around g_menu_model_is_mutable().
func (m *GMenuModel) IsMutable() bool {
	r := glib2_0.SysCall("g_menu_model_is_mutable", m.Instance())
	return ToGoBool(r)
}

// GetNItems is a wrapper around g_menu_model_get_n_items().
func (m *GMenuModel) GetNItems() int {
	r := glib2_0.SysCall("g_menu_model_get_n_items", m.Instance())
	return int(r)
}

// GetItemLink is a wrapper around g_menu_model_get_item_link().
func (m *GMenuModel) GetItemLink(index int, link string) *GMenuModel {
	cstr := CStr(link)
	r := glib2_0.SysCall("g_menu_model_get_item_link", m.Instance(), uintptr(index), cstr)
	if r == 0 {
		return nil
	}
	return AsGMenuModel(unsafe.Pointer(r))
}

// ItemsChanged is a wrapper around g_menu_model_items_changed().
func (m *GMenuModel) ItemsChanged(position, removed, added int) {
	glib2_0.SysCall("g_menu_model_items_changed", m.Instance(),
		uintptr(position), uintptr(removed), uintptr(added))
}
