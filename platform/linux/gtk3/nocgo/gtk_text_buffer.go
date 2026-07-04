//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextBuffer is a representation of GTK's GtkTextBuffer.
type TextBuffer struct {
	Object
}

func AsTextBuffer(ptr unsafe.Pointer) *TextBuffer {
	if ptr == nil {
		return nil
	}
	m := new(TextBuffer)
	m.instance = ptr
	return m
}

// NewTextBuffer is a wrapper around gtk_text_buffer_new().
func NewTextBuffer() *TextBuffer {
	r := gtk3.SysCall("gtk_text_buffer_new", 0)
	if r == 0 {
		return nil
	}
	return AsTextBuffer(unsafe.Pointer(r))
}

// SetText is a wrapper around gtk_text_buffer_set_text().
func (m *TextBuffer) SetText(text string) {
	cstr := CStr(text)
	gtk3.SysCall("gtk_text_buffer_set_text", m.Instance(), cstr, uintptr(len(text)))
}

// GetText is a wrapper around gtk_text_buffer_get_text().
func (m *TextBuffer) GetText(start, end ITextIter, includeHiddenChars bool) string {
	s := start.(*TextIter)
	e := end.(*TextIter)
	r := gtk3.SysCall("gtk_text_buffer_get_text", m.Instance(),
		s.native(), e.native(), ToCBool(includeHiddenChars))
	return GoStr(r)
}

// GetBounds is a wrapper around gtk_text_buffer_get_bounds().
func (m *TextBuffer) GetBounds() (start, end ITextIter) {
	s := new(TextIter)
	e := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_bounds", m.Instance(), s.native(), e.native())
	return s, e
}

// GetCharCount is a wrapper around gtk_text_buffer_get_char_count().
func (m *TextBuffer) GetCharCount() int {
	r := gtk3.SysCall("gtk_text_buffer_get_char_count", m.Instance())
	return int(r)
}

// GetLineCount is a wrapper around gtk_text_buffer_get_line_count().
func (m *TextBuffer) GetLineCount() int {
	r := gtk3.SysCall("gtk_text_buffer_get_line_count", m.Instance())
	return int(r)
}

// Insert is a wrapper around gtk_text_buffer_insert().
func (m *TextBuffer) Insert(iter ITextIter, text string) {
	i := iter.(*TextIter)
	cstr := CStr(text)
	gtk3.SysCall("gtk_text_buffer_insert", m.Instance(), i.native(), cstr, uintptr(len(text)))
}

// InsertAtCursor is a wrapper around gtk_text_buffer_insert_at_cursor().
func (m *TextBuffer) InsertAtCursor(text string) {
	cstr := CStr(text)
	gtk3.SysCall("gtk_text_buffer_insert_at_cursor", m.Instance(), cstr, uintptr(len(text)))
}

// Delete is a wrapper around gtk_text_buffer_delete().
func (m *TextBuffer) Delete(start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	gtk3.SysCall("gtk_text_buffer_delete", m.Instance(), s.native(), e.native())
}

// GetStartIter is a wrapper around gtk_text_buffer_get_start_iter().
func (m *TextBuffer) GetStartIter() ITextIter {
	iter := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_start_iter", m.Instance(), iter.native())
	return iter
}

// GetEndIter is a wrapper around gtk_text_buffer_get_end_iter().
func (m *TextBuffer) GetEndIter() ITextIter {
	iter := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_end_iter", m.Instance(), iter.native())
	return iter
}

// GetIterAtOffset is a wrapper around gtk_text_buffer_get_iter_at_offset().
func (m *TextBuffer) GetIterAtOffset(charOffset int) ITextIter {
	iter := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_iter_at_offset", m.Instance(), iter.native(), uintptr(charOffset))
	return iter
}

// GetIterAtLine is a wrapper around gtk_text_buffer_get_iter_at_line().
func (m *TextBuffer) GetIterAtLine(lineNumber int) ITextIter {
	iter := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_iter_at_line", m.Instance(), iter.native(), uintptr(lineNumber))
	return iter
}

// GetModified is a wrapper around gtk_text_buffer_get_modified().
func (m *TextBuffer) GetModified() bool {
	r := gtk3.SysCall("gtk_text_buffer_get_modified", m.Instance())
	return ToGoBool(r)
}

// SetModified is a wrapper around gtk_text_buffer_set_modified().
func (m *TextBuffer) SetModified(setting bool) {
	gtk3.SysCall("gtk_text_buffer_set_modified", m.Instance(), ToCBool(setting))
}

// PlaceCursor is a wrapper around gtk_text_buffer_place_cursor().
func (m *TextBuffer) PlaceCursor(iter ITextIter) {
	i := iter.(*TextIter)
	gtk3.SysCall("gtk_text_buffer_place_cursor", m.Instance(), i.native())
}

// GetSelectionBounds is a wrapper around gtk_text_buffer_get_selection_bounds().
func (m *TextBuffer) GetSelectionBounds() (start, end ITextIter, ok bool) {
	s := new(TextIter)
	e := new(TextIter)
	r := gtk3.SysCall("gtk_text_buffer_get_selection_bounds", m.Instance(), s.native(), e.native())
	return s, e, ToGoBool(r)
}

// DeleteSelection is a wrapper around gtk_text_buffer_delete_selection().
func (m *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	r := gtk3.SysCall("gtk_text_buffer_delete_selection", m.Instance(),
		ToCBool(interactive), ToCBool(defaultEditable))
	return ToGoBool(r)
}

// GetInsert is a wrapper around gtk_text_buffer_get_insert().
func (m *TextBuffer) GetInsert() ITextMark {
	r := gtk3.SysCall("gtk_text_buffer_get_insert", m.Instance())
	if r == 0 {
		return nil
	}
	mark := new(TextMark)
	mark.instance = unsafe.Pointer(r)
	return mark
}

// CreateMark is a wrapper around gtk_text_buffer_create_mark().
func (m *TextBuffer) CreateMark(markName string, where ITextIter, leftGravity bool) ITextMark {
	cstr := CStr(markName)
	var wherePtr uintptr
	if where != nil {
		wherePtr = where.(*TextIter).native()
	}
	r := gtk3.SysCall("gtk_text_buffer_create_mark", m.Instance(), cstr, wherePtr, ToCBool(leftGravity))
	if r == 0 {
		return nil
	}
	mark := new(TextMark)
	mark.instance = unsafe.Pointer(r)
	return mark
}

// GetMark is a wrapper around gtk_text_buffer_get_mark().
func (m *TextBuffer) GetMark(name string) ITextMark {
	r := gtk3.SysCall("gtk_text_buffer_get_mark", m.Instance(), CStr(name))
	if r == 0 {
		return nil
	}
	mark := new(TextMark)
	mark.instance = unsafe.Pointer(r)
	return mark
}

// DeleteMark is a wrapper around gtk_text_buffer_delete_mark().
func (m *TextBuffer) DeleteMark(mark ITextMark) {
	gtk3.SysCall("gtk_text_buffer_delete_mark", m.Instance(), mark.(*TextMark).Instance())
}

// GetIterAtMark is a wrapper around gtk_text_buffer_get_iter_at_mark().
func (m *TextBuffer) GetIterAtMark(mark ITextMark) ITextIter {
	iter := new(TextIter)
	gtk3.SysCall("gtk_text_buffer_get_iter_at_mark", m.Instance(), iter.native(), mark.(*TextMark).Instance())
	return iter
}

// ApplyTagByName is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (m *TextBuffer) ApplyTagByName(name string, start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	gtk3.SysCall("gtk_text_buffer_apply_tag_by_name", m.Instance(), CStr(name), s.native(), e.native())
}

// RemoveTagByName is a wrapper around gtk_text_buffer_remove_tag_by_name().
func (m *TextBuffer) RemoveTagByName(name string, start, end ITextIter) {
	s := start.(*TextIter)
	e := end.(*TextIter)
	gtk3.SysCall("gtk_text_buffer_remove_tag_by_name", m.Instance(), CStr(name), s.native(), e.native())
}

func (m *TextBuffer) SetOnChanged(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnChanged, callback.C_trampoline_2_void, fn, 0)
}
