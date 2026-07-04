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

// NewRadioButton is a wrapper around gtk_radio_button_new(NULL) — creates the first button in a new group.
func NewRadioButton() *RadioButton {
	c := C.gtk_radio_button_new(nil)
	if c == nil {
		return nil
	}
	return wrapRadioButton(ToGoObject(unsafe.Pointer(c)))
}

// NewRadioButtonFromWidget is a wrapper around gtk_radio_button_new_from_widget().
func NewRadioButtonFromWidget(radioGroupMember *RadioButton) *RadioButton {
	var member *C.GtkRadioButton
	if radioGroupMember != nil {
		member = radioGroupMember.native()
	}
	c := C.gtk_radio_button_new_from_widget(member)
	if c == nil {
		return nil
	}
	return wrapRadioButton(ToGoObject(unsafe.Pointer(c)))
}

// NewRadioButtonWithLabel is a wrapper around gtk_radio_button_new_with_label(NULL, label).
func NewRadioButtonWithLabel(label string) *RadioButton {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label(nil, (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapRadioButton(ToGoObject(unsafe.Pointer(c)))
}
