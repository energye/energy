package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// Range is a representation of GTK's GtkRange.
type Range struct {
	Widget
}

// native returns a pointer to the underlying GtkRange.
func (v *Range) native() *C.GtkRange {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRange(p)
}

func marshalRange(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapRange(obj), nil
}

func wrapRange(obj *Object) *Range {
	if obj == nil {
		return nil
	}
	return &Range{Widget{InitiallyUnowned{obj}}}
}

// GetFillLevel is a wrapper around gtk_range_get_fill_level().
func (v *Range) GetFillLevel() float64 {
	return float64(C.gtk_range_get_fill_level(v.native()))
}

// GetRestrictToFillLevel is a wrapper around gtk_range_get_restrict_to_fill_level().
func (v *Range) GetRestrictToFillLevel() bool {
	return GoBool(C.gtk_range_get_restrict_to_fill_level(v.native()))
}

// GetShowFillLevel is a wrapper around gtk_range_get_show_fill_level().
func (v *Range) GetShowFillLevel() bool {
	return GoBool(C.gtk_range_get_show_fill_level(v.native()))
}

// SetFillLevel is a wrapper around gtk_range_set_fill_level().
func (v *Range) SetFillLevel(fill_level float64) {
	C.gtk_range_set_fill_level(v.native(), C.gdouble(fill_level))
}

// RestrictToFillLevel is a wrapper around gtk_range_set_restrict_to_fill_level().
func (v *Range) RestrictToFillLevel(restrict_to_fill_level bool) {
	C.gtk_range_set_restrict_to_fill_level(v.native(), CBool(restrict_to_fill_level))
}

// SetShowFillLevel is a wrapper around gtk_range_set_show_fill_level().
func (v *Range) SetShowFillLevel(show_fill_level bool) {
	C.gtk_range_set_show_fill_level(v.native(), CBool(show_fill_level))
}

// GetAdjustment is a wrapper around gtk_range_get_adjustment().
func (v *Range) GetAdjustment() *Adjustment {

	c := C.gtk_range_get_adjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetAdjustment is a wrapper around gtk_range_set_adjustment().
func (v *Range) SetAdjustment(adjustment *Adjustment) {
	C.gtk_range_set_adjustment(v.native(), adjustment.native())
}

// GetValue is a wrapper around gtk_range_get_value().
func (v *Range) GetValue() float64 {
	c := C.gtk_range_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_range_set_value().
func (v *Range) SetValue(value float64) {
	C.gtk_range_set_value(v.native(), C.gdouble(value))
}

// SetIncrements is a wrapper around gtk_range_set_increments().
func (v *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

// SetRange is a wrapper around gtk_range_set_range().
func (v *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// GetInverted is a wrapper around gtk_range_get_inverted().
func (v *Range) GetInverted() bool {
	c := C.gtk_range_get_inverted(v.native())
	return GoBool(c)
}

// SetInverted is a wrapper around gtk_range_set_inverted().
func (v *Range) SetInverted(inverted bool) {
	C.gtk_range_set_inverted(v.native(), CBool(inverted))
}

// GetRoundDigits is a wrapper around gtk_range_get_round_digits().
func (v *Range) GetRoundDigits() int {
	return int(C.gtk_range_get_round_digits(v.native()))
}

// SetRoundDigits is a wrapper around gtk_range_set_round_digits().
func (v *Range) SetRoundDigits(round_digits int) {
	C.gtk_range_set_round_digits(v.native(), C.gint(round_digits))
}

// SetLowerStepperSensitivity is a wrapper around gtk_range_set_lower_stepper_sensitivity().
func (v *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	C.gtk_range_set_lower_stepper_sensitivity(
		v.native(),
		C.GtkSensitivityType(sensitivity))
}

// GetLowerStepperSensitivity is a wrapper around gtk_range_get_lower_stepper_sensitivity().
func (v *Range) GetLowerStepperSensitivity() SensitivityType {
	return SensitivityType(C.gtk_range_get_lower_stepper_sensitivity(
		v.native()))
}

// SetUpperStepperSensitivity is a wrapper around gtk_range_set_upper_stepper_sensitivity().
func (v *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	C.gtk_range_set_upper_stepper_sensitivity(
		v.native(),
		C.GtkSensitivityType(sensitivity))
}

// GetUpperStepperSensitivity is a wrapper around gtk_range_get_upper_stepper_sensitivity().
func (v *Range) GetUpperStepperSensitivity() SensitivityType {
	return SensitivityType(C.gtk_range_get_upper_stepper_sensitivity(
		v.native()))
}

// GetFlippable is a wrapper around gtk_range_get_flippable().
func (v *Range) GetFlippable() bool {
	return GoBool(C.gtk_range_get_flippable(v.native()))
}

// SetFlippable is a wrapper around gtk_range_set_flippable().
func (v *Range) SetFlippable(flippable bool) {
	C.gtk_range_set_flippable(v.native(), CBool(flippable))
}

// GetRangeRect is a wrapper around gtk_range_get_range_rect().
func (v *Range) GetRangeRect() *Rectangle {
	var cRect *C.GdkRectangle
	C.gtk_range_get_range_rect(v.native(), cRect)
	return WrapRectangle(uintptr(unsafe.Pointer(cRect)))
}

// GetSliderRange is a wrapper around gtk_range_get_slider_range().
func (v *Range) GetSliderRange() (int, int) {
	var cStart, cEnd C.gint
	C.gtk_range_get_slider_range(v.native(), &cStart, &cEnd)
	return int(cStart), int(cEnd)
}

// GetSliderFixedSize is a wrapper gtk_range_get_slider_size_fixed().
func (v *Range) GetSliderFixedSize() bool {
	return GoBool(C.gtk_range_get_slider_size_fixed(v.native()))
}

// SetSliderFixedSize is a wrapper around gtk_range_set_slider_size_fixed().
func (v *Range) SetSliderFixedSize(size_fixed bool) {
	C.gtk_range_set_slider_size_fixed(v.native(), CBool(size_fixed))
}
