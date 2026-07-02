package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextBuffer is a representation of GTK's GtkTextBuffer.
type TextBuffer struct {
	*Object
}

func (v *TextBuffer) native() *C.GtkTextBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTextBuffer(unsafe.Pointer(v.GObject))
}

func wrapTextBuffer(obj *Object) *TextBuffer {
	return &TextBuffer{obj}
}

// NewTextBuffer is a wrapper around gtk_text_buffer_new().
func NewTextBuffer() *TextBuffer {
	c := C.gtk_text_buffer_new(nil)
	if c == nil {
		return nil
	}
	return wrapTextBuffer(ToGoObject(unsafe.Pointer(c)))
}

// SetText is a wrapper around gtk_text_buffer_set_text().
func (v *TextBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// GetText is a wrapper around gtk_text_buffer_get_text().
func (v *TextBuffer) GetText(start, end ITextIter, includeHiddenChars bool) string {
	s := start.(*TextIter)
	e := end.(*TextIter)
	c := C.gtk_text_buffer_get_text(v.native(), &s.nativeIter, &e.nativeIter, CBool(includeHiddenChars))
	return C.GoString((*C.char)(c))
}

// GetBounds is a wrapper around gtk_text_buffer_get_bounds().
func (v *TextBuffer) GetBounds() (start, end ITextIter) {
	s := &TextIter{}
	e := &TextIter{}
	C.gtk_text_buffer_get_bounds(v.native(), &s.nativeIter, &e.nativeIter)
	return s, e
}

// GetCharCount is a wrapper around gtk_text_buffer_get_char_count().
func (v *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.native()))
}

// GetLineCount is a wrapper around gtk_text_buffer_get_line_count().
func (v *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.native()))
}

// Insert is a wrapper around gtk_text_buffer_insert().
func (v *TextBuffer) Insert(iter ITextIter, text string) {
	i := iter.(*TextIter)
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert(v.native(), &i.nativeIter, (*C.gchar)(cstr), C.gint(len(text)))
}

// InsertAtCursor is a wrapper around gtk_text_buffer_insert_at_cursor().
func (v *TextBuffer) InsertAtCursor(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_at_cursor(v.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// Delete is a wrapper around gtk_text_buffer_delete().
func (v *TextBuffer) Delete(start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	C.gtk_text_buffer_delete(v.native(), &s.nativeIter, &e.nativeIter)
}

// GetStartIter is a wrapper around gtk_text_buffer_get_start_iter().
func (v *TextBuffer) GetStartIter() ITextIter {
	iter := &TextIter{}
	C.gtk_text_buffer_get_start_iter(v.native(), &iter.nativeIter)
	return iter
}

// GetEndIter is a wrapper around gtk_text_buffer_get_end_iter().
func (v *TextBuffer) GetEndIter() ITextIter {
	iter := &TextIter{}
	C.gtk_text_buffer_get_end_iter(v.native(), &iter.nativeIter)
	return iter
}

// GetIterAtOffset is a wrapper around gtk_text_buffer_get_iter_at_offset().
func (v *TextBuffer) GetIterAtOffset(charOffset int) ITextIter {
	iter := &TextIter{}
	C.gtk_text_buffer_get_iter_at_offset(v.native(), &iter.nativeIter, C.gint(charOffset))
	return iter
}

// GetIterAtLine is a wrapper around gtk_text_buffer_get_iter_at_line().
func (v *TextBuffer) GetIterAtLine(lineNumber int) ITextIter {
	iter := &TextIter{}
	C.gtk_text_buffer_get_iter_at_line(v.native(), &iter.nativeIter, C.gint(lineNumber))
	return iter
}

// GetModified is a wrapper around gtk_text_buffer_get_modified().
func (v *TextBuffer) GetModified() bool {
	return GoBool(C.gtk_text_buffer_get_modified(v.native()))
}

// SetModified is a wrapper around gtk_text_buffer_set_modified().
func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.native(), CBool(setting))
}

// PlaceCursor is a wrapper around gtk_text_buffer_place_cursor().
func (v *TextBuffer) PlaceCursor(iter ITextIter) {
	i := iter.(*TextIter)
	C.gtk_text_buffer_place_cursor(v.native(), &i.nativeIter)
}

// GetSelectionBounds is a wrapper around gtk_text_buffer_get_selection_bounds().
func (v *TextBuffer) GetSelectionBounds() (start, end ITextIter, ok bool) {
	s := &TextIter{}
	e := &TextIter{}
	c := C.gtk_text_buffer_get_selection_bounds(v.native(), &s.nativeIter, &e.nativeIter)
	return s, e, GoBool(c)
}

// DeleteSelection is a wrapper around gtk_text_buffer_delete_selection().
func (v *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	return GoBool(C.gtk_text_buffer_delete_selection(v.native(), CBool(interactive), CBool(defaultEditable)))
}

// GetInsert is a wrapper around gtk_text_buffer_get_insert().
func (v *TextBuffer) GetInsert() ITextMark {
	c := C.gtk_text_buffer_get_insert(v.native())
	if c == nil {
		return nil
	}
	return &TextMark{ToGoObject(unsafe.Pointer(c))}
}

// CreateMark is a wrapper around gtk_text_buffer_create_mark().
func (v *TextBuffer) CreateMark(markName string, where ITextIter, leftGravity bool) ITextMark {
	cstr := C.CString(markName)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkTextIter
	if where != nil {
		w = &(where.(*TextIter)).nativeIter
	}
	c := C.gtk_text_buffer_create_mark(v.native(), (*C.gchar)(cstr), w, CBool(leftGravity))
	if c == nil {
		return nil
	}
	return &TextMark{ToGoObject(unsafe.Pointer(c))}
}

// GetMark is a wrapper around gtk_text_buffer_get_mark().
func (v *TextBuffer) GetMark(name string) ITextMark {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_text_buffer_get_mark(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return &TextMark{ToGoObject(unsafe.Pointer(c))}
}

// DeleteMark is a wrapper around gtk_text_buffer_delete_mark().
func (v *TextBuffer) DeleteMark(mark ITextMark) {
	m := mark.(*TextMark)
	C.gtk_text_buffer_delete_mark(v.native(), (*C.GtkTextMark)(unsafe.Pointer(m.GObject)))
}

// GetIterAtMark is a wrapper around gtk_text_buffer_get_iter_at_mark().
func (v *TextBuffer) GetIterAtMark(mark ITextMark) ITextIter {
	m := mark.(*TextMark)
	iter := &TextIter{}
	C.gtk_text_buffer_get_iter_at_mark(v.native(), &iter.nativeIter, (*C.GtkTextMark)(unsafe.Pointer(m.GObject)))
	return iter
}

// ApplyTagByName is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (v *TextBuffer) ApplyTagByName(name string, start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_apply_tag_by_name(v.native(), (*C.gchar)(cstr), &s.nativeIter, &e.nativeIter)
}

// RemoveTagByName is a wrapper around gtk_text_buffer_remove_tag_by_name().
func (v *TextBuffer) RemoveTagByName(name string, start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_remove_tag_by_name(v.native(), (*C.gchar)(cstr), &s.nativeIter, &e.nativeIter)
}
