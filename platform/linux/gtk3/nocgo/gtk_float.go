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
	"sync"

	"github.com/ebitengine/purego"
)

var gtkFloatOnce sync.Once

var (
	gtkWidgetSetOpacity func(uintptr, float64)
	gtkWidgetGetOpacity func(uintptr) float64

	gtkLabelSetXAlign func(uintptr, float32)
	gtkLabelGetXAlign func(uintptr) float32
	gtkLabelSetYAlign func(uintptr, float32)
	gtkLabelGetYAlign func(uintptr) float32
	gtkLabelGetAngle  func(uintptr) float64
	gtkLabelSetAngle  func(uintptr, float64)

	gtkAdjustmentNew                 func(float64, float64, float64, float64, float64, float64) uintptr
	gtkAdjustmentGetValue            func(uintptr) float64
	gtkAdjustmentSetValue            func(uintptr, float64)
	gtkAdjustmentGetLower            func(uintptr) float64
	gtkAdjustmentSetLower            func(uintptr, float64)
	gtkAdjustmentGetUpper            func(uintptr) float64
	gtkAdjustmentSetUpper            func(uintptr, float64)
	gtkAdjustmentGetPageSize         func(uintptr) float64
	gtkAdjustmentSetPageSize         func(uintptr, float64)
	gtkAdjustmentGetPageIncrement    func(uintptr) float64
	gtkAdjustmentSetPageIncrement    func(uintptr, float64)
	gtkAdjustmentGetStepIncrement    func(uintptr) float64
	gtkAdjustmentSetStepIncrement    func(uintptr, float64)
	gtkAdjustmentGetMinimumIncrement func(uintptr) float64
	gtkAdjustmentConfigure           func(uintptr, float64, float64, float64, float64, float64, float64)

	gtkRangeGetValue      func(uintptr) float64
	gtkRangeSetValue      func(uintptr, float64)
	gtkRangeSetIncrements func(uintptr, float64, float64)
	gtkRangeSetRange      func(uintptr, float64, float64)
	gtkRangeGetFillLevel  func(uintptr) float64
	gtkRangeSetFillLevel  func(uintptr, float64)

	gtkEntrySetAlignment        func(uintptr, float32)
	gtkEntryGetAlignment        func(uintptr) float32
	gtkEntrySetProgressFraction func(uintptr, float64)
	gtkEntryGetProgressFraction func(uintptr) float64
	gtkEntrySetProgressPulseStep func(uintptr, float64)
	gtkEntryGetProgressPulseStep func(uintptr) float64

	gtkSpinButtonGetValue       func(uintptr) float64
	gtkSpinButtonSetValue       func(uintptr, float64)
	gtkSpinButtonSetRange       func(uintptr, float64, float64)
	gtkSpinButtonSetIncrements  func(uintptr, float64, float64)
	gtkSpinButtonNew            func(uintptr, float64, uintptr) uintptr
)

func registerGtkFloatFuncs() {
	gtkFloatOnce.Do(func() {
		if gtk3 == nil || gtk3.Dll == 0 {
			return
		}
		lib := uintptr(gtk3.Dll)

		purego.RegisterLibFunc(&gtkWidgetSetOpacity, lib, "gtk_widget_set_opacity")
		purego.RegisterLibFunc(&gtkWidgetGetOpacity, lib, "gtk_widget_get_opacity")

		purego.RegisterLibFunc(&gtkLabelSetXAlign, lib, "gtk_label_set_xalign")
		purego.RegisterLibFunc(&gtkLabelGetXAlign, lib, "gtk_label_get_xalign")
		purego.RegisterLibFunc(&gtkLabelSetYAlign, lib, "gtk_label_set_yalign")
		purego.RegisterLibFunc(&gtkLabelGetYAlign, lib, "gtk_label_get_yalign")
		purego.RegisterLibFunc(&gtkLabelGetAngle, lib, "gtk_label_get_angle")
		purego.RegisterLibFunc(&gtkLabelSetAngle, lib, "gtk_label_set_angle")

		purego.RegisterLibFunc(&gtkAdjustmentNew, lib, "gtk_adjustment_new")
		purego.RegisterLibFunc(&gtkAdjustmentGetValue, lib, "gtk_adjustment_get_value")
		purego.RegisterLibFunc(&gtkAdjustmentSetValue, lib, "gtk_adjustment_set_value")
		purego.RegisterLibFunc(&gtkAdjustmentGetLower, lib, "gtk_adjustment_get_lower")
		purego.RegisterLibFunc(&gtkAdjustmentSetLower, lib, "gtk_adjustment_set_lower")
		purego.RegisterLibFunc(&gtkAdjustmentGetUpper, lib, "gtk_adjustment_get_upper")
		purego.RegisterLibFunc(&gtkAdjustmentSetUpper, lib, "gtk_adjustment_set_upper")
		purego.RegisterLibFunc(&gtkAdjustmentGetPageSize, lib, "gtk_adjustment_get_page_size")
		purego.RegisterLibFunc(&gtkAdjustmentSetPageSize, lib, "gtk_adjustment_set_page_size")
		purego.RegisterLibFunc(&gtkAdjustmentGetPageIncrement, lib, "gtk_adjustment_get_page_increment")
		purego.RegisterLibFunc(&gtkAdjustmentSetPageIncrement, lib, "gtk_adjustment_set_page_increment")
		purego.RegisterLibFunc(&gtkAdjustmentGetStepIncrement, lib, "gtk_adjustment_get_step_increment")
		purego.RegisterLibFunc(&gtkAdjustmentSetStepIncrement, lib, "gtk_adjustment_set_step_increment")
		purego.RegisterLibFunc(&gtkAdjustmentGetMinimumIncrement, lib, "gtk_adjustment_get_minimum_increment")
		purego.RegisterLibFunc(&gtkAdjustmentConfigure, lib, "gtk_adjustment_configure")

		purego.RegisterLibFunc(&gtkRangeGetValue, lib, "gtk_range_get_value")
		purego.RegisterLibFunc(&gtkRangeSetValue, lib, "gtk_range_set_value")
		purego.RegisterLibFunc(&gtkRangeSetIncrements, lib, "gtk_range_set_increments")
		purego.RegisterLibFunc(&gtkRangeSetRange, lib, "gtk_range_set_range")
		purego.RegisterLibFunc(&gtkRangeGetFillLevel, lib, "gtk_range_get_fill_level")
		purego.RegisterLibFunc(&gtkRangeSetFillLevel, lib, "gtk_range_set_fill_level")

		purego.RegisterLibFunc(&gtkEntrySetAlignment, lib, "gtk_entry_set_alignment")
		purego.RegisterLibFunc(&gtkEntryGetAlignment, lib, "gtk_entry_get_alignment")
		purego.RegisterLibFunc(&gtkEntrySetProgressFraction, lib, "gtk_entry_set_progress_fraction")
		purego.RegisterLibFunc(&gtkEntryGetProgressFraction, lib, "gtk_entry_get_progress_fraction")
		purego.RegisterLibFunc(&gtkEntrySetProgressPulseStep, lib, "gtk_entry_set_progress_pulse_step")
		purego.RegisterLibFunc(&gtkEntryGetProgressPulseStep, lib, "gtk_entry_get_progress_pulse_step")

		purego.RegisterLibFunc(&gtkSpinButtonGetValue, lib, "gtk_spin_button_get_value")
		purego.RegisterLibFunc(&gtkSpinButtonSetValue, lib, "gtk_spin_button_set_value")
		purego.RegisterLibFunc(&gtkSpinButtonSetRange, lib, "gtk_spin_button_set_range")
		purego.RegisterLibFunc(&gtkSpinButtonSetIncrements, lib, "gtk_spin_button_set_increments")
		purego.RegisterLibFunc(&gtkSpinButtonNew, lib, "gtk_spin_button_new")
	})
}
