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

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

func AsLabel(ptr unsafe.Pointer) ILabel {
	if ptr == nil {
		return nil
	}
	m := new(Label)
	m.instance = ptr
	return m
}

// NewLabel is a wrapper around gtk_label_new().
func NewLabel(str string) ILabel {
	r := gtk3.SysCall("gtk_label_new", CStr(str))
	if r == 0 {
		return nil
	}
	return AsLabel(unsafe.Pointer(r))
}

// SetText is a wrapper around gtk_label_set_text().
func (m *Label) SetText(str string) {
	gtk3.SysCall("gtk_label_set_text", m.Instance(), CStr(str))
}

// GetText is a wrapper around gtk_label_get_text().
func (m *Label) GetText() (string, error) {
	r := gtk3.SysCall("gtk_label_get_text", m.Instance())
	return GoStr(r), nil
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (m *Label) SetMarkup(str string) {
	gtk3.SysCall("gtk_label_set_markup", m.Instance(), CStr(str))
}

// SetMarkupWithMnemonic is a wrapper around gtk_label_set_markup_with_mnemonic().
func (m *Label) SetMarkupWithMnemonic(str string) {
	gtk3.SysCall("gtk_label_set_markup_with_mnemonic", m.Instance(), CStr(str))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (m *Label) SetPattern(pattern string) {
	gtk3.SysCall("gtk_label_set_pattern", m.Instance(), CStr(pattern))
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (m *Label) SetJustify(jtype Justification) {
	gtk3.SysCall("gtk_label_set_justify", m.Instance(), uintptr(jtype))
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (m *Label) GetJustify() Justification {
	r := gtk3.SysCall("gtk_label_get_justify", m.Instance())
	return Justification(r)
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (m *Label) GetCurrentUri() string {
	r := gtk3.SysCall("gtk_label_get_current_uri", m.Instance())
	return GoStr(r)
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (m *Label) GetTrackVisitedLinks() bool {
	r := gtk3.SysCall("gtk_label_get_track_visited_links", m.Instance())
	return ToGoBool(r)
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (m *Label) SetTrackVisitedLinks(trackLinks bool) {
	gtk3.SysCall("gtk_label_set_track_visited_links", m.Instance(), ToCBool(trackLinks))
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (m *Label) GetAngle() float64 {
	registerGtkFloatFuncs()
	return gtkLabelGetAngle(m.Instance())
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (m *Label) SetAngle(angle float64) {
	registerGtkFloatFuncs()
	gtkLabelSetAngle(m.Instance(), angle)
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (m *Label) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend int32
	r := gtk3.SysCall("gtk_label_get_selection_bounds", m.Instance(), uintptr(unsafe.Pointer(&cstart)), uintptr(unsafe.Pointer(&cend)))
	return int(cstart), int(cend), ToGoBool(r)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (m *Label) GetSingleLineMode() bool {
	r := gtk3.SysCall("gtk_label_get_single_line_mode", m.Instance())
	return ToGoBool(r)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (m *Label) SetSingleLineMode(mode bool) {
	gtk3.SysCall("gtk_label_set_single_line_mode", m.Instance(), ToCBool(mode))
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (m *Label) GetUseMarkup() bool {
	r := gtk3.SysCall("gtk_label_get_use_markup", m.Instance())
	return ToGoBool(r)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (m *Label) SetUseMarkup(use bool) {
	gtk3.SysCall("gtk_label_set_use_markup", m.Instance(), ToCBool(use))
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (m *Label) GetUseUnderline() bool {
	r := gtk3.SysCall("gtk_label_get_use_underline", m.Instance())
	return ToGoBool(r)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (m *Label) SetUseUnderline(use bool) {
	gtk3.SysCall("gtk_label_set_use_underline", m.Instance(), ToCBool(use))
}

// LabelNewWithMnemonic is a wrapper around gtk_label_new_with_mnemonic().
func LabelNewWithMnemonic(str string) (*Label, error) {
	r := gtk3.SysCall("gtk_label_new_with_mnemonic", CStr(str))
	if r == 0 {
		return nil, errNilPtr
	}
	return &Label{Widget{Object{instance: unsafe.Pointer(r)}}}, nil
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (m *Label) SetEllipsize(mode EllipsizeMode) {
	gtk3.SysCall("gtk_label_set_ellipsize", m.Instance(), uintptr(mode))
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (m *Label) GetEllipsize() EllipsizeMode {
	r := gtk3.SysCall("gtk_label_get_ellipsize", m.Instance())
	return EllipsizeMode(r)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (m *Label) SetWidthChars(nChars int) {
	gtk3.SysCall("gtk_label_set_width_chars", m.Instance(), uintptr(nChars))
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (m *Label) GetWidthChars() int {
	r := gtk3.SysCall("gtk_label_get_width_chars", m.Instance())
	return int(r)
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (m *Label) SetMaxWidthChars(nChars int) {
	gtk3.SysCall("gtk_label_set_max_width_chars", m.Instance(), uintptr(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (m *Label) GetMaxWidthChars() int {
	r := gtk3.SysCall("gtk_label_get_max_width_chars", m.Instance())
	return int(r)
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (m *Label) SetLineWrap(wrap bool) {
	gtk3.SysCall("gtk_label_set_line_wrap", m.Instance(), ToCBool(wrap))
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (m *Label) GetLineWrap() bool {
	r := gtk3.SysCall("gtk_label_get_line_wrap", m.Instance())
	return ToGoBool(r)
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (m *Label) SetSelectable(setting bool) {
	gtk3.SysCall("gtk_label_set_selectable", m.Instance(), ToCBool(setting))
}

// GetSelectable is a wrapper around gtk_label_get_selectable().
func (m *Label) GetSelectable() bool {
	r := gtk3.SysCall("gtk_label_get_selectable", m.Instance())
	return ToGoBool(r)
}

// SelectRegion is a wrapper around gtk_label_select_region().
func (m *Label) SelectRegion(startOffset, endOffset int) {
	gtk3.SysCall("gtk_label_select_region", m.Instance(), uintptr(startOffset), uintptr(endOffset))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (m *Label) SetLabel(str string) {
	gtk3.SysCall("gtk_label_set_label", m.Instance(), CStr(str))
}

// GetLabel is a wrapper around gtk_label_get_label().
func (m *Label) GetLabel() string {
	r := gtk3.SysCall("gtk_label_get_label", m.Instance())
	return GoStr(r)
}

// GetMnemonicKeyval is a wrapper around gtk_label_get_mnemonic_keyval().
func (m *Label) GetMnemonicKeyval() uint {
	r := gtk3.SysCall("gtk_label_get_mnemonic_keyval", m.Instance())
	return uint(r)
}

// SetMnemonicWidget is a wrapper around gtk_label_set_mnemonic_widget().
func (m *Label) SetMnemonicWidget(widget IWidget) {
	var widgetPtr uintptr
	if widget != nil {
		widgetPtr = widget.Instance()
	}
	gtk3.SysCall("gtk_label_set_mnemonic_widget", m.Instance(), widgetPtr)
}

// SetXAlign is a wrapper around gtk_label_set_xalign().
func (m *Label) SetXAlign(n float64) {
	registerGtkFloatFuncs()
	gtkLabelSetXAlign(m.Instance(), float32(n))
}

// GetXAlign is a wrapper around gtk_label_get_xalign().
func (m *Label) GetXAlign() float64 {
	registerGtkFloatFuncs()
	return float64(gtkLabelGetXAlign(m.Instance()))
}

// SetYAlign is a wrapper around gtk_label_set_yalign().
func (m *Label) SetYAlign(n float64) {
	registerGtkFloatFuncs()
	gtkLabelSetYAlign(m.Instance(), float32(n))
}

// GetYAlign is a wrapper around gtk_label_get_yalign().
func (m *Label) GetYAlign() float64 {
	registerGtkFloatFuncs()
	return float64(gtkLabelGetYAlign(m.Instance()))
}
