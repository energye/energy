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
