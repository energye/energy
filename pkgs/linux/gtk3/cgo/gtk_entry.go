package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

type _IEntry interface {
	toEntry() *C.GtkEntry
}

// Entry is a representation of GTK's GtkEntry.
type Entry struct {
	Widget
	// Interfaces
	Editable
	CellEditable
}

func AsEntry(p unsafe.Pointer) *Entry {
	return wrapEntry(ToGoObject(p))
}

func (v *Entry) toEntry() *C.GtkEntry {
	return v.native()
}

// native returns a pointer to the underlying GtkEntry.
func (v *Entry) native() *C.GtkEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEntry(p)
}

func wrapEntry(obj *Object) *Entry {
	if obj == nil {
		return nil
	}

	e := wrapEditable(obj)
	ce := wrapCellEditable(obj)
	return &Entry{Widget{InitiallyUnowned{obj}}, *e, *ce}
}

// NewEntry is a wrapper around gtk_entry_new().
func NewEntry() IEntry {
	c := C.gtk_entry_new()
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapEntry(obj)
}

// NewEntryWithBuffer is a wrapper around gtk_entry_new_with_buffer().
func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	c := C.gtk_entry_new_with_buffer(buffer.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapEntry(obj)
}

// GetBuffer is a wrapper around gtk_entry_get_buffer().
func (v *Entry) GetBuffer() (*EntryBuffer, error) {
	c := C.gtk_entry_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return &EntryBuffer{obj}, nil
}

// SetBuffer is a wrapper around gtk_entry_set_buffer().
func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(v.native(), buffer.native())
}

// SetText is a wrapper around gtk_entry_set_text().
func (v *Entry) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText is a wrapper around gtk_entry_get_text().
func (v *Entry) GetText() string {
	c := C.gtk_entry_get_text(v.native())
	if c == nil {
		return ""
	}
	return GoString(c)
}

// GetTextLength is a wrapper around gtk_entry_get_text_length().
func (v *Entry) GetTextLength() uint16 {
	c := C.gtk_entry_get_text_length(v.native())
	return uint16(c)
}

// SetVisibility is a wrapper around gtk_entry_set_visibility().
func (v *Entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(v.native(), CBool(visible))
}

// SetInvisibleChar is a wrapper around gtk_entry_set_invisible_char().
func (v *Entry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(v.native(), C.gunichar(ch))
}

// UnsetInvisibleChar is a wrapper around gtk_entry_unset_invisible_char().
func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(v.native())
}

// SetMaxLength is a wrapper around gtk_entry_set_max_length().
func (v *Entry) SetMaxLength(len int) {
	C.gtk_entry_set_max_length(v.native(), C.gint(len))
}

// GetActivatesDefault is a wrapper around gtk_entry_get_activates_default().
func (v *Entry) GetActivatesDefault() bool {
	c := C.gtk_entry_get_activates_default(v.native())
	return GoBool(c)
}

// GetHasFrame is a wrapper around gtk_entry_get_has_frame().
func (v *Entry) GetHasFrame() bool {
	c := C.gtk_entry_get_has_frame(v.native())
	return GoBool(c)
}

// GetWidthChars is a wrapper around gtk_entry_get_width_chars().
func (v *Entry) GetWidthChars() int {
	c := C.gtk_entry_get_width_chars(v.native())
	return int(c)
}

// SetActivatesDefault is a wrapper around gtk_entry_set_activates_default().
func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(v.native(), CBool(setting))
}

// SetHasFrame is a wrapper around gtk_entry_set_has_frame().
func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(v.native(), CBool(setting))
}

// SetWidthChars is a wrapper around gtk_entry_set_width_chars().
func (v *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(v.native(), C.gint(nChars))
}

// GetInvisibleChar is a wrapper around gtk_entry_get_invisible_char().
func (v *Entry) GetInvisibleChar() rune {
	c := C.gtk_entry_get_invisible_char(v.native())
	return rune(c)
}

// SetAlignment is a wrapper around gtk_entry_set_alignment().
func (v *Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(v.native(), C.gfloat(xalign))
}

// GetAlignment is a wrapper around gtk_entry_get_alignment().
func (v *Entry) GetAlignment() float32 {
	c := C.gtk_entry_get_alignment(v.native())
	return float32(c)
}

// SetPlaceholderText is a wrapper around gtk_entry_set_placeholder_text().
func (v *Entry) SetPlaceholderText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_placeholder_text(v.native(), (*C.gchar)(cstr))
}

// GetPlaceholderText is a wrapper around gtk_entry_get_placeholder_text().
func (v *Entry) GetPlaceholderText() string {
	c := C.gtk_entry_get_placeholder_text(v.native())
	if c == nil {
		return ""
	}
	return GoString(c)
}

// SetOverwriteMode is a wrapper around gtk_entry_set_overwrite_mode().
func (v *Entry) SetOverwriteMode(overwrite bool) {
	C.gtk_entry_set_overwrite_mode(v.native(), CBool(overwrite))
}

// GetOverwriteMode is a wrapper around gtk_entry_get_overwrite_mode().
func (v *Entry) GetOverwriteMode() bool {
	c := C.gtk_entry_get_overwrite_mode(v.native())
	return GoBool(c)
}

// GetLayout is a wrapper around gtk_entry_get_layout().
//func (v *Entry) GetLayout() *pango.Layout {
//	c := C.gtk_entry_get_layout(v.native())
//	return pango.WrapLayout(uintptr(unsafe.Pointer(c)))
//}

// GetLayoutOffsets is a wrapper around gtk_entry_get_layout_offsets().
func (v *Entry) GetLayoutOffsets() (x, y int) {
	var gx, gy C.gint
	C.gtk_entry_get_layout_offsets(v.native(), &gx, &gy)
	return int(gx), int(gy)
}

// LayoutIndexToTextIndex is a wrapper around gtk_entry_layout_index_to_text_index().
func (v *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	c := C.gtk_entry_layout_index_to_text_index(v.native(),
		C.gint(layoutIndex))
	return int(c)
}

// TextIndexToLayoutIndex is a wrapper around gtk_entry_text_index_to_layout_index().
func (v *Entry) TextIndexToLayoutIndex(textIndex int) int {
	c := C.gtk_entry_text_index_to_layout_index(v.native(),
		C.gint(textIndex))
	return int(c)
}

// GetMaxLength is a wrapper around gtk_entry_get_max_length().
func (v *Entry) GetMaxLength() int {
	c := C.gtk_entry_get_max_length(v.native())
	return int(c)
}

// GetVisibility is a wrapper around gtk_entry_get_visibility().
func (v *Entry) GetVisibility() bool {
	c := C.gtk_entry_get_visibility(v.native())
	return GoBool(c)
}

// SetCompletion is a wrapper around gtk_entry_set_completion().
func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(v.native(), completion.native())
}

// GetCompletion is a wrapper around gtk_entry_get_completion().
func (v *Entry) GetCompletion() *EntryCompletion {
	c := C.gtk_entry_get_completion(v.native())
	if c == nil {
		return nil
	}

	e := &EntryCompletion{ToGoObject(unsafe.Pointer(c))}
	return e
}

// SetCursorHAdjustment is a wrapper around gtk_entry_set_cursor_hadjustment().
func (v *Entry) SetCursorHAdjustment(adjustment *Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(v.native(), adjustment.native())
}

// GetCursorHAdjustment is a wrapper around gtk_entry_get_cursor_hadjustment().
func (v *Entry) GetCursorHAdjustment() *Adjustment {
	c := C.gtk_entry_get_cursor_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return &Adjustment{InitiallyUnowned{obj}}
}

// SetProgressFraction is a wrapper around gtk_entry_set_progress_fraction().
func (v *Entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(v.native(), C.gdouble(fraction))
}

// GetProgressFraction is a wrapper around gtk_entry_get_progress_fraction().
func (v *Entry) GetProgressFraction() float64 {
	c := C.gtk_entry_get_progress_fraction(v.native())
	return float64(c)
}

// SetProgressPulseStep is a wrapper around gtk_entry_set_progress_pulse_step().
func (v *Entry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(v.native(), C.gdouble(fraction))
}

// GetProgressPulseStep is a wrapper around gtk_entry_get_progress_pulse_step().
func (v *Entry) GetProgressPulseStep() float64 {
	c := C.gtk_entry_get_progress_pulse_step(v.native())
	return float64(c)
}

// ProgressPulse is a wrapper around gtk_entry_progress_pulse().
func (v *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(v.native())
}

// ResetIMContext is a wrapper around gtk_entry_reset_im_context().
func (v *Entry) ResetIMContext() {
	C.gtk_entry_reset_im_context(v.native())
}

// SetIconFromIconName is a wrapper around gtk_entry_set_icon_from_icon_name().
func (v *Entry) SetIconFromIconName(iconPos EntryIconPosition, name string) {
	var icon *C.gchar
	if name != "" {
		n := C.CString(name)
		defer C.free(unsafe.Pointer(n))
		icon = (*C.gchar)(n)
	}
	C.gtk_entry_set_icon_from_icon_name(v.native(), C.GtkEntryIconPosition(iconPos), icon)
}

// RemoveIcon is a convenience func to set a nil pointer to the icon name.
func (v *Entry) RemoveIcon(iconPos EntryIconPosition) {
	C.gtk_entry_set_icon_from_icon_name(v.native(), C.GtkEntryIconPosition(iconPos), nil)
}

// TODO: Needs gio/GIcon implemented first
// SetIconFromGIcon is a wrapper around gtk_entry_set_icon_from_gicon().
func (v *Entry) SetIconFromGIcon(iconPos EntryIconPosition, icon *Icon) {
	C.gtk_entry_set_icon_from_gicon(v.native(), C.GtkEntryIconPosition(iconPos), (*C.GIcon)(icon.NativePrivate()))
}

// GetIconStorageType is a wrapper around gtk_entry_get_icon_storage_type().
func (v *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	c := C.gtk_entry_get_icon_storage_type(v.native(), C.GtkEntryIconPosition(iconPos))
	return ImageType(c)
}

// GetIconName is a wrapper around gtk_entry_get_icon_name().
func (v *Entry) GetIconName(iconPos EntryIconPosition) string {
	c := C.gtk_entry_get_icon_name(v.native(), C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return ""
	}
	return GoString(c)
}

// GetIconGIcon is a wrapper around gtk_entry_get_icon_gicon().
func (v *Entry) GetIconGIcon(iconPos EntryIconPosition) *Icon {
	c := C.gtk_entry_get_icon_gicon(v.native(), C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return nil
	}
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	i := &Icon{obj}
	return i
}

// SetIconActivatable is a wrapper around gtk_entry_set_icon_activatable().
func (v *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	C.gtk_entry_set_icon_activatable(v.native(), C.GtkEntryIconPosition(iconPos), CBool(activatable))
}

// GetIconActivatable is a wrapper around gtk_entry_get_icon_activatable().
func (v *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_activatable(v.native(), C.GtkEntryIconPosition(iconPos))
	return GoBool(c)
}

// SetIconSensitive is a wrapper around gtk_entry_set_icon_sensitive().
func (v *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	C.gtk_entry_set_icon_sensitive(v.native(), C.GtkEntryIconPosition(iconPos), CBool(sensitive))
}

// GetIconSensitive is a wrapper around gtk_entry_get_icon_sensitive().
func (v *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_sensitive(v.native(), C.GtkEntryIconPosition(iconPos))
	return GoBool(c)
}

// GetIconAtPos is a wrapper around gtk_entry_get_icon_at_pos().
func (v *Entry) GetIconAtPos(x, y int) int {
	c := C.gtk_entry_get_icon_at_pos(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetIconTooltipText is a wrapper around gtk_entry_set_icon_tooltip_text().
func (v *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	var text *C.gchar
	if tooltip != "" {
		cstr := C.CString(tooltip)
		defer C.free(unsafe.Pointer(cstr))
		text = cstr
	}

	C.gtk_entry_set_icon_tooltip_text(v.native(), C.GtkEntryIconPosition(iconPos), text)
}

// GetIconTooltipText is a wrapper around gtk_entry_get_icon_tooltip_text().
func (v *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	c := C.gtk_entry_get_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return ""
	}
	return GoString(c)
}

// SetIconTooltipMarkup is a wrapper around gtk_entry_set_icon_tooltip_markup().
func (v *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	var text *C.gchar
	if tooltip != "" {
		cstr := C.CString(tooltip)
		defer C.free(unsafe.Pointer(cstr))
		text = cstr
	}

	C.gtk_entry_set_icon_tooltip_markup(v.native(), C.GtkEntryIconPosition(iconPos), text)
}

// GetIconTooltipMarkup is a wrapper around gtk_entry_get_icon_tooltip_markup().
func (v *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	c := C.gtk_entry_get_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return ""
	}
	return GoString(c)
}

// GetCurrentIconDragSource is a wrapper around gtk_entry_get_current_icon_drag_source().
func (v *Entry) GetCurrentIconDragSource() int {
	c := C.gtk_entry_get_current_icon_drag_source(v.native())
	return int(c)
}

func (v *Entry) SetOnChanged(fn TTextChangedEvent) *SignalHandler {
	return registerAction(v, EsnChanged, MakeTextChangedEvent(fn))
}

func (v *Entry) SetOnCommit(fn TTextCommitEvent) *SignalHandler {
	return registerAction(v, EsnActivate, MakeTextCommitEvent(fn))
}

func (v *Entry) SetOnKeyPress(fn TTextKeyEvent) *SignalHandler {
	return registerAction(v, EsnKeyPressEvent, MakeTextKeyEvent(fn))
}

func (v *Entry) SetOnKeyRelease(fn TTextKeyEvent) *SignalHandler {
	return registerAction(v, EsnKeyReleaseEvent, MakeTextKeyEvent(fn))
}
