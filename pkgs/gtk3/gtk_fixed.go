package gtk3

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "gtk_fixed.go.h"
import "C"
import (
	"unsafe"
)

// Fixed is a representation of GTK's GtkFixed.
type Fixed struct {
	Container
}

func ToFixed(p unsafe.Pointer) *Fixed {
	return &Fixed{Container{Widget{InitiallyUnowned{ToGoObject(p)}}}}
}

func (v *Fixed) native() *C.GtkFixed {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFixed(p)
}

func wrapFixed(obj *Object) *Fixed {
	if obj == nil {
		return nil
	}
	return &Fixed{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewFixed is a wrapper around gtk_fixed_new().
func NewFixed() (*Fixed, error) {
	c := C.gtk_fixed_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapFixed(obj), nil
}

// Put is a wrapper around gtk_fixed_put().
func (v *Fixed) Put(w IWidget, x, y int) {
	C.gtk_fixed_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Move is a wrapper around gtk_fixed_move().
func (v *Fixed) Move(w IWidget, x, y int) {
	C.gtk_fixed_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}
