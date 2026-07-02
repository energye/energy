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

// SpinButton is a representation of GTK's GtkSpinButton.
type SpinButton struct {
	Entry
}

func AsSpinButton(ptr unsafe.Pointer) *SpinButton {
	if ptr == nil {
		return nil
	}
	m := new(SpinButton)
	m.instance = ptr
	return m
}

// NewSpinButton is a wrapper around gtk_spin_button_new().
func NewSpinButton(adjustment IAdjustment, climbRate float64, digits uint) *SpinButton {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	// Encode climbRate as uint64 bits for SysCall float64 passing
	r := gtk3.SysCall("gtk_spin_button_new", adj,
		uintptr(uint64(climbRate)), uintptr(digits))
	if r == 0 {
		return nil
	}
	return AsSpinButton(unsafe.Pointer(r))
}

// GetValue is a wrapper around gtk_spin_button_get_value().
func (m *SpinButton) GetValue() float64 {
	r := gtk3.SysCall("gtk_spin_button_get_value", m.Instance())
	return float64(r)
}

// SetValue is a wrapper around gtk_spin_button_set_value().
func (m *SpinButton) SetValue(value float64) {
	gtk3.SysCall("gtk_spin_button_set_value", m.Instance(), uintptr(uint64(value)))
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (m *SpinButton) SetRange(min, max float64) {
	gtk3.SysCall("gtk_spin_button_set_range", m.Instance(),
		uintptr(uint64(min)), uintptr(uint64(max)))
}

// SetIncrements is a wrapper around gtk_spin_button_set_increments().
func (m *SpinButton) SetIncrements(step, page float64) {
	gtk3.SysCall("gtk_spin_button_set_increments", m.Instance(),
		uintptr(uint64(step)), uintptr(uint64(page)))
}

// SetDigits is a wrapper around gtk_spin_button_set_digits().
func (m *SpinButton) SetDigits(digits uint) {
	gtk3.SysCall("gtk_spin_button_set_digits", m.Instance(), uintptr(digits))
}
