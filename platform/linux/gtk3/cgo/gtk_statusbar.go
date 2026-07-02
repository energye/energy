package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// Statusbar is a representation of GTK's GtkStatusbar.
type Statusbar struct {
	Box
}

func (v *Statusbar) native() *C.GtkStatusbar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusbar(p)
}

func wrapStatusbar(obj *Object) *Statusbar {
	return &Statusbar{Box{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewStatusbar is a wrapper around gtk_statusbar_new().
func NewStatusbar() *Statusbar {
	c := C.gtk_statusbar_new()
	if c == nil {
		return nil
	}
	return wrapStatusbar(ToGoObject(unsafe.Pointer(c)))
}

// GetContextId is a wrapper around gtk_statusbar_get_context_id().
func (v *Statusbar) GetContextId(contextDescription string) uint {
	cstr := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(cstr))
	return uint(C.gtk_statusbar_get_context_id(v.native(), (*C.gchar)(cstr)))
}

// Push is a wrapper around gtk_statusbar_push().
func (v *Statusbar) Push(contextId uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	return uint(C.gtk_statusbar_push(v.native(), C.guint(contextId), (*C.gchar)(cstr)))
}

// Pop is a wrapper around gtk_statusbar_pop().
func (v *Statusbar) Pop(contextId uint) {
	C.gtk_statusbar_pop(v.native(), C.guint(contextId))
}

// RemoveAll is a wrapper around gtk_statusbar_remove_all().
func (v *Statusbar) RemoveAll(contextId uint) {
	C.gtk_statusbar_remove_all(v.native(), C.guint(contextId))
}
