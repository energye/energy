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
	. "github.com/energye/energy/v3/pkgs/gtk3/types"
	"unsafe"
)

type Widget struct {
	Object
}

func AsWidget(ptr unsafe.Pointer) IWidget {
	if ptr == nil {
		return nil
	}
	m := new(Widget)
	m.instance = ptr
	return m
}

// GetScreen is a wrapper around gtk_widget_get_screen().
func (m *Widget) GetScreen() IScreen {
	r := gtk3.SysCall("gtk_widget_get_screen", m.Instance())
	if r == 0 {
		return nil
	}
	s := &Screen{}
	s.instance = unsafe.Pointer(r)
	return s
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (m *Widget) SetVisual(visual IVisual) {
	gtk3.SysCall("gtk_widget_set_visual", m.Instance(), visual.Instance())
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (m *Widget) SetAppPaintable(paintable bool) {
	gtk3.SysCall("gtk_widget_set_app_paintable", m.Instance(), ToCBool(paintable))
}

// GetName is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (m *Widget) GetName() string {
	r := gtk3.SysCall("gtk_widget_get_name", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (m *Widget) GetAllocation() IRectangle {
	var rect Rectangle
	gtk3.SysCall("gtk_widget_get_allocation", m.Instance(), uintptr(unsafe.Pointer(&rect)))
	return &rect
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (m *Widget) SetSizeRequest(width, height int) {
	gtk3.SysCall("gtk_widget_set_size_request", m.Instance(), uintptr(width), uintptr(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (m *Widget) GetSizeRequest() (width, height int) {
	gtk3.SysCall("gtk_widget_get_size_request", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (m *Widget) GetStyleContext() IStyleContext {
	r := gtk3.SysCall("gtk_widget_get_style_context", m.Instance())
	return AsStyleContext(unsafe.Pointer(r))
}
