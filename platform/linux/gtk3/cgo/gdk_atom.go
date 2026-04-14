package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Atom is a representation of GDK's GdkAtom.
type Atom TAtom

func AsAtom(v unsafe.Pointer) IAtom {
	return Atom(v)
}

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

func (v Atom) Name() string {
	c := C.gdk_atom_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

func (m Atom) Atom() TAtom {
	return TAtom(m)
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) IAtom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), CBool(onlyIfExists))
	return Atom(unsafe.Pointer(c))
}
