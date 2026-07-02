package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextMark is a representation of GTK's GtkTextMark.
type TextMark struct {
	*Object
}

func wrapTextMark(obj *Object) *TextMark {
	return &TextMark{obj}
}

// SetVisible is a wrapper around gtk_text_mark_set_visible().
func (v *TextMark) SetVisible(setting bool) {
	C.gtk_text_mark_set_visible((*C.GtkTextMark)(unsafe.Pointer(v.GObject)), CBool(setting))
}

// GetVisible is a wrapper around gtk_text_mark_get_visible().
func (v *TextMark) GetVisible() bool {
	return GoBool(C.gtk_text_mark_get_visible((*C.GtkTextMark)(unsafe.Pointer(v.GObject))))
}

// GetDeleted is a wrapper around gtk_text_mark_get_deleted().
func (v *TextMark) GetDeleted() bool {
	return GoBool(C.gtk_text_mark_get_deleted((*C.GtkTextMark)(unsafe.Pointer(v.GObject))))
}

// GetName is a wrapper around gtk_text_mark_get_name().
func (v *TextMark) GetName() string {
	c := C.gtk_text_mark_get_name((*C.GtkTextMark)(unsafe.Pointer(v.GObject)))
	return C.GoString((*C.char)(c))
}

// GetBuffer is a wrapper around gtk_text_mark_get_buffer().
func (v *TextMark) GetBuffer() ITextBuffer {
	buf := C.gtk_text_mark_get_buffer((*C.GtkTextMark)(unsafe.Pointer(v.GObject)))
	if buf == nil {
		return nil
	}
	return wrapTextBuffer(ToGoObject(unsafe.Pointer(buf)))
}
