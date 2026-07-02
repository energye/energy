package cgo

/*
#cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
#include <gio/gio.h>
#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>
#include "glib.go.h"
*/
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.native())))
	return Type(c)
}
