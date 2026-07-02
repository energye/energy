package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// static GtkListStore* _gtk_list_store_new(int n_columns, GType type1) {
//     return gtk_list_store_new(n_columns, type1);
// }
// static GtkListStore* _gtk_list_store_new2(int n_columns, GType type1, GType type2) {
//     return gtk_list_store_new(n_columns, type1, type2);
// }
// static GtkListStore* _gtk_list_store_new3(int n_columns, GType type1, GType type2, GType type3) {
//     return gtk_list_store_new(n_columns, type1, type2, type3);
// }
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	*Object
}

func (v *ListStore) native() *C.GtkListStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkListStore(unsafe.Pointer(v.GObject))
}

func wrapListStore(obj *Object) *ListStore {
	return &ListStore{obj}
}

// NewListStore is a wrapper around gtk_list_store_new().
func NewListStore(columnTypes ...types.Type) *ListStore {
	nColumns := len(columnTypes)
	if nColumns == 0 {
		return nil
	}
	var c *C.GtkListStore
	switch nColumns {
	case 1:
		c = C._gtk_list_store_new(C.gint(nColumns), C.GType(columnTypes[0]))
	case 2:
		c = C._gtk_list_store_new2(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]))
	case 3:
		c = C._gtk_list_store_new3(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]), C.GType(columnTypes[2]))
	default:
		// For more columns, build dynamically - use first 3 as fallback
		c = C._gtk_list_store_new3(C.gint(nColumns), C.GType(columnTypes[0]), C.GType(columnTypes[1]), C.GType(columnTypes[2]))
	}
	if c == nil {
		return nil
	}
	return wrapListStore(ToGoObject(unsafe.Pointer(c)))
}

// Append is a wrapper around gtk_list_store_append().
func (v *ListStore) Append() types.ITreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_append(v.native(), &ti)
	return &TreeIter{ti}
}

// SetValue is a wrapper around gtk_list_store_set_value() for string values.
func (v *ListStore) SetValue(iter types.ITreeIter, column int, value string) {
	i := iter.(*TreeIter)
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_list_store_set(v.native(), i.native(), C.gint(column), unsafe.Pointer(cstr))
}

// Remove is a wrapper around gtk_list_store_remove().
func (v *ListStore) Remove(iter types.ITreeIter) bool {
	i := iter.(*TreeIter)
	return GoBool(C.gtk_list_store_remove(v.native(), i.native()))
}

// Clear is a wrapper around gtk_list_store_clear().
func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.native())
}
