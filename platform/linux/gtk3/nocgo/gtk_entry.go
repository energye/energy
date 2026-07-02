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

// Layout is a representation of GTK's GtkLayout.
type Entry struct {
	Widget
}

func AsEntry(ptr unsafe.Pointer) IEntry {
	if ptr == nil {
		return nil
	}
	m := new(Entry)
	m.instance = ptr
	return m
}

// NewEntry is a wrapper around gtk_entry_new().
func NewEntry() IEntry {
	r := gtk3.SysCall("gtk_entry_new")
	return AsEntry(unsafe.Pointer(r))
}

// SetText is a wrapper around gtk_entry_set_text().
func (m *Entry) SetText(text string) {
	gtk3.SysCall("gtk_entry_set_text", m.Instance(), CStr(text))
}

// GetText is a wrapper around gtk_entry_get_text().
func (m *Entry) GetText() string {
	c := gtk3.SysCall("gtk_entry_get_text", m.Instance())
	return GoStr(c)
}

// GetTextLength is a wrapper around gtk_entry_get_text_length().
func (m *Entry) GetTextLength() uint16 {
	c := gtk3.SysCall("gtk_entry_get_text_length", m.Instance())
	return uint16(c)
}

// NewEntryWithBuffer is a wrapper around gtk_entry_new_with_buffer().
func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	r := gtk3.SysCall("gtk_entry_new_with_buffer", buffer.Instance())
	if r == 0 {
		return nil
	}
	m := new(Entry)
	m.instance = unsafe.Pointer(r)
	return m
}

// GetBuffer is a wrapper around gtk_entry_get_buffer().
func (m *Entry) GetBuffer() (*EntryBuffer, error) {
	r := gtk3.SysCall("gtk_entry_get_buffer", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	buf := new(EntryBuffer)
	buf.instance = unsafe.Pointer(r)
	return buf, nil
}

// SetBuffer is a wrapper around gtk_entry_set_buffer().
func (m *Entry) SetBuffer(buffer *EntryBuffer) {
	gtk3.SysCall("gtk_entry_set_buffer", m.Instance(), buffer.Instance())
}

// SetVisibility is a wrapper around gtk_entry_set_visibility().
func (m *Entry) SetVisibility(visible bool) {
	gtk3.SysCall("gtk_entry_set_visibility", m.Instance(), ToCBool(visible))
}

// GetVisibility is a wrapper around gtk_entry_get_visibility().
func (m *Entry) GetVisibility() bool {
	r := gtk3.SysCall("gtk_entry_get_visibility", m.Instance())
	return ToGoBool(r)
}

// SetMaxLength is a wrapper around gtk_entry_set_max_length().
func (m *Entry) SetMaxLength(len int) {
	gtk3.SysCall("gtk_entry_set_max_length", m.Instance(), uintptr(len))
}

// GetMaxLength is a wrapper around gtk_entry_get_max_length().
func (m *Entry) GetMaxLength() int {
	r := gtk3.SysCall("gtk_entry_get_max_length", m.Instance())
	return int(r)
}

// SetHasFrame is a wrapper around gtk_entry_set_has_frame().
func (m *Entry) SetHasFrame(setting bool) {
	gtk3.SysCall("gtk_entry_set_has_frame", m.Instance(), ToCBool(setting))
}

// GetHasFrame is a wrapper around gtk_entry_get_has_frame().
func (m *Entry) GetHasFrame() bool {
	r := gtk3.SysCall("gtk_entry_get_has_frame", m.Instance())
	return ToGoBool(r)
}

// SetWidthChars is a wrapper around gtk_entry_set_width_chars().
func (m *Entry) SetWidthChars(nChars int) {
	gtk3.SysCall("gtk_entry_set_width_chars", m.Instance(), uintptr(nChars))
}

// GetWidthChars is a wrapper around gtk_entry_get_width_chars().
func (m *Entry) GetWidthChars() int {
	r := gtk3.SysCall("gtk_entry_get_width_chars", m.Instance())
	return int(r)
}

// SetActivatesDefault is a wrapper around gtk_entry_set_activates_default().
func (m *Entry) SetActivatesDefault(setting bool) {
	gtk3.SysCall("gtk_entry_set_activates_default", m.Instance(), ToCBool(setting))
}

// GetActivatesDefault is a wrapper around gtk_entry_get_activates_default().
func (m *Entry) GetActivatesDefault() bool {
	r := gtk3.SysCall("gtk_entry_get_activates_default", m.Instance())
	return ToGoBool(r)
}

// SetAlignment is a wrapper around gtk_entry_set_alignment().
func (m *Entry) SetAlignment(xalign float32) {
	gtk3.SysCall("gtk_entry_set_alignment", m.Instance(), uintptr(xalign))
}

// GetAlignment is a wrapper around gtk_entry_get_alignment().
func (m *Entry) GetAlignment() float32 {
	r := gtk3.SysCall("gtk_entry_get_alignment", m.Instance())
	return float32(r)
}

// SetPlaceholderText is a wrapper around gtk_entry_set_placeholder_text().
func (m *Entry) SetPlaceholderText(text string) {
	gtk3.SysCall("gtk_entry_set_placeholder_text", m.Instance(), CStr(text))
}

// GetPlaceholderText is a wrapper around gtk_entry_get_placeholder_text().
func (m *Entry) GetPlaceholderText() string {
	r := gtk3.SysCall("gtk_entry_get_placeholder_text", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// SetInvisibleChar is a wrapper around gtk_entry_set_invisible_char().
func (m *Entry) SetInvisibleChar(ch rune) {
	gtk3.SysCall("gtk_entry_set_invisible_char", m.Instance(), uintptr(ch))
}

// UnsetInvisibleChar is a wrapper around gtk_entry_unset_invisible_char().
func (m *Entry) UnsetInvisibleChar() {
	gtk3.SysCall("gtk_entry_unset_invisible_char", m.Instance())
}

// GetInvisibleChar is a wrapper around gtk_entry_get_invisible_char().
func (m *Entry) GetInvisibleChar() rune {
	r := gtk3.SysCall("gtk_entry_get_invisible_char", m.Instance())
	return rune(r)
}

// SetOverwriteMode is a wrapper around gtk_entry_set_overwrite_mode().
func (m *Entry) SetOverwriteMode(overwrite bool) {
	gtk3.SysCall("gtk_entry_set_overwrite_mode", m.Instance(), ToCBool(overwrite))
}

// GetOverwriteMode is a wrapper around gtk_entry_get_overwrite_mode().
func (m *Entry) GetOverwriteMode() bool {
	r := gtk3.SysCall("gtk_entry_get_overwrite_mode", m.Instance())
	return ToGoBool(r)
}

// GetLayoutOffsets is a wrapper around gtk_entry_get_layout_offsets().
func (m *Entry) GetLayoutOffsets() (x, y int) {
	var gx, gy uintptr
	gtk3.SysCall("gtk_entry_get_layout_offsets", m.Instance(),
		uintptr(unsafe.Pointer(&gx)), uintptr(unsafe.Pointer(&gy)))
	return int(gx), int(gy)
}

// LayoutIndexToTextIndex is a wrapper around gtk_entry_layout_index_to_text_index().
func (m *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	r := gtk3.SysCall("gtk_entry_layout_index_to_text_index", m.Instance(), uintptr(layoutIndex))
	return int(r)
}

// TextIndexToLayoutIndex is a wrapper around gtk_entry_text_index_to_layout_index().
func (m *Entry) TextIndexToLayoutIndex(textIndex int) int {
	r := gtk3.SysCall("gtk_entry_text_index_to_layout_index", m.Instance(), uintptr(textIndex))
	return int(r)
}

// ProgressPulse is a wrapper around gtk_entry_progress_pulse().
func (m *Entry) ProgressPulse() {
	gtk3.SysCall("gtk_entry_progress_pulse", m.Instance())
}

// ResetIMContext is a wrapper around gtk_entry_reset_im_context().
func (m *Entry) ResetIMContext() {
	gtk3.SysCall("gtk_entry_reset_im_context", m.Instance())
}

// SetIconFromIconName is a wrapper around gtk_entry_set_icon_from_icon_name().
func (m *Entry) SetIconFromIconName(iconPos EntryIconPosition, name string) {
	var icon uintptr
	if name != "" {
		icon = CStr(name)
	}
	gtk3.SysCall("gtk_entry_set_icon_from_icon_name", m.Instance(), uintptr(iconPos), icon)
}

// RemoveIcon is a convenience func to set a nil pointer to the icon name.
func (m *Entry) RemoveIcon(iconPos EntryIconPosition) {
	gtk3.SysCall("gtk_entry_set_icon_from_icon_name", m.Instance(), uintptr(iconPos), 0)
}

// GetIconStorageType is a wrapper around gtk_entry_get_icon_storage_type().
func (m *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	r := gtk3.SysCall("gtk_entry_get_icon_storage_type", m.Instance(), uintptr(iconPos))
	return ImageType(r)
}

// GetIconName is a wrapper around gtk_entry_get_icon_name().
func (m *Entry) GetIconName(iconPos EntryIconPosition) string {
	r := gtk3.SysCall("gtk_entry_get_icon_name", m.Instance(), uintptr(iconPos))
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// SetIconActivatable is a wrapper around gtk_entry_set_icon_activatable().
func (m *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	gtk3.SysCall("gtk_entry_set_icon_activatable", m.Instance(), uintptr(iconPos), ToCBool(activatable))
}

// GetIconActivatable is a wrapper around gtk_entry_get_icon_activatable().
func (m *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	r := gtk3.SysCall("gtk_entry_get_icon_activatable", m.Instance(), uintptr(iconPos))
	return ToGoBool(r)
}

// SetIconSensitive is a wrapper around gtk_entry_set_icon_sensitive().
func (m *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	gtk3.SysCall("gtk_entry_set_icon_sensitive", m.Instance(), uintptr(iconPos), ToCBool(sensitive))
}

// GetIconSensitive is a wrapper around gtk_entry_get_icon_sensitive().
func (m *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	r := gtk3.SysCall("gtk_entry_get_icon_sensitive", m.Instance(), uintptr(iconPos))
	return ToGoBool(r)
}

// GetIconAtPos is a wrapper around gtk_entry_get_icon_at_pos().
func (m *Entry) GetIconAtPos(x, y int) int {
	r := gtk3.SysCall("gtk_entry_get_icon_at_pos", m.Instance(), uintptr(x), uintptr(y))
	return int(r)
}

// SetIconTooltipText is a wrapper around gtk_entry_set_icon_tooltip_text().
func (m *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	var text uintptr
	if tooltip != "" {
		text = CStr(tooltip)
	}
	gtk3.SysCall("gtk_entry_set_icon_tooltip_text", m.Instance(), uintptr(iconPos), text)
}

// GetIconTooltipText is a wrapper around gtk_entry_get_icon_tooltip_text().
func (m *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	r := gtk3.SysCall("gtk_entry_get_icon_tooltip_text", m.Instance(), uintptr(iconPos))
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// SetIconTooltipMarkup is a wrapper around gtk_entry_set_icon_tooltip_markup().
func (m *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	var text uintptr
	if tooltip != "" {
		text = CStr(tooltip)
	}
	gtk3.SysCall("gtk_entry_set_icon_tooltip_markup", m.Instance(), uintptr(iconPos), text)
}

// GetIconTooltipMarkup is a wrapper around gtk_entry_get_icon_tooltip_markup().
func (m *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	r := gtk3.SysCall("gtk_entry_get_icon_tooltip_markup", m.Instance(), uintptr(iconPos))
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetCurrentIconDragSource is a wrapper around gtk_entry_get_current_icon_drag_source().
func (m *Entry) GetCurrentIconDragSource() int {
	r := gtk3.SysCall("gtk_entry_get_current_icon_drag_source", m.Instance())
	return int(r)
}

// SetCompletion is a wrapper around gtk_entry_set_completion().
func (m *Entry) SetCompletion(completion *EntryCompletion) {
	gtk3.SysCall("gtk_entry_set_completion", m.Instance(), completion.Instance())
}

// GetCompletion is a wrapper around gtk_entry_get_completion().
func (m *Entry) GetCompletion() *EntryCompletion {
	r := gtk3.SysCall("gtk_entry_get_completion", m.Instance())
	if r == 0 {
		return nil
	}
	e := new(EntryCompletion)
	e.instance = unsafe.Pointer(r)
	return e
}

// SetProgressFraction is a wrapper around gtk_entry_set_progress_fraction().
func (m *Entry) SetProgressFraction(fraction float64) {
	gtk3.SysCall("gtk_entry_set_progress_fraction", m.Instance(), uintptr(fraction))
}

// GetProgressFraction is a wrapper around gtk_entry_get_progress_fraction().
func (m *Entry) GetProgressFraction() float64 {
	r := gtk3.SysCall("gtk_entry_get_progress_fraction", m.Instance())
	return float64(r)
}

// SetProgressPulseStep is a wrapper around gtk_entry_set_progress_pulse_step().
func (m *Entry) SetProgressPulseStep(fraction float64) {
	gtk3.SysCall("gtk_entry_set_progress_pulse_step", m.Instance(), uintptr(fraction))
}

// GetProgressPulseStep is a wrapper around gtk_entry_get_progress_pulse_step().
func (m *Entry) GetProgressPulseStep() float64 {
	r := gtk3.SysCall("gtk_entry_get_progress_pulse_step", m.Instance())
	return float64(r)
}

func (m *Entry) SetOnChanged(fn TTextChangedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnChanged, callback.C_trampoline_2_void, fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnCommit(fn TTextCommitEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnActivate, callback.C_trampoline_2_void, fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnKeyPress(fn TTextKeyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnKeyPressEvent, callback.C_trampoline_3_gboolean,
		fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnKeyRelease(fn TTextKeyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnKeyReleaseEvent, callback.C_trampoline_3_gboolean, fn, 0)
	return signalHandlerID
}

// SetIconFromGIcon is a wrapper around gtk_entry_set_icon_from_gicon().
func (m *Entry) SetIconFromGIcon(iconPos EntryIconPosition, icon *Icon) {
	gtk3.SysCall("gtk_entry_set_icon_from_gicon", m.Instance(), uintptr(iconPos), icon.NativePrivate())
}

// GetIconGIcon is a wrapper around gtk_entry_get_icon_gicon().
func (m *Entry) GetIconGIcon(iconPos EntryIconPosition) *Icon {
	r := gtk3.SysCall("gtk_entry_get_icon_gicon", m.Instance(), uintptr(iconPos))
	if r == 0 {
		return nil
	}
	return AsIcon(unsafe.Pointer(r))
}

// SetCursorHAdjustment is a wrapper around gtk_entry_set_cursor_hadjustment().
func (m *Entry) SetCursorHAdjustment(adjustment *Adjustment) {
	gtk3.SysCall("gtk_entry_set_cursor_hadjustment", m.Instance(), adjustment.Instance())
}

// GetCursorHAdjustment is a wrapper around gtk_entry_get_cursor_hadjustment().
func (m *Entry) GetCursorHAdjustment() *Adjustment {
	r := gtk3.SysCall("gtk_entry_get_cursor_hadjustment", m.Instance())
	if r == 0 {
		return nil
	}
	adj := new(Adjustment)
	adj.instance = unsafe.Pointer(r)
	return adj
}
