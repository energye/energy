package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"
import (
	. "github.com/energye/energy/v3/pkgs/gtk3/types"
)

// Overlay is a representation of GTK's GtkOverlay.
type Overlay struct {
	Bin
}

// native returns a pointer to the underlying GtkOverlay.
func (v *Overlay) native() *C.GtkOverlay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOverlay(p)
}

func wrapOverlay(obj *Object) *Overlay {
	if obj == nil {
		return nil
	}

	return &Overlay{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewOverlay is a wrapper around gtk_overlay_new().
func NewOverlay() *Overlay {
	c := C.gtk_overlay_new()
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapOverlay(obj)
}

// AddOverlay is a wrapper around gtk_overlay_add_overlay().
func (v *Overlay) AddOverlay(widget IWidget) {
	C.gtk_overlay_add_overlay(v.native(), widget.(_IWidget).toWidget())
}

// ReorderOverlay is a wrapper around gtk_overlay_reorder_overlay().
func (v *Overlay) ReorderOverlay(child IWidget, position int) {
	C.gtk_overlay_reorder_overlay(v.native(), child.(_IWidget).toWidget(), C.int(position))
}

// GetOverlayPassThrough is a wrapper around gtk_overlay_get_overlay_pass_through().
func (v *Overlay) GetOverlayPassThrough(widget IWidget) bool {
	c := C.gtk_overlay_get_overlay_pass_through(v.native(), widget.(_IWidget).toWidget())
	return GoBool(c)
}

// SetOverlayPassThrough is a wrapper around gtk_overlay_set_overlay_pass_through().
func (v *Overlay) SetOverlayPassThrough(widget IWidget, passThrough bool) {
	C.gtk_overlay_set_overlay_pass_through(v.native(), widget.(_IWidget).toWidget(), CBool(passThrough))
}
