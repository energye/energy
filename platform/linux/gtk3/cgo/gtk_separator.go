package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Separator is a representation of GTK's GtkSeparator.
type Separator struct {
	Widget
}

func (v *Separator) native() *C.GtkSeparator {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkSeparator(unsafe.Pointer(v.GObject))
}

func wrapSeparator(obj *Object) *Separator {
	return &Separator{Widget{InitiallyUnowned{obj}}}
}

// NewSeparator is a wrapper around gtk_separator_new().
func NewSeparator(orientation Orientation) *Separator {
	c := C.gtk_separator_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil
	}
	return wrapSeparator(ToGoObject(unsafe.Pointer(c)))
}
