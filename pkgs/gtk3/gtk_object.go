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
	"reflect"
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

func toGoStringArray(c **C.gchar) []string {
	if c == nil {
		return nil
	}
	defer C.g_strfreev(c)
	strsLen := 0
	scanPtr := c
	for ; *scanPtr != nil; scanPtr = nextGCharPtr(scanPtr) {
		strsLen++
	}
	strs := make([]string, strsLen)
	currentPtr := c
	for i := range strs {
		strs[i] = C.GoString((*C.char)(*currentPtr))
		currentPtr = nextGCharPtr(currentPtr)
	}
	return strs
}

func nextGCharPtr(current **C.gchar) **C.gchar {
	ptrSize := unsafe.Sizeof(*current)
	nextAddr := uintptr(unsafe.Pointer(current)) + ptrSize
	return (**C.gchar)(unsafe.Pointer(nextAddr))
}

func ucharString(guchar *C.guchar) string {
	// Seek and find the string length.
	var strlen int
	for ptr := guchar; *ptr != 0; ptr = nextguchar(ptr) {
		strlen++
	}

	// Array of unsigned char means GoString is unavailable, so maybe this is
	// fine.
	var data []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Len = strlen
	sliceHeader.Cap = strlen
	sliceHeader.Data = uintptr(unsafe.Pointer(guchar))

	// Return a copy of the string.
	return string(data)
}
func nextguchar(guchar *C.guchar) *C.guchar {
	return (*C.guchar)(unsafe.Pointer(uintptr(unsafe.Pointer(guchar)) + 1))
}

func ToCObject(p unsafe.Pointer) *C.GObject {
	return (*C.GObject)(p)
}
func ToGoObject(instance unsafe.Pointer) *Object {
	cObj := ToCObject(instance)
	return &Object{GObject: cObj}
}

func ToCInt(v int) C.int {
	return C.int(v)
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
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject), (*C.gchar)(cstr))
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return GoBool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

func (v *Object) SetData(key string, value unsafe.Pointer) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cData := C.gpointer(uintptr(value))
	C.g_object_set_data(v.native(), cKey, cData)
}

func (v *Object) GetData(key string) unsafe.Pointer {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cData := C.g_object_get_data(v.native(), cKey)
	return unsafe.Pointer(cData)
}
