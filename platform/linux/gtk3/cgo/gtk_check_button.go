package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// CheckButton is a representation of GTK's GtkCheckButton.
type CheckButton struct {
	Button
}

func (v *CheckButton) native() *C.GtkCheckButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkCheckButton(unsafe.Pointer(v.GObject))
}

func wrapCheckButton(obj *Object) *CheckButton {
	return &CheckButton{Button{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}
}

// NewCheckButton is a wrapper around gtk_check_button_new().
func NewCheckButton() *CheckButton {
	c := C.gtk_check_button_new()
	if c == nil {
		return nil
	}
	return wrapCheckButton(ToGoObject(unsafe.Pointer(c)))
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *CheckButton) GetActive() bool {
	return GoBool(C.gtk_toggle_button_get_active(C.toGtkToggleButton(unsafe.Pointer(v.GObject))))
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *CheckButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(C.toGtkToggleButton(unsafe.Pointer(v.GObject)), CBool(isActive))
}
