package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TreeView is a representation of GTK's GtkTreeView.
type TreeView struct {
	Container
}

func (v *TreeView) native() *C.GtkTreeView {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTreeView(unsafe.Pointer(v.GObject))
}

func wrapTreeView(obj *Object) *TreeView {
	return &TreeView{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewTreeView is a wrapper around gtk_tree_view_new().
func NewTreeView() *TreeView {
	c := C.gtk_tree_view_new()
	if c == nil {
		return nil
	}
	return wrapTreeView(ToGoObject(unsafe.Pointer(c)))
}

// NewTreeViewWithModel is a wrapper around gtk_tree_view_new_with_model().
func NewTreeViewWithModel(model IListStore) *TreeView {
	m := model.(*ListStore)
	c := C.gtk_tree_view_new_with_model(C.toGtkTreeModel(unsafe.Pointer(m.GObject)))
	if c == nil {
		return nil
	}
	return wrapTreeView(ToGoObject(unsafe.Pointer(c)))
}

// GetModel is a wrapper around gtk_tree_view_get_model().
func (v *TreeView) GetModel() IListStore {
	c := C.gtk_tree_view_get_model(v.native())
	if c == nil {
		return nil
	}
	return wrapListStore(ToGoObject(unsafe.Pointer(c)))
}

// SetModel is a wrapper around gtk_tree_view_set_model().
func (v *TreeView) SetModel(model IListStore) {
	var mptr *C.GtkTreeModel
	if model != nil {
		m := model.(*ListStore)
		mptr = C.toGtkTreeModel(unsafe.Pointer(m.GObject))
	}
	C.gtk_tree_view_set_model(v.native(), mptr)
}

// GetSelection is a wrapper around gtk_tree_view_get_selection().
func (v *TreeView) GetSelection() ITreeSelection {
	c := C.gtk_tree_view_get_selection(v.native())
	if c == nil {
		return nil
	}
	return wrapTreeSelection(ToGoObject(unsafe.Pointer(c)))
}

// AppendColumn is a wrapper around gtk_tree_view_append_column().
func (v *TreeView) AppendColumn(column ITreeViewColumn) int {
	c := column.(*TreeViewColumn)
	return int(C.gtk_tree_view_append_column(v.native(), c.native()))
}

// SetHeadersVisible is a wrapper around gtk_tree_view_set_headers_visible().
func (v *TreeView) SetHeadersVisible(show bool) {
	C.gtk_tree_view_set_headers_visible(v.native(), CBool(show))
}

// GetHeadersVisible is a wrapper around gtk_tree_view_get_headers_visible().
func (v *TreeView) GetHeadersVisible() bool {
	return GoBool(C.gtk_tree_view_get_headers_visible(v.native()))
}

// ExpandAll is a wrapper around gtk_tree_view_expand_all().
func (v *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(v.native())
}

// CollapseAll is a wrapper around gtk_tree_view_collapse_all().
func (v *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(v.native())
}
