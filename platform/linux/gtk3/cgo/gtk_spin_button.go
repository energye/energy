package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// SpinButton is a representation of GTK's GtkSpinButton.
type SpinButton struct {
	Entry
}

func (v *SpinButton) native() *C.GtkSpinButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkSpinButton(unsafe.Pointer(v.GObject))
}

func wrapSpinButton(obj *Object) *SpinButton {
	return &SpinButton{Entry{Widget{InitiallyUnowned{obj}}, Editable{obj}, CellEditable{InitiallyUnowned{obj}}}}
}

// NewSpinButton is a wrapper around gtk_spin_button_new().
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint) *SpinButton {
	var adj *C.GtkAdjustment
	if adjustment != nil {
		adj = adjustment.native()
	}
	c := C.gtk_spin_button_new(adj, C.gdouble(climbRate), C.guint(digits))
	if c == nil {
		return nil
	}
	return wrapSpinButton(ToGoObject(unsafe.Pointer(c)))
}

// GetValue is a wrapper around gtk_spin_button_get_value().
func (v *SpinButton) GetValue() float64 {
	return float64(C.gtk_spin_button_get_value(v.native()))
}

// SetValue is a wrapper around gtk_spin_button_set_value().
func (v *SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(v.native(), C.gdouble(value))
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// SetIncrements is a wrapper around gtk_spin_button_set_increments().
func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

// SetDigits is a wrapper around gtk_spin_button_set_digits().
func (v *SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(v.native(), C.guint(digits))
}

func (m *SpinButton) SetOnValueChanged(fn TValueChangedEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnValueChanged, callback.C_trampoline_2_void, fn, 0)
}

// GetAdjustment is a wrapper around gtk_spin_button_get_adjustment().
func (v *SpinButton) GetAdjustment() IAdjustment {
	c := C.gtk_spin_button_get_adjustment(v.native())
	if c == nil {
		return nil
	}
	return wrapAdjustment(ToGoObject(unsafe.Pointer(c)))
}

// SetNumeric is a wrapper around gtk_spin_button_set_numeric().
func (v *SpinButton) SetNumeric(numeric bool) {
	C.gtk_spin_button_set_numeric(v.native(), CBool(numeric))
}

// GetNumeric is a wrapper around gtk_spin_button_get_numeric().
func (v *SpinButton) GetNumeric() bool {
	return GoBool(C.gtk_spin_button_get_numeric(v.native()))
}

// SetSnapToTicks is a wrapper around gtk_spin_button_set_snap_to_ticks().
func (v *SpinButton) SetSnapToTicks(snap bool) {
	C.gtk_spin_button_set_snap_to_ticks(v.native(), CBool(snap))
}

// GetSnapToTicks is a wrapper around gtk_spin_button_get_snap_to_ticks().
func (v *SpinButton) GetSnapToTicks() bool {
	return GoBool(C.gtk_spin_button_get_snap_to_ticks(v.native()))
}

// SetWrap is a wrapper around gtk_spin_button_set_wrap().
func (v *SpinButton) SetWrap(wrap bool) {
	C.gtk_spin_button_set_wrap(v.native(), CBool(wrap))
}

// GetWrap is a wrapper around gtk_spin_button_get_wrap().
func (v *SpinButton) GetWrap() bool {
	return GoBool(C.gtk_spin_button_get_wrap(v.native()))
}

// Spin is a wrapper around gtk_spin_button_spin().
func (v *SpinButton) Spin(direction SpinDirection, increment float64) {
	C.gtk_spin_button_spin(v.native(), C.GtkSpinType(direction), C.gdouble(increment))
}
