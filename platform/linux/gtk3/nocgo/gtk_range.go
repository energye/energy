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

// Range is a representation of GTK's GtkRange.
type Range struct {
	Widget
}

func AsRange(ptr unsafe.Pointer) IRange {
	if ptr == nil {
		return nil
	}
	m := new(Range)
	m.instance = ptr
	return m
}

// GetFillLevel is a wrapper around gtk_range_get_fill_level().
func (m *Range) GetFillLevel() float64 {
	registerGtkFloatFuncs()
	return gtkRangeGetFillLevel(m.Instance())
}

// GetRestrictToFillLevel is a wrapper around gtk_range_get_restrict_to_fill_level().
func (m *Range) GetRestrictToFillLevel() bool {
	r := gtk3.SysCall("gtk_range_get_restrict_to_fill_level", m.Instance())
	return ToGoBool(r)
}

// GetShowFillLevel is a wrapper around gtk_range_get_show_fill_level().
func (m *Range) GetShowFillLevel() bool {
	r := gtk3.SysCall("gtk_range_get_show_fill_level", m.Instance())
	return ToGoBool(r)
}

// SetFillLevel is a wrapper around gtk_range_set_fill_level().
func (m *Range) SetFillLevel(fillLevel float64) {
	registerGtkFloatFuncs()
	gtkRangeSetFillLevel(m.Instance(), fillLevel)
}

// RestrictToFillLevel is a wrapper around gtk_range_set_restrict_to_fill_level().
func (m *Range) RestrictToFillLevel(restrictToFillLevel bool) {
	gtk3.SysCall("gtk_range_set_restrict_to_fill_level", m.Instance(), ToCBool(restrictToFillLevel))
}

// SetShowFillLevel is a wrapper around gtk_range_set_show_fill_level().
func (m *Range) SetShowFillLevel(showFillLevel bool) {
	gtk3.SysCall("gtk_range_set_show_fill_level", m.Instance(), ToCBool(showFillLevel))
}

// GetAdjustment is a wrapper around gtk_range_get_adjustment().
func (m *Range) GetAdjustment() IAdjustment {
	r := gtk3.SysCall("gtk_range_get_adjustment", m.Instance())
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// SetAdjustment is a wrapper around gtk_range_set_adjustment().
func (m *Range) SetAdjustment(adj IAdjustment) {
	var adjPtr uintptr
	if adj != nil {
		adjPtr = adj.Instance()
	}
	gtk3.SysCall("gtk_range_set_adjustment", m.Instance(), adjPtr)
}

// GetValue is a wrapper around gtk_range_get_value().
func (m *Range) GetValue() float64 {
	registerGtkFloatFuncs()
	return gtkRangeGetValue(m.Instance())
}

// SetValue is a wrapper around gtk_range_set_value().
func (m *Range) SetValue(value float64) {
	registerGtkFloatFuncs()
	gtkRangeSetValue(m.Instance(), value)
}

// SetIncrements is a wrapper around gtk_range_set_increments().
func (m *Range) SetIncrements(step, page float64) {
	registerGtkFloatFuncs()
	gtkRangeSetIncrements(m.Instance(), step, page)
}

// SetRange is a wrapper around gtk_range_set_range().
func (m *Range) SetRange(min, max float64) {
	registerGtkFloatFuncs()
	gtkRangeSetRange(m.Instance(), min, max)
}

// GetInverted is a wrapper around gtk_range_get_inverted().
func (m *Range) GetInverted() bool {
	r := gtk3.SysCall("gtk_range_get_inverted", m.Instance())
	return ToGoBool(r)
}

// SetInverted is a wrapper around gtk_range_set_inverted().
func (m *Range) SetInverted(setting bool) {
	gtk3.SysCall("gtk_range_set_inverted", m.Instance(), ToCBool(setting))
}

// GetRoundDigits is a wrapper around gtk_range_get_round_digits().
func (m *Range) GetRoundDigits() int {
	r := gtk3.SysCall("gtk_range_get_round_digits", m.Instance())
	return int(r)
}

// SetRoundDigits is a wrapper around gtk_range_set_round_digits().
func (m *Range) SetRoundDigits(roundDigits int) {
	gtk3.SysCall("gtk_range_set_round_digits", m.Instance(), uintptr(roundDigits))
}

// SetLowerStepperSensitivity is a wrapper around gtk_range_set_lower_stepper_sensitivity().
func (m *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	gtk3.SysCall("gtk_range_set_lower_stepper_sensitivity", m.Instance(), uintptr(sensitivity))
}

// GetLowerStepperSensitivity is a wrapper around gtk_range_get_lower_stepper_sensitivity().
func (m *Range) GetLowerStepperSensitivity() SensitivityType {
	r := gtk3.SysCall("gtk_range_get_lower_stepper_sensitivity", m.Instance())
	return SensitivityType(r)
}

// SetUpperStepperSensitivity is a wrapper around gtk_range_set_upper_stepper_sensitivity().
func (m *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	gtk3.SysCall("gtk_range_set_upper_stepper_sensitivity", m.Instance(), uintptr(sensitivity))
}

// GetUpperStepperSensitivity is a wrapper around gtk_range_get_upper_stepper_sensitivity().
func (m *Range) GetUpperStepperSensitivity() SensitivityType {
	r := gtk3.SysCall("gtk_range_get_upper_stepper_sensitivity", m.Instance())
	return SensitivityType(r)
}

// GetFlippable is a wrapper around gtk_range_get_flippable().
func (m *Range) GetFlippable() bool {
	r := gtk3.SysCall("gtk_range_get_flippable", m.Instance())
	return ToGoBool(r)
}

// SetFlippable is a wrapper around gtk_range_set_flippable().
func (m *Range) SetFlippable(flippable bool) {
	gtk3.SysCall("gtk_range_set_flippable", m.Instance(), ToCBool(flippable))
}

// GetRangeRect is a wrapper around gtk_range_get_range_rect().
func (m *Range) GetRangeRect() *Rectangle {
	var rect Rectangle
	gtk3.SysCall("gtk_range_get_range_rect", m.Instance(), uintptr(unsafe.Pointer(&rect)))
	return &rect
}

// GetSliderRange is a wrapper around gtk_range_get_slider_range().
func (m *Range) GetSliderRange() (int, int) {
	var cStart, cEnd int32
	gtk3.SysCall("gtk_range_get_slider_range", m.Instance(), uintptr(unsafe.Pointer(&cStart)), uintptr(unsafe.Pointer(&cEnd)))
	return int(cStart), int(cEnd)
}

// GetSliderFixedSize is a wrapper around gtk_range_get_slider_size_fixed().
func (m *Range) GetSliderFixedSize() bool {
	r := gtk3.SysCall("gtk_range_get_slider_size_fixed", m.Instance())
	return ToGoBool(r)
}

// SetSliderFixedSize is a wrapper around gtk_range_set_slider_size_fixed().
func (m *Range) SetSliderFixedSize(sizeFixed bool) {
	gtk3.SysCall("gtk_range_set_slider_size_fixed", m.Instance(), ToCBool(sizeFixed))
}

// SetOnValueChanged is a callback for the "value-changed" signal.
func (m *Range) SetOnValueChanged(fn TValueChangedEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnValueChanged, callback.C_trampoline_2_void, fn, 0)
}
