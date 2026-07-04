package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TreeSelection is a representation of GTK's GtkTreeSelection.
type TreeSelection struct {
	*Object
}

func (v *TreeSelection) native() *C.GtkTreeSelection {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTreeSelection(unsafe.Pointer(v.GObject))
}

func wrapTreeSelection(obj *Object) *TreeSelection {
	return &TreeSelection{obj}
}

// SetMode is a wrapper around gtk_tree_selection_set_mode().
func (v *TreeSelection) SetMode(mode SelectionMode) {
	C.gtk_tree_selection_set_mode(v.native(), C.GtkSelectionMode(mode))
}

// GetMode is a wrapper around gtk_tree_selection_get_mode().
func (v *TreeSelection) GetMode() SelectionMode {
	return SelectionMode(C.gtk_tree_selection_get_mode(v.native()))
}

func (m *TreeSelection) SetOnChanged(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnChanged, callback.C_trampoline_2_void, fn, 0)
}

// GetSelected is a wrapper around gtk_tree_selection_get_selected().
// Returns the TreeModel and TreeIter for the currently selected row (SINGLE/BROWSE mode).
// If nothing is selected, returns nil, nil.
func (v *TreeSelection) GetSelected() (ITreeModel, ITreeIter) {
	var model *C.GtkTreeModel
	var iter C.GtkTreeIter
	ok := C.gtk_tree_selection_get_selected(v.native(), &model, &iter)
	if ok == 0 || model == nil {
		return nil, nil
	}
	obj := ToGoObject(unsafe.Pointer(model))
	return &TreeStore{Object: obj}, &TreeIter{GtkTreeIter: iter}
}

// CountSelectedRows is a wrapper around gtk_tree_selection_count_selected_rows().
func (v *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.native()))
}

// SelectAll is a wrapper around gtk_tree_selection_select_all().
func (v *TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(v.native())
}

// UnselectAll is a wrapper around gtk_tree_selection_unselect_all().
func (v *TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(v.native())
}
