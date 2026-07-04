package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	*Object
}

func (v *TreeStore) native() *C.GtkTreeStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTreeStore(unsafe.Pointer(v.GObject))
}

func wrapTreeStore(obj *Object) *TreeStore {
	return &TreeStore{obj}
}

// NewTreeStore is a wrapper around gtk_tree_store_new().
func NewTreeStore(columnTypes ...types.Type) *TreeStore {
	nColumns := len(columnTypes)
	if nColumns == 0 {
		return nil
	}
	var c *C.GtkTreeStore
	switch nColumns {
	case 1:
		c = C._gtk_tree_store_new(C.gint(nColumns), C.GType(columnTypes[0]))
	case 2:
		c = C._gtk_tree_store_new2(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]))
	case 3:
		c = C._gtk_tree_store_new3(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]), C.GType(columnTypes[2]))
	default:
		c = C._gtk_tree_store_new3(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]), C.GType(columnTypes[2]))
	}
	if c == nil {
		return nil
	}
	return wrapTreeStore(ToGoObject(unsafe.Pointer(c)))
}

// Append is a wrapper around gtk_tree_store_append().
// parent can be nil to append a top-level row.
func (v *TreeStore) Append(parent types.ITreeIter) types.ITreeIter {
	var ti C.GtkTreeIter
	var p *C.GtkTreeIter
	if parent != nil {
		p = parent.(*TreeIter).native()
	}
	C.gtk_tree_store_append(v.native(), &ti, p)
	return &TreeIter{ti}
}

// SetValue is a wrapper around gtk_tree_store_set_value() for string values.
func (v *TreeStore) SetValue(iter types.ITreeIter, column int, value string) {
	i := iter.(*TreeIter)
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_tree_store_set(v.native(), i.native(), C.gint(column), unsafe.Pointer(cstr))
}

// Remove is a wrapper around gtk_tree_store_remove().
func (v *TreeStore) Remove(iter types.ITreeIter) bool {
	i := iter.(*TreeIter)
	return GoBool(C.gtk_tree_store_remove(v.native(), i.native()))
}

// Clear is a wrapper around gtk_tree_store_clear().
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}

// Insert is a wrapper around gtk_tree_store_insert().
func (v *TreeStore) Insert(parent types.ITreeIter, position int) types.ITreeIter {
	var ti C.GtkTreeIter
	var p *C.GtkTreeIter
	if parent != nil {
		p = parent.(*TreeIter).native()
	}
	C.gtk_tree_store_insert(v.native(), &ti, p, C.gint(position))
	return &TreeIter{ti}
}

// InsertBefore is a wrapper around gtk_tree_store_insert_before().
func (v *TreeStore) InsertBefore(parent, sibling types.ITreeIter) types.ITreeIter {
	var ti C.GtkTreeIter
	var p *C.GtkTreeIter
	var s *C.GtkTreeIter
	if parent != nil {
		p = parent.(*TreeIter).native()
	}
	if sibling != nil {
		s = sibling.(*TreeIter).native()
	}
	C.gtk_tree_store_insert_before(v.native(), &ti, p, s)
	return &TreeIter{ti}
}

// InsertAfter is a wrapper around gtk_tree_store_insert_after().
func (v *TreeStore) InsertAfter(parent, sibling types.ITreeIter) types.ITreeIter {
	var ti C.GtkTreeIter
	var p *C.GtkTreeIter
	var s *C.GtkTreeIter
	if parent != nil {
		p = parent.(*TreeIter).native()
	}
	if sibling != nil {
		s = sibling.(*TreeIter).native()
	}
	C.gtk_tree_store_insert_after(v.native(), &ti, p, s)
	return &TreeIter{ti}
}

// MoveBefore is a wrapper around gtk_tree_store_move_before().
func (v *TreeStore) MoveBefore(iter, position types.ITreeIter) {
	i := iter.(*TreeIter)
	p := position.(*TreeIter)
	C.gtk_tree_store_move_before(v.native(), i.native(), p.native())
}

// MoveAfter is a wrapper around gtk_tree_store_move_after().
func (v *TreeStore) MoveAfter(iter, position types.ITreeIter) {
	i := iter.(*TreeIter)
	p := position.(*TreeIter)
	C.gtk_tree_store_move_after(v.native(), i.native(), p.native())
}

// Swap is a wrapper around gtk_tree_store_swap().
func (v *TreeStore) Swap(a, b types.ITreeIter) {
	ia := a.(*TreeIter)
	ib := b.(*TreeIter)
	C.gtk_tree_store_swap(v.native(), ia.native(), ib.native())
}

// IterIsValid is a wrapper around gtk_tree_store_iter_is_valid().
func (v *TreeStore) IterIsValid(iter types.ITreeIter) bool {
	i := iter.(*TreeIter)
	return GoBool(C.gtk_tree_store_iter_is_valid(v.native(), i.native()))
}

// IsAncestor is a wrapper around gtk_tree_store_is_ancestor().
func (v *TreeStore) IsAncestor(iter, descendant types.ITreeIter) bool {
	i := iter.(*TreeIter)
	d := descendant.(*TreeIter)
	return GoBool(C.gtk_tree_store_is_ancestor(v.native(), i.native(), d.native()))
}

// GetDepth is a wrapper around gtk_tree_store_iter_depth().
func (v *TreeStore) GetDepth(iter types.ITreeIter) int {
	i := iter.(*TreeIter)
	return int(C.gtk_tree_store_iter_depth(v.native(), i.native()))
}
