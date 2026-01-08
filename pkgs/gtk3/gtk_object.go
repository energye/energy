package gtk3

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <gtk/gtk.h>
#include <gio/gio.h>
#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>
#include "gtk.go.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

// Container is a representation of GTK's GtkContainer.
type Container struct {
	Widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *Container) native() *C.GtkContainer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkContainer(p)
}

func marshalContainer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapContainer(obj), nil
}

func wrapContainer(obj *Object) *Container {
	if obj == nil {
		return nil
	}

	return &Container{Widget{InitiallyUnowned{obj}}}
}

// Add is a wrapper around gtk_container_add().
func (v *Container) Add(w IWidget) {
	C.gtk_container_add(v.native(), w.toWidget())
}

// Remove is a wrapper around gtk_container_remove().
func (v *Container) Remove(w IWidget) {
	C.gtk_container_remove(v.native(), w.toWidget())
}

// CheckResize is a wrapper around gtk_container_check_resize().
func (v *Container) CheckResize() {
	C.gtk_container_check_resize(v.native())
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (v *Container) SetFocusChild(child IWidget) {
	C.gtk_container_set_focus_child(v.native(), child.toWidget())
}

// GetFocusVAdjustment is a wrapper around gtk_container_get_focus_vadjustment().
func (v *Container) GetFocusVAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusVAdjustment is a wrapper around gtk_container_set_focus_vadjustment().
func (v *Container) SetFocusVAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(v.native(), adjustment.native())
}

// GetFocusHAdjustment is a wrapper around gtk_container_get_focus_hadjustment().
func (v *Container) GetFocusHAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusHAdjustment is a wrapper around gtk_container_set_focus_hadjustment().
func (v *Container) SetFocusHAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(v.native(), adjustment.native())
}

// ChildNotify is a wrapper around gtk_container_child_notify().
func (v *Container) ChildNotify(child IWidget, childProperty string) {
	cstr := C.CString(childProperty)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_container_child_notify(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (v *Container) GetBorderWidth() uint {
	c := C.gtk_container_get_border_width(v.native())
	return uint(c)
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (v *Container) SetBorderWidth(borderWidth uint) {
	C.gtk_container_set_border_width(v.native(), C.guint(borderWidth))
}

// Bin is a representation of GTK's GtkBin.
type Bin struct {
	Container
}

// InitiallyUnowned is a representation of GLib's GInitiallyUnowned.
type InitiallyUnowned struct {
	// This must be a pointer so copies of the ref-sinked object
	// do not outlive the original object, causing an unref
	// finalizer to prematurely run.
	*Object
}

// Object is a representation of GLib's GObject.
type Object struct {
	GObject *C.GObject
}

func CBool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func GoBool(b C.gboolean) bool {
	return b != C.FALSE
}

func ToCObject(p unsafe.Pointer) *C.GObject {
	return (*C.GObject)(p)
}
func ToGoObject(instance unsafe.Pointer) *Object {
	cObj := ToCObject(instance)
	return &Object{GObject: cObj}
}

func GoString(cStr *C.gchar) string {
	return C.GoString((*C.char)(cStr))
}

// native returns a pointer to the underlying GObject.
func (v *Object) native() *C.GObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGObject(p)
}

// Ref is a wrapper around g_object_ref().
func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject))
}

// Unref is a wrapper around g_object_unref().
func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

// RefSink is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.GObject))
}

// IsFloating is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return GoBool(c)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.GObject)
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return GoBool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}
