package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextView is a representation of GTK's GtkTextView.
type TextView struct {
	Widget
}

func (v *TextView) native() *C.GtkTextView {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTextView(unsafe.Pointer(v.GObject))
}

func wrapTextView(obj *Object) *TextView {
	return &TextView{Widget{InitiallyUnowned{obj}}}
}

// NewTextView is a wrapper around gtk_text_view_new().
func NewTextView() *TextView {
	c := C.gtk_text_view_new()
	if c == nil {
		return nil
	}
	return wrapTextView(ToGoObject(unsafe.Pointer(c)))
}

// NewTextViewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func NewTextViewWithBuffer(buffer ITextBuffer) *TextView {
	b := buffer.(*TextBuffer)
	c := C.gtk_text_view_new_with_buffer(b.native())
	if c == nil {
		return nil
	}
	return wrapTextView(ToGoObject(unsafe.Pointer(c)))
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (v *TextView) GetBuffer() ITextBuffer {
	c := C.gtk_text_view_get_buffer(v.native())
	if c == nil {
		return nil
	}
	return wrapTextBuffer(ToGoObject(unsafe.Pointer(c)))
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (v *TextView) SetBuffer(buffer ITextBuffer) {
	b := buffer.(*TextBuffer)
	C.gtk_text_view_set_buffer(v.native(), b.native())
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (v *TextView) SetEditable(editable bool) {
	C.gtk_text_view_set_editable(v.native(), CBool(editable))
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (v *TextView) GetEditable() bool {
	return GoBool(C.gtk_text_view_get_editable(v.native()))
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (v *TextView) SetWrapMode(wrapMode WrapMode) {
	C.gtk_text_view_set_wrap_mode(v.native(), C.GtkWrapMode(wrapMode))
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(v.native()))
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (v *TextView) SetCursorVisible(visible bool) {
	C.gtk_text_view_set_cursor_visible(v.native(), CBool(visible))
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (v *TextView) GetCursorVisible() bool {
	return GoBool(C.gtk_text_view_get_cursor_visible(v.native()))
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(v.native(), CBool(overwrite))
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (v *TextView) GetOverwrite() bool {
	return GoBool(C.gtk_text_view_get_overwrite(v.native()))
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (v *TextView) SetJustification(justify Justification) {
	C.gtk_text_view_set_justification(v.native(), C.GtkJustification(justify))
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (v *TextView) GetJustification() Justification {
	return Justification(C.gtk_text_view_get_justification(v.native()))
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (v *TextView) SetAcceptsTab(acceptsTab bool) {
	C.gtk_text_view_set_accepts_tab(v.native(), CBool(acceptsTab))
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (v *TextView) GetAcceptsTab() bool {
	return GoBool(C.gtk_text_view_get_accepts_tab(v.native()))
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (v *TextView) SetLeftMargin(margin int) {
	C.gtk_text_view_set_left_margin(v.native(), C.gint(margin))
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (v *TextView) GetLeftMargin() int {
	return int(C.gtk_text_view_get_left_margin(v.native()))
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (v *TextView) SetRightMargin(margin int) {
	C.gtk_text_view_set_right_margin(v.native(), C.gint(margin))
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (v *TextView) GetRightMargin() int {
	return int(C.gtk_text_view_get_right_margin(v.native()))
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (v *TextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(v.native(), C.gint(indent))
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (v *TextView) GetIndent() int {
	return int(C.gtk_text_view_get_indent(v.native()))
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (v *TextView) SetPixelsAboveLines(px int) {
	C.gtk_text_view_set_pixels_above_lines(v.native(), C.gint(px))
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (v *TextView) GetPixelsAboveLines() int {
	return int(C.gtk_text_view_get_pixels_above_lines(v.native()))
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (v *TextView) SetPixelsBelowLines(px int) {
	C.gtk_text_view_set_pixels_below_lines(v.native(), C.gint(px))
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (v *TextView) GetPixelsBelowLines() int {
	return int(C.gtk_text_view_get_pixels_below_lines(v.native()))
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (v *TextView) SetPixelsInsideWrap(px int) {
	C.gtk_text_view_set_pixels_inside_wrap(v.native(), C.gint(px))
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (v *TextView) GetPixelsInsideWrap() int {
	return int(C.gtk_text_view_get_pixels_inside_wrap(v.native()))
}

// ScrollToIter is a wrapper around gtk_text_view_scroll_to_iter().
func (v *TextView) ScrollToIter(iter ITextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool {
	i := iter.(*TextIter)
	return GoBool(C.gtk_text_view_scroll_to_iter(v.native(), &i.nativeIter,
		C.gdouble(withinMargin), CBool(useAlign), C.gdouble(xalign), C.gdouble(yalign)))
}

// PlaceCursorOnscreen is a wrapper around gtk_text_view_place_cursor_onscreen().
func (v *TextView) PlaceCursorOnscreen() bool {
	return GoBool(C.gtk_text_view_place_cursor_onscreen(v.native()))
}

// ResetImContext is a wrapper around gtk_text_view_reset_im_context().
func (v *TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context(v.native())
}
