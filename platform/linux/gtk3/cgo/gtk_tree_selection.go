package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
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
