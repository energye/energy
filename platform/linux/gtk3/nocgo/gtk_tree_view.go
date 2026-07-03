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

// TreeView is a representation of GTK's GtkTreeView.
type TreeView struct {
	Container
}

func AsTreeView(ptr unsafe.Pointer) *TreeView {
	if ptr == nil {
		return nil
	}
	m := new(TreeView)
	m.instance = ptr
	return m
}

// NewTreeView is a wrapper around gtk_tree_view_new().
func NewTreeView() *TreeView {
	r := gtk3.SysCall("gtk_tree_view_new")
	if r == 0 {
		return nil
	}
	return AsTreeView(unsafe.Pointer(r))
}

// NewTreeViewWithModel is a wrapper around gtk_tree_view_new_with_model().
func NewTreeViewWithModel(model IListStore) *TreeView {
	r := gtk3.SysCall("gtk_tree_view_new_with_model", model.(*ListStore).Instance())
	if r == 0 {
		return nil
	}
	return AsTreeView(unsafe.Pointer(r))
}

// GetModel is a wrapper around gtk_tree_view_get_model().
func (m *TreeView) GetModel() IListStore {
	r := gtk3.SysCall("gtk_tree_view_get_model", m.Instance())
	if r == 0 {
		return nil
	}
	return AsListStore(unsafe.Pointer(r))
}

// SetModel is a wrapper around gtk_tree_view_set_model().
func (m *TreeView) SetModel(model IListStore) {
	var modelPtr uintptr
	if model != nil {
		modelPtr = model.(*ListStore).Instance()
	}
	gtk3.SysCall("gtk_tree_view_set_model", m.Instance(), modelPtr)
}

// GetTreeModel returns the model as ITreeModel.
func (m *TreeView) GetTreeModel() ITreeModel {
	r := gtk3.SysCall("gtk_tree_view_get_model", m.Instance())
	if r == 0 {
		return nil
	}
	return AsTreeStore(unsafe.Pointer(r))
}

// SetTreeModel sets the model from an ITreeModel.
func (m *TreeView) SetTreeModel(model ITreeModel) {
	var modelPtr uintptr
	if model != nil {
		modelPtr = model.Instance()
	}
	gtk3.SysCall("gtk_tree_view_set_model", m.Instance(), modelPtr)
}

// GetSelection is a wrapper around gtk_tree_view_get_selection().
func (m *TreeView) GetSelection() ITreeSelection {
	r := gtk3.SysCall("gtk_tree_view_get_selection", m.Instance())
	if r == 0 {
		return nil
	}
	return AsTreeSelection(unsafe.Pointer(r))
}

// AppendColumn is a wrapper around gtk_tree_view_append_column().
func (m *TreeView) AppendColumn(column ITreeViewColumn) int {
	r := gtk3.SysCall("gtk_tree_view_append_column", m.Instance(), column.(*TreeViewColumn).Instance())
	return int(r)
}

// SetHeadersVisible is a wrapper around gtk_tree_view_set_headers_visible().
func (m *TreeView) SetHeadersVisible(show bool) {
	gtk3.SysCall("gtk_tree_view_set_headers_visible", m.Instance(), ToCBool(show))
}

// GetHeadersVisible is a wrapper around gtk_tree_view_get_headers_visible().
func (m *TreeView) GetHeadersVisible() bool {
	r := gtk3.SysCall("gtk_tree_view_get_headers_visible", m.Instance())
	return ToGoBool(r)
}

// ExpandAll is a wrapper around gtk_tree_view_expand_all().
func (m *TreeView) ExpandAll() {
	gtk3.SysCall("gtk_tree_view_expand_all", m.Instance())
}

// CollapseAll is a wrapper around gtk_tree_view_collapse_all().
func (m *TreeView) CollapseAll() {
	gtk3.SysCall("gtk_tree_view_collapse_all", m.Instance())
}
