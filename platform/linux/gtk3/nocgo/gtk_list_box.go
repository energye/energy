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

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type ListBox struct {
	Container
}

func AsListBox(ptr unsafe.Pointer) *ListBox {
	if ptr == nil {
		return nil
	}
	m := new(ListBox)
	m.instance = ptr
	return m
}

func NewListBox() *ListBox {
	r := gtk3.SysCall("gtk_list_box_new")
	if r == 0 {
		return nil
	}
	return AsListBox(unsafe.Pointer(r))
}

func (m *ListBox) Prepend(child IWidget) {
	gtk3.SysCall("gtk_list_box_prepend", m.Instance(), child.Instance())
}

func (m *ListBox) Insert(child IWidget, position int) {
	gtk3.SysCall("gtk_list_box_insert", m.Instance(), child.Instance(), uintptr(position))
}

func (m *ListBox) SelectRow(row IWidget) {
	gtk3.SysCall("gtk_list_box_select_row", m.Instance(), row.Instance())
}

func (m *ListBox) GetSelectedRow() IWidget {
	r := gtk3.SysCall("gtk_list_box_get_selected_row", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

func (m *ListBox) SetSelectionMode(mode SelectionMode) {
	gtk3.SysCall("gtk_list_box_set_selection_mode", m.Instance(), uintptr(mode))
}

func (m *ListBox) GetSelectionMode() SelectionMode {
	r := gtk3.SysCall("gtk_list_box_get_selection_mode", m.Instance())
	return SelectionMode(r)
}
