package cgo

// #include <gtk/gtk.h>
// #include "gtk_header_bar.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Revealer struct {
	Bin
}

func (v *Revealer) native() *C.GtkRevealer {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkRevealer(unsafe.Pointer(v.GObject))
}

func wrapRevealer(obj *Object) *Revealer {
	return &Revealer{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

func NewRevealer() *Revealer {
	c := C.gtk_revealer_new()
	if c == nil {
		return nil
	}
	return wrapRevealer(ToGoObject(unsafe.Pointer(c)))
}

func (v *Revealer) SetRevealChild(reveal bool) {
	C.gtk_revealer_set_reveal_child(v.native(), CBool(reveal))
}

func (v *Revealer) GetRevealChild() bool {
	return GoBool(C.gtk_revealer_get_reveal_child(v.native()))
}

func (v *Revealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(v.native(), C.guint(duration))
}

func (v *Revealer) GetTransitionDuration() uint {
	return uint(C.gtk_revealer_get_transition_duration(v.native()))
}

func (v *Revealer) SetTransitionType(t RevealerTransitionType) {
	C.gtk_revealer_set_transition_type(v.native(), C.GtkRevealerTransitionType(t))
}

func (v *Revealer) GetTransitionType() RevealerTransitionType {
	return RevealerTransitionType(C.gtk_revealer_get_transition_type(v.native()))
}
