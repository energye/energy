package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextTag is a representation of GTK's GtkTextTag.
type TextTag struct {
	*Object
}

func wrapTextTag(obj *Object) *TextTag {
	return &TextTag{obj}
}

// GetPriority is a wrapper around gtk_text_tag_get_priority().
func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority((*C.GtkTextTag)(unsafe.Pointer(v.GObject))))
}

// SetPriority is a wrapper around gtk_text_tag_set_priority().
func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority((*C.GtkTextTag)(unsafe.Pointer(v.GObject)), C.gint(priority))
}

// TextTagTable is a representation of GTK's GtkTextTagTable.
type TextTagTable struct {
	*Object
}

func wrapTextTagTable(obj *Object) *TextTagTable {
	return &TextTagTable{obj}
}

// NewTextTagTable is a wrapper around gtk_text_tag_table_new().
func NewTextTagTable() *TextTagTable {
	c := C.gtk_text_tag_table_new()
	if c == nil {
		return nil
	}
	return wrapTextTagTable(ToGoObject(unsafe.Pointer(c)))
}

// Add is a wrapper around gtk_text_tag_table_add().
func (v *TextTagTable) Add(tag ITextTag) bool {
	t := tag.(*TextTag)
	return GoBool(C.gtk_text_tag_table_add(
		(*C.GtkTextTagTable)(unsafe.Pointer(v.GObject)),
		(*C.GtkTextTag)(unsafe.Pointer(t.GObject))))
}

// Lookup is a wrapper around gtk_text_tag_table_lookup().
func (v *TextTagTable) Lookup(name string) ITextTag {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_text_tag_table_lookup((*C.GtkTextTagTable)(unsafe.Pointer(v.GObject)), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapTextTag(ToGoObject(unsafe.Pointer(c)))
}

// Remove is a wrapper around gtk_text_tag_table_remove().
func (v *TextTagTable) Remove(tag ITextTag) {
	t := tag.(*TextTag)
	C.gtk_text_tag_table_remove(
		(*C.GtkTextTagTable)(unsafe.Pointer(v.GObject)),
		(*C.GtkTextTag)(unsafe.Pointer(t.GObject)))
}
