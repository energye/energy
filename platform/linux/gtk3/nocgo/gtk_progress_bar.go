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

import "unsafe"

// ProgressBar is a representation of GTK's GtkProgressBar.
type ProgressBar struct {
	Widget
}

func AsProgressBar(ptr unsafe.Pointer) *ProgressBar {
	if ptr == nil {
		return nil
	}
	m := new(ProgressBar)
	m.instance = ptr
	return m
}

// NewProgressBar is a wrapper around gtk_progress_bar_new().
func NewProgressBar() *ProgressBar {
	r := gtk3.SysCall("gtk_progress_bar_new")
	if r == 0 {
		return nil
	}
	return AsProgressBar(unsafe.Pointer(r))
}

// SetFraction is a wrapper around gtk_progress_bar_set_fraction().
func (m *ProgressBar) SetFraction(fraction float64) {
	registerGtkFloatFuncs()
	gtkProgressBarSetFraction(m.Instance(), fraction)
}

// GetFraction is a wrapper around gtk_progress_bar_get_fraction().
func (m *ProgressBar) GetFraction() float64 {
	registerGtkFloatFuncs()
	return gtkProgressBarGetFraction(m.Instance())
}

// Pulse is a wrapper around gtk_progress_bar_pulse().
func (m *ProgressBar) Pulse() {
	gtk3.SysCall("gtk_progress_bar_pulse", m.Instance())
}

// SetText is a wrapper around gtk_progress_bar_set_text().
func (m *ProgressBar) SetText(text string) {
	gtk3.SysCall("gtk_progress_bar_set_text", m.Instance(), CStr(text))
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (m *ProgressBar) SetShowText(showText bool) {
	gtk3.SysCall("gtk_progress_bar_set_show_text", m.Instance(), ToCBool(showText))
}

// SetPulseStep is a wrapper around gtk_progress_bar_set_pulse_step().
func (m *ProgressBar) SetPulseStep(fraction float64) {
	registerGtkFloatFuncs()
	gtkProgressBarSetPulseStep(m.Instance(), fraction)
}
