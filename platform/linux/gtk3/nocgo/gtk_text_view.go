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
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TextView is a representation of GTK's GtkTextView.
type TextView struct {
	Widget
}

func AsTextView(ptr unsafe.Pointer) *TextView {
	if ptr == nil {
		return nil
	}
	m := new(TextView)
	m.instance = ptr
	return m
}

// NewTextView is a wrapper around gtk_text_view_new().
func NewTextView() *TextView {
	r := gtk3.SysCall("gtk_text_view_new")
	if r == 0 {
		return nil
	}
	return AsTextView(unsafe.Pointer(r))
}

// NewTextViewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func NewTextViewWithBuffer(buffer ITextBuffer) *TextView {
	r := gtk3.SysCall("gtk_text_view_new_with_buffer", buffer.(*TextBuffer).Instance())
	if r == 0 {
		return nil
	}
	return AsTextView(unsafe.Pointer(r))
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (m *TextView) GetBuffer() ITextBuffer {
	r := gtk3.SysCall("gtk_text_view_get_buffer", m.Instance())
	if r == 0 {
		return nil
	}
	return AsTextBuffer(unsafe.Pointer(r))
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (m *TextView) SetBuffer(buffer ITextBuffer) {
	gtk3.SysCall("gtk_text_view_set_buffer", m.Instance(), buffer.(*TextBuffer).Instance())
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (m *TextView) SetEditable(editable bool) {
	gtk3.SysCall("gtk_text_view_set_editable", m.Instance(), ToCBool(editable))
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (m *TextView) GetEditable() bool {
	r := gtk3.SysCall("gtk_text_view_get_editable", m.Instance())
	return ToGoBool(r)
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (m *TextView) SetWrapMode(wrapMode WrapMode) {
	gtk3.SysCall("gtk_text_view_set_wrap_mode", m.Instance(), uintptr(wrapMode))
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (m *TextView) GetWrapMode() WrapMode {
	r := gtk3.SysCall("gtk_text_view_get_wrap_mode", m.Instance())
	return WrapMode(r)
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (m *TextView) SetCursorVisible(visible bool) {
	gtk3.SysCall("gtk_text_view_set_cursor_visible", m.Instance(), ToCBool(visible))
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (m *TextView) GetCursorVisible() bool {
	r := gtk3.SysCall("gtk_text_view_get_cursor_visible", m.Instance())
	return ToGoBool(r)
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (m *TextView) SetOverwrite(overwrite bool) {
	gtk3.SysCall("gtk_text_view_set_overwrite", m.Instance(), ToCBool(overwrite))
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (m *TextView) GetOverwrite() bool {
	r := gtk3.SysCall("gtk_text_view_get_overwrite", m.Instance())
	return ToGoBool(r)
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (m *TextView) SetJustification(justify Justification) {
	gtk3.SysCall("gtk_text_view_set_justification", m.Instance(), uintptr(justify))
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (m *TextView) GetJustification() Justification {
	r := gtk3.SysCall("gtk_text_view_get_justification", m.Instance())
	return Justification(r)
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (m *TextView) SetAcceptsTab(acceptsTab bool) {
	gtk3.SysCall("gtk_text_view_set_accepts_tab", m.Instance(), ToCBool(acceptsTab))
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (m *TextView) GetAcceptsTab() bool {
	r := gtk3.SysCall("gtk_text_view_get_accepts_tab", m.Instance())
	return ToGoBool(r)
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (m *TextView) SetLeftMargin(margin int) {
	gtk3.SysCall("gtk_text_view_set_left_margin", m.Instance(), uintptr(margin))
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (m *TextView) GetLeftMargin() int {
	r := gtk3.SysCall("gtk_text_view_get_left_margin", m.Instance())
	return int(r)
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (m *TextView) SetRightMargin(margin int) {
	gtk3.SysCall("gtk_text_view_set_right_margin", m.Instance(), uintptr(margin))
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (m *TextView) GetRightMargin() int {
	r := gtk3.SysCall("gtk_text_view_get_right_margin", m.Instance())
	return int(r)
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (m *TextView) SetIndent(indent int) {
	gtk3.SysCall("gtk_text_view_set_indent", m.Instance(), uintptr(indent))
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (m *TextView) GetIndent() int {
	r := gtk3.SysCall("gtk_text_view_get_indent", m.Instance())
	return int(r)
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (m *TextView) SetPixelsAboveLines(px int) {
	gtk3.SysCall("gtk_text_view_set_pixels_above_lines", m.Instance(), uintptr(px))
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (m *TextView) GetPixelsAboveLines() int {
	r := gtk3.SysCall("gtk_text_view_get_pixels_above_lines", m.Instance())
	return int(r)
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (m *TextView) SetPixelsBelowLines(px int) {
	gtk3.SysCall("gtk_text_view_set_pixels_below_lines", m.Instance(), uintptr(px))
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (m *TextView) GetPixelsBelowLines() int {
	r := gtk3.SysCall("gtk_text_view_get_pixels_below_lines", m.Instance())
	return int(r)
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (m *TextView) SetPixelsInsideWrap(px int) {
	gtk3.SysCall("gtk_text_view_set_pixels_inside_wrap", m.Instance(), uintptr(px))
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (m *TextView) GetPixelsInsideWrap() int {
	r := gtk3.SysCall("gtk_text_view_get_pixels_inside_wrap", m.Instance())
	return int(r)
}

// ScrollToIter is a wrapper around gtk_text_view_scroll_to_iter().
func (m *TextView) ScrollToIter(iter ITextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool {
	i := iter.(*TextIter)
	r := gtk3.SysCall("gtk_text_view_scroll_to_iter", m.Instance(),
		i.native(), uintptr(uint64(withinMargin)), ToCBool(useAlign),
		uintptr(uint64(xalign)), uintptr(uint64(yalign)))
	return ToGoBool(r)
}

// PlaceCursorOnscreen is a wrapper around gtk_text_view_place_cursor_onscreen().
func (m *TextView) PlaceCursorOnscreen() bool {
	r := gtk3.SysCall("gtk_text_view_place_cursor_onscreen", m.Instance())
	return ToGoBool(r)
}

// ResetImContext is a wrapper around gtk_text_view_reset_im_context().
func (m *TextView) ResetImContext() {
	gtk3.SysCall("gtk_text_view_reset_im_context", m.Instance())
}
