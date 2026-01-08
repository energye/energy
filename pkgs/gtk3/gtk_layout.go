package gtk3

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <gio/gio.h>
#include <gtk/gtk.h>
#include "gtk.go.h"
*/
import "C"
import (
	"unsafe"
)

// Layout is a representation of GTK's GtkLayout.
type Layout struct {
	Container
}

func ToLayout(p unsafe.Pointer) *Layout {
	return &Layout{Container{Widget{InitiallyUnowned{ToGoObject(p)}}}}
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *Layout) native() *C.GtkLayout {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLayout(p)
}

func wrapLayout(obj *Object) *Layout {
	if obj == nil {
		return nil
	}

	return &Layout{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewLayout is a wrapper around gtk_layout_new().
func NewLayout(hadjustment, vadjustment *Adjustment) (*Layout, error) {
	c := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

// Layout.Put is a wrapper around gtk_layout_put().
func (v *Layout) Put(w IWidget, x, y int) {
	C.gtk_layout_put(v.native(), GtkWidget(w), C.gint(x), C.gint(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (v *Layout) Move(w IWidget, x, y int) {
	C.gtk_layout_move(v.native(), GtkWidget(w), C.gint(x), C.gint(y))
}

// Layout.SetSize is a wrapper around gtk_layout_set_size
func (v *Layout) SetSize(width, height uint) {
	C.gtk_layout_set_size(v.native(), C.guint(width), C.guint(height))
}

// Layout.GetSize is a wrapper around gtk_layout_get_size
func (v *Layout) GetSize() (width, height uint) {
	var w, h C.guint
	C.gtk_layout_get_size(v.native(), &w, &h)
	return uint(w), uint(h)
}
