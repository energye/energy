package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type ComboBoxText struct {
	Widget
}

func (v *ComboBoxText) native() *C.GtkComboBoxText {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkComboBoxText(unsafe.Pointer(v.GObject))
}

func wrapComboBoxText(obj *Object) *ComboBoxText {
	return &ComboBoxText{Widget{InitiallyUnowned{obj}}}
}

// NewComboBoxText is a wrapper around gtk_combo_box_text_new().
func NewComboBoxText() *ComboBoxText {
	c := C.gtk_combo_box_text_new()
	if c == nil {
		return nil
	}
	return wrapComboBoxText(ToGoObject(unsafe.Pointer(c)))
}

// NewComboBoxTextWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func NewComboBoxTextWithEntry() *ComboBoxText {
	c := C.gtk_combo_box_text_new_with_entry()
	if c == nil {
		return nil
	}
	return wrapComboBoxText(ToGoObject(unsafe.Pointer(c)))
}

// Append is a wrapper around gtk_combo_box_text_append().
func (v *ComboBoxText) Append(id string, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_append(v.native(), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (v *ComboBoxText) AppendText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_append_text(v.native(), (*C.gchar)(cstr))
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (v *ComboBoxText) GetActiveText() string {
	c := C.gtk_combo_box_text_get_active_text(v.native())
	if c == nil {
		return ""
	}
	s := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return s
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (v *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(v.native())
}

// GetActive is a wrapper around gtk_combo_box_get_active().
func (v *ComboBoxText) GetActive() int {
	return int(C.gtk_combo_box_get_active(C.toGtkComboBox(unsafe.Pointer(v.GObject))))
}

// SetActive is a wrapper around gtk_combo_box_set_active().
func (v *ComboBoxText) SetActive(index int) {
	C.gtk_combo_box_set_active(C.toGtkComboBox(unsafe.Pointer(v.GObject)), C.gint(index))
}
