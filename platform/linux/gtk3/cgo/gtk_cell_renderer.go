package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// cellRenderer is a cgo-internal interface for getting the C GtkCellRenderer pointer.
type cellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	*Object
}

func (v *CellRendererText) toCellRenderer() *C.GtkCellRenderer {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkCellRenderer(unsafe.Pointer(v.GObject))
}

func wrapCellRendererText(obj *Object) *CellRendererText {
	return &CellRendererText{obj}
}

// NewCellRendererText is a wrapper around gtk_cell_renderer_text_new().
func NewCellRendererText() *CellRendererText {
	c := C.gtk_cell_renderer_text_new()
	if c == nil {
		return nil
	}
	return wrapCellRendererText(ToGoObject(unsafe.Pointer(c)))
}
