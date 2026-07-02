//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build amd64 || arm64 || loong64

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// GtkTextIter is 80 bytes on x86_64 (verified via sizeof).
const textIterSize = 80

// TextIter is a representation of GTK's GtkTextIter.
// It is a value type (stack-allocated), not a GObject.
type TextIter struct {
	data [textIterSize]byte
}

func (m *TextIter) native() uintptr {
	return uintptr(unsafe.Pointer(&m.data[0]))
}

// GetOffset is a wrapper around gtk_text_iter_get_offset().
func (m *TextIter) GetOffset() int {
	r := gtk3.SysCall("gtk_text_iter_get_offset", m.native())
	return int(r)
}

// GetLine is a wrapper around gtk_text_iter_get_line().
func (m *TextIter) GetLine() int {
	r := gtk3.SysCall("gtk_text_iter_get_line", m.native())
	return int(r)
}

// GetLineOffset is a wrapper around gtk_text_iter_get_line_offset().
func (m *TextIter) GetLineOffset() int {
	r := gtk3.SysCall("gtk_text_iter_get_line_offset", m.native())
	return int(r)
}

// SetOffset is a wrapper around gtk_text_iter_set_offset().
func (m *TextIter) SetOffset(charOffset int) {
	gtk3.SysCall("gtk_text_iter_set_offset", m.native(), uintptr(charOffset))
}

// SetLine is a wrapper around gtk_text_iter_set_line().
func (m *TextIter) SetLine(lineNumber int) {
	gtk3.SysCall("gtk_text_iter_set_line", m.native(), uintptr(lineNumber))
}

// ForwardChar is a wrapper around gtk_text_iter_forward_char().
func (m *TextIter) ForwardChar() bool {
	r := gtk3.SysCall("gtk_text_iter_forward_char", m.native())
	return ToGoBool(r)
}

// BackwardChar is a wrapper around gtk_text_iter_backward_char().
func (m *TextIter) BackwardChar() bool {
	r := gtk3.SysCall("gtk_text_iter_backward_char", m.native())
	return ToGoBool(r)
}

// ForwardLine is a wrapper around gtk_text_iter_forward_line().
func (m *TextIter) ForwardLine() bool {
	r := gtk3.SysCall("gtk_text_iter_forward_line", m.native())
	return ToGoBool(r)
}

// BackwardLine is a wrapper around gtk_text_iter_backward_line().
func (m *TextIter) BackwardLine() bool {
	r := gtk3.SysCall("gtk_text_iter_backward_line", m.native())
	return ToGoBool(r)
}

// IsEnd is a wrapper around gtk_text_iter_is_end().
func (m *TextIter) IsEnd() bool {
	r := gtk3.SysCall("gtk_text_iter_is_end", m.native())
	return ToGoBool(r)
}

// IsStart is a wrapper around gtk_text_iter_is_start().
func (m *TextIter) IsStart() bool {
	r := gtk3.SysCall("gtk_text_iter_is_start", m.native())
	return ToGoBool(r)
}

// Equal is a wrapper around gtk_text_iter_equal().
func (m *TextIter) Equal(other ITextIter) bool {
	r := gtk3.SysCall("gtk_text_iter_equal", m.native(), other.(*TextIter).native())
	return ToGoBool(r)
}
