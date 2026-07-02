package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Grid is a representation of GTK's GtkGrid.
type Grid struct {
	Container
}

func (v *Grid) native() *C.GtkGrid {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkGrid(unsafe.Pointer(v.GObject))
}

func wrapGrid(obj *Object) *Grid {
	return &Grid{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewGrid is a wrapper around gtk_grid_new().
func NewGrid() *Grid {
	c := C.gtk_grid_new()
	if c == nil {
		return nil
	}
	return wrapGrid(ToGoObject(unsafe.Pointer(c)))
}

// Attach is a wrapper around gtk_grid_attach().
func (v *Grid) Attach(child IWidget, left, top, width, height int) {
	C.gtk_grid_attach(v.native(), GtkWidget(child), C.gint(left), C.gint(top), C.gint(width), C.gint(height))
}

// AttachNextTo is a wrapper around gtk_grid_attach_next_to().
func (v *Grid) AttachNextTo(child, sibling IWidget, side PositionType, width, height int) {
	C.gtk_grid_attach_next_to(v.native(), GtkWidget(child), GtkWidget(sibling), C.GtkPositionType(side), C.gint(width), C.gint(height))
}

// SetRowSpacing is a wrapper around gtk_grid_set_row_spacing().
func (v *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing is a wrapper around gtk_grid_get_row_spacing().
func (v *Grid) GetRowSpacing() uint {
	return uint(C.gtk_grid_get_row_spacing(v.native()))
}

// SetColumnSpacing is a wrapper around gtk_grid_set_column_spacing().
func (v *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing is a wrapper around gtk_grid_get_column_spacing().
func (v *Grid) GetColumnSpacing() uint {
	return uint(C.gtk_grid_get_column_spacing(v.native()))
}

// SetRowHomogeneous is a wrapper around gtk_grid_set_row_homogeneous().
func (v *Grid) SetRowHomogeneous(homogeneous bool) {
	C.gtk_grid_set_row_homogeneous(v.native(), CBool(homogeneous))
}

// GetRowHomogeneous is a wrapper around gtk_grid_get_row_homogeneous().
func (v *Grid) GetRowHomogeneous() bool {
	return GoBool(C.gtk_grid_get_row_homogeneous(v.native()))
}

// SetColumnHomogeneous is a wrapper around gtk_grid_set_column_homogeneous().
func (v *Grid) SetColumnHomogeneous(homogeneous bool) {
	C.gtk_grid_set_column_homogeneous(v.native(), CBool(homogeneous))
}

// GetColumnHomogeneous is a wrapper around gtk_grid_get_column_homogeneous().
func (v *Grid) GetColumnHomogeneous() bool {
	return GoBool(C.gtk_grid_get_column_homogeneous(v.native()))
}
