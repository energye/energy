package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	*Object
}

func (v *TreeViewColumn) native() *C.GtkTreeViewColumn {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTreeViewColumn(unsafe.Pointer(v.GObject))
}

func wrapTreeViewColumn(obj *Object) *TreeViewColumn {
	return &TreeViewColumn{obj}
}

// NewTreeViewColumn is a wrapper around gtk_tree_view_column_new().
func NewTreeViewColumn() *TreeViewColumn {
	c := C.gtk_tree_view_column_new()
	if c == nil {
		return nil
	}
	return wrapTreeViewColumn(ToGoObject(unsafe.Pointer(c)))
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (v *TreeViewColumn) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tree_view_column_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (v *TreeViewColumn) GetTitle() string {
	c := C.gtk_tree_view_column_get_title(v.native())
	return C.GoString((*C.char)(c))
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (v *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	cr := cell.(cellRenderer)
	C.gtk_tree_view_column_pack_start(v.native(), cr.toCellRenderer(), CBool(expand))
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (v *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	cr := renderer.(cellRenderer)
	C.gtk_tree_view_column_add_attribute(v.native(), cr.toCellRenderer(), (*C.gchar)(cstr), C.gint(column))
}

func (m *TreeViewColumn) SetOnClicked(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnClicked, callback.C_trampoline_2_void, fn, 0)
}
