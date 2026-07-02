package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// ProgressBar is a representation of GTK's GtkProgressBar.
type ProgressBar struct {
	Widget
}

func (v *ProgressBar) native() *C.GtkProgressBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkProgressBar(p)
}

func wrapProgressBar(obj *Object) *ProgressBar {
	return &ProgressBar{Widget{InitiallyUnowned{obj}}}
}

// NewProgressBar is a wrapper around gtk_progress_bar_new().
func NewProgressBar() *ProgressBar {
	c := C.gtk_progress_bar_new()
	if c == nil {
		return nil
	}
	return wrapProgressBar(ToGoObject(unsafe.Pointer(c)))
}

// SetFraction is a wrapper around gtk_progress_bar_set_fraction().
func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(v.native(), C.gdouble(fraction))
}

// GetFraction is a wrapper around gtk_progress_bar_get_fraction().
func (v *ProgressBar) GetFraction() float64 {
	return float64(C.gtk_progress_bar_get_fraction(v.native()))
}

// Pulse is a wrapper around gtk_progress_bar_pulse().
func (v *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(v.native())
}

// SetText is a wrapper around gtk_progress_bar_set_text().
func (v *ProgressBar) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_progress_bar_set_text(v.native(), (*C.gchar)(cstr))
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (v *ProgressBar) SetShowText(showText bool) {
	C.gtk_progress_bar_set_show_text(v.native(), CBool(showText))
}

// SetPulseStep is a wrapper around gtk_progress_bar_set_pulse_step().
func (v *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(v.native(), C.gdouble(fraction))
}
