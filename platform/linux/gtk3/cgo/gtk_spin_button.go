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
