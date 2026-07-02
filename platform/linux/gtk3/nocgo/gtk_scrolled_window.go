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

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

func AsScrolledWindow(ptr unsafe.Pointer) IScrolledWindow {
	if ptr == nil {
		return nil
	}
	m := new(ScrolledWindow)
	m.instance = ptr
	return m
}

// NewScrolledWindow is a wrapper around gtk_scrolled_window_new().
func NewScrolledWindow(hadjustment, vadjustment IAdjustment) IScrolledWindow {
	var hadj, vadj uintptr
	if hadjustment != nil {
		hadj = hadjustment.Instance()
	}
	if vadjustment != nil {
		vadj = vadjustment.Instance()
	}
	r := gtk3.SysCall("gtk_scrolled_window_new", hadj, vadj)
	if r == 0 {
		return nil
	}
	return AsScrolledWindow(unsafe.Pointer(r))
}

// GetPolicy is a wrapper around gtk_scrolled_window_get_policy().
func (m *ScrolledWindow) GetPolicy() (hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	var hScrPol, vScrPol int32
	gtk3.SysCall("gtk_scrolled_window_get_policy", m.Instance(),
		uintptr(unsafe.Pointer(&hScrPol)), uintptr(unsafe.Pointer(&vScrPol)))
	hScrollbarPolicy = PolicyType(hScrPol)
	vScrollbarPolicy = PolicyType(vScrPol)
	return
}

// SetPolicy is a wrapper around gtk_scrolled_window_set_policy().
func (m *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	gtk3.SysCall("gtk_scrolled_window_set_policy", m.Instance(),
		uintptr(hScrollbarPolicy), uintptr(vScrollbarPolicy))
}

// GetHAdjustment is a wrapper around gtk_scrolled_window_get_hadjustment().
func (m *ScrolledWindow) GetHAdjustment() IAdjustment {
	r := gtk3.SysCall("gtk_scrolled_window_get_hadjustment", m.Instance())
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// SetHAdjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (m *ScrolledWindow) SetHAdjustment(adjustment IAdjustment) {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	gtk3.SysCall("gtk_scrolled_window_set_hadjustment", m.Instance(), adj)
}

// GetVAdjustment is a wrapper around gtk_scrolled_window_get_vadjustment().
func (m *ScrolledWindow) GetVAdjustment() IAdjustment {
	r := gtk3.SysCall("gtk_scrolled_window_get_vadjustment", m.Instance())
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// SetVAdjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (m *ScrolledWindow) SetVAdjustment(adjustment IAdjustment) {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	gtk3.SysCall("gtk_scrolled_window_set_vadjustment", m.Instance(), adj)
}

// GetHScrollbar is a wrapper around gtk_scrolled_window_get_hscrollbar().
func (m *ScrolledWindow) GetHScrollbar() IScrollbar {
	r := gtk3.SysCall("gtk_scrolled_window_get_hscrollbar", m.Instance())
	if r == 0 {
		return nil
	}
	return AsScrollbar(unsafe.Pointer(r))
}

// GetVScrollbar is a wrapper around gtk_scrolled_window_get_vscrollbar().
func (m *ScrolledWindow) GetVScrollbar() IScrollbar {
	r := gtk3.SysCall("gtk_scrolled_window_get_vscrollbar", m.Instance())
	if r == 0 {
		return nil
	}
	return AsScrollbar(unsafe.Pointer(r))
}

// GetPlacement is a wrapper around gtk_scrolled_window_get_placement().
func (m *ScrolledWindow) GetPlacement() CornerType {
	r := gtk3.SysCall("gtk_scrolled_window_get_placement", m.Instance())
	return CornerType(r)
}

// SetPlacement is a wrapper around gtk_scrolled_window_set_placement().
func (m *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	gtk3.SysCall("gtk_scrolled_window_set_placement", m.Instance(), uintptr(windowPlacement))
}

// UnsetPlacement is a wrapper around gtk_scrolled_window_unset_placement().
func (m *ScrolledWindow) UnsetPlacement() {
	gtk3.SysCall("gtk_scrolled_window_unset_placement", m.Instance())
}

// GetShadowType is a wrapper around gtk_scrolled_window_get_shadow_type().
func (m *ScrolledWindow) GetShadowType() ShadowType {
	r := gtk3.SysCall("gtk_scrolled_window_get_shadow_type", m.Instance())
	return ShadowType(r)
}

// SetShadowType is a wrapper around gtk_scrolled_window_set_shadow_type().
func (m *ScrolledWindow) SetShadowType(t ShadowType) {
	gtk3.SysCall("gtk_scrolled_window_set_shadow_type", m.Instance(), uintptr(t))
}

// GetKineticScrolling is a wrapper around gtk_scrolled_window_get_kinetic_scrolling().
func (m *ScrolledWindow) GetKineticScrolling() bool {
	r := gtk3.SysCall("gtk_scrolled_window_get_kinetic_scrolling", m.Instance())
	return ToGoBool(r)
}

// SetKineticScrolling is a wrapper around gtk_scrolled_window_set_kinetic_scrolling().
func (m *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	gtk3.SysCall("gtk_scrolled_window_set_kinetic_scrolling", m.Instance(), ToCBool(kineticScrolling))
}

// GetCaptureButtonPress is a wrapper around gtk_scrolled_window_get_capture_button_press().
func (m *ScrolledWindow) GetCaptureButtonPress() bool {
	r := gtk3.SysCall("gtk_scrolled_window_get_capture_button_press", m.Instance())
	return ToGoBool(r)
}

// SetCaptureButtonPress is a wrapper around gtk_scrolled_window_set_capture_button_press().
func (m *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	gtk3.SysCall("gtk_scrolled_window_set_capture_button_press", m.Instance(), ToCBool(captureButtonPress))
}

// GetMinContentWidth is a wrapper around gtk_scrolled_window_get_min_content_width().
func (m *ScrolledWindow) GetMinContentWidth() int {
	r := gtk3.SysCall("gtk_scrolled_window_get_min_content_width", m.Instance())
	return int(r)
}

// SetMinContentWidth is a wrapper around gtk_scrolled_window_set_min_content_width().
func (m *ScrolledWindow) SetMinContentWidth(width int) {
	gtk3.SysCall("gtk_scrolled_window_set_min_content_width", m.Instance(), uintptr(width))
}

// GetMinContentHeight is a wrapper around gtk_scrolled_window_get_min_content_height().
func (m *ScrolledWindow) GetMinContentHeight() int {
	r := gtk3.SysCall("gtk_scrolled_window_get_min_content_height", m.Instance())
	return int(r)
}

// SetMinContentHeight is a wrapper around gtk_scrolled_window_set_min_content_height().
func (m *ScrolledWindow) SetMinContentHeight(height int) {
	gtk3.SysCall("gtk_scrolled_window_set_min_content_height", m.Instance(), uintptr(height))
}
