package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// RadioButton is a representation of GTK's GtkRadioButton.
type RadioButton struct {
	CheckButton
}

func (v *RadioButton) native() *C.GtkRadioButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkRadioButton(unsafe.Pointer(v.GObject))
}

func wrapRadioButton(obj *Object) *RadioButton {
	return &RadioButton{CheckButton{Button{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewRadioButtonWithLabelFromWidget is a wrapper around gtk_radio_button_new_with_label_from_widget().
func NewRadioButtonWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	var member *C.GtkRadioButton
	if radioGroupMember != nil {
		member = radioGroupMember.native()
	}
	c := C.gtk_radio_button_new_with_label_from_widget(member, (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapRadioButton(ToGoObject(unsafe.Pointer(c)))
}
