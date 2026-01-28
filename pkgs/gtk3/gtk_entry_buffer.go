package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// EntryBuffer is a representation of GTK's GtkEntryBuffer.
type EntryBuffer struct {
	*Object
}

// native returns a pointer to the underlying GtkEntryBuffer.
func (v *EntryBuffer) native() *C.GtkEntryBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntryBuffer(p)
}

func wrapEntryBuffer(obj *Object) *EntryBuffer {
	if obj == nil {
		return nil
	}

	return &EntryBuffer{obj}
}

// NewEntryBuffer is a wrapper around gtk_entry_buffer_new().
func NewEntryBuffer(initialChars string, nInitialChars int) *EntryBuffer {
	cstr := C.CString(initialChars)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_new((*C.gchar)(cstr), C.gint(nInitialChars))
	if c == nil {
		return nil
	}

	e := wrapEntryBuffer(ToGoObject(unsafe.Pointer(c)))
	return e
}

// GetText is a wrapper around gtk_entry_buffer_get_text().  A
// non-nil error is returned in the case that gtk_entry_buffer_get_text
// returns NULL to differentiate between NULL and an empty string.
func (v *EntryBuffer) GetText() (string, error) {
	c := C.gtk_entry_buffer_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return GoString(c), nil
}

// SetText is a wrapper around gtk_entry_buffer_set_text().
func (v *EntryBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// GetBytes is a wrapper around gtk_entry_buffer_get_bytes().
func (v *EntryBuffer) GetBytes() uint {
	c := C.gtk_entry_buffer_get_bytes(v.native())
	return uint(c)
}

// GetLength is a wrapper around gtk_entry_buffer_get_length().
func (v *EntryBuffer) GetLength() uint {
	c := C.gtk_entry_buffer_get_length(v.native())
	return uint(c)
}

// GetMaxLength is a wrapper around gtk_entry_buffer_get_max_length().
func (v *EntryBuffer) GetMaxLength() int {
	c := C.gtk_entry_buffer_get_max_length(v.native())
	return int(c)
}

// SetMaxLength is a wrapper around gtk_entry_buffer_set_max_length().
func (v *EntryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(v.native(), C.gint(maxLength))
}

// InsertText is a wrapper around gtk_entry_buffer_insert_text().
func (v *EntryBuffer) InsertText(position uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_insert_text(v.native(), C.guint(position), (*C.gchar)(cstr), C.gint(len(text)))
	return uint(c)
}

// DeleteText is a wrapper around gtk_entry_buffer_delete_text().
func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	c := C.gtk_entry_buffer_delete_text(v.native(), C.guint(position),
		C.gint(nChars))
	return uint(c)
}

// EmitDeletedText is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (v *EntryBuffer) EmitDeletedText(pos, nChars uint) {
	C.gtk_entry_buffer_emit_deleted_text(v.native(), C.guint(pos),
		C.guint(nChars))
}

// EmitInsertedText is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (v *EntryBuffer) EmitInsertedText(pos uint, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_emit_inserted_text(v.native(), C.guint(pos),
		(*C.gchar)(cstr), C.guint(len(text)))
}
