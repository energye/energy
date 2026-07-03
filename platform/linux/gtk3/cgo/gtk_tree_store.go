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
