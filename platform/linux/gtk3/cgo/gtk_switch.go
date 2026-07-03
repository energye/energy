package cgo

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Switch is a representation of GTK's GtkSwitch.
type Switch struct {
	Widget
}

func (v *Switch) native() *C.GtkSwitch {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkSwitch)(unsafe.Pointer(v.GObject))
}

func wrapSwitch(obj *Object) *Switch {
	return &Switch{Widget{InitiallyUnowned{obj}}}
}

func AsSwitch(ptr unsafe.Pointer) *Switch {
	return wrapSwitch(ToGoObject(ptr))
}

func NewSwitch() *Switch {
	c := C.gtk_switch_new()
	if c == nil {
		return nil
	}
	return wrapSwitch(ToGoObject(unsafe.Pointer(c)))
}

func (v *Switch) GetActive() bool {
	return GoBool(C.gtk_switch_get_active(v.native()))
}

func (v *Switch) SetActive(isActive bool) {
	C.gtk_switch_set_active(v.native(), CBool(isActive))
}

func (v *Switch) GetState() bool {
	return GoBool(C.gtk_switch_get_state(v.native()))
}

func (v *Switch) SetState(state bool) {
	C.gtk_switch_set_state(v.native(), CBool(state))
}
