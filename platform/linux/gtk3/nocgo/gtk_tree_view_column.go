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

// TreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	Object
}

func AsTreeViewColumn(ptr unsafe.Pointer) *TreeViewColumn {
	if ptr == nil {
		return nil
	}
	m := new(TreeViewColumn)
	m.instance = ptr
	return m
}

// NewTreeViewColumn is a wrapper around gtk_tree_view_column_new().
func NewTreeViewColumn() *TreeViewColumn {
	r := gtk3.SysCall("gtk_tree_view_column_new")
	if r == 0 {
		return nil
	}
	return AsTreeViewColumn(unsafe.Pointer(r))
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (m *TreeViewColumn) SetTitle(title string) {
	gtk3.SysCall("gtk_tree_view_column_set_title", m.Instance(), CStr(title))
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (m *TreeViewColumn) GetTitle() string {
	r := gtk3.SysCall("gtk_tree_view_column_get_title", m.Instance())
	return GoStr(r)
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (m *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	gtk3.SysCall("gtk_tree_view_column_pack_start", m.Instance(), cell.Instance(), ToCBool(expand))
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (m *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	gtk3.SysCall("gtk_tree_view_column_add_attribute", m.Instance(),
		renderer.Instance(), CStr(attribute), uintptr(column))
}
