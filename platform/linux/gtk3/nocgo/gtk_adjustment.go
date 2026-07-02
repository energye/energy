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

// Adjustment is a representation of GTK's GtkAdjustment.
type Adjustment struct {
	Object
}

func AsAdjustment(ptr unsafe.Pointer) IAdjustment {
	if ptr == nil {
		return nil
	}
	m := new(Adjustment)
	m.instance = ptr
	return m
}

// NewAdjustment is a wrapper around gtk_adjustment_new().
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) IAdjustment {
	registerGtkFloatFuncs()
	if gtkAdjustmentNew == nil {
		return nil
	}
	r := gtkAdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize)
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// GetValue is a wrapper around gtk_adjustment_get_value().
func (m *Adjustment) GetValue() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetValue(m.Instance())
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (m *Adjustment) SetValue(value float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetValue(m.Instance(), value)
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (m *Adjustment) GetLower() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetLower(m.Instance())
}

// SetLower is a wrapper around gtk_adjustment_set_lower().
func (m *Adjustment) SetLower(lower float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetLower(m.Instance(), lower)
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (m *Adjustment) GetUpper() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetUpper(m.Instance())
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (m *Adjustment) SetUpper(upper float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetUpper(m.Instance(), upper)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (m *Adjustment) GetPageSize() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetPageSize(m.Instance())
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (m *Adjustment) SetPageSize(pageSize float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetPageSize(m.Instance(), pageSize)
}

// Configure is a wrapper around gtk_adjustment_configure().
func (m *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentConfigure(m.Instance(), value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

// GetPageIncrement is a wrapper around gtk_adjustment_get_page_increment().
func (m *Adjustment) GetPageIncrement() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetPageIncrement(m.Instance())
}

// SetPageIncrement is a wrapper around gtk_adjustment_set_page_increment().
func (m *Adjustment) SetPageIncrement(pageIncrement float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetPageIncrement(m.Instance(), pageIncrement)
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (m *Adjustment) GetStepIncrement() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetStepIncrement(m.Instance())
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (m *Adjustment) SetStepIncrement(stepIncrement float64) {
	registerGtkFloatFuncs()
	gtkAdjustmentSetStepIncrement(m.Instance(), stepIncrement)
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (m *Adjustment) GetMinimumIncrement() float64 {
	registerGtkFloatFuncs()
	return gtkAdjustmentGetMinimumIncrement(m.Instance())
}
