package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

// native() returns a pointer to the underlying GtkBox.
func (v *Box) native() *C.GtkBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkBox(p)
}

func marshalBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

func wrapBox(obj *Object) *Box {
	if obj == nil {
		return nil
	}

	return &Box{Container{Widget{InitiallyUnowned{obj}}}}
}

func (v *Box) toOrientable() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return C.toGtkOrientable(unsafe.Pointer(v.GObject))
}

// GetOrientation is a wrapper around C.gtk_orientable_get_orientation() for a GtkBox
func (v *Box) GetOrientation() Orientation {
	return Orientation(C.gtk_orientable_get_orientation(v.toOrientable()))
}

// SetOrientation is a wrapper around C.gtk_orientable_set_orientation() for a GtkBox
func (v *Box) SetOrientation(o Orientation) {
	C.gtk_orientable_set_orientation(v.toOrientable(), C.GtkOrientation(o))
}

// NewBox is a wrapper around gtk_box_new().
func NewBox(orientation Orientation, spacing int) *Box {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapBox(obj)
}

// PackStart is a wrapper around gtk_box_pack_start().
func (v *Box) PackStart(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_start(v.native(), child.toWidget(), CBool(expand),
		CBool(fill), C.guint(padding))
}

// PackEnd is a wrapper around gtk_box_pack_end().
func (v *Box) PackEnd(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_end(v.native(), child.toWidget(), CBool(expand),
		CBool(fill), C.guint(padding))
}

// GetHomogeneous is a wrapper around gtk_box_get_homogeneous().
func (v *Box) GetHomogeneous() bool {
	c := C.gtk_box_get_homogeneous(v.native())
	return GoBool(c)
}

// SetHomogeneous is a wrapper around gtk_box_set_homogeneous().
func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(v.native(), CBool(homogeneous))
}

// GetSpacing is a wrapper around gtk_box_get_spacing().
func (v *Box) GetSpacing() int {
	c := C.gtk_box_get_spacing(v.native())
	return int(c)
}

// SetSpacing is a wrapper around gtk_box_set_spacing()
func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(v.native(), C.gint(spacing))
}

// ReorderChild is a wrapper around gtk_box_reorder_child().
func (v *Box) ReorderChild(child IWidget, position int) {
	C.gtk_box_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// QueryChildPacking is a wrapper around gtk_box_query_child_packing().
func (v *Box) QueryChildPacking(child IWidget) (expand, fill bool, padding uint, packType PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpackType C.GtkPackType

	C.gtk_box_query_child_packing(v.native(), child.toWidget(), &cexpand,
		&cfill, &cpadding, &cpackType)
	return GoBool(cexpand), GoBool(cfill), uint(cpadding), PackType(cpackType)
}

// SetChildPacking is a wrapper around gtk_box_set_child_packing().
func (v *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, packType PackType) {
	C.gtk_box_set_child_packing(v.native(), child.toWidget(), CBool(expand),
		CBool(fill), C.guint(padding), C.GtkPackType(packType))
}

// Orientable is a representation of GTK's GtkOrientable GInterface.
type Orientable struct {
	*Object
}

// IOrientable is an interface type implemented by all structs
// embedding an Orientable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkOrientable.
type IOrientable interface {
	toOrientable() *C.GtkOrientable
}

// native returns a pointer to the underlying GObject as a GtkOrientable.
func (v *Orientable) native() *C.GtkOrientable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkOrientable(p)
}

func wrapOrientable(obj *Object) *Orientable {
	if obj == nil {
		return nil
	}
	return &Orientable{obj}
}

// GetOrientation is a wrapper around gtk_orientable_get_orientation().
func (v *Orientable) GetOrientation() Orientation {
	c := C.gtk_orientable_get_orientation(v.native())
	return Orientation(c)
}

// SetOrientation is a wrapper around gtk_orientable_set_orientation().
func (v *Orientable) SetOrientation(orientation Orientation) {
	C.gtk_orientable_set_orientation(v.native(),
		C.GtkOrientation(orientation))
}
