package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// Adjustment is a representation of GTK's GtkAdjustment.
type Adjustment struct {
	InitiallyUnowned
}

// native returns a pointer to the underlying GtkAdjustment.
func (v *Adjustment) native() *C.GtkAdjustment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAdjustment(p)
}

func marshalAdjustment(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

func wrapAdjustment(obj *Object) *Adjustment {
	if obj == nil {
		return nil
	}

	return &Adjustment{InitiallyUnowned{obj}}
}

// NewAdjustment is a wrapper around gtk_adjustment_new().
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) (*Adjustment, error) {
	c := C.gtk_adjustment_new(C.gdouble(value),
		C.gdouble(lower),
		C.gdouble(upper),
		C.gdouble(stepIncrement),
		C.gdouble(pageIncrement),
		C.gdouble(pageSize))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

// GetValue is a wrapper around gtk_adjustment_get_value().
func (v *Adjustment) GetValue() float64 {
	c := C.gtk_adjustment_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (v *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.native(), C.gdouble(value))
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (v *Adjustment) GetLower() float64 {
	c := C.gtk_adjustment_get_lower(v.native())
	return float64(c)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (v *Adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(v.native()))
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (v *Adjustment) SetPageSize(value float64) {
	C.gtk_adjustment_set_page_size(v.native(), C.gdouble(value))
}

// Configure is a wrapper around gtk_adjustment_configure().
func (v *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(v.native(), C.gdouble(value),
		C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
}

// SetLower is a wrapper around gtk_adjustment_set_lower().
func (v *Adjustment) SetLower(value float64) {
	C.gtk_adjustment_set_lower(v.native(), C.gdouble(value))
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (v *Adjustment) GetUpper() float64 {
	c := C.gtk_adjustment_get_upper(v.native())
	return float64(c)
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (v *Adjustment) SetUpper(value float64) {
	C.gtk_adjustment_set_upper(v.native(), C.gdouble(value))
}

// GetPageIncrement is a wrapper around gtk_adjustment_get_page_increment().
func (v *Adjustment) GetPageIncrement() float64 {
	c := C.gtk_adjustment_get_page_increment(v.native())
	return float64(c)
}

// SetPageIncrement is a wrapper around gtk_adjustment_set_page_increment().
func (v *Adjustment) SetPageIncrement(value float64) {
	C.gtk_adjustment_set_page_increment(v.native(), C.gdouble(value))
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (v *Adjustment) GetStepIncrement() float64 {
	c := C.gtk_adjustment_get_step_increment(v.native())
	return float64(c)
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (v *Adjustment) SetStepIncrement(value float64) {
	C.gtk_adjustment_set_step_increment(v.native(), C.gdouble(value))
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (v *Adjustment) GetMinimumIncrement() float64 {
	c := C.gtk_adjustment_get_minimum_increment(v.native())
	return float64(c)
}
