package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// EntryCompletion is a representation of GTK's GtkEntryCompletion.
type EntryCompletion struct {
	*Object
}

// native returns a pointer to the underlying GtkEntryCompletion.
func (v *EntryCompletion) native() *C.GtkEntryCompletion {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryCompletion(p)
}

func wrapEntryCompletion(obj *Object) *EntryCompletion {
	if obj == nil {
		return nil
	}

	return &EntryCompletion{obj}
}

// NewEntryCompletion is a wrapper around gtk_entry_completion_new
func NewEntryCompletion() *EntryCompletion {
	c := C.gtk_entry_completion_new()
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapEntryCompletion(obj)
}

// SetMinimumKeyLength is a wrapper around gtk_entry_completion_set_minimum_key_length
func (v *EntryCompletion) SetMinimumKeyLength(minimumLength int) {
	C.gtk_entry_completion_set_minimum_key_length(v.native(), C.gint(minimumLength))
}

// GetMinimumKeyLength is a wrapper around gtk_entry_completion_get_minimum_key_length
func (v *EntryCompletion) GetMinimumKeyLength() int {
	c := C.gtk_entry_completion_get_minimum_key_length(v.native())
	return int(c)
}

// SetTextColumn is a wrapper around gtk_entry_completion_set_text_column
func (v *EntryCompletion) SetTextColumn(textColumn int) {
	C.gtk_entry_completion_set_text_column(v.native(), C.gint(textColumn))
}

// GetTextColumn is a wrapper around gtk_entry_completion_get_text_column
func (v *EntryCompletion) GetTextColumn() int {
	c := C.gtk_entry_completion_get_text_column(v.native())
	return int(c)
}

// SetInlineCompletion is a wrapper around gtk_entry_completion_set_inline_completion
func (v *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	C.gtk_entry_completion_set_inline_completion(v.native(), CBool(inlineCompletion))
}

// GetInlineCompletion is a wrapper around gtk_entry_completion_get_inline_completion
func (v *EntryCompletion) GetInlineCompletion() bool {
	c := C.gtk_entry_completion_get_inline_completion(v.native())
	return GoBool(c)
}

// SetPopupCompletion is a wrapper around gtk_entry_completion_set_popup_completion
func (v *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	C.gtk_entry_completion_set_popup_completion(v.native(), CBool(popupCompletion))
}

// GetPopupCompletion is a wrapper around gtk_entry_completion_get_popup_completion
func (v *EntryCompletion) GetPopupCompletion() bool {
	c := C.gtk_entry_completion_get_popup_completion(v.native())
	return GoBool(c)
}

// SetPopupSetWidth is a wrapper around gtk_entry_completion_set_popup_set_width
func (v *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	C.gtk_entry_completion_set_popup_set_width(v.native(), CBool(popupSetWidth))
}

// GetPopupSetWidth is a wrapper around gtk_entry_completion_get_popup_set_width
func (v *EntryCompletion) GetPopupSetWidth() bool {
	c := C.gtk_entry_completion_get_popup_set_width(v.native())
	return GoBool(c)
}
