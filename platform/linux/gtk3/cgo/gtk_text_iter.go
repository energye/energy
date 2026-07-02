package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import . "github.com/energye/energy/v3/platform/linux/types"

// TextIter is a representation of GTK's GtkTextIter.
type TextIter struct {
	nativeIter C.GtkTextIter
}

func newTextIterFromC(cIter C.GtkTextIter) *TextIter {
	return &TextIter{nativeIter: cIter}
}

// GetOffset is a wrapper around gtk_text_iter_get_offset().
func (v *TextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&v.nativeIter))
}

// GetLine is a wrapper around gtk_text_iter_get_line().
func (v *TextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&v.nativeIter))
}

// GetLineOffset is a wrapper around gtk_text_iter_get_line_offset().
func (v *TextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&v.nativeIter))
}

// SetOffset is a wrapper around gtk_text_iter_set_offset().
func (v *TextIter) SetOffset(charOffset int) {
	C.gtk_text_iter_set_offset(&v.nativeIter, C.gint(charOffset))
}

// SetLine is a wrapper around gtk_text_iter_set_line().
func (v *TextIter) SetLine(lineNumber int) {
	C.gtk_text_iter_set_line(&v.nativeIter, C.gint(lineNumber))
}

// ForwardChar is a wrapper around gtk_text_iter_forward_char().
func (v *TextIter) ForwardChar() bool {
	return GoBool(C.gtk_text_iter_forward_char(&v.nativeIter))
}

// BackwardChar is a wrapper around gtk_text_iter_backward_char().
func (v *TextIter) BackwardChar() bool {
	return GoBool(C.gtk_text_iter_backward_char(&v.nativeIter))
}

// ForwardLine is a wrapper around gtk_text_iter_forward_line().
func (v *TextIter) ForwardLine() bool {
	return GoBool(C.gtk_text_iter_forward_line(&v.nativeIter))
}

// BackwardLine is a wrapper around gtk_text_iter_backward_line().
func (v *TextIter) BackwardLine() bool {
	return GoBool(C.gtk_text_iter_backward_line(&v.nativeIter))
}

// IsEnd is a wrapper around gtk_text_iter_is_end().
func (v *TextIter) IsEnd() bool {
	return GoBool(C.gtk_text_iter_is_end(&v.nativeIter))
}

// IsStart is a wrapper around gtk_text_iter_is_start().
func (v *TextIter) IsStart() bool {
	return GoBool(C.gtk_text_iter_is_start(&v.nativeIter))
}

// Equal is a wrapper around gtk_text_iter_equal().
func (v *TextIter) Equal(other ITextIter) bool {
	o := other.(*TextIter)
	return GoBool(C.gtk_text_iter_equal(&v.nativeIter, &o.nativeIter))
}
