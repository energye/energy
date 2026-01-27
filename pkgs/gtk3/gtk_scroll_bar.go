package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// Scrollbar is a representation of GTK's GtkScrollbar.
type Scrollbar struct {
	Range
}

// native returns a pointer to the underlying GtkScrollbar.
func (v *Scrollbar) native() *C.GtkScrollbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrollbar(p)
}

func marshalScrollbar(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapScrollbar(obj), nil
}

func wrapScrollbar(obj *Object) *Scrollbar {
	if obj == nil {
		return nil
	}
	return &Scrollbar{Range{Widget{InitiallyUnowned{obj}}}}
}

// NewScrollbar is a wrapper around gtk_scrollbar_new().
func NewScrollbar(orientation Orientation, adjustment *Adjustment) (*Scrollbar, error) {
	c := C.gtk_scrollbar_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapScrollbar(obj), nil
}
