package cgo

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

type Spinner struct {
	Widget
}

func (v *Spinner) native() *C.GtkSpinner {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkSpinner)(unsafe.Pointer(v.GObject))
}

func wrapSpinner(obj *Object) *Spinner {
	return &Spinner{Widget{InitiallyUnowned{obj}}}
}

func NewSpinner() *Spinner {
	c := C.gtk_spinner_new()
	if c == nil {
		return nil
	}
	return wrapSpinner(ToGoObject(unsafe.Pointer(c)))
}

func (v *Spinner) Start() {
	C.gtk_spinner_start(v.native())
}

func (v *Spinner) Stop() {
	C.gtk_spinner_stop(v.native())
}
